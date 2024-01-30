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

func AddHabitLog(c echo.Context) error {
	habitLog := c.Get("habit_log").(models.HabitLog)
	queryResult := db.DB_CONNECTION.GetDB().Create(&habitLog)
	return queryResult.Error
}

// to be improved a lot
func GetHabitWithLogs(c echo.Context) error {
	var habit models.Habit
	var habitLog []models.HabitLog // Add type []models.HabitLog
	queryResult := db.DB_CONNECTION.GetDB().Where("id = ?", c.Get("id")).First(&habit)
	if queryResult.Error != nil {
		return middleware.HandleQueryResult(queryResult, c, middleware.RequestResponse{Message: "Habit not found", Data: habit}, false)
	}
	queryResult = db.DB_CONNECTION.GetDB().Where("habit_id = ?", c.Get("id")).Find(&habitLog)
	return middleware.HandleQueryResult(queryResult, c, middleware.RequestResponse{Message: "Success", Data: map[string]interface{}{"habit": habit, "logs": habitLog}}, true)
}

func CreateConsumable(c echo.Context) error {
	consumable := c.Get("consumable").(models.Consumable)
	queryResult := db.DB_CONNECTION.GetDB().Create(&consumable)
	return middleware.HandleQueryResult(queryResult, c, middleware.RequestResponse{Message: "Created Successfully", Data: consumable}, false)
}

func GetConsumables(c echo.Context) error {
	var consumables []models.Consumable
	queryResult := db.DB_CONNECTION.GetDB().Find(&consumables)
	return middleware.HandleQueryResult(queryResult, c, middleware.RequestResponse{Message: "Success", Data: consumables}, true)
}
