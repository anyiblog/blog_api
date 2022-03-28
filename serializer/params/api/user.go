package api

// UserWeChatSilentLoginParam 静默登录
type UserWeChatSilentLoginParam struct {
	Code string `form:"code" json:"code" binding:"required"`
}

// UserGetWechatUserInfoParam 获取微信用户信息
type UserGetWechatUserInfoParam struct {
	EncryptedData string `form:"encryptedData" json:"encryptedData" binding:"required"`
	Iv            string `form:"iv" json:"iv" binding:"required"`
}

// UserBindWechatPhoneParam 获取微信用户手机号
type UserBindWechatPhoneParam struct {
	EncryptedData string `form:"encryptedData" json:"encryptedData" binding:"required"`
	Iv            string `form:"iv" json:"iv" binding:"required"`
}

// UserBindPhoneAndPwdParam 仅适用第三方登录后无手机信息，绑定手机和密码
type UserBindPhoneAndPwdParam struct {
	Phone   string `form:"phone" json:"phone" binding:"required"`
	SmsCode string `form:"smsCode" json:"smsCode" binding:"required"`
	Pwd     string `form:"pwd" json:"pwd" binding:"required"`
}

// UserSmsLoginParam 验证码登录
type UserSmsLoginParam struct {
	Phone   string `form:"phone" json:"phone" binding:"required"`
	SmsCode string `form:"smsCode" json:"smsCode" binding:"required"`
}

// UserPwdLoginParam 密码登录
type UserPwdLoginParam struct {
	Phone string `form:"phone" json:"phone" binding:"required"`
	Pwd   string `form:"pwd" json:"pwd" binding:"required"`
}

// UserRegisterParam 注册
type UserRegisterParam struct {
	Phone   string `form:"phone" json:"phone" binding:"required"`
	Pwd     string `form:"pwd" json:"pwd" binding:"required"`
	SmsCode string `form:"smsCode" json:"smsCode" binding:"required"`
}

// UserUpdatePwdParam 修改密码
type UserUpdatePwdParam struct {
	Phone   string `form:"phone" json:"phone" binding:"required"`
	NewPwd  string `form:"newPwd" json:"newPwd" binding:"required"`
	SmsCode string `form:"smsCode" json:"smsCode" binding:"required"`
}

// UserResetPwdParam 重置密码
type UserResetPwdParam struct {
	Phone   string `form:"phone" json:"phone" binding:"required"`
	NewPwd  string `form:"newPwd" json:"newPwd" binding:"required"`
	SmsCode string `form:"smsCode" json:"smsCode" binding:"required"`
}
