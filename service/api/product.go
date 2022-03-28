package api

import (
	"anyiblog/model"
	"anyiblog/serializer"
	ApiParams "anyiblog/serializer/params/api"
	"anyiblog/service"
	"anyiblog/util"
	"github.com/guonaihong/gout"
	"github.com/tidwall/gjson"
)

// QueryProductService 获取产品
func QueryProductService(params *ApiParams.QueryProductParam) serializer.Response {
	queryParams := make(map[string]interface{}) // 定值查询
	likeParams := make(map[string]interface{})  // 模糊查询
	if len(params.ProductName) > 0 {
		likeParams["spu_name"] = params.ProductName
	}
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}
	return serializer.Response{
		Code: serializer.SystemOk,
		Msg:  "获取成功",
		Data: model.GetAllGoods(queryParams, likeParams, params.Limit, params.Page, false),
	}
}

// CosImgInfo 获取cos对应的图片信息
type CosImgInfo struct {
	Format   string `json:"format"`
	Width    string `json:"width"`
	Height   string `json:"height"`
	Size     string `json:"size"`
	Md5      string `json:"md5"`
	PhotoRgb string `json:"photo_rgb"`
}

// ResDetailContent 商品详情图片信息
type ResDetailContent struct {
	Url    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

// GetProductDetail 获取某个产品详情
func GetProductDetail(spuID string) serializer.Response {
	var imgInfo CosImgInfo
	var detailContent []ResDetailContent
	isOk, detail := model.GetGoodsDetail(spuID)
	jsonArr := gjson.Parse(detail.DetailContent)
	for _, v := range jsonArr.Array() {
		_ = gout.GET(v.String() + "?imageInfo").Debug(true).BindJSON(&imgInfo).Do()
		detailContent = append(detailContent, ResDetailContent{
			Url:    v.String(),
			Width:  imgInfo.Width,
			Height: imgInfo.Height,
		})
	}
	detail.DetailContent = util.StructToJson(detailContent)
	if isOk {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "获取商品详情成功",
			Data: detail,
		}
	}
	return serializer.Response{
		Code: serializer.SystemError,
		Msg:  "获取商品详情错误",
	}
}

// AgainPayService 重新付款
func AgainPayService(userID, orderInfoID string) serializer.Response {
	openID := model.GetOpenIdByUserId(userID)
	orderInfoQueryFields := []string{"order_no", "order_total_price"} // 查询字段
	orderInfoQueryParam := make(map[string]interface{})                // 查询条件
	orderInfoQueryParam["order_info_id"] = orderInfoID
	orderInfoRes := model.GetOrderInfoByQuery(orderInfoQueryFields, orderInfoQueryParam).(map[string]interface{})
	miniProgramPayRes := service.AgainPay(orderInfoRes["order_no"].(string), openID, orderInfoRes["order_total_price"].(float64), orderInfoID)
	return serializer.Response{
		Code: serializer.SystemOk,
		Data: miniProgramPayRes,
	}
}

//BuyProduct 购买产品，创建订单
func BuyProduct(params *ApiParams.BuyProductParam, userId string) serializer.Response {
	orderInfoID := util.GenerateUUID() // 订单信息ID
	var allOrderCount int              // 商品数量
	var allProductPrice float64        // 商品总额
	var allOrderFreightPrice float64   // 运费
	var allOrderTotalPrice float64     // 应付总额
	//var allOrderActualPrice float64 // 实付总额

	skuFields := []string{"sku_user_price", "SkuWeight"} // 查询字段
	skuQueryParam := make(map[string]interface{})        // 查询条件
	pList := params.ProductList
	for _, productItem := range pList {
		skuQueryParam["sku_id"] = productItem.SkuID
		allOrderCount = allOrderCount + productItem.BuyCount
		skuInfoRes := model.GetSkuInfoByQuery(skuFields, skuQueryParam).(map[string]interface{})
		allProductPrice = allProductPrice + (skuInfoRes["sku_user_price"].(float64) * float64(productItem.BuyCount))
		allOrderFreightPrice = 0
		model.CreateOrderGoods(orderInfoID, productItem.SpuID, productItem.SkuID, skuInfoRes["sku_user_price"].(float64), allOrderTotalPrice, productItem.BuyCount)
	}
	allOrderTotalPrice = allOrderFreightPrice + allProductPrice // 应付总额 = 运费+商品总额
	createOrderPrams := model.CreateOrderPrams{
		OrderInfoID:       orderInfoID,
		UserID:            userId,
		OrderCount:        allOrderCount,
		OrderProductPrice: allProductPrice,
		OrderFreightPrice: allOrderFreightPrice,
		OrderTotalPrice:   allOrderTotalPrice,
		OrderAddressID:    params.AddressID,
		OtherJSON:         params.OtherInfo,
	}
	ok, orderNo := model.CreateOrder(createOrderPrams)
	payRes := service.InitWechatPay(orderNo, model.GetOpenIdByUserId(userId), allOrderTotalPrice, orderInfoID)
	if ok {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "创建订单成功",
			Data: payRes,
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "提交订单错误",
		}
	}
}
