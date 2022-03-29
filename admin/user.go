package admin

import (
	"anyiblog/serializer"
	"github.com/gin-gonic/gin"
)

// UserPwdLogin 后台系统对外API接口
func UserPwdLogin(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: serializer.SystemError,
		Msg:  "ok",
		Data: "ok",
	})
}

//func Logout(c *gin.Context) {
//	token, _ := c.Get("token")
//	res := adminService.Logout(token.(string))
//	c.JSON(200, res)
//}
