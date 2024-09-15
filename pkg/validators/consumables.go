package validators

import (
	"time"
	"todo/pkg/constants"
	"todo/pkg/models"
	utils "todo/pkg/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

var ValidationArrayForCreateBook = ValidationArray{
	ValidationStruct{Field: "title", Kind: KIND_STRING, Required: true},
	ValidationStruct{Field: "author", Kind: KIND_STRING, Required: true},
	ValidationStruct{Field: "pages", Kind: KIND_INT, Required: true},
}

var ValidationArrayForCreateFood = ValidationArray{
	ValidationStruct{Field: "name", Kind: KIND_STRING, Required: true, Source: FROM_FORM},
	ValidationStruct{Field: "kcal", Kind: KIND_FLOAT, Required: true},
	ValidationStruct{Field: "protein", Kind: KIND_FLOAT, Required: true},
	ValidationStruct{Field: "fiber", Kind: KIND_FLOAT, Required: true},
	ValidationStruct{Field: "fat", Kind: KIND_FLOAT, Required: true},
	ValidationStruct{Field: "carbs", Kind: KIND_FLOAT, Required: true},
}

func CreateBookValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		objMap, err := handleValidationArray(ValidationArrayForCreateBook, c)
		if err != nil {
			return utils.HandleEchoError(c, err)
		}
		c.Set("book", models.Book{
			Title:  objMap["title"].(string),
			Author: objMap["author"].(string),
			Pages:  objMap["pages"].(uint),
		})
		return next(c)
	}
}

func CreateFoodValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		objMap, err := handleValidationArray(ValidationArrayForCreateFood, c)
		if err != nil {
			return utils.HandleEchoError(c, err)
		}
		c.Set("food", models.Food_Item{
			Name:    objMap["name"].(string),
			Kcal:    objMap["kcal"].(float32),
			Protein: objMap["protein"].(float32),
			Fiber:   objMap["fiber"].(float32),
			Fat:     objMap["fat"].(float32),
			Carbs:   objMap["carbs"].(float32),
		})
		return next(c)
	}
}

func FoodConsumedValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		objMap, err := handleValidationArray(ValidationArray{
			ValidationStruct{Field: "food_item_id", Kind: KIND_INT, Required: true},
			ValidationStruct{Field: "quantity", Kind: KIND_FLOAT, Required: true},
			ValidationStruct{Field: "date", Kind: KIND_DATE, Required: false, Default: datatypes.Date(time.Now())}, // TODO: need to convert UTC for servers
		}, c)
		if err != nil {
			return utils.HandleEchoError(c, err)
		}
		c.Set("food_consumed", models.FoodConsumption{
			Food_Item_ID: objMap["food_item_id"].(uint),
			Quantity:     objMap["quantity"].(float32),
			Date:         objMap["date"].(datatypes.Date),
			UserID:       c.Get("user_id").(uint),
		})
		return next(c)
	}
}

func FoodConsumptionByDateValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		objMap, err := handleValidationArray(ValidationArray{
			ValidationStruct{Field: "date", Kind: KIND_DATE, Required: false, Source: FROM_QUERY, Default: datatypes.Date(time.Now())},
		}, c)
		if err != nil {
			return utils.HandleEchoError(c, err)
		}
		c.Set("date", objMap["date"].(datatypes.Date))
		return next(c)
	}
}

func GetFoodLogByModeValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		objMap, err := handleValidationArray(ValidationArray{
			ValidationStruct{
				Field:        "mode",
				Kind:         KIND_INT,
				Required:     false,
				ShouldBeFrom: []uint{constants.FOOD_LOG_DAY_MODE, constants.FOOD_LOG_WEEK_MODE},
				Default:      constants.FOOD_LOG_DAY_MODE,
				Source:       FROM_QUERY,
			},
		}, c)
		if err != nil {
			return utils.HandleEchoError(c, err)
		}
		c.Set("mode", objMap["mode"].(uint))
		return next(c)
	}
}
