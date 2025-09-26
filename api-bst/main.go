package main

import (
	"context"

	"github.com/gin-gonic/gin"
)

func pingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "pong"})
}

func main() {
	context.Context

	r := gin.Default()

	r.GET("/", pingHandler)
	r.Run(":8080")
}
