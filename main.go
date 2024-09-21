package main

import (
	"cocktail/controllers" // Корректный путь
	"cocktail/database"    // Корректный путь

	"github.com/gin-gonic/gin"
)

func main() {
	// Подключение к базе данных
	database.Connect()

	// Инициализация роутера Gin
	r := gin.Default()

	// CRUD маршруты для коктейлей
	r.POST("/cocktails", controllers.CreateCocktail)
	r.GET("/cocktails", controllers.GetCocktails)
	r.GET("/cocktails/:id", controllers.GetCocktail)
	r.PUT("/cocktails/:id", controllers.UpdateCocktail)
	r.DELETE("/cocktails/:id", controllers.DeleteCocktail)

	// Запуск сервера
	r.Run() // По умолчанию на порту 8080
}
