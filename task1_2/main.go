package main

import (
	"fmt"
	"log"
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

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.HandleMethodNotAllowed = true
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

	return r
}

func main() {
	r := SetupRouter()
	log.Fatal(r.Run(":8080"))
}