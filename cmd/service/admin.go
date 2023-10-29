package service

import (
	"todo/pkg/constants"
	db "todo/pkg/database"
	"todo/pkg/models"

	"github.com/labstack/echo/v4"
)

func migrateDB(c echo.Context) error {
	err := db.DB_CONNECTION.GetDB().AutoMigrate(&models.Task{})
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
		{Title: "Task 1", Status: int(constants.Backlog)},
		{Title: "Task 2", Status: int(constants.Todo)},
		{Title: "Task 3", Status: int(constants.Done)},
		{Title: "Task 4", Status: int(constants.Done)},
		{Title: "Task 5", Status: int(constants.InProgress)},
		{Title: "Task 6", Status: int(constants.InProgress)},
		{Title: "Task 7", Status: int(constants.InProgress)},
		{Title: "Task 8", Status: int(constants.Todo)},
		{Title: "Task 9", Status: int(constants.InProgress)},
		{Title: "Task 10", Status: int(constants.Backlog)},
	}

	transactionResult := db.DB_CONNECTION.GetDB().Create(&tasks)
	if transactionResult.Error != nil {
		return c.JSON(500, "Error seeding tasks")
	} else {
		return c.JSON(200, "Seeded Successfully")
	}
}
