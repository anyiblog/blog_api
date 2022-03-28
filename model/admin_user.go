package model

import (
	"anyiblog/conf"
	"anyiblog/util"
)

// AdminUser 后台用户表
type AdminUser struct {
	AdminUserID string `gorm:"primary_key;unique;column:admin_user_id;type:char(36);not null" json:"admin_user_id"`
	Phone       string `gorm:"column:phone;type:char(11);not null" json:"phone"`
	Password    string `gorm:"column:password;type:char(32);not null" json:"password"`
	AdminRole   int    `gorm:"column:admin_role;type:int;not null" json:"admin_role"` // 角色
	ShopID      string `gorm:"column:shop_id;type:char(36);not null" json:"shop_id"`
}

// ResAdminUser 后台用户返回信息
type ResAdminUser struct {
	Token       string `json:"token"`
	AdminUserID string `json:"admin_user_id"`
	Phone       string `json:"phone"`
	AdminRole   int    `json:"admin_role"`
	ShopInfo    Shops  `json:"shop_info"`
}

// LoginCheck 登录
func LoginCheck(phone, pwd string) (ResAdminUser, bool) {
	var adminUserInfo AdminUser
	var resAdminUser ResAdminUser
	var i int64
	i = conf.DB.Model(AdminUser{}).Where("phone = ? AND password = ?", phone, util.Md5(pwd)).Find(&adminUserInfo).RowsAffected
	if i > 0 {
		resAdminUser.Token = "" // 给空值 后续登录逻辑赋值
		resAdminUser.AdminUserID = adminUserInfo.AdminUserID
		resAdminUser.Phone = adminUserInfo.Phone
		resAdminUser.AdminRole = adminUserInfo.AdminRole
		resAdminUser.ShopInfo, _ = GetShopInfoById(adminUserInfo.ShopID)
		return resAdminUser, true
	} else {
		return resAdminUser, false
	}
}
