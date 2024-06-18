package service

import (
	"todo/cmd/dao"
	db "todo/pkg/database"
	"todo/pkg/models"
	types "todo/pkg/types"
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
	foodItem := c.Get("food").(models.Food_Item)
	return InsertIntoDB(c, &foodItem)
}

func CreateFoodConsumed(c echo.Context) error {
	foodConsumption := c.Get("food_consumed").(models.FoodConsumption)
	return InsertIntoDB(c, &foodConsumption)
}

func GetFoodItems(c echo.Context) error {
	foodItems := []models.Food_Item{}
	queryResult := db.DB_CONNECTION.GetDB().Find(&foodItems)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Success", Data: foodItems}, true)
}

func GetFoodConsumed(c echo.Context) error {
	foodConsumed := []types.DayLevelFoodConsumption{}
	queryResult := db.DB_CONNECTION.GetDB().Raw(dao.GetNutrientsConsumedForDate, c.Get("date")).Scan(&foodConsumed)
	// sum of all nutrients consumed = sum of all nutrients in foodConsumed
	// totalFoodConsumed := types.DayLevelFoodConsumption{}
	// for _, food := range foodConsumed {
	// 	totalFoodConsumed.Kcal += food.Kcal
	// 	totalFoodConsumed.Protein += food.Protein
	// 	totalFoodConsumed.Fiber += food.Fiber
	// }
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Success", Data: map[string]interface{}{
		"food_consumed": foodConsumed,
		// "total_food_consumed": totalFoodConsumed,
	}}, true)
}

func GetDailyFoodLogs(c echo.Context) error {
	foodConsumed := []models.UserFoodRequirements{}
	queryResult := db.DB_CONNECTION.GetDB().Raw(dao.GetFoodConsumptionLogs).Scan(&foodConsumed)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Success", Data: foodConsumed}, true)
}
