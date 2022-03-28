package model

import (
	"anyiblog/conf"
	"anyiblog/util"
	"time"
)

// Brand 品牌表
type Brand struct {
	BrandID   string    `gorm:"primary_key;column:brand_id;type:char(36);not null" json:"BrandID"` // 品牌ID
	BrandName string    `gorm:"column:brand_name;type:varchar(50);not null" json:"BrandName"`      // 品牌名称
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"CreatedAt"`                 // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp" json:"UpdatedAt"`                 // 更新时间
}

// ResBrandList 所有品牌列表
type ResBrandList struct {
	Count     int64   `json:"Count"`
	BrandList []Brand `json:"BrandList"`
}

func GetAllBrand(limit, page int) ResBrandList {
	var resBrandList ResBrandList
	conf.DB.Model(Brand{}).Order("created_at desc").Limit(limit).Offset((page - 1) * limit).Find(&resBrandList.BrandList)
	conf.DB.Model(Brand{}).Count(&resBrandList.Count)
	return resBrandList
}

func GetBrandName(BrandID string) (bool, string) {
	var brandName struct {
		BrandName string
	}
	if conf.DB.Model(Brand{}).Select("brand_name").Where("brand_id = ?", BrandID).Scan(&brandName).RowsAffected > 0 {
		return true, brandName.BrandName
	} else {
		return false, ""
	}
}

func GetBrandID(BrandName string) (bool, string) {
	var brandID struct {
		BrandID string
	}
	if conf.DB.Model(Brand{}).Select("brand_id").Where("brand_name = ?", BrandName).Scan(&brandID).RowsAffected > 0 {
		return false, brandID.BrandID
	} else {
		return false, ""
	}
}

func CreateCakeBrand(brandName string) bool {
	branID := util.GenerateUUID()
	b := Brand{
		BrandID:   branID,
		BrandName: brandName,
		CreatedAt: util.NowTime(),
		UpdatedAt: util.NowTime(),
	}
	if conf.DB.Create(b).RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func UpdateCakeBrand(brandName, brandId string) bool {
	data := make(map[string]interface{})
	data["brand_name"] = brandName
	data["updated_at"] = util.NowTime()
	if conf.DB.Model(Brand{}).Where("brand_id = ?", brandId).Updates(data).RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func DeleteCakeBrand(brandId string) bool {
	if conf.DB.Where("brand_id = ?", brandId).Delete(Brand{}).RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
