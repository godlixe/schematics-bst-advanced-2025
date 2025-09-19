package controller

import (
	"context"
	"net/http"

	"contoh-3/dto"
	"contoh-3/model"
	"contoh-3/service"
	apix "contoh-3/utils/api"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(ctx context.Context, user *model.User)
	Register(ctx context.Context, user *model.User)
}

type AuthController struct {
	authService AuthService
}

func NewAuthController(a service.AuthService) *AuthController {
	return &AuthController{
		authService: a,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req dto.UserLoginDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := c.authService.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}
	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "login successful",
		Data:    token,
	})
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req dto.UserRegisterDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.authService.Login(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	ctx.JSON(http.StatusOK, dto.LoginResponse{Token: token})
}
