package model

import (
	"anyiblog/conf"
	"anyiblog/serializer/params/api"
	"anyiblog/util"
)

// UserAddress 用户收货地址表
type UserAddress struct {
	AddressID  string `gorm:"primary_key;column:address_id;type:char(36);not null" json:"AddressID"` // 地址ID
	UserID     string `gorm:"column:user_id;type:char(36);not null" json:"-"`                        // 用户ID
	UserName   string `gorm:"column:user_name;type:varchar(20);not null" json:"UserName"`            // 收货人
	Tel        string `gorm:"column:tel;type:varchar(20);not null" json:"Tel"`                       // 电话
	Province   string `gorm:"column:province;type:varchar(20);not null" json:"Province"`             // 一级地址
	City       string `gorm:"column:city;type:varchar(20);not null" json:"City"`                     // 二级地址
	County     string `gorm:"column:county;type:varchar(20);not null" json:"County"`                 // 三级地址
	DetailInfo string `gorm:"column:detail_info;type:varchar(255);not null" json:"DetailInfo"`       // 详细地址
	IsDefault  int    `gorm:"column:is_default;type:int;not null" json:"-"`                          // 是否默认（0不是，1是）
}

// CreateAddress 从参数结构体新增地址，如果存在则不添加
func CreateAddress(addressInfo api.SetAddressParams, userId string) UserAddress {
	addressID := util.GenerateUUID()
	address := UserAddress{
		AddressID:  addressID,
		UserID:     userId,
		UserName:   addressInfo.UserName,
		Tel:        addressInfo.TelNumber,
		Province:   addressInfo.ProvinceName,
		City:       addressInfo.CityName,
		County:     addressInfo.CountyName,
		DetailInfo: addressInfo.DetailInfo,
		IsDefault:  0,
	}
	if conf.DB.Create(&address).RowsAffected > 0 {
		return address
	} else {
		return UserAddress{}
	}
}

func GetAddressInfoByAddressID(userId, addressID string) UserAddress {
	var resAddress UserAddress
	conf.DB.Model(UserAddress{}).Where("user_id = ? AND address_id = ? ", userId, addressID).Find(&resAddress)
	return resAddress
}

// GetDefaultAddress 获取默认地址
func GetDefaultAddress(userId string) UserAddress {
	var resAddress UserAddress
	conf.DB.Model(UserAddress{}).Where("user_id = ? AND is_default = ? ", userId, 1).Find(&resAddress)
	return resAddress
}

// SetDefaultAddress 设置默认地址
func SetDefaultAddress(userId, addressId string) bool {
	oldAddress := GetDefaultAddress(userId)
	conf.DB.Model(UserAddress{}).Where("user_id = ? AND address_id = ?  ", userId, oldAddress.AddressID).Update("is_default", 0)
	dbRes := conf.DB.Model(UserAddress{}).Where("user_id = ? AND address_id = ?  ", userId, addressId).Update("is_default", 1)
	if dbRes.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
