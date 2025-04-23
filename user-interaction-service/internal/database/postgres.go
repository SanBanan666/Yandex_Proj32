package database

import (
	"fmt"
	"log"
	"os"

	"user-interaction-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Автоматическая миграция схемы
	err = DB.AutoMigrate(&models.Review{}, &models.EventRegistration{})
	if err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}

	log.Println("Успешное подключение к базе данных")
}
