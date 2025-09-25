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

type CommentService interface {
	Create(comment *model.Comment) (*model.Comment, error)
	GetByUserID(userID int) ([]model.Comment, error)
	Update(comment *model.Comment) (*model.Comment, error)
}

type CommentController struct {
	commentService CommentService
}

func NewCommentController(commentService CommentService) *CommentController {
	return &CommentController{
		commentService: commentService,
	}
}

func (c *CommentController) GetByUserID(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")

	comments, err := c.commentService.GetByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, apix.HTTPResponse{
			Message: "comment not found",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully get comments",
		Data:    comments,
	})
}

func (c *CommentController) Create(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	var input dto.CreateCommentDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "input data invalid",
			Data:    ve,
		})
		return
	}

	comment := model.Comment{
		UserID:  userID,
		BlogID:  input.BlogID,
		Content: input.Content,
	}

	created, err := c.commentService.Create(&comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to create comment",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully created comment",
		Data:    created,
	})
}

func (c *CommentController) Update(ctx *gin.Context) {
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

	comment := model.Comment{
		ID:      id,
		UserID:  userID,
		Content: input.Content,
	}

	updated, err := c.commentService.Update(&comment)
	if err != nil {
		ctx.JSON(http.StatusNotFound, apix.HTTPResponse{
			Message: "failed to update blog",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "update comment successful",
		Data:    updated,
	})
}
