package service

import (
	"todo/cmd/dao"
	db "todo/pkg/database"
	"todo/pkg/models"
	utils "todo/pkg/utils"

	"github.com/labstack/echo/v4"
)

func CreateHabit(c echo.Context) error {
	habit := c.Get("habit").(models.Habit)
	return db.InsertIntoDB(c, &habit)
}

func GetHabits(c echo.Context) error {
	var habits []models.Habit
	userId := c.Get("user_id").(uint)
	queryResult := db.DB_CONNECTION.GetDB().Where("user_id = ?", userId).Order("created_at desc").Find(&habits)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Success", Data: habits}, true)
}

func AddHabitLog(c echo.Context) error {
	habitLog := c.Get("habit_log").(models.HabitLog)
	err := db.InsertIntoDB(c, &habitLog)
	if err2 := updateHabitUsage(habitLog.HabitID); err2 != nil {
		return utils.HandleEchoError(c, err2)
	}
	return err
}

// will also need to add CRON to update habit usage
func updateHabitUsage(habitId uint) error {
	queryResult := db.DB_CONNECTION.GetDB().Exec(dao.UpdateHabitFromLogs, habitId)
	return queryResult.Error
}

// to be improved a lot
func GetHabitWithLogs(c echo.Context) error {
	var habit models.Habit
	var habitLog []models.HabitLog
	queryResult := db.DB_CONNECTION.GetDB().Where("id = ?", c.Get("id")).First(&habit)
	if queryResult.Error != nil {
		return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Habit not found", Data: habit}, false)
	}
	queryResult = db.DB_CONNECTION.GetDB().Where("habit_id = ?", c.Get("id")).Find(&habitLog)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Success", Data: map[string]interface{}{"habit": habit, "logs": habitLog}}, true)
}
