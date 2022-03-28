package model

import (
	"anyiblog/conf"
	adminParams "anyiblog/serializer/params/admin"
	"anyiblog/util"
)

// GoodsSku Sku表（具体的某个商品）
type GoodsSku struct {
	SkuID        string  `gorm:"primary_key;column:sku_id;type:char(36);not null" json:"SkuID"`         // 商品主键ID
	SkuNo        string  `gorm:"column:sku_no;type:char(20);not null" json:"SkuNo"`                     // 商品编号，唯一（11位）
	SpuID        string  `gorm:"column:spu_id;type:char(36);not null" json:"-"`                         // 产品ID（spuId）
	SkuImg       string  `gorm:"column:sku_img;type:char(36);not null" json:"SkuImg"`                   // 商品图片（imgId）
	SkuVipPrice  float64 `gorm:"column:sku_vip_price;type:decimal(18,2);not null" json:"SkuVipPrice"`   // 会员售价
	SkuUserPrice float64 `gorm:"column:sku_user_price;type:decimal(18,2);not null" json:"SkuUserPrice"` // 普通用户售价
	SkuPrice     float64 `gorm:"column:sku_price;type:decimal(18,2);not null" json:"SkuPrice"`          // 市场价，营销用
	SkuSales     int     `gorm:"column:sku_sales;type:int unsigned;not null" json:"SkuSales"`           // 商品销量
	SkuStock     int     `gorm:"column:sku_stock;type:int;not null" json:"SkuStock"`                    // 商品库存
	SkuWeight    float64 `gorm:"column:sku_weight;type:float(10,2);not null" json:"SkuWeight"`          // 商品重量
	SkuSpec      string  `gorm:"column:sku_spec;type:longtext;not null" json:"sku"`                     // sku属性
}


// GetGoodsSkuByQuery 根据查询字段名和查询参数获取结果
func GetGoodsSkuByQuery(Fields interface{}, QueryParams map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{},0)
	conf.DB.Model(GoodsSku{}).Select(Fields).Where(QueryParams).Find(&result)
	return result
}

func AddGoodsSku(skuInfo []adminParams.SkuInfo, SpuID string) bool {
	var skuID string
	var counter int
	for _, v := range skuInfo {
		skuID = util.GenerateUUID()
		goodsSku := GoodsSku{
			SkuID:        skuID,
			SkuNo:        util.RandStr(11),
			SpuID:        SpuID,
			SkuImg:       v.SkuImg,
			SkuVipPrice:  v.SkuVipPrice,
			SkuUserPrice: v.SkuUserPrice,
			SkuPrice:     v.SkuPrice,
			SkuSales:     0,
			SkuStock:     v.SkuStock,
			SkuWeight:    v.SkuWeight,
			SkuSpec:      v.SkuSpec,
		}
		if conf.DB.Create(goodsSku).RowsAffected > 0 {
			counter = counter + 1
		}
	}
	if counter == len(skuInfo) {
		return true
	} else {
		return false
	}
}

// GetAllSku 根据产品ID获取该产品的所有sku
func GetAllSku(SpuID string) []GoodsSku {
	var skus []GoodsSku
	conf.DB.Model(GoodsSku{}).Where("spu_id = ?", SpuID).Find(&skus)
	return skus
}

// GetSkuInfoByQuery 根据查询字段名和查询参数获取结果
func GetSkuInfoByQuery(Fields interface{}, QueryParams map[string]interface{}) interface{} {
	result := map[string]interface{}{}
	conf.DB.Model(GoodsSku{}).Select(Fields).Where(QueryParams).First(&result)
	return result
}

// SkuMinusStock 减库存，更新销量
func SkuMinusStock(skuId string, num int) bool {
	var stockAndSales struct {
		SkuSales int
		SkuStock int
	}
	conf.DB.Model(GoodsSku{}).Select("sku_sales,sku_stock").Where("sku_id = ?", skuId).Find(&stockAndSales)
	upSales := conf.DB.Model(GoodsSku{}).Where("sku_id = ?", skuId).Update("sku_sales", stockAndSales.SkuSales+num)
	upStock := conf.DB.Model(GoodsSku{}).Where("sku_id = ?", skuId).Update("sku_stock", stockAndSales.SkuStock-num)
	if upSales.RowsAffected > 0 && upStock.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
