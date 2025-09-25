package routes

import (
	"contoh-2/controller"
	"contoh-2/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userController *controller.UserController, jwtService middleware.JWTService) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.GET("", middleware.Authenticate(jwtService), userController.GetUser)
	}
}
