package service

import (
	db "todo/pkg/database"
	middleware "todo/pkg/middlewares"
	"todo/pkg/models"

	"github.com/labstack/echo/v4"
)

func CreateHabit(c echo.Context) error {
	habit := c.Get("habit").(models.Habit)
	queryResult := db.DB_CONNECTION.GetDB().Create(&habit)
	return middleware.HandleQueryResult(queryResult, c, middleware.RequestResponse{Message: "Created Successfully", Data: habit}, false)
}

func GetHabits(c echo.Context) error {
	var habits []models.Habit
	queryResult := db.DB_CONNECTION.GetDB().Find(&habits)
	return middleware.HandleQueryResult(queryResult, c, middleware.RequestResponse{Message: "Success", Data: habits}, true)
}
