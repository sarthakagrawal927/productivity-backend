package service

import (
	db "todo/pkg/database"
	"todo/pkg/models"
	validators "todo/pkg/utils"

	"github.com/labstack/echo/v4"
)

func CreateTodo(c echo.Context) error {
	task := c.Get("task").(models.Task)
	queryResult := db.DB_CONNECTION.GetDB().Create(&task)
	return validators.HandleQueryResult(queryResult, c, validators.RequestResponse{Message: "Created Successfully", Data: task}, false)
}

func GetTodo(c echo.Context) error {
	var tasks []models.Task
	userId := c.Get("user_id").(uint)
	// status := c.Get("status").(uint)
	queryResult := db.DB_CONNECTION.GetDB().Where("user_id = ?", userId).Find(&tasks)
	return validators.HandleQueryResult(queryResult, c, validators.RequestResponse{Message: "Success", Data: tasks}, true)
}

func DeleteTodo(c echo.Context) error {
	var task models.Task
	id := c.Get("id").(uint)
	userId := c.Get("user_id").(uint)
	queryResult := db.DB_CONNECTION.GetDB().Where("id = ?", id).Where("user_id = ?", userId).Delete(&task)
	return validators.HandleQueryResult(queryResult, c, validators.RequestResponse{Message: "Deleted Successfully"}, false)
}

func UpdateTodo(c echo.Context) error {
	updateObj := c.Get("updateObj").(map[string]interface{})
	queryResult := db.DB_CONNECTION.GetDB().Model(&models.Task{}).Where("id = ?", updateObj["id"]).Updates(updateObj)
	return validators.HandleQueryResult(queryResult, c, validators.RequestResponse{Message: "Updated Successfully"}, false)
}
