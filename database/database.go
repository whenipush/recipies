package database

import (
	"cocktail/models"
	"log"

	"gorm.io/driver/sqlite" // Импорт драйвера для SQLite
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Подключение к базе данных SQLite (используйте имя файла базы данных)
	var err error
	DB, err = gorm.Open(sqlite.Open("cocktails.db"), &gorm.Config{}) // Используйте SQLite
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	} else {
		log.Println("Database connected successfully!")
	}

	// Автоматическая миграция модели Cocktail
	DB.AutoMigrate(&models.Cocktail{}, &models.Ingredient{}, &models.User{}, &models.Rating{})
}
