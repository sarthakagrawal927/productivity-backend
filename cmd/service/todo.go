package service

import (
	db "todo/pkg/database"
	"todo/pkg/models"

	"github.com/labstack/echo"
)

func CreateTodo(c echo.Context) error {
	var task models.Task = models.Task{
		Title: "test",
	}
	db.DB_CONNECTION.GetDB().Create(&task)
	return c.String(200, "Created Successfully")
}

func GetTodo(c echo.Context) error {
	var tasks []models.Task
	db.DB_CONNECTION.GetDB().Find(&tasks)
	return c.JSON(200, tasks)
}
