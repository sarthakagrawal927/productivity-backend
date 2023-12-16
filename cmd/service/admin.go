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
			DueDate:    "2023-12-16T19:57:29.675203+05:30",
			Priority:   1,
			Complexity: 1,
			Source:     1,
			SourceId:   1,
			TagIds:     []int64{},
		},
		{
			Model:      gorm.Model{},
			Meta:       models.Meta{Title: "Task 1", Desc: "Task 1 Desc"},
			Status:     2,
			DueDate:    "2023-13-16T19:57:29.675203+05:30",
			Priority:   2,
			Complexity: 2,
			Source:     2,
			SourceId:   2,
			TagIds:     []int64{},
		},
		{
			Model:      gorm.Model{},
			Meta:       models.Meta{Title: "Task 1", Desc: "Task 1 Desc"},
			Status:     3,
			DueDate:    "2023-14-16T19:57:29.675203+05:30",
			Priority:   3,
			Complexity: 3,
			Source:     3,
			SourceId:   3,
			TagIds:     []int64{},
		},
		{
			Model:      gorm.Model{},
			Meta:       models.Meta{Title: "Task 1", Desc: "Task 1 Desc"},
			Status:     4,
			DueDate:    "2023-15-16T19:57:29.675203+05:30",
			Priority:   4,
			Complexity: 4,
			Source:     4,
			SourceId:   4,
			TagIds:     []int64{},
		},
		{
			Model:      gorm.Model{},
			Meta:       models.Meta{Title: "Task 1", Desc: "Task 1 Desc"},
			Status:     5,
			DueDate:    "2023-16-16T19:57:29.675203+05:30",
			Priority:   5,
			Complexity: 5,
			Source:     5,
			SourceId:   5,
			TagIds:     []int64{},
		},
		{
			Model:      gorm.Model{},
			Meta:       models.Meta{Title: "Task 1", Desc: "Task 1 Desc"},
			Status:     6,
			DueDate:    "2023-17-16T19:57:29.675203+05:30",
			Priority:   6,
			Complexity: 6,
			Source:     6,
			SourceId:   6,
			TagIds:     []int64{},
		},
		{
			Model:      gorm.Model{},
			Meta:       models.Meta{Title: "Task 1", Desc: "Task 1 Desc"},
			Status:     7,
			DueDate:    "2023-18-16T19:57:29.675203+05:30",
			Priority:   7,
			Complexity: 7,
			Source:     7,
			SourceId:   7,
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
