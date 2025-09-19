package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"contoh-3/dto"
	"contoh-3/model"
	apix "contoh-3/utils/api"

	"github.com/gin-gonic/gin"
)

type BlogService interface {
	Create(blog model.Blog) (model.Blog, error)
	GetByID(id int) (model.Blog, error)
	GetAll() ([]model.Blog, error)
	Update(blog model.Blog) (model.Blog, error)
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

func (c *BlogController) Create(ctx *gin.Context) {
	var input dto.BlogDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: fmt.Sprintf("input data invalid %v", err),
			Data:    nil,
		})
		return
	}

	blog := model.Blog{}
	created, err := c.blogService.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to create blog",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusCreated, created)
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
			Data:    nil,
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
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "get all blogs success",
		Data:    blogs,
	})
}

func (c *BlogController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid id",
			Data:    nil,
		})
		return
	}
	var input dto.BlogDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: fmt.Sprintf("input data invalid %v", err),
			Data:    nil,
		})
		return
	}
	updated, err := c.blogService.Update(id, input)
	if err != nil {
		ctx.JSON(http.StatusNotFound, apix.HTTPResponse{
			Message: "failed to update blog",
			Data:    nil,
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
			Data:    nil,
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}
