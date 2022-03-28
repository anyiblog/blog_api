package model

import (
	"anyiblog/conf"
	adminParams "anyiblog/serializer/params/admin"
	"anyiblog/util"
	"time"
)

// Shops 线下门店表
type Shops struct {
	ShopID           string    `gorm:"primary_key;column:shop_id;type:char(36);not null" json:"shop_id"`             // 门店ID
	ShopName         string    `gorm:"column:shop_name;type:char(25);not null" json:"shop_name"`                     // 门店名称
	ShopLogo         string    `gorm:"column:shop_logo;type:char(36);not null" json:"shop_logo"`                     // 门店logo（关联图片表ID）
	ShopAddress      string    `gorm:"column:shop_address;type:varchar(255);not null" json:"shop_address"`           // 门店地址
	ShopContactName  string    `gorm:"column:shop_contact_name;type:varchar(255);not null" json:"shop_contact_name"` // 门店联系人
	ShopContactPhone string    `gorm:"column:shop_contact_phone;type:char(11);not null" json:"shop_contact_phone"`   // 门店联系电话
	ShopTime         string    `gorm:"column:shop_time;type:varchar(255);not null" json:"shop_time"`                 // 营业时间
	ShopOtherInfo    string    `gorm:"column:shop_other_info;type:varchar(255)" json:"shop_other_info"`              // 备注信息
	CreatedAt        time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`                           // 店铺注册时间
}

func GetShopInfoById(id string) (Shops, bool) {
	var resShopInfo Shops
	var i int64
	i = conf.DB.Model(Shops{}).Where("shop_id = ?", id).Find(&resShopInfo).RowsAffected
	resShopInfo.ShopLogo = GetImgUrl(resShopInfo.ShopLogo)
	if i > 0 {
		return resShopInfo, true
	} else {
		return resShopInfo, false
	}
}

func CreateShop(params *adminParams.CreateShopParams) (Shops, bool) {
	shopId := util.GenerateUUID()
	shop := Shops{
		ShopID:           shopId,
		ShopName:         params.ShopName,
		ShopLogo:         params.ShopLogo,
		ShopAddress:      params.ShopAddress,
		ShopContactName:  params.ShopContactName,
		ShopContactPhone: params.ShopContactPhone,
		ShopTime:         params.ShopTime,
		ShopOtherInfo:    params.ShopOtherInfo,
		CreatedAt:        util.NowTime(),
	}
	if conf.DB.Create(shop).RowsAffected > 0 {
		shop.ShopLogo = GetImgUrl(shop.ShopLogo)
		return shop, true
	} else {
		return shop, false
	}
}
