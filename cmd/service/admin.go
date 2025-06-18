package service

import (
	db "todo/pkg/database"
	"todo/pkg/models"

	"github.com/labstack/echo/v4"
)

func migrateDB(c echo.Context) error {
	err := db.DB_CONNECTION.GetDB().AutoMigrate(
		&models.Task{},
		&models.Habit{},
		&models.HabitLog{},

		&models.JournalEntry{},
		&models.JournalPrompt{},

		&models.Book{},

		&models.Food_Item{},
		&models.FoodConsumption{},
	)
	if err != nil {
		return c.JSON(500, "Error migrating models")
	} else {
		return c.JSON(200, "Migrated models Successfully")
	}
}
