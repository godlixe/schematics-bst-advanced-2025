package middleware

import (
	apix "contoh-2/utils/api"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	ValidateToken(token string) (*jwt.Token, error)
	GetUserIDByToken(token string) (int, error)
}

func Authenticate(jwtService JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, apix.HTTPResponse{
				Message: "invalid auth header",
				Data:    nil,
			})
			return
		}

		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, apix.HTTPResponse{
				Message: "invalid token",
				Data:    nil,
			})
			return
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, apix.HTTPResponse{
				Message: "invalid token",
				Data:    nil,
			})
			return
		}

		userID, err := jwtService.GetUserIDByToken(authHeader)
		ctx.Set("user_id", userID)
		ctx.Next()
	}
}
