package middleware

import (
	"anyiblog/conf"
	"anyiblog/model"
	"anyiblog/serializer"
	"anyiblog/service"
	"github.com/gin-gonic/gin"
)

// TokenAuth 用户端 Token鉴权 每次都检查Redis是否有同等Token，
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		checkUser := c.GetHeader("checkUser")
		if len(token) <= 0 {
			c.Abort()
			c.JSON(200, serializer.Response{
				Code: serializer.AuthFailed,
				Msg:  "Illegal request",
			})
		} else {
			if checkUser == "0" { // 如果是获取用户信息页面，则不需要校验用户是否存在数据库
				c.Set("token", token)
				c.Next()
			} else { // 从缓存里取出微信解密信息，然后去数据查找对应openId的用户
				weChatUserInfo, getRedisWechatInfoErr := service.GetRedisWechatInfo(token)
				var userIdAndToken struct {
					UserId string
				}
				if getRedisWechatInfoErr == nil {
					i := conf.DB.Model(model.User{}).Select("user_id").Where("open_id = ?", weChatUserInfo.OpenId).Find(&userIdAndToken)
					if i.RowsAffected > 0 {
						c.Set("token", token)
						c.Set("userId", userIdAndToken.UserId)
						c.Set("openId", weChatUserInfo.OpenId)
						c.Next()
					}
				} else {
					c.Abort()
					c.JSON(200, serializer.Response{
						Code: serializer.AuthFailed,
						Msg:  "Authorization Failed",
					})
				}
			}
		}
	}
}

func BindPhoneAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userId")
		var userPhone struct {
			Phone string
		}
		conf.DB.Model(model.User{}).Select("phone").Where("user_id = ?", userID).Find(&userPhone)
		if len(userPhone.Phone) > 0 {
			c.Next()
		} else {
			c.Abort()
			c.JSON(200, serializer.Response{
				Code: serializer.NotBindPhone,
				Msg:  "Not bind phone",
			})
		}
	}
}

// AdminTokenAuth 后台系统用户鉴权
func AdminTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if len(token) <= 0 {
			c.Abort()
			c.JSON(200, serializer.Response{
				Code: serializer.AuthFailed,
				Msg:  "Illegal request",
			})
		} else {
			// 从缓存里取出微信解密信息，然后去数据查找对应openId的用户
			//if cacheRedis.IsExists(token) == 1 {
			//	userId := cacheRedis.Get(token)
			//	c.Set("token", token)
			//	c.Set("userId", userId)
			//	c.Next()
			//} else {
			//	c.Abort()
			//	c.JSON(200, serializer.Response{
			//		Code: serializer.AuthFailed,
			//		Msg:  "Authorization Failed",
			//	})
			//}
		}
	}
}
