package controllers

import (
	"cocktail/database" // Обновлено на корректный путь
	"cocktail/models"   // Обновлено на корректный путь
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCocktail создает новый коктейль
func CreateCocktail(c *gin.Context) {
	var cocktail models.Cocktail

	// Чтение JSON-данных о коктейле
	if err := c.ShouldBindJSON(&cocktail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Для каждого ингредиента проверяем его наличие в базе данных
	for i, ingredient := range cocktail.Ingredients {
		var existingIngredient models.Ingredient
		if err := database.DB.Where("name = ?", ingredient.Name).First(&existingIngredient).Error; err == nil {
			// Если ингредиент уже существует, используем его
			cocktail.Ingredients[i] = existingIngredient
		} else {
			// Если ингредиент не существует, создаем новый
			database.DB.Create(&ingredient)
		}
	}

	// Сохранение коктейля с ингредиентами
	if err := database.DB.Create(&cocktail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cocktail created successfully!", "cocktail": cocktail})
}

// GetCocktails возвращает список всех коктейлей
func GetCocktails(c *gin.Context) {
	var cocktails []models.Cocktail
	database.DB.Find(&cocktails)
	c.JSON(http.StatusOK, cocktails)
}

// GetCocktail возвращает коктейль по ID
func GetCocktail(c *gin.Context) {
	id := c.Param("id")
	var cocktail models.Cocktail
	if err := database.DB.First(&cocktail, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cocktail not found"})
		return
	}
	c.JSON(http.StatusOK, cocktail)
}

// UpdateCocktail обновляет коктейль по ID
func UpdateCocktail(c *gin.Context) {
	id := c.Param("id")
	var cocktail models.Cocktail
	if err := database.DB.First(&cocktail, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cocktail not found"})
		return
	}

	if err := c.ShouldBindJSON(&cocktail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&cocktail)
	c.JSON(http.StatusOK, cocktail)
}

// DeleteCocktail удаляет коктейль по ID
func DeleteCocktail(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Cocktail{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cocktail not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cocktail deleted"})
}
