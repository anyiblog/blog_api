package middleware

import (
	"anyiblog/model"
	"anyiblog/serializer"
	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Get("token")
		if model.IsAdminByToken(token.(string)) {
			c.Next()
		} else {
			c.JSON(200, serializer.Response{
				Code: serializer.SystemError,
				Msg:  "该用户暂未拥有管理员权限",
			})
			c.Abort()
		}
	}
}
