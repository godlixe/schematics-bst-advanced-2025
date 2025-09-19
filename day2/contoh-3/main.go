package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"contoh-3/controller"
	"contoh-3/repository"
	"contoh-3/routes"
	"contoh-3/service"
)

func main() {
	blogRepo := repository.NewBlogRepository()
	blogSvc := service.NewBlogService(blogRepo)
	blogCtrl := controller.NewBlogController(blogSvc)

	r := gin.Default()

	routes.BlogRoutes(r, blogCtrl)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
