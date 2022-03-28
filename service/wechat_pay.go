package service

import (
	"anyiblog/model"
	"anyiblog/serializer"
	"anyiblog/util"
	"fmt"
	"github.com/go-pay/gopay"
	goPayUtil "github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/wechat"
	"os"
	"strconv"
	"time"
)

// MiniProgramRes InitWechatPay NewClientV3 初始化微信客户端 V3
//	appid：appid 或者服务商模式的 sp_appid
//	mchid：商户ID 或者服务商模式的 sp_mchid
// 	serialNo：商户证书的证书序列号
//	apiV3Key：apiV3Key，商户平台获取
//	pkContent：私钥 apiclient_key.pem 读取后的内容

type MiniProgramRes struct {
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

// InitWechatPay 统一下单
func InitWechatPay(orderNo, openID string, totalPrice float64, orderInfoID string) MiniProgramRes {
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	client := wechat.NewClient(os.Getenv("WeChat_AppId"), os.Getenv("Wechat_Pay_mch_ID"), os.Getenv("Wechat_Pay_APIkey"), true)
	_ = client.AddCertPkcs12FilePath("../cert/apiclient_cert.p12")
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", goPayUtil.GetRandomString(32)).
		Set("openid", openID).
		Set("timeStamp", timeStamp).
		Set("sign_type", wechat.SignType_MD5).
		Set("body", "伴宠行-小程序商品购买").
		Set("attach", orderInfoID). // 关联订单商品表
		Set("out_trade_no", orderNo).
		Set("total_fee", util.Yuan2Fen(totalPrice)).
		//Set("total_fee", 1).
		Set("spbill_create_ip", "127.0.0.1").
		Set("notify_url", os.Getenv("Wechat_Pay_Notify_Url")).
		Set("trade_type", "JSAPI")
	order, orderErr := client.UnifiedOrder(bm)
	fmt.Println(order)
	fmt.Println(orderErr)
	Packages := "prepay_id=" + order.PrepayId // 预下单后返回的ID
	// 构建小程序PaySign
	MiniProgramPaySign := wechat.GetMiniPaySign(os.Getenv("WeChat_AppId"), order.NonceStr, Packages, "MD5", timeStamp, os.Getenv("Wechat_Pay_APIkey"))
	return MiniProgramRes{
		TimeStamp: timeStamp,
		NonceStr:  order.NonceStr,
		Package:   Packages,
		SignType:  "MD5",
		PaySign:   MiniProgramPaySign,
	}
}

// AgainPay 再次支付
func AgainPay(orderNo, openID string, totalPrice float64, orderInfoID string) MiniProgramRes {
	fmt.Println(totalPrice)
	return InitWechatPay(orderNo, openID, totalPrice, orderInfoID)
}

type PaidServiceParams struct { //商品支付完成后续所需参数
	TimeEnd       string
	TotalFee      string
	OutTradeNo    string
	Attach        string
	TransactionId string
	IsSubscribe   string
	Openid        string
}

// PaidService 商品支付完成后处理事宜
// 更新订单信息表，更改SKU库存,记录SPU销量，创建微信支付表记录
func PaidService(wxRes PaidServiceParams) {
	// 结束支付时间
	payTimeEnd := util.StrTimeFormat(util.TimeFormat(util.StrTimeFormat(wxRes.TimeEnd, "20060102150405")), "2006-01-02 15:04:05")
	// 支付总费用（元）
	payPrice := util.Fen2Yuan(wxRes.TotalFee)

	// 更新订单信息表
	queryParams := make(map[string]interface{})
	queryParams["order_no"] = wxRes.OutTradeNo
	updateFields := make(map[string]interface{})
	updateFields["order_actual_price"] = payPrice
	updateFields["order_pay_time"] = payTimeEnd
	updateFields["order_status"] = model.OrderStatusPayOk
	model.UpdateOrderInfo(queryParams, updateFields)

	// 更改SKU库存，记录SPU销量
	orderInfoId := wxRes.Attach
	orderGoodsList := model.QueryOrderGoods(orderInfoId)
	for _, item := range orderGoodsList {
		model.SetProductSales(item.SpuID, item.ProductCount)
		model.SkuMinusStock(item.SkuID, item.ProductCount)
	}

	// 记录微信支付信息
	wechatPayInfo := model.WechatPay{
		WechatPayID: wxRes.TransactionId,
		OrderNo:     wxRes.OutTradeNo,
		PayAmount:   payPrice,
		IsSubscribe: wxRes.IsSubscribe,
		UserOpenID:  wxRes.Openid,
		PayTimeEnd:  payTimeEnd,
	}
	model.CreateWechatPay(wechatPayInfo)
}

// WechatPayQueryOrder 手动查询订单
func WechatPayQueryOrder(orderNo string) serializer.Response {
	client := wechat.NewClient(os.Getenv("WeChat_AppId"), os.Getenv("Wechat_Pay_mch_ID"), os.Getenv("Wechat_Pay_APIkey"), true)
	_ = client.AddCertPkcs12FilePath("../cert/apiclient_cert.p12")
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", orderNo).
		Set("nonce_str", goPayUtil.GetRandomString(32)).
		Set("sign_type", wechat.SignType_MD5)
	Resp, _, err := client.QueryOrder(bm)
	if err != nil {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  err.Error(),
		}
	}
	if Resp.ReturnCode == "FAIL" {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  Resp.ReturnMsg,
		}
	}
	if Resp.ReturnCode == "SUCCESS" && Resp.ResultCode == "FAIL" {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  Resp.ErrCode + Resp.ErrCodeDes,
		}
	}
	// 返回的数据字段只有在return_code 、result_code、trade_state都为SUCCESS时有返回
	if Resp.ReturnCode == "SUCCESS" && Resp.ResultCode == "SUCCESS" && Resp.TradeState == "SUCCESS" {
		// 查询订单是否已经支付，并且微信返回订单金额是否与数据库一致
		if model.QueryOrderNotPay(Resp.TransactionId, Resp.OutTradeNo) && model.TotalPriceEqualWxResPrice(Resp.OutTradeNo, util.Fen2Yuan(Resp.TotalFee)) {
			PaidService(PaidServiceParams{
				TimeEnd:       Resp.TimeEnd,
				TotalFee:      Resp.TotalFee,
				OutTradeNo:    Resp.OutTradeNo,
				Attach:        Resp.Attach,
				TransactionId: Resp.TransactionId,
				IsSubscribe:   Resp.IsSubscribe,
				Openid:        Resp.Openid,
			})
			return serializer.Response{
				Code: serializer.SystemError,
				Msg:  "当前订单已更新",
			}
		} else {
			return serializer.Response{
				Code: serializer.SystemError,
				Msg:  "当前订单无需更新",
			}
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "交易状态：" + Resp.TradeState,
		}
	}
}
