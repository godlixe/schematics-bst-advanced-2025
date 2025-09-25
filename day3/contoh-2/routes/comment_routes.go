package routes

import (
	"contoh-2/controller"
	"contoh-2/middleware"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.Engine, commentController *controller.CommentController, jwtService middleware.JWTService) {
	commentRoutes := router.Group("/comments")
	{
		commentRoutes.GET("", middleware.Authenticate(jwtService), commentController.GetByUserID)
		commentRoutes.POST("", middleware.Authenticate(jwtService), commentController.Create)
		commentRoutes.PUT("/:id", middleware.Authenticate(jwtService), commentController.Update)
	}
}
