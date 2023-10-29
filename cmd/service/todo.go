package service

import (
	db "todo/pkg/database"
	"todo/pkg/models"

	"github.com/labstack/echo/v4"
)

func CreateTodo(c echo.Context) error {
	task, err := validateAndSanitizeCreateTodo(c)
	if err != nil {
		return c.String(400, err.Error())
	}
	queryResult := db.DB_CONNECTION.GetDB().Create(&task)
	if queryResult.Error != nil {
		return c.String(400, queryResult.Error.Error())
	}
	return c.String(200, "Created Successfully")
}

func GetTodo(c echo.Context) error {
	var tasks []models.Task
	status, err := validateAndSanitizeGetTodo(c)
	if err != nil {
		return c.String(400, err.Error())
	}
	if status == 0 {
		db.DB_CONNECTION.GetDB().Find(&tasks)
	} else {
		db.DB_CONNECTION.GetDB().Where("status = ?", status).Find(&tasks)
	}
	return c.JSON(200, tasks)
}
