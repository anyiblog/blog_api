package model

// GoodsSpec 规格表
type GoodsSpec struct {
	GoodsSpecID string `gorm:"primary_key;column:goods_spec_id;type:char(36);not null" json:"GoodsSpecID"`
	SpecNameID  string `gorm:"column:spec_name_id;type:char(36);not null" json:"SpecNameID"`   // 规格
	SpecValueID string `gorm:"column:spec_value_id;type:char(36);not null" json:"SpecValueID"` // 规格值
	SkuID       string `gorm:"column:sku_id;type:char(36);not null" json:"SkuID"`              // 关联商品
}


