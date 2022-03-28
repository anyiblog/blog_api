package model

// SpecValue 规格值表
type SpecValue struct {
	SpecValueID string `gorm:"primary_key;column:spec_value_id;type:char(36);not null" json:"SpecValueID"`
	SpecNameID  string `gorm:"column:spec_name_id;type:char(36);not null" json:"specNameID"` // 规格名ID
	SpecValue   string `gorm:"column:spec_value;type:varchar(50);not null" json:"SpecValue"` // 规格值
}