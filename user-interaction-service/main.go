package main

import (
	"log"

	"user-interaction-service/internal/database"
	"user-interaction-service/internal/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация подключений к базам данных
	database.InitDB()
	redis.InitRedis()

	r := gin.Default()

	// Роуты для отзывов
	reviews := r.Group("/reviews")
	{
		reviews.POST("/", createReviewHandler)
		reviews.GET("/event/:id", getEventReviewsHandler)
		reviews.PUT("/:id", updateReviewHandler)
		reviews.DELETE("/:id", deleteReviewHandler)
	}

	// Роуты для регистрации на события
	registration := r.Group("/registration")
	{
		registration.POST("/event/:id", registerForEventHandler)
		registration.DELETE("/event/:id", unregisterFromEventHandler)
		registration.GET("/event/:id", getEventRegistrationsHandler)
	}

	// Запуск сервера на порту 8083
	if err := r.Run(":8083"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func createReviewHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Создание отзыва",
	})
}

func getEventReviewsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Получение отзывов о событии",
	})
}

func updateReviewHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Обновление отзыва",
	})
}

func deleteReviewHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Удаление отзыва",
	})
}

func registerForEventHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Регистрация на событие",
	})
}

func unregisterFromEventHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Отмена регистрации на событие",
	})
}

func getEventRegistrationsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Получение списка регистраций на событие",
	})
}
