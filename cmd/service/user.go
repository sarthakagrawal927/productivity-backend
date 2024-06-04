package service

import (
	"net/http"
	"time"
	db "todo/pkg/database"
	utils "todo/pkg/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type HabitLogWithHabitType struct {
	HabitID     uint           `json:"habit_id"`
	ResultCount uint           `json:"result_count"`
	ResultDate  datatypes.Date `json:"result_time"`
	Comment     string         `json:"comment"`

	Title         string `json:"title"`
	Mode          uint   `json:"mode"`
	FrequencyType uint   `json:"frequency_type"`
	Anti          bool   `json:"anti"`
	gorm.Model
}

func GetDailyLogs(c echo.Context) error {
	var habitLog []HabitLogWithHabitType = []HabitLogWithHabitType{}
	currentTime := time.Now()
	formattedTodayDate := currentTime.Format("2006-01-02 15:04:05")
	formattedYesterdayDate := currentTime.AddDate(0, 0, -1).Format("2006-01-02 15:04:05")

	queryResult := db.DB_CONNECTION.GetDB().Table("habit_logs").Select("habit_logs.*, habits.mode, habits.frequency_type, habits.title").Joins("LEFT JOIN habits on habits.id = habit_logs.habit_id").Where("habit_logs.result_date IN (?, ?)", formattedTodayDate, formattedYesterdayDate).Scan(&habitLog)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Success", Data: habitLog}, true)
}

func GetTodaySchedule(c echo.Context) error {
	formattedSchedule, schedule, taskEntries := getFormattedSchedule()
	return c.JSON(http.StatusOK, utils.RequestResponse{Message: "Success", Data: map[string]interface{}{
		"formatted_schedule": formattedSchedule,
		"schedule":           schedule,
		"task_entries":       taskEntries,
	}})
}
