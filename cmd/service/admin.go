package service

import (
	db "todo/pkg/database"
	"todo/pkg/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func migrateDB(c echo.Context) error {
	err := db.DB_CONNECTION.GetDB().AutoMigrate(&models.Task{}, &models.JournalEntry{})
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
			Model:      gorm.Model{},
			Meta:       models.Meta{Title: "Task 1", Desc: "Task 1 Desc"},
			Status:     1,
			DueDate:    "01012024",
			Priority:   1,
			Complexity: 1,
			Source:     1,
			SourceId:   1,
			TagIds:     []int64{},
		},
	}
	// AutoMigrate will automatically create the table based on the struct definition.
	db.DB_CONNECTION.GetDB().AutoMigrate(&models.Task{})

	transactionResult := db.DB_CONNECTION.GetDB().Create(&tasks)
	if transactionResult.Error != nil {
		return c.JSON(500, "Error seeding tasks")
	} else {
		return c.JSON(200, "Seeded Successfully")
	}
}
