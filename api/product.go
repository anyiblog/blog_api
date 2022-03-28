package api

import (
	"anyiblog/model"
	"anyiblog/serializer"
	apiParams "anyiblog/serializer/params/api"
	"anyiblog/service"
	apiService "anyiblog/service/api"
	"anyiblog/util"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat"
	"net/http"
	"os"
)

// GetProducts 获取产品(根据查询参数)
func GetProducts(c *gin.Context) {
	queryProductParam := &apiParams.QueryProductParam{}
	if err := c.ShouldBindJSON(&queryProductParam); err == nil {
		res := apiService.QueryProductService(queryProductParam)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: gin.H{"error": err.Error()},
		})
	}
}

// GetProductDetail 获取产品详情
func GetProductDetail(c *gin.Context) {
	spuID := c.Query("SpuID")
	res := apiService.GetProductDetail(spuID)
	c.JSON(200, res)
}

// BuyProduct 购买产品
func BuyProduct(c *gin.Context) {
	userId, _ := c.Get("userId")
	buyProductParam := apiParams.BuyProductParam{}
	if err := c.ShouldBindJSON(&buyProductParam); err == nil {
		res := apiService.BuyProduct(&buyProductParam, userId.(string))
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

func AgainPay(c *gin.Context) {
	userId, _ := c.Get("userId")
	orderInfoID := c.Query("OrderInfoID")
	res := apiService.AgainPayService(userId.(string), orderInfoID)
	c.JSON(200, res)
}

// WechatPayNotifyUrl 支付后，异步回调
// 对于支付结果通知的内容一定要做签名验证,并校验返回的订单金额是否与商户侧的订单金额一致，防止数据泄漏导致出现“假通知”，造成资金损失。
func WechatPayNotifyUrl(c *gin.Context) {
	notifyReq, _ := wechat.ParseNotifyToBodyMap(c.Request)
	signOk, _ := wechat.VerifySign(os.Getenv("Wechat_Pay_APIkey"), "MD5", notifyReq)
	rsp := new(wechat.NotifyResponse)
	if signOk { // 验签成功，回复微信支付成功
		wxRes := wechat.NotifyRequest{}
		if err := c.ShouldBindXML(&wxRes); err == nil {
			wechatPayId := wxRes.TransactionId
			orderNo := wxRes.OutTradeNo
			if model.QueryOrderNotPay(wechatPayId, orderNo) && model.TotalPriceEqualWxResPrice(wxRes.OutTradeNo, util.Fen2Yuan(wxRes.TotalFee)) { // 查询订单是否已经支付，并且微信返回订单金额是否与数据库一致
				service.PaidService(service.PaidServiceParams{
					TimeEnd:       wxRes.TimeEnd,
					TotalFee:      wxRes.TotalFee,
					OutTradeNo:    wxRes.OutTradeNo,
					Attach:        wxRes.Attach,
					TransactionId: wxRes.TransactionId,
					IsSubscribe:   wxRes.IsSubscribe,
					Openid:        wxRes.Openid,
				})
				rsp.ReturnCode = gopay.SUCCESS
				rsp.ReturnMsg = gopay.OK
				c.Abort()
				c.String(http.StatusOK, "%s", rsp.ToXmlString())
			}
		}
	} else {
		rsp.ReturnCode = gopay.FAIL
		rsp.ReturnMsg = gopay.NULL
		c.Abort()
		c.String(http.StatusOK, "%s", rsp.ToXmlString())
	}
}

func PayQueryOrder(c *gin.Context) {
	orderNo := c.Query("OrderNo")
	res := service.WechatPayQueryOrder(orderNo)
	c.JSON(200, res)
}

func GetDiscountList(c *gin.Context) {

}

func RecommendProductList(c *gin.Context) {

}
