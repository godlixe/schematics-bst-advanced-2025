package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"contoh-1/controller"
	"contoh-1/middleware"
	"contoh-1/model"
	"contoh-1/repository"
	"contoh-1/routes"
	"contoh-1/service"
	"contoh-1/utils/database"
)

func main() {
	godotenv.Load()
	db := database.ConnectDatabase()
	db.AutoMigrate(model.Blog{})
	blogRepo := repository.NewBlogRepository(db)
	blogSvc := service.NewBlogService(blogRepo)
	blogCtrl := controller.NewBlogController(blogSvc)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.CustomLogger())

	routes.BlogRoutes(r, blogCtrl)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
