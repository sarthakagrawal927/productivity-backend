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
		&models.FoodConsumed{},
	)
	if err != nil {
		return c.JSON(500, "Error migrating models")
	} else {
		return c.JSON(200, "Migrated models Successfully")
	}
}

func deleteAllTasks(c echo.Context) error {
	transactionResult := db.DB_CONNECTION.GetDB().Where("1=1").Delete(&models.Task{})
	if transactionResult.Error != nil {
		return c.JSON(500, "Error deleting tasks")
	} else {
		return c.JSON(200, "Deleted Successfully")
	}
}

func seedTasks(c echo.Context) error {
	tasks := []models.Task{
		{
			Meta: models.Meta{
				Title: "Task 1",
				Desc:  "Task 1 Desc",
			},
		},
	}

	transactionResult := db.DB_CONNECTION.GetDB().Create(&tasks)
	if transactionResult.Error != nil {
		return c.JSON(500, "Error seeding tasks")
	} else {
		return c.JSON(200, "Seeded Successfully")
	}
}
