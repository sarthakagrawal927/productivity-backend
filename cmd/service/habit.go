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
	return InsertIntoDB(c, &habit)
}

func GetHabits(c echo.Context) error {
	var habits []models.Habit
	queryResult := db.DB_CONNECTION.GetDB().Find(&habits)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Success", Data: habits}, true)
}

func AddHabitLog(c echo.Context) error {
	habitLog := c.Get("habit_log").(models.HabitLog)
	err := InsertIntoDB(c, &habitLog)
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

func InsertIntoDB(c echo.Context, item interface{}) error {
	queryResult := db.DB_CONNECTION.GetDB().Create(item)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Created Successfully", Data: item}, false)
}

func CreateBookConsumable(c echo.Context) error {
	book := c.Get("book").(models.Book)
	return InsertIntoDB(c, &book)
}

func CreateFoodConsumable(c echo.Context) error {
	food := c.Get("food").(models.Food_Item)
	return InsertIntoDB(c, &food)
}

func GetConsumables(c echo.Context) error {
	var consumables []models.Consumable
	queryResult := db.DB_CONNECTION.GetDB().Find(&consumables)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Success", Data: consumables}, true)
}
