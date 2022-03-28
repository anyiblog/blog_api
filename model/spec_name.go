package model

// SpecName 规格名表
type SpecName struct {
	SpecNameID string `gorm:"primary_key;column:spec_name_id;type:char(36);not null" json:"SpecNameID"` // 规格ID
	SpecName   string `gorm:"column:spec_name;type:varchar(50);not null" json:"SpecName"`               // 规格名称
}
