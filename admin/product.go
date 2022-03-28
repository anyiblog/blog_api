package admin

import (
	"anyiblog/serializer"
	adminParams "anyiblog/serializer/params/admin"
	adminService "anyiblog/service/admin"
	"github.com/gin-gonic/gin"
)

// AddProduct 添加产品
func AddProduct(c *gin.Context) {
	addProductParam := &adminParams.AddProductParam{}
	if err := c.ShouldBindJSON(&addProductParam); err == nil {
		res := adminService.AddProductService(addProductParam)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// GetAllProduct 获取所有产品
func GetAllProduct(c *gin.Context) {
	getAllProductParam := &adminParams.GetAllProductParam{}
	if err := c.ShouldBindJSON(&getAllProductParam); err == nil {
		res := adminService.GetAllProductService(getAllProductParam)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// ProductOn 上架产品
func ProductOn(c *gin.Context) {
	setProductStatusParam := &adminParams.SetProductStatusParam{}
	if err := c.ShouldBindJSON(&setProductStatusParam); err == nil {
		res := adminService.ProductOnService(setProductStatusParam.ProductIDArray)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// ProductOff 下架产品
func ProductOff(c *gin.Context) {
	setProductStatusParam := &adminParams.SetProductStatusParam{}
	if err := c.ShouldBindJSON(&setProductStatusParam); err == nil {
		res := adminService.ProductOffService(setProductStatusParam.ProductIDArray)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// DeleteProduct 删除产品
func DeleteProduct(c *gin.Context) {
	deleteProductParam := &adminParams.SetProductStatusParam{}
	if err := c.ShouldBindJSON(&deleteProductParam); err == nil {
		res := adminService.DeleteProductService(deleteProductParam.ProductIDArray)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}
