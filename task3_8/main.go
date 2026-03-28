package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go_service/docs"
)

// @title           Go Gin API
// @version         1.0
// @description     Сервер на Gin с поддержкой Swagger (Задание 8).
// @host            localhost:8080
// @BasePath        /

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		fmt.Printf("[LOG] %s %s %d %v\n", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), latency)
	}
}

// @Summary      Проверка статуса
// @Tags         System
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /status [get]
func StatusHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

// @Summary      Информация о сервисе
// @Tags         System
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /info [get]
func InfoHandler(c *gin.Context) {
	c.JSON(200, gin.H{"service": "go-gin", "version": "1.0"})
}

// DataHandler godoc
// @Summary      Отправка данных
// @Tags         Data
// @Accept       json
// @Produce      json
// @Param        payload  body      object  true  "JSON данные для отправки"
// @Success      200      {object}  map[string]string
// @Router       /data [post]
func DataHandler(c *gin.Context) {
    c.JSON(200, gin.H{"message": "data received"})
}

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.Use(LoggerMiddleware())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/status", StatusHandler)
	r.GET("/info", InfoHandler)
	r.POST("/data", DataHandler)

	return r
}

func main() {
	r := SetupRouter()
	r.Run(":8080")
}