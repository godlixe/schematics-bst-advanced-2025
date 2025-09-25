package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		log.Printf("[GIN-CUSTOM] %s | %3d | %13v | %15s | %-7s  %s",
			start.Format("2006/01/02 - 15:04:05"),
			status,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
