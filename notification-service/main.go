package main

import (
	"log"

	"notification-service/internal/database"
	"notification-service/internal/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация подключений к базам данных
	database.InitDB()
	redis.InitRedis()

	r := gin.Default()

	// Роуты для уведомлений
	notifications := r.Group("/notifications")
	{
		notifications.POST("/send", sendNotificationHandler)
		notifications.GET("/user/:id", getUserNotificationsHandler)
		notifications.PUT("/:id/read", markNotificationAsReadHandler)
		notifications.DELETE("/:id", deleteNotificationHandler)
	}

	// Роуты для управления подписками
	subscriptions := r.Group("/subscriptions")
	{
		subscriptions.POST("/", createSubscriptionHandler)
		subscriptions.DELETE("/:id", deleteSubscriptionHandler)
		subscriptions.GET("/user/:id", getUserSubscriptionsHandler)
	}

	// Запуск сервера на порту 8084
	if err := r.Run(":8084"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func sendNotificationHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Отправка уведомления",
	})
}

func getUserNotificationsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Получение уведомлений пользователя",
	})
}

func markNotificationAsReadHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Отметка уведомления как прочитанного",
	})
}

func deleteNotificationHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Удаление уведомления",
	})
}

func createSubscriptionHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Создание подписки",
	})
}

func deleteSubscriptionHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Удаление подписки",
	})
}

func getUserSubscriptionsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Получение подписок пользователя",
	})
}
