package routes

import (
	"contoh-2/controller"
	"contoh-2/middleware"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(router *gin.Engine, blogController *controller.BlogController, jwtService middleware.JWTService) {
	blogRoutes := router.Group("/blogs")
	{
		blogRoutes.GET("", middleware.Authenticate(jwtService), blogController.GetAll)
		blogRoutes.GET("/:id", middleware.Authenticate(jwtService), blogController.GetByID)
		blogRoutes.POST("", middleware.Authenticate(jwtService), blogController.Create)
		blogRoutes.PUT("/:id", middleware.Authenticate(jwtService), blogController.Update)
		blogRoutes.DELETE("/:id", middleware.Authenticate(jwtService), blogController.Delete)
	}
}
