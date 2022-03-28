package admin

import (
	"anyiblog/serializer"
	adminParams "anyiblog/serializer/params/admin"
	adminService "anyiblog/service/admin"
	"github.com/gin-gonic/gin"
)

// GetReserve 获取宠物预约列表
func GetReserve(c *gin.Context) {
	getReserveParam := &adminParams.GetReserveParam{}
	if err := c.ShouldBindJSON(&getReserveParam); err == nil {
		res := adminService.GetReserve(getReserveParam)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// GetAllServiceProduct 获取所有服务类别
func GetAllServiceProduct(c *gin.Context) {
	res := adminService.GetAllServiceProductService()
	c.JSON(200, res)
}

// GetAllServiceItemInfo 获取所有服务子项目信息
func GetAllServiceItemInfo(c *gin.Context) {
	serviceProductId := c.Query("ServiceProductId")
	res := adminService.GetAllServiceItemInfo(serviceProductId)
	c.JSON(200, res)
}

// AddServiceProduct 添加服务类别
func AddServiceProduct(c *gin.Context) {
	addServiceProductParam := &adminParams.AddServiceProductParam{}
	if err := c.ShouldBindJSON(&addServiceProductParam); err == nil {
		res := adminService.AddServiceProductService(addServiceProductParam.ServiceProductName)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// AddServiceItem 添加服务类别子项目
func AddServiceItem(c *gin.Context) {
	addServiceItemParam := &adminParams.AddServiceItemParam{}
	if err := c.ShouldBindJSON(&addServiceItemParam); err == nil {
		res := adminService.AddServiceItemService(addServiceItemParam)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// DeleteServiceItem 删除服务类别子项目
func DeleteServiceItem(c *gin.Context) {
	deleteServiceItemParam := &adminParams.DeleteServiceItemParam{}
	if err := c.ShouldBindJSON(&deleteServiceItemParam); err == nil {
		res := adminService.DeleteServiceItemService(deleteServiceItemParam.ServiceItemID)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// AcceptReserve 接受预约订单
func AcceptReserve(c *gin.Context) {
	acceptReserveParam := &adminParams.AcceptReserveParam{}
	if err := c.ShouldBindJSON(&acceptReserveParam); err == nil {
		res := adminService.AcceptReserve(acceptReserveParam)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// CancelReserve 取消预约订单
func CancelReserve(c *gin.Context) {
	reserveId := c.Query("ReserveID")
	res := adminService.CancelReserve(reserveId)
	c.JSON(200, res)
}

// ActionReserve 开始服务预约订单
func ActionReserve(c *gin.Context) {
	actionReserveParam := &adminParams.ActionReserveParam{}
	if err := c.ShouldBindJSON(&actionReserveParam); err == nil {
		res := adminService.ActionReserve(actionReserveParam)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// UpdateReserve 修改预约订单信息
func UpdateReserve(c *gin.Context) {
	updateReserveParam := &adminParams.UpdateReserveParam{}
	if err := c.ShouldBindJSON(&updateReserveParam); err == nil {
		res := adminService.UpdateReserve(updateReserveParam)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// EndReserve 取消预约订单
func EndReserve(c *gin.Context) {
	reserveId := c.Query("ReserveID")
	userId := c.Query("UserID")
	res := adminService.EndReserve(reserveId,userId)
	c.JSON(200, res)
}
