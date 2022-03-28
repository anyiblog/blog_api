package api

import (
	"anyiblog/serializer"
	"github.com/gin-gonic/gin"
)

func GetSiteBaseInfo(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "",
		Data: nil,
	})
}

func GetCategoryList(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "",
		Data: nil,
	})
}


func GetBannerList(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "",
		Data: nil,
	})
}

func SearchArticle(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "",
		Data: nil,
	})
}

func GetArticleList(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "",
		Data: nil,
	})
}

func GetArticleDetail(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "",
		Data: nil,
	})
}

func GetCommentList(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "",
		Data: nil,
	})
}

func ReleaseComment(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "",
		Data: nil,
	})
}

func GetRecommendArticle(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "",
		Data: nil,
	})
}
