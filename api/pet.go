package api

import (
	"anyiblog/serializer"
	userParams "anyiblog/serializer/params/api"
	"anyiblog/service"
	adminService "anyiblog/service/admin"
	"github.com/gin-gonic/gin"
)

// AddMyPet 添加宠物
func AddMyPet(c *gin.Context) {
	userId, _ := c.Get("userId")
	addMyPetParam := userParams.AddMyPetParam{}
	if err := c.ShouldBindJSON(&addMyPetParam); err == nil {
		res := service.AddMyPetService(&addMyPetParam, userId.(string))
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// MyPetList 获取我的宠物列表
func MyPetList(c *gin.Context) {
	userId, _ := c.Get("userId")
	res := service.MyPetList(userId.(string))
	c.JSON(200, res)
}

// AddPetReserve 添加宠物预约
func AddPetReserve(c *gin.Context) {
	userId, _ := c.Get("userId")
	openId, _ := c.Get("openId")
	addPetReserveParam := &userParams.AddPetReserveParam{}
	if err := c.ShouldBindJSON(&addPetReserveParam); err == nil {
		res := service.AddPetReserve(addPetReserveParam, userId.(string), openId.(string))
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// SelectPetReserve 查询宠物预约记录
func SelectPetReserve(c *gin.Context) {
	userId, _ := c.Get("userId")
	selectPetReserveParam := &userParams.SelectPetReserveParam{}
	if err := c.ShouldBindJSON(&selectPetReserveParam); err == nil {
		res := service.SelectPetReserve(selectPetReserveParam.PetID, userId.(string), selectPetReserveParam.Limit, selectPetReserveParam.Page)
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
