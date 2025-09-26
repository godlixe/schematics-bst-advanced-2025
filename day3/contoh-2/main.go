package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"contoh-2/controller"
	"contoh-2/middleware"
	"contoh-2/model"
	"contoh-2/repository"
	"contoh-2/routes"
	"contoh-2/service"
	"contoh-2/utils/database"
)

func main() {
	godotenv.Load()
	db := database.ConnectDatabase()

	db.AutoMigrate(
		model.Blog{},
		model.User{},
		model.Comment{},
		model.Tag{},
		model.BlogsTags{},
	)

	blogRepo := repository.NewBlogRepository(db)
	userRepo := repository.NewUserRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	tagRepo := repository.NewTagRepository(db)

	blogSvc := service.NewBlogService(userRepo, blogRepo)
	jwtSvc := service.NewJWTService()
	userSvc := service.NewUserService(jwtSvc, userRepo)
	commentSvc := service.NewCommentService(commentRepo)
	tagSvc := service.NewTagService(tagRepo)

	blogCtrl := controller.NewBlogController(blogSvc)
	userCtrl := controller.NewUserController(userSvc)
	commentCtrl := controller.NewCommentController(commentSvc)
	tagCtrl := controller.NewTagController(tagSvc)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.CustomLogger())

	routes.BlogRoutes(r, blogCtrl, jwtSvc)
	routes.UserRoutes(r, userCtrl, jwtSvc)
	routes.CommentRoutes(r, commentCtrl, jwtSvc)
	routes.TagRoutes(r, tagCtrl, jwtSvc)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
