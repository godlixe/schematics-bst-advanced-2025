package controller

import (
	"contoh-2/dto"
	"contoh-2/model"
	apix "contoh-2/utils/api"
	validatorx "contoh-2/utils/validator"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Register(user *model.User) (*model.User, error)
	Login(user *model.User) (string, error)
	GetByID(id int) (*model.User, error)
	Update(user *model.User) (*model.User, error)
}

type UserController struct {
	userService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterUserDTO
	err := ctx.ShouldBindJSON(&registerDTO)
	if err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "input data invalid",
			Data:    ve,
		})
		return
	}

	user := model.User{
		Email:    registerDTO.Email,
		Password: registerDTO.Password,
		Name:     registerDTO.Name,
	}

	res, err := c.userService.Register(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "failed to create user",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, apix.HTTPResponse{
		Message: "successfully created user",
		Data:    res,
	})
}

func (c *UserController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginUserDTO
	err := ctx.ShouldBindJSON(&loginDTO)
	if err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "input data invalid",
			Data:    ve,
		})
		return
	}

	user := model.User{
		Email:    loginDTO.Email,
		Password: loginDTO.Password,
	}

	res, err := c.userService.Login(&user)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "failed to login",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, apix.HTTPResponse{
		Message: "successfully logged in",
		Data:    res,
	})
}

func (c *UserController) GetUser(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	res, err := c.userService.GetByID(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "failed to get user",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, apix.HTTPResponse{
		Message: "successfully get user",
		Data:    res,
	})
}
