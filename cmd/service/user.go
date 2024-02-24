package service

import (
	"sync"
	"time"
	db "todo/pkg/database"
	"todo/pkg/models"
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
	var habitLog []HabitLogWithHabitType
	currentTime := time.Now()
	formattedTodayDate := currentTime.Format("2006-01-02 15:04:05")
	formattedYesterdayDate := currentTime.AddDate(0, 0, -1).Format("2006-01-02 15:04:05")

	queryResult := db.DB_CONNECTION.GetDB().Table("habit_logs").Select("habit_logs.*, habits.mode, habits.frequency_type, habits.title").Joins("LEFT JOIN habits on habits.id = habit_logs.habit_id").Where("habit_logs.result_date IN (?, ?)", formattedTodayDate, formattedYesterdayDate).Scan(&habitLog)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Success", Data: habitLog}, true)
}

func GetAllDataForUser(c echo.Context) error {
	var (
		Habits []models.Habit
		Tasks  []models.Task
	)

	dbInstance := db.DB_CONNECTION.GetDB()
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		dbInstance.Where("anti = ?", false).Find(&Habits)
	}()

	go func() {
		defer wg.Done()
		dbInstance.Order("deadline DESC").Find(&Tasks)
	}()

	wg.Wait()

	return c.JSON(200, map[string]interface{}{
		"habits": Habits,
		"tasks":  Tasks,
	})
}
