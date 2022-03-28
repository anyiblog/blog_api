package api

import (
	"anyiblog/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserOrderList(c *gin.Context) {
	userId, _ := c.Get("userId")
	page, _ := strconv.Atoi(c.Query("Page"))
	limit, _ := strconv.Atoi(c.Query("Limit"))
	queryType, _ := strconv.Atoi(c.Query("Type"))
	res := service.GetUserOrderListService(userId.(string), page, limit, queryType)
	c.JSON(200, res)
}

func GetUserOrderDetail(c *gin.Context) {
	userId, _ := c.Get("userId")
	orderInfoID := c.Query("OrderInfoID")
	res := service.GetUserOrderDetailService(userId.(string), orderInfoID)
	c.JSON(200, res)
}
