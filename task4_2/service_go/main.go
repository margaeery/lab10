package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go_service/docs"
)

type Item struct {
	ID    string  `json:"id" example:"1"`
	Name  string  `json:"name" example:"Ноутбук"`
	Price float64 `json:"price" example:"50000"`
}

var database = []Item{
	{ID: "1", Name: "Ноутбук", Price: 50000},
	{ID: "2", Name: "Мышь", Price: 1500},
}

// @Summary Список товаров
// @Description Возвращает массив всех товаров из памяти
// @Tags items
// @Produce json
// @Success 200 {array} Item
// @Router /items [get]
func GetItems(c *gin.Context) {
	c.JSON(200, database)
}

// @Summary Создать товар
// @Description Добавляет новый товар в список в памяти
// @Tags items
// @Accept json
// @Produce json
// @Param item body Item true "Данные товара"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Неверный формат JSON"
// @Failure 422 {object} map[string]string "Ошибка валидации"
// @Router /items [post]
func CreateItem(c *gin.Context) {
	var newItem Item
	
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(400, gin.H{"error": "Неверный формат данных"})
		return
	}

	if newItem.Price <= 0 {
		c.JSON(422, gin.H{"error": "Цена должна быть больше нуля"})
		return
	}

	database = append(database, newItem)

	c.JSON(201, gin.H{
		"status": "создано",
		"item":   newItem,
		"total":  len(database),
	})
}

// @title Go Data Service API
// @version 1.0
// @description Сервис управления товарами с хранением в памяти
// @host localhost:8002
// @BasePath /
func main() {
	r := gin.New()
	r.Use(gin.Logger())

	r.GET("/items", GetItems)
	r.POST("/items", CreateItem)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(r.Run(":8002"))
}