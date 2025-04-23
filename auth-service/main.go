package main

import (
	"log"

	"auth-service/internal/database"
	"auth-service/internal/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация подключений к базам данных
	database.InitDB()
	redis.InitRedis()

	r := gin.Default()

	// Роуты для аутентификации
	auth := r.Group("/auth")
	{
		auth.POST("/register", registerHandler)
		auth.POST("/login", loginHandler)
		auth.POST("/refresh", refreshTokenHandler)
	}

	// Запуск сервера на порту 8081
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func registerHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Регистрация пользователя",
	})
}

func loginHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Вход пользователя",
	})
}

func refreshTokenHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Обновление токена",
	})
}
