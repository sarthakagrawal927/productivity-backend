package validators

import (
	middleware "todo/pkg/middlewares"
	"todo/pkg/models"

	"github.com/labstack/echo/v4"
)

var ValidationArrayForCreateConsumable = ValidationArray{
	getSingleValidationObj(ValidationStruct{
		Field:    "habit_id",
		Source:   FROM_FORM,
		Kind:     KIND_INT,
		Required: true,
	}),
	getSingleValidationObj(ValidationStruct{
		Field:    "smallest_unit_label",
		Source:   FROM_FORM,
		Kind:     KIND_STRING,
		Required: true,
	}),
	getSingleValidationObj(ValidationStruct{
		Field:    "num_total_unit",
		Source:   FROM_FORM,
		Kind:     KIND_INT,
		Required: true,
	}),
	getSingleValidationObj(ValidationStruct{
		Field:    "time_per_unit",
		Source:   FROM_FORM,
		Kind:     KIND_INT,
		Required: true,
	}),
	getSingleValidationObj(ValidationStruct{
		Field:    "num_remaining_unit",
		Source:   FROM_FORM,
		Kind:     KIND_INT,
		Required: true,
	}),
}

func CreateConsumableValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		completeValidationArray := append(ValidationArrayForCreateConsumable, ValidationArrayForMeta...)
		objMap, err := handleValidationArray(completeValidationArray, c)
		if err != nil {
			return middleware.HandleEchoError(c, err)
		}
		c.Set("consumable", models.Consumable{
			HabitID: objMap["habit_id"].(uint),
			Meta: models.Meta{
				Title: objMap["title"].(string),
				Desc:  objMap["desc"].(string),
			},
			SmallestUnitLabel: objMap["smallest_unit_label"].(string),
			NumTotalUnit:      objMap["num_total_unit"].(uint),
			TimePerUnit:       objMap["time_per_unit"].(uint),
			NumRemainingUnit:  objMap["num_remaining_unit"].(uint),
		})
		return next(c)
	}
}
