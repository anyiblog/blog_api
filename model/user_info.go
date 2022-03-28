package model

import (
	"anyiblog/conf"
	"anyiblog/util"
	"os"
)

// UserInfo 用户信息表
type UserInfo struct {
	UserInfoID     string  `gorm:"-,primary_key;column:user_info_id;type:char(36);not null" json:"-"`
	UserID         string  `gorm:"-,column:user_id;type:char(36);not null" json:"-"`                     // 用户id
	NickName       string  `gorm:"column:nick_name;type:char(20);not null" json:"nick_name"`             // 昵称
	Avatar         string  `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`               // 头像
	Gender         int     `gorm:"column:gender;type:int;not null" json:"gender"`                        // 性别
	Balance        float64 `gorm:"column:balance;type:decimal(18,2) unsigned;not null" json:"balance"`   // 余额
	Integral       uint    `gorm:"column:integral;type:int unsigned;not null" json:"integral"`           // 积分
	DefaultAddress string  `gorm:"column:default_address;type:char(36);not null" json:"default_address"` // 默认收货地址（地址表ID）
}

// CreateUserInfoByUserId 创建普通用户信息
func CreateUserInfoByUserId(userId string) UserInfo {
	var userInfo = struct {
		NickName  string `json:"nick_name"`
		Gender    int    `json:"gender"`
		AvatarURL string `json:"avatar_url"`
	}{
		"伴宠行_" + util.RandStr(5),
		1,
		os.Getenv("Default_User_Avatar"),
	}
	userInfoId := util.GenerateUUID()

	ui := UserInfo{
		UserInfoID:     userInfoId,
		UserID:         userId,
		NickName:       userInfo.NickName,
		Gender:         userInfo.Gender,
		Avatar:         userInfo.AvatarURL,
		Balance:        0.00,
		Integral:       0,
		DefaultAddress: "",
	}
	conf.DB.Create(ui)
	return UserInfo{
		NickName:       userInfo.NickName,
		Gender:         userInfo.Gender,
		Avatar:         userInfo.AvatarURL,
		Balance:        0.00,
		Integral:       0,
		DefaultAddress: "",
	}
}

// CreateWxUserInfo 创建微信用户信息
func CreateWxUserInfo(userId string, wxUserInfo util.WxUserInfo) UserInfo {
	userInfoId := util.GenerateUUID()
	ui := UserInfo{
		UserInfoID:     userInfoId,
		UserID:         userId,
		NickName:       wxUserInfo.NickName,
		Gender:         wxUserInfo.Gender,
		Avatar:         wxUserInfo.AvatarURL,
		Balance:        0.00,
		Integral:       0,
		DefaultAddress: "",
	}
	conf.DB.Create(ui)
	return UserInfo{
		NickName:       wxUserInfo.NickName,
		Gender:         wxUserInfo.Gender,
		Avatar:         wxUserInfo.AvatarURL,
		Balance:        0.00,
		Integral:       0,
		DefaultAddress: "",
	}
}

// GetUserInfoByUserId 获取用户信息
func GetUserInfoByUserId(userId string) UserInfo {
	ResUserInfo := UserInfo{}
	//获取用户信息
	conf.DB.Model(UserInfo{}).Where("user_id = ?", userId).Scan(&ResUserInfo)
	if len(ResUserInfo.Avatar) == 36 {
		ResUserInfo.Avatar = GetImgUrl(ResUserInfo.Avatar)
	}
	return ResUserInfo
}
