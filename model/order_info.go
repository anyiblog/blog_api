package model

import (
	"anyiblog/conf"
	"anyiblog/util"
	"time"
)

const (
	OrderStatusFiled        = 0 // 待付款
	OrderStatusPayOk        = 1 // 已支付待发货
	OrderStatusToBeReceived = 2 // 待收货
	OrderStatusSuccess      = 3 // 已完成

	//// NotDelivery 配送方式
	//NotDelivery = 0 // 未配送
	//NotDelivery = 0 // 未配送
)

// OrderInfo 订单信息表
type OrderInfo struct {
	OrderInfoID       string    `gorm:"primary_key;column:order_info_id;type:char(36);not null" json:"OrderInfoID"`                        // 订单信息ID
	OrderNo           string    `gorm:"column:order_no;type:char(20);not null" json:"OrderNo"`                                             // 订单编号（10位）
	UserID            string    `gorm:"column:user_id;type:char(36);not null" json:"-"`                                                    // 用户ID
	OrderCount        int       `gorm:"column:order_count;type:int(11);not null" json:"OrderCount"`                                        // 订单商品的总数量
	OrderTotalPrice   float64   `gorm:"column:order_total_price;type:decimal(18,2) unsigned zerofill;not null" json:"OrderTotalPrice"`     // 订单总额（应付款）
	OrderActualPrice  float64   `gorm:"column:order_actual_price;type:decimal(18,2) unsigned zerofill" json:"OrderActualPrice"`            // 实付总额
	OrderProductPrice float64   `gorm:"column:order_product_price;type:decimal(18,2) unsigned zerofill;not null" json:"OrderProductPrice"` // 商品总额
	OrderFreightPrice float64   `gorm:"column:order_freight_price;type:decimal(18,2) unsigned zerofill;not null" json:"OrderFreightPrice"` // 订单运费
	OrderCreateTime   time.Time `gorm:"column:order_create_time;type:timestamp;not null" json:"OrderCreateTime"`                           // 生成订单时间
	OrderAddressID    string    `gorm:"column:order_address_id;type:char(36);not null" json:"-"`                                           // 订单收货地址（外键ID）
	OrderPayMethod    int       `gorm:"column:order_pay_method;type:int(11)" json:"OrderPayMethod"`                                        // 订单支付方式
	OrderPayTime      time.Time `gorm:"column:order_pay_time;type:timestamp" json:"OrderPayTime"`                                          // 订单支付时间
	OrderDelivery     int       `gorm:"column:order_delivery;type:int(11)" json:"OrderDelivery"`                                           // 订单配送方式
	OrderStatus       int       `gorm:"column:order_status;type:int(10) unsigned;not null" json:"OrderStatus"`                             // 订单状态
	OrderEndTime      time.Time `gorm:"column:order_end_time;type:timestamp" json:"OrderEndTime"`                                          // 订单完成时间
	OrderCloseTime    time.Time `gorm:"column:order_close_time;type:timestamp" json:"OrderCloseTime"`                                      // 订单关闭时间
	OtherJSON         string    `gorm:"column:other_json;type:json;not null" json:"OtherJSON"`                                             // 订单备注信息，json格式
}

// GetOrderInfoByQuery 根据查询字段名和查询参数获取结果
func GetOrderInfoByQuery(Fields interface{}, QueryParams map[string]interface{}) interface{} {
	result := map[string]interface{}{}
	conf.DB.Model(OrderInfo{}).Select(Fields).Where(QueryParams).First(&result)
	return result
}

type CreateOrderPrams struct {
	OrderInfoID       string
	UserID            string
	OrderCount        int
	OrderProductPrice float64 // 商品总额
	//OrderActualPrice float64 // 实付总额
	OrderFreightPrice float64 // 订单运费
	OrderTotalPrice   float64 // 应付总额
	OrderAddressID    string
	OtherJSON         string
}

