package api

import (
	"anyiblog/serializer"
	apiParams "anyiblog/serializer/params/api"
	apiService "anyiblog/service/api"
	"github.com/gin-gonic/gin"
)

// SetAddress 添加收货地址
func SetAddress(c *gin.Context) {
	userId, _ := c.Get("userId")
	setAddressParam := apiParams.SetAddressParams{}
	if err := c.ShouldBindJSON(&setAddressParam); err == nil {
		res := apiService.SetAddressService(setAddressParam, userId.(string))
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

func GetDefaultAddress(c *gin.Context) {
	userId, _ := c.Get("userId")
	res := apiService.GetDefaultAddressService(userId.(string))
	c.JSON(200, res)
}

func SetDefaultAddress(c *gin.Context) {
	userId, _ := c.Get("userId")
	addressId := c.Query("AddressID")
	res := apiService.SetDefaultAddressService(userId.(string), addressId)
	c.JSON(200, res)
}
