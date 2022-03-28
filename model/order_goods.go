package model

import (
	"anyiblog/conf"
	"anyiblog/util"
)

// OrderGoods 订单商品表
type OrderGoods struct {
	OrderGoodsID string  `gorm:"primary_key;column:order_goods_id;type:char(36);not null" json:"order_goods_id"` // 订单商品ID
	OrderInfoID  string  `gorm:"column:order_info_id;type:char(36);not null" json:"order_no"`                    // 订单ID（关联订单信息表）
	SpuID        string  `gorm:"column:spu_id;type:char(36);not null" json:"spu_id"`                             // 产品ID
	SkuID        string  `gorm:"column:sku_id;type:char(36);not null" json:"sku_id"`                             // 商品ID
	UintPrice    float64 `gorm:"column:uint_price;type:decimal(18,2);not null" json:"uint_price"`                // 商品单价
	ProductCount int     `gorm:"column:product_count;type:int(10);not null" json:"product_count"`                // 商品数量
	TotalPrice   float64 `gorm:"column:total_price;type:decimal(10,2);not null" json:"total_price"`              // 商品总价
}

// GetOrderGoodsByQuery 根据查询字段名和查询参数获取结果
func GetOrderGoodsByQuery(Fields interface{}, QueryParams map[string]interface{}) []map[string]interface{} {
	result := make([]map[string]interface{},0)
	conf.DB.Model(OrderGoods{}).Select(Fields).Where(QueryParams).Find(&result)
	return result
}

func CreateOrderGoods(OrderInfoID, SpuID, SkuID string, UintPrice, TotalPrice float64, ProductCount int) {
	orderGoodsId := util.GenerateUUID()
	orderGoods := OrderGoods{
		OrderGoodsID: orderGoodsId,
		OrderInfoID:  OrderInfoID,
		SpuID:        SpuID,
		SkuID:        SkuID,
		UintPrice:    UintPrice,
		ProductCount: ProductCount,
		TotalPrice:   TotalPrice,
	}
	conf.DB.Create(orderGoods)
}

func QueryOrderGoods(orderInfoID string) []OrderGoods {
	var orderGoodsList []OrderGoods
	conf.DB.Where("order_info_id = ?", orderInfoID).Find(&orderGoodsList)
	return orderGoodsList
}
