package user

import (
	"anyiblog/cache"
	"anyiblog/model"
	"anyiblog/service"
	"anyiblog/util"
	"errors"
	"fmt"
	"github.com/guonaihong/gout"
	"os"
	"strings"
	"time"
)

var cacheRedis = cache.NewCache(0, "")

// WechatSilentLoginService 静默登录 return token
func WechatSilentLoginService(jsCode string) string {
	var wechatOpenInfo service.WechatOpenInfo
	err := gout.GET("https://api.weixin.qq.com/sns/jscode2session").Debug(true).SetQuery(gout.H{
		"grant_type": "authorize",
		"appid":      os.Getenv("WeChat_AppId"),
		"secret":     os.Getenv("WeChat_Secret"),
		"js_code":    jsCode,
	}).BindJSON(&wechatOpenInfo).Do()
	if err != nil {
		return ""
	}
	// TODO 会有重复存入的问题
	// 以token为键，用户的sessionKey拼上[BCX]用户的openId 为值，存入Redis
	cacheTokenKey := util.GenerateToken()
	cacheValue := strings.Join([]string{wechatOpenInfo.SessionKey, wechatOpenInfo.OpenId}, "[BCX]")
	cacheRedis.Set(cacheTokenKey, cacheValue, 5*time.Minute) // 临时存入微信用户Session和openId 3分钟,如果获取用户信息就永久化
	return cacheTokenKey
}

// Logout 退出登录
func Logout(token string) bool {
	if len(cacheRedis.Get(token)) != 0 {
		if cacheRedis.Delete(token) > 0 {
			return true
		}
	}
	return false
}

// GetWechatUserInfoService 获取微信用户信息 如果未注册自动注册，默认手机号为空
func GetWechatUserInfoService(token, encryptedData, iv string) (model.ResUser, error) {
	wechatOpenInfo, err := service.GetRedisWechatInfo(token) // 根据token从redis获取解密后用户的key和openid
	var resErrMsg error
	if err != nil {
		resErrMsg = errors.New("获取用户信息失败")
		return model.ResUser{}, resErrMsg
	}
	wechatUserInfo := util.DeWxUserInfo(encryptedData, wechatOpenInfo.SessionKey, iv)
	// 如果获取用户信息，表示用户已经成功登录，此时把Token设为不过期
	cacheRedis.SetKeyNotExpire(token)
	if model.IsExistByOpenId(wechatOpenInfo.OpenId) {
		return model.GetUserByOpenId(wechatOpenInfo.OpenId), resErrMsg
	} else {
		return model.CreateWxUser(wechatOpenInfo.OpenId, wechatUserInfo), resErrMsg
	}
}

// BindWechatPhoneService 绑定微信用户手机号
func BindWechatPhoneService(token, encryptedData, iv string) (util.WxPhoneInfo, error) {
	wechatOpenInfo, err := service.GetRedisWechatInfo(token) // 根据token从redis获取解密后用户的key和openid
	var resErrMsg error
	if err != nil {
		resErrMsg = errors.New("获取用户手机号失败")
		return util.WxPhoneInfo{}, resErrMsg
	}
	wechatPhoneInfo := util.DeWxPhoneInfo(encryptedData, wechatOpenInfo.SessionKey, iv) // 手机号信息
	fmt.Println(wechatPhoneInfo)
	if model.BindPhoneByOpenId(wechatOpenInfo.OpenId, wechatPhoneInfo.PhoneNumber) {
		return wechatPhoneInfo, resErrMsg
	} else {
		resErrMsg = errors.New("绑定用户手机号失败")
		return util.WxPhoneInfo{}, resErrMsg
	}
}
