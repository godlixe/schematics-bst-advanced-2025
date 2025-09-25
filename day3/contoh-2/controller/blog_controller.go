package controller

import (
	"net/http"
	"strconv"

	"contoh-2/dto"
	"contoh-2/model"
	apix "contoh-2/utils/api"
	validatorx "contoh-2/utils/validator"

	"github.com/gin-gonic/gin"
)

type BlogService interface {
	Create(blog *model.Blog) (*model.Blog, error)
	GetByID(id int) (*model.Blog, error)
	GetAll() ([]model.Blog, error)
	Update(blog *model.Blog) (*model.Blog, error)
	Delete(id int) error
}

type BlogController struct {
	blogService BlogService
}

func NewBlogController(blogService BlogService) *BlogController {
	return &BlogController{
		blogService: blogService,
	}
}

func (c *BlogController) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid id",
			Data:    nil,
		})
		return
	}
	blog, err := c.blogService.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, apix.HTTPResponse{
			Message: "blog not found",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, blog)
}

func (c *BlogController) GetAll(ctx *gin.Context) {
	blogs, err := c.blogService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to get blogs",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "get all blogs success",
		Data:    blogs,
	})
}

func (c *BlogController) Create(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	var input dto.CreateBlogDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "input data invalid",
			Data:    ve,
		})
		return
	}

	blog := model.Blog{
		UserID:  userID,
		Title:   input.Title,
		Content: input.Content,
		Author:  input.Author,
	}

	created, err := c.blogService.Create(&blog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to create blog",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *BlogController) Update(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid id",
			Data:    nil,
		})
		return
	}
	var input dto.UpdateBlogDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "input data invalid",
			Data:    ve,
		})
		return
	}

	blog := model.Blog{
		ID:      id,
		UserID:  userID,
		Title:   input.Title,
		Content: input.Content,
		Author:  input.Author,
	}

	updated, err := c.blogService.Update(&blog)
	if err != nil {
		ctx.JSON(http.StatusNotFound, apix.HTTPResponse{
			Message: "failed to update blog",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "update blog successful",
		Data:    updated,
	})
}

func (c *BlogController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid id",
			Data:    nil,
		})
		return
	}
	if err := c.blogService.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, apix.HTTPResponse{
			Message: "failed to delete blog",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "delete blog successful",
		Data:    nil,
	})
}
