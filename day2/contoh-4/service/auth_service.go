package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Login(username, password string) (string, error)
	ParseToken(tokenString string) (*jwt.RegisteredClaims, error)
}

type authService struct {
	secret        []byte
	validUsername string
	validPassword string
}

func NewAuthService(secret string) AuthService {
	return &authService{secret: []byte(secret), validUsername: "admin", validPassword: "password"}
}

func (a *authService) Login(username, password string) (string, error) {
	if username != a.validUsername || password != a.validPassword {
		return "", errors.New("invalid credentials")
	}
	claims := jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.secret)
}

func (a *authService) ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return a.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
