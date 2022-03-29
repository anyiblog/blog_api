package middleware

import (
	"anyiblog/serializer"
	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if model.IsAdminByToken(token.(string)) {
		//	c.Next()
		//} else {
		//	c.JSON(200, serializer.Response{
		//		Code: serializer.SystemError,
		//		Msg:  "该用户暂未拥有管理员权限",
		//	})
		//	c.Abort()
		//}
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "该用户暂未拥有管理员权限",
		})
		c.Abort()
	}
}
