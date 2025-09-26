package controller

import (
	"net/http"
	"strconv"

	"contoh-2/dto"
	"contoh-2/model"
	apix "contoh-2/utils/api"

	"github.com/gin-gonic/gin"
)

type TagService interface {
	CreateBatch(tags []model.Tag) ([]model.Tag, error)
	GetAll() ([]model.Tag, error)
	GetByBlogID(blogID int) ([]model.Tag, error)
	Delete(id int) error
	AddTagsToBlog(blogID int, tagIDs []int) error
	RemoveTagsFromBlog(blogID int, tagIDs []int) error
}

type TagController struct {
	tagService TagService
}

func NewTagController(tagService TagService) *TagController {
	return &TagController{tagService: tagService}
}

func (c *TagController) CreateBatch(ctx *gin.Context) {
	var input dto.BatchCreateTagsDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "input data invalid",
			Data:    err.Error(),
		})
		return
	}

	tags := make([]model.Tag, 0, len(input.Tags))
	for _, t := range input.Tags {
		tags = append(tags, model.Tag{Name: t.Name})
	}

	created, err := c.tagService.CreateBatch(tags)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to create tags",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, apix.HTTPResponse{Message: "created tags", Data: created})
}

func (c *TagController) GetAll(ctx *gin.Context) {
	tags, err := c.tagService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{Message: "failed to get tags", Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "get tags success",
		Data:    tags,
	})
}

func (c *TagController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid id",
			Data:    nil,
		})
		return
	}

	if err := c.tagService.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, apix.HTTPResponse{
			Message: "failed to delete tag",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "delete tag successful",
		Data:    nil,
	})
}

func (c *TagController) GetByBlogID(ctx *gin.Context) {
	blogIDStr := ctx.Param("id")
	blogID, err := strconv.Atoi(blogIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid blog id",
			Data:    nil,
		})
		return
	}

	tags, err := c.tagService.GetByBlogID(blogID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to get tags",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "get blog tags success",
		Data:    tags,
	})
}

func (c *TagController) AddTagsToBlog(ctx *gin.Context) {
	blogIDStr := ctx.Param("id")
	blogID, err := strconv.Atoi(blogIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid blog id",
			Data:    nil,
		})
		return
	}
	var input dto.AddBlogTagsDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "input data invalid",
			Data:    err.Error(),
		})
		return
	}

	if err := c.tagService.AddTagsToBlog(blogID, input.TagIDs); err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to add tags",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "added tags to blog",
		Data:    nil,
	})
}

func (c *TagController) RemoveTagsFromBlog(ctx *gin.Context) {
	blogIDStr := ctx.Param("id")
	blogID, err := strconv.Atoi(blogIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid blog id",
			Data:    nil,
		})
		return
	}
	var input dto.DeleteBlogTagsDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "input data invalid",
			Data:    err.Error(),
		})
		return
	}

	if err := c.tagService.RemoveTagsFromBlog(blogID, input.TagIDs); err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to remove tags",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "removed tags from blog",
		Data:    nil,
	})
}
