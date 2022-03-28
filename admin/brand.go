package admin

import (
	"anyiblog/serializer"
	adminParams "anyiblog/serializer/params/admin"
	adminService "anyiblog/service/admin"
	"github.com/gin-gonic/gin"
)

func GetAllBrand(c *gin.Context)  {
	getAllBrandParam := &adminParams.GetAllBrandParams{}
	if err := c.ShouldBindJSON(&getAllBrandParam); err == nil {
		res := adminService.GetAllBrandService(getAllBrandParam)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}
