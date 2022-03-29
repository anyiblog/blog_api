package middleware

import (
	"anyiblog/serializer"
	"github.com/gin-gonic/gin"
)

// TokenAuth 用户端 Token鉴权 每次都检查Redis是否有同等Token，
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if len(token) <= 0 {
			c.Abort()
			c.JSON(200, serializer.Response{
				Code: serializer.AuthFailed,
				Msg:  "Illegal request",
			})
		} else {
			c.Abort()
			c.JSON(200, serializer.Response{
				Code: serializer.AuthFailed,
				Msg:  "Authorization Failed",
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
