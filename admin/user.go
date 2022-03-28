package admin

import (
	"anyiblog/serializer"
	adminParams "anyiblog/serializer/params/admin"
	adminService "anyiblog/service/admin"
	"github.com/gin-gonic/gin"
)

// UserPwdLogin 后台系统对外API接口
func UserPwdLogin(c *gin.Context) {
	userLoginParam := &adminParams.UserLoginParam{}
	if err := c.ShouldBindJSON(&userLoginParam); err == nil {
		res := adminService.PwdLogin(userLoginParam.Phone, userLoginParam.Pwd)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//func Logout(c *gin.Context) {
//	token, _ := c.Get("token")
//	res := adminService.Logout(token.(string))
//	c.JSON(200, res)
//}