// CreateOrder 创建订单
func CreateOrder(orderInfoParams CreateOrderPrams) (bool, string) {
	OrderNo := util.GenerateOrderNo()
	orderInfo := OrderInfo{
		OrderInfoID:       orderInfoParams.OrderInfoID,
		OrderNo:           OrderNo,
		UserID:            orderInfoParams.UserID,
		OrderCount:        orderInfoParams.OrderCount,
		OrderTotalPrice:   orderInfoParams.OrderTotalPrice,
		OrderProductPrice: orderInfoParams.OrderProductPrice,
		OrderFreightPrice: orderInfoParams.OrderFreightPrice,
		OrderCreateTime:   util.NowTime(),
		OrderAddressID:    orderInfoParams.OrderAddressID,
		OrderPayMethod:    0,
		OrderDelivery:     0,
		OrderStatus:       OrderStatusFiled,
		OtherJSON:         orderInfoParams.OtherJSON,
	}
	if conf.DB.Create(orderInfo).RowsAffected > 0 {
		return true, OrderNo
	} else {
		return false, ""
	}
}

// UpdateOrderInfo 更改订单信息字段
func UpdateOrderInfo(QueryParams map[string]interface{}, Fields map[string]interface{}) bool {
	if conf.DB.Model(OrderInfo{}).Where(QueryParams).Updates(Fields).RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// TotalPriceEqualWxResPrice
// 校验应付总额是否等于微信返回的已支付金额，如果等于代表订单支付成功，并更改订单状态
func TotalPriceEqualWxResPrice(OrderNo string, WxResPrice float64) bool {
	var oderPrice struct {
		OrderTotalPrice float64
	}
	conf.DB.Model(OrderInfo{}).Select("order_total_price").Where("order_no = ?", OrderNo).Find(&oderPrice)
	if oderPrice.OrderTotalPrice == WxResPrice {
		return true
	} else {
		return false
	}
}

type OrderInfoClient struct {
	OrderInfoID       string  `json:"orderInfoID"` // 订单信息ID
	OrderNo           string  `json:"OrderNo"`
	OrderStatus       int     `json:"OrderStatus"`
	OrderCount        int     `json:"OrderCount"`         // 订单商品的总数量
	OrderTotalPrice   float64 `json:"OrderTotalPrice"`    // 订单总额（应付款）
	OrderActualPrice  float64 ` json:"OrderActualPrice"`  // 实付总额
	OrderProductPrice float64 ` json:"OrderProductPrice"` // 商品总额
	OrderFreightPrice float64 `json:"OrderFreightPrice"`  // 订单运费
}
type OrderProductItemClient struct {
	ProductImg   string  `json:"ProductImg"`
	ProductName  string  `json:"ProductName"`
	ProductSpec  string  `json:"ProductSpec"`
	ProductPrice float64 `json:"ProductPrice"`
	ProductCount int     `json:"ProductCount"`
}
type ResOrderListClient struct {
	OrderInfoClient
	ProductList []OrderProductItemClient `json:"ProductList" gorm:"-"`
}

type ResOrderListCount struct {
	Count     int64                `json:"Count"`
	OrderList []ResOrderListClient `json:"OrderList"`
}

// QueryOrderBriefInfoByUserID 查询用户订单列表
func QueryOrderBriefInfoByUserID(userId string, page, limit, queryType int) (ResOrderListCount, bool) {
	var resOrderInfo []ResOrderListClient    // 数据库数据结构体
	var returnOrderInfo []ResOrderListClient // 订单信息结构体
	var allOrderListCount int64
	if queryType == -1 { // -1 查询全部
		db := conf.DB.Model(OrderInfo{}).Where("user_id = ?", userId).Order("order_create_time desc").Offset((page - 1) * limit).Limit(limit).Find(&resOrderInfo)
		if db.RowsAffected <= 0 {
			return ResOrderListCount{}, false
		}
	} else {
		db := conf.DB.Model(OrderInfo{}).Where("user_id = ? AND order_status = ?", userId, queryType).Order("order_create_time desc").Offset((page - 1) * limit).Limit(limit).Find(&resOrderInfo)
		if db.RowsAffected <= 0 {
			return ResOrderListCount{}, false
		}
	}
	conf.DB.Model(OrderInfo{}).Where("user_id = ?", userId).Count(&allOrderListCount)
	for _, orderItem := range resOrderInfo {
		orderGoodsQueryFields := []string{"spu_id", "sku_id", "uint_price", "product_count"} // 查询字段
		orderGoodsQueryParam := make(map[string]interface{})                                 // 查询条件
		orderGoodsQueryParam["order_info_id"] = orderItem.OrderInfoID
		orderGoodsRes := GetOrderGoodsByQuery(orderGoodsQueryFields, orderGoodsQueryParam)
		for _, orderGoodsResItem := range orderGoodsRes {
			goodsSkuQueryFields := []string{"sku_spec"}        // 查询字段
			goodsSkuQueryParam := make(map[string]interface{}) // 查询条件
			goodsSkuQueryParam["sku_id"] = orderGoodsResItem["sku_id"]
			goodsSkuRes := GetGoodsSkuByQuery(goodsSkuQueryFields, goodsSkuQueryParam)

			goodsSpuQueryFields := []string{"spu_img", "spu_name"} // 查询字段
			goodsSpuQueryParam := make(map[string]interface{})     // 查询条件
			goodsSpuQueryParam["spu_id"] = orderGoodsResItem["spu_id"]
			goodsSpuRes := GetSpuInfoByQuery(goodsSpuQueryFields, goodsSpuQueryParam)

			orderItem.ProductList = append(orderItem.ProductList, OrderProductItemClient{
				ProductImg:   goodsSpuRes["spu_img"].(string),
				ProductName:  goodsSpuRes["spu_name"].(string),
				ProductSpec:  goodsSkuRes["sku_spec"].(string),
				ProductPrice: orderGoodsResItem["uint_price"].(float64),
				ProductCount: orderGoodsResItem["product_count"].(int),
			})
		}
		returnOrderInfo = append(returnOrderInfo, orderItem)
	}
	return ResOrderListCount{
		Count:     allOrderListCount,
		OrderList: returnOrderInfo,
	}, true
}

type ResOrderDetailClient struct {
	OrderInfo
	WechatPayID string
	AddressInfo UserAddress
	ProductList []OrderProductItemClient `json:"ProductList" gorm:"-"`
}

// QueryOrderDetailInfoByUserID 查询用户订单详情信息
func QueryOrderDetailInfoByUserID(userId, orderInfoID string) ResOrderDetailClient {
	var resOrderDetail OrderInfo // 数据库数据结构体
	var resOrderAddress UserAddress
	var productItem []OrderProductItemClient // 数据库数据结构体
	var returnOrderInfo ResOrderDetailClient // 返回结构体

	conf.DB.Model(OrderInfo{}).Where("user_id = ? AND order_info_id = ?", userId, orderInfoID).Find(&resOrderDetail)

	wechatPayID := GetPayIDByOrderNo(resOrderDetail.OrderNo)
	resOrderAddress = GetAddressInfoByAddressID(userId, resOrderDetail.OrderAddressID)

	orderGoodsQueryFields := []string{"spu_id", "sku_id", "uint_price", "product_count"} // 查询字段
	orderGoodsQueryParam := make(map[string]interface{})                                 // 查询条件
	orderGoodsQueryParam["order_info_id"] = resOrderDetail.OrderInfoID
	orderGoodsRes := GetOrderGoodsByQuery(orderGoodsQueryFields, orderGoodsQueryParam)
	for _, orderGoodsResItem := range orderGoodsRes {
		goodsSkuQueryFields := []string{"sku_spec"}        // 查询字段
		goodsSkuQueryParam := make(map[string]interface{}) // 查询条件
		goodsSkuQueryParam["sku_id"] = orderGoodsResItem["sku_id"]
		goodsSkuRes := GetGoodsSkuByQuery(goodsSkuQueryFields, goodsSkuQueryParam)

		goodsSpuQueryFields := []string{"spu_img", "spu_name"} // 查询字段
		goodsSpuQueryParam := make(map[string]interface{})     // 查询条件
		goodsSpuQueryParam["spu_id"] = orderGoodsResItem["spu_id"]
		goodsSpuRes := GetSpuInfoByQuery(goodsSpuQueryFields, goodsSpuQueryParam)

		productItem = append(productItem, OrderProductItemClient{
			ProductImg:   GetImgUrl(goodsSpuRes["spu_img"].(string)),
			ProductName:  goodsSpuRes["spu_name"].(string),
			ProductSpec:  goodsSkuRes["sku_spec"].(string),
			ProductPrice: orderGoodsResItem["uint_price"].(float64),
			ProductCount: orderGoodsResItem["product_count"].(int),
		})
	}
	returnOrderInfo = ResOrderDetailClient{resOrderDetail, wechatPayID, resOrderAddress, productItem}
	return returnOrderInfo
}
