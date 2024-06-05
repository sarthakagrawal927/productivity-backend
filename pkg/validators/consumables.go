package validators

import (
	"todo/pkg/models"
	utils "todo/pkg/utils"

	"github.com/labstack/echo/v4"
)

var ValidationArrayForCreateBook = ValidationArray{
	ValidationStruct{Field: "title", Kind: "string", Required: true},
	ValidationStruct{Field: "author", Kind: "string", Required: true},
	ValidationStruct{Field: "pages", Kind: "int", Required: true},
}

var ValidationArrayForCreateFood = ValidationArray{
	ValidationStruct{Field: "name", Kind: "string", Required: true},
	ValidationStruct{Field: "kcal", Kind: "int", Required: true},
	ValidationStruct{Field: "protein", Kind: "int", Required: true},
	ValidationStruct{Field: "fiber", Kind: "int", Required: true},
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
			Pages:  objMap["pages"].(int),
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
			Kcal:    objMap["kcal"].(int),
			Protein: objMap["protein"].(int),
			Fiber:   objMap["fiber"].(int),
		})
		return next(c)
	}
}
