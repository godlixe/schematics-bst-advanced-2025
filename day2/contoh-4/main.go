package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"contoh-3/controller"
	"contoh-3/middleware"
	"contoh-3/repository"
	"contoh-3/service"
)

func main() {
	repo := repository.NewInMemoryBlogRepository()
	blogSvc := service.NewBlogService(repo)
	authSvc := service.NewAuthService("secret-key-change-me")

	blogCtrl := controller.NewBlogController(blogSvc)
	authCtrl := controller.NewAuthController(authSvc)

	r := gin.Default()

	r.POST("/login", authCtrl.Login)

	api := r.Group("/api")
	api.Use(middleware.JWTAuth(authSvc))
	{
		api.POST("/blogs", blogCtrl.Create)
		api.GET("/blogs", blogCtrl.GetAll)
		api.GET("/blogs/:id", blogCtrl.GetByID)
		api.PUT("/blogs/:id", blogCtrl.Update)
		api.DELETE("/blogs/:id", blogCtrl.Delete)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
