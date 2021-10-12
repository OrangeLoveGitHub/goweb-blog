package routes

import (
	v1 "goweb-blog/api/v1"
	"goweb-blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// User model router
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		// category model router
		router.POST("category/add", v1.AddCategory)
		router.GET("category", v1.GetCategory)
		router.PUT("category/:id", v1.EditCategory)
		router.DELETE("category/:id", v1.DeleteCategory)
		// article model  router
		router.POST("article/add", v1.AddArticle)
		router.GET("article", v1.GetArticle)
		router.PUT("article/:id", v1.EditArticle)
		router.DELETE("article/:id", v1.DeleteArticle)

	}
	r.Run(utils.HttpPort)
}
