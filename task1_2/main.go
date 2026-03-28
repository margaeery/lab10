package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		fmt.Printf("[LOG] %s %s %d %v\n", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), latency)
	}
}

func main() {
	r := gin.New()
	r.Use(LoggerMiddleware())
	r.Use(gin.Recovery())

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/info", func(c *gin.Context) {
		c.JSON(200, gin.H{"service": "go-gin", "version": "1.0"})
	})

	r.POST("/data", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "data received"})
	})

	r.Run(":8080")
}