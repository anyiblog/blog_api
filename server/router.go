package server

import (
	"anyiblog/admin"
	"anyiblog/api"
	"anyiblog/middleware"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	// https支持
	r.Use(middleware.TlsHandler())
	// 全局中间件
	r.Use(middleware.Cors())
	v1 := r.Group("/v1")
	{
		{
			//v1.GET("Ping", api.Ping)
			//v1.POST("UploadFile", api.UploadFile)
			//v1.POST("DeleteFile", api.DeleteFile)
		}
		{
			c := v1.Group("/client") // 前端接口分组

			c.GET("GetSiteBaseInfo", api.GetSiteBaseInfo) // 获取网站基本信息
			c.GET("GetCategoryList", api.GetCategoryList) // 获取分类列表
			c.GET("GetBannerList", api.GetBannerList) // 获取Banner列表
			c.GET("SearchArticle", api.SearchArticle) // 搜索文章
			c.GET("GetArticleList", api.GetArticleList) // 获取文章列表
			c.GET("GetArticleDetail", api.GetArticleDetail) // 获取文章详情
			c.GET("GetCommentList", api.GetCommentList) // 获取评论列表
			c.GET("ReleaseComment", api.ReleaseComment) // 发布评论
			c.GET("GetRecommendArticle", api.GetRecommendArticle) // 获取推荐文章
		}
		{ //后台接口
			s := v1.Group("/admin")
			{

				s.POST("Login", admin.UserPwdLogin)

				s.Use(middleware.AdminTokenAuth())
				{ // Token必须是数据库已有用户，并且是管理员

				}
			}
		}
	}
	return r
}
