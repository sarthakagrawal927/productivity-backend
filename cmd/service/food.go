package service

import (
	"fmt"
	"todo/cmd/dao"
	"todo/pkg/constants"
	db "todo/pkg/database"
	"todo/pkg/models"
	types "todo/pkg/types"
	utils "todo/pkg/utils"

	"github.com/labstack/echo/v4"
)

func CreateFoodConsumable(c echo.Context) error {
	foodItem := c.Get("food").(models.Food_Item)
	return db.InsertIntoDB(c, &foodItem)
}

func CreateFoodConsumed(c echo.Context) error {
	foodConsumption := c.Get("food_consumed").(models.FoodConsumption)
	return db.InsertIntoDB(c, &foodConsumption)
}

func GetFoodItems(c echo.Context) error {
	foodItems := []models.Food_Item{}
	userId := c.Get("user_id").(uint)
	queryResult := db.DB_CONNECTION.GetDB().Where("user_id = ?", userId).Order("created_at desc").Find(&foodItems)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Success", Data: foodItems}, true)
}

func GetFoodConsumed(c echo.Context) error {
	foodConsumed := []types.DayLevelFoodConsumption{}
	queryResult := db.DB_CONNECTION.GetDB().Raw(dao.GetNutrientsConsumedForDate, c.Get("date"), c.Get("user_id")).Scan(&foodConsumed)
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
	mode := c.Get("mode").(uint)
	userId := c.Get("user_id").(uint)
	var dateGroup string
	if mode == constants.FOOD_LOG_WEEK_MODE {
		dateGroup = "date_trunc('week', fc.\"date\")"
	} else {
		dateGroup = "fc.\"date\""
	}
	queryResult := db.DB_CONNECTION.GetDB().Raw(fmt.Sprintf(dao.GetFoodConsumptionLogs, dateGroup), userId).Scan(&foodConsumed)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Success", Data: foodConsumed}, true)
}
