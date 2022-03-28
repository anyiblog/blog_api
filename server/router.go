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
			v1.GET("Ping", api.Ping)
			v1.POST("UploadFile", api.UploadFile)
			v1.POST("DeleteFile", api.DeleteFile)
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
					// 店铺模块
					s.POST("CreateShop", admin.CreateShop)
					// 查询预约记录
					s.POST("GetReserve", admin.GetReserve)
					//获取所有服务项目
					s.GET("GetAllServiceProduct", admin.GetAllServiceProduct)
					// 获取所有服务项目的子项目
					s.GET("GetAllServiceItemInfo", admin.GetAllServiceItemInfo)
					// 添加服务类别
					s.POST("AddServiceProduct", admin.AddServiceProduct)
					// 添加服务类别的子项目
					s.POST("AddServiceItem", admin.AddServiceItem)
					// 删除服务类别的子项目
					s.POST("DeleteServiceItem", admin.DeleteServiceItem)

					// 接受预约订单
					s.POST("AcceptReserve", admin.AcceptReserve)
					// 取消预约订单
					s.GET("CancelReserve", admin.CancelReserve)
					// 修改预约订单信息
					s.POST("UpdateReserve", admin.UpdateReserve)
					// 服务预约订单
					s.POST("ActionReserve", admin.ActionReserve)
					// 结束预约订单
					s.GET("EndReserve", admin.EndReserve)

					// 品牌
					s.GET("GetAllBrand", admin.GetAllBrand)

					// 商品模块
					// 新增产品
					s.POST("AddProduct", admin.AddProduct)
					// 获取所有产品
					s.POST("GetAllProduct", admin.GetAllProduct)
					// 上架产品
					s.POST("SetProductOn", admin.ProductOn)
					// 下架产品
					s.POST("SetProductOff", admin.ProductOff)
					// 删除产品
					s.POST("DeleteProduct", admin.DeleteProduct)
				}
			}
		}
	}
	return r
}
