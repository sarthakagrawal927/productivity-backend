package validators

import (
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
	ValidationStruct{Field: "kcal", Kind: KIND_INT, Required: true},
	ValidationStruct{Field: "protein", Kind: KIND_INT, Required: true},
	ValidationStruct{Field: "fiber", Kind: KIND_INT, Required: true},
}

var ValidationArrayForFoodConsumed = ValidationArray{
	ValidationStruct{Field: "food_item_id", Kind: KIND_INT, Required: true},
	ValidationStruct{Field: "quantity", Kind: KIND_FLOAT, Required: true},
	ValidationStruct{Field: "date", Kind: KIND_DATE, Required: true},
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
			Kcal:    objMap["kcal"].(uint),
			Protein: objMap["protein"].(uint),
			Fiber:   objMap["fiber"].(uint),
		})
		return next(c)
	}
}

func FoodConsumedValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		objMap, err := handleValidationArray(ValidationArrayForFoodConsumed, c)
		if err != nil {
			return utils.HandleEchoError(c, err)
		}
		c.Set("food_consumed", models.FoodConsumption{
			Food_Item_ID: objMap["food_item_id"].(uint),
			Quantity:     objMap["quantity"].(float32),
			Date:         objMap["date"].(datatypes.Date),
		})
		return next(c)
	}
}

func FoodConsumptionByDateValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		objMap, err := handleValidationArray(ValidationArray{
			ValidationStruct{Field: "date", Kind: KIND_DATE, Required: true, Source: FROM_QUERY},
		}, c)
		if err != nil {
			return utils.HandleEchoError(c, err)
		}
		c.Set("date", objMap["date"].(datatypes.Date))
		return next(c)
	}
}
