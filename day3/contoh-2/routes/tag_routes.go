package routes

import (
	"contoh-2/controller"
	"contoh-2/middleware"

	"github.com/gin-gonic/gin"
)

func TagRoutes(router *gin.Engine, tagController *controller.TagController, jwtService middleware.JWTService) {
	tagRoutes := router.Group("/tags")
	{
		tagRoutes.GET("", middleware.Authenticate(jwtService), tagController.GetAll)
		tagRoutes.POST("/batch", middleware.Authenticate(jwtService), tagController.CreateBatch)
		tagRoutes.DELETE("/:id", middleware.Authenticate(jwtService), tagController.Delete)
	}

	blogTagRoutes := router.Group("/blogs/:id/tags")
	{
		blogTagRoutes.GET("", middleware.Authenticate(jwtService), tagController.GetByBlogID)
		blogTagRoutes.POST("", middleware.Authenticate(jwtService), tagController.AddTagsToBlog)
		blogTagRoutes.DELETE("", middleware.Authenticate(jwtService), tagController.RemoveTagsFromBlog)
	}
}
