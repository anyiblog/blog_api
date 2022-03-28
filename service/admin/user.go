package admin

import (
	"anyiblog/cache"
	"anyiblog/model"
	"anyiblog/serializer"
	"anyiblog/util"
	"os"
	"strings"
	"time"
)

var cacheRedis = cache.NewCache(0, "")

// PwdLogin 用户登录
// redis缓存Token key为将userId和手机号AES加密并用 [BCX] 拼接 value为userId
func PwdLogin(phone, pwd string) serializer.Response {
	var tokenExpireTime time.Duration = (time.Hour * 24) * serializer.TokenRedisTime // 5天
	resAdminUser, isOk := model.LoginCheck(phone, pwd)
	if isOk { // 登录成功
		// 将userId和手机号AES加密后，去查找redis是否存在相同的key，如果存在即删除当前，从新生成key
		strText := strings.Join([]string{resAdminUser.AdminUserID, resAdminUser.Phone}, "[BCX]")
		// 当前用户加密后的Token
		encryptTokenKey, _ := util.AesEncrypt(strText, os.Getenv("TokenKey"), os.Getenv("TokenIv"))

		if cacheRedis.IsExists(encryptTokenKey) == 1 {
			oldTokenVal := cacheRedis.Get(encryptTokenKey)
			if oldTokenVal != resAdminUser.AdminUserID {
				// 已存在旧token的值不等于当前登录的值 设置当前用户的token
				cacheRedis.Set(encryptTokenKey, resAdminUser.AdminUserID, tokenExpireTime)
			}
			// 新旧token key和value都想等的话给token续期
			cacheRedis.SetKeyExpire(encryptTokenKey, tokenExpireTime)
		} else {
			cacheRedis.Set(encryptTokenKey, resAdminUser.AdminUserID, tokenExpireTime)
		}

		resAdminUser.Token = encryptTokenKey
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "登录成功",
			Data: resAdminUser,
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "用户不存在或密码错误",
		}
	}
}
