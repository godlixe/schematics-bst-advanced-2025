package routes

import (
	"contoh-3/controller"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(router *gin.Engine, blogController *controller.BlogController) {
	blogRoutes := router.Group("/blogs")
	{
		blogRoutes.GET("", blogController.GetAll)
		blogRoutes.GET("/:id", blogController.GetByID)
		blogRoutes.POST("", blogController.Create)
		blogRoutes.PUT("/:id", blogController.Update)
		blogRoutes.DELETE("/:id", blogController.Delete)
	}
}
