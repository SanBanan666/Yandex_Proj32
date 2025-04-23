package main

import (
	"log"

	"event-service/internal/database"
	"event-service/internal/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация подключений к базам данных
	database.InitDB()
	redis.InitRedis()

	r := gin.Default()

	// Роуты для управления событиями
	events := r.Group("/events")
	{
		events.POST("/", createEventHandler)
		events.GET("/", getEventsHandler)
		events.GET("/:id", getEventHandler)
		events.PUT("/:id", updateEventHandler)
		events.DELETE("/:id", deleteEventHandler)
	}

	// Запуск сервера на порту 8082
	if err := r.Run(":8082"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func createEventHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Создание события",
	})
}

func getEventsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Получение списка событий",
	})
}

func getEventHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Получение информации о событии",
	})
}

func updateEventHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Обновление события",
	})
}

func deleteEventHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Удаление события",
	})
}
