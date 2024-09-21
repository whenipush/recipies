package models

type Cocktail struct {
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"type:varchar(100);not null"`
	Description string       `gorm:"type:text"`
	Ingredients []Ingredient `gorm:"many2many:cocktail_ingredients"`
	Ratings     []Rating     `gorm:"foreignKey:CocktailID"`
}

type Ingredient struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);not null;unique"`
}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(100);not null;unique"`
	Ratings  []Rating
}

type Rating struct {
	ID         uint `gorm:"primaryKey"`
	Score      int  `gorm:"type:int;not null"` // Рейтинг от 1 до 5
	CocktailID uint `gorm:"not null"`          // Внешний ключ для коктейля
	UserID     uint `gorm:"not null"`          // Внешний ключ для пользователя
	User       User `gorm:"foreignKey:UserID"`
}
