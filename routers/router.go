package routers

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	{
		//用户模块的路由接口
		auth.GET("admin/users",v1.GetUsers)
		auth.PUT("user/:id",v1.EditUser)
		auth.DELETE("user/:id",v1.DeleteUser)
		auth.PUT("admin/changepw/:id",v1.ChangeUserPassword)

		// 分类模块的路由接口
		auth.GET("admin/category", v1.GetCate)
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
	}

	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add",v1.AddUser)
		router.GET("users",v1.GetUsers)
		router.GET("user/:id",v1.GetUserInfo)

		// 文章分类信息模块
		router.GET("category", v1.GetCate)
		router.GET("category/:id", v1.GetCateInfo)

		// 文章模块
	}

	r.Run(utils.HttpPort)
}
