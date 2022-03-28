package model

import (
	"anyiblog/conf"
	"anyiblog/util"
	"time"
)

// User 用户表
type User struct {
	UserID    string    `gorm:"primary_key;column:user_id;type:char(36);not null" json:"user_id"`
	Phone     string    `gorm:"column:phone;type:char(11)" json:"phone"`                        // 手机号
	Password  string    `gorm:"column:password;type:char(32)" json:"password"`                  // 密码
	OpenID    string    `gorm:"column:open_id;type:varchar(255);not null" json:"open_id"`       // 微信openid
	AdminRole uint      `gorm:"column:admin_role;type:int unsigned;not null" json:"admin_role"` // 后台管理权限分组
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`             // 注册时间
}

// ResUser 返回用户信息
type ResUser struct {
	UserID    string   `json:"user_id"`
	Phone     string   `json:"phone"`
	AdminRole uint     `json:"admin_role"`
	UserInfo  UserInfo `json:"user_info" gorm:"-"`
}

// GetOpenIdByUserId 根据UserId获取OpenId
func GetOpenIdByUserId(userId string) string {
	var open struct {
		OpenID string
	}
	conf.DB.Model(User{}).Debug().Select("open_id").Where("user_id = ?", userId).Scan(&open)
	return open.OpenID
}

// GetUserPhoneByUserId 根据UserId获取用户
func GetUserPhoneByUserId(userId string) string {
	var phone struct {
		Phone string
	}
	conf.DB.Model(User{}).Select("phone").Where("user_id = ?", userId).Scan(&phone)
	return phone.Phone
}

// GetUserByOpenId 根据openId获取用户
func GetUserByOpenId(openId string) ResUser {
	ResUser := ResUser{}
	conf.DB.Model(User{}).Select("user_id,phone,admin_role").Where("open_id = ?", openId).Find(&ResUser)
	if len(ResUser.UserID) == 36 {
		ResUser.UserInfo = GetUserInfoByUserId(ResUser.UserID)
	}
	return ResUser
}

// CreateWxUser 创建微信小程序用户
func CreateWxUser(openId string, userInfo util.WxUserInfo) ResUser {
	userId := util.GenerateUUID()
	u := User{
		UserID:    userId,
		Phone:     "",
		Password:  "",
		OpenID:    openId,
		AdminRole: 0,
		CreatedAt: util.NowTime(),
	}
	conf.DB.Create(u) //创建用户
	return ResUser{
		UserID:    u.UserID,
		Phone:     u.Phone,
		AdminRole: u.AdminRole,
		UserInfo:  CreateWxUserInfo(u.UserID, userInfo),
	}
}

// IsExistByOpenId 是否存在用户 true 存在 false 不存在
func IsExistByOpenId(openId string) bool {
	var i int64
	conf.DB.Model(User{}).Where("open_id = ?", openId).Count(&i)
	if i > 0 {
		return true
	} else {
		return false
	}
}

func BindPhoneByOpenId(openId, phone string) bool {
	if conf.DB.Model(User{}).Where("open_id = ?", openId).Updates(map[string]interface{}{"phone": phone}).RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// IsAdmin 是否是管理员 true 是管理员 false 不是
func IsAdmin(phone string) bool {
	var admin struct {
		Admin int // 0 为拥有管理员权限，大于 0 即拥有管理权限 （埋坑，后续可能会有权限系统）
	}
	conf.DB.Model(User{}).Select("admin").Where("phone = ?", phone).Scan(&admin)
	if admin.Admin > 0 {
		return true
	} else {
		return false
	}
}

// IsAdminByToken 是否是管理员 true 是管理员 false 不是
func IsAdminByToken(token string) bool {
	var admin struct {
		Admin int // 0 为拥有管理员权限，大于 0 即拥有管理权限 （埋坑，后续可能会有权限系统）
	}
	conf.DB.Model(User{}).Select("admin").Where("token = ?", token).Scan(&admin)
	if admin.Admin > 0 {
		return true
	} else {
		return false
	}
}

// GetUserIdForPwdAndAdmin 手机+密码获取userId 并且必须是管理员
func GetUserIdForPwdAndAdmin(phone, pwd string) string {
	var userId struct {
		UserId string
	}
	conf.DB.Model(User{}).Select("user_id").Where("phone = ? and password = ? ", phone, util.Md5(pwd)).Scan(&userId)
	return userId.UserId
}

// GetUserIdForPwd 手机+密码获取userId
func GetUserIdForPwd(phone, pwd string) string {
	var userId struct {
		UserId string
	}
	conf.DB.Model(User{}).Select("user_id").Where("phone = ? and password = ?", phone, util.Md5(pwd)).Scan(&userId)
	return userId.UserId
}

// GetUserIdForPhone 手机获取userId
func GetUserIdForPhone(phone string) string {
	var userId struct {
		UserId string
	}
	conf.DB.Model(User{}).Select("user_id").Where("phone = ?", phone).Scan(&userId)
	return userId.UserId
}

// GetUser 获取用户token和手机号
func GetUser(userId string) (token, phone string) {
	var user struct {
		Token string
		Phone string
	}
	conf.DB.Model(User{}).Select("phone,token").Where("user_id = ?", userId).Scan(&user)
	return user.Token, user.Phone
}

// BindPhoneAndPwd 如果手机号是随机字符串则绑定新手机号（仅适用第三方平台无手机号）
func BindPhoneAndPwd(userId, phone, pwd string) bool {
	if conf.DB.Model(User{}).Where("user_id = ?", userId).Updates(map[string]interface{}{"phone": phone, "password": util.Md5(pwd)}).RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// UpdatePwd 通过验证码更新密码
func UpdatePwd(phone, newPwd string) bool {
	if conf.DB.Model(User{}).Where("phone = ?", phone).Updates(map[string]interface{}{"password": util.Md5(newPwd)}).RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
