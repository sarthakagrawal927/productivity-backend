package validators

import (
	"todo/pkg/constants"
	"todo/pkg/models"
	utils "todo/pkg/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

// ValidationArrayForCreateHabit defines validation rules for creating a habit
var ValidationArrayForCreateHabit = ValidationArray{
	ValidationStruct{Field: "title", Kind: KIND_STRING, Required: true},
	ValidationStruct{Field: "description", Kind: KIND_STRING, Required: false},

	ValidationStruct{Field: "frequency_type", Required: true, ShouldBeFrom: constants.HabitFreqTypeList},
	ValidationStruct{Field: "upper_limit", Required: false},
	ValidationStruct{Field: "lower_limit", Required: false},
	ValidationStruct{Field: "priority", Required: false, ShouldBeFrom: constants.PriorityTypeList},
	ValidationStruct{Field: "mode", Required: false, ShouldBeFrom: constants.HabitModeList},
}

// ValidationArrayForCreateHabitLog defines validation rules for logging a habit result
var ValidationArrayForCreateHabitLog = ValidationArray{
	ValidationStruct{Field: "habit_id", Kind: KIND_INT, Required: true},
	ValidationStruct{Field: "count", Kind: KIND_INT, Required: true},
	ValidationStruct{Field: "logged_for_date", Kind: KIND_DATE, Required: true},
	ValidationStruct{Field: "mood_rating", Kind: KIND_INT, Required: false, ShouldBeFrom: constants.MoodRatingList, Default: constants.MoodRatingNeutral},
	ValidationStruct{Field: "comment", Kind: KIND_STRING, Required: false},
}

func CreateHabitValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		objMap, err := handleValidationArray(ValidationArrayForCreateHabit, c)
		if err != nil {
			return utils.HandleEchoError(c, err)
		}

		habit := models.Habit{
			UserId:      c.Get("user_id").(uint),
			Title:       objMap["title"].(string),
			Description: objMap["description"].(string),

			FrequencyType: objMap["frequency_type"].(uint),
			UpperLimit:    objMap["upper_limit"].(uint),
			LowerLimit:    objMap["lower_limit"].(uint),
			Priority:      objMap["priority"].(uint),
			Mode:          objMap["mode"].(uint),
			Status:        constants.HabitActive,
		}

		c.Set("habit", habit)
		return next(c)
	}
}

func CreateHabitLogValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		objMap, err := handleValidationArray(ValidationArrayForCreateHabitLog, c)
		if err != nil {
			return utils.HandleEchoError(c, err)
		}

		habitLog := models.HabitLog{
			UserId:        c.Get("user_id").(uint),
			HabitID:       objMap["habit_id"].(uint),
			Count:         objMap["count"].(uint),
			LoggedForDate: objMap["logged_for_date"].(datatypes.Date),
			MoodRating:    objMap["mood_rating"].(uint),
			Comment:       objMap["comment"].(string),
		}

		c.Set("habit_log", habitLog)
		return next(c)
	}
}

func GetSingleHabitValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := validateAndGetId(c.Param("id"))
		if err != nil {
			return utils.HandleEchoError(c, err)
		}

		c.Set("id", int(id))
		return next(c)
	}
}
