package admin

import (
	"anyiblog/serializer"
	adminParams "anyiblog/serializer/params/admin"
	adminService "anyiblog/service/admin"
	"github.com/gin-gonic/gin"
)

// CreateShop 创建店铺
func CreateShop(c *gin.Context) {
	createShopParams := &adminParams.CreateShopParams{}
	if err := c.ShouldBindJSON(&createShopParams); err == nil {
		res := adminService.CreateShopService(createShopParams)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}
