package validators

import (
	"time"
	"todo/pkg/constants"
	"todo/pkg/models"
	utils "todo/pkg/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

func CreateHabitValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		habit := models.Habit{}
		var err error

		if habit.Title, err = validateStringFromForm(c, "title"); err != nil {
			return utils.HandleEchoError(c, err)
		}

		if habit.Desc, err = validateStringFromForm(c, "desc"); err != nil {
			return utils.HandleEchoError(c, err)
		}

		if habit.Target, err = validateInt("target", c.FormValue("target")); err != nil {
			return utils.HandleEchoError(c, err)
		}

		if habit.FrequencyType, err = validateIntFromArrayFromForm(c, "frequency_type", constants.HabitFreqTypeList); err != nil {
			return utils.HandleEchoError(c, err)
		}

		if habit.Mode, err = validateIntFromArrayFromForm(c, "mode", constants.HabitModeList); err != nil {
			return utils.HandleEchoError(c, err)
		}

		if habit.Status, err = validateIntFromArrayFromForm(c, "status", constants.HabitStatusList, constants.HabitActive); err != nil {
			return utils.HandleEchoError(c, err)
		}

		if habit.Anti, err = validateBool("anti", c.FormValue("anti")); err != nil {
			return utils.HandleEchoError(c, err)
		}

		if habit.ApproxTimeNeeded, err = validateInt("approx_time_needed", c.FormValue("approx_time_needed")); err != nil {
			return utils.HandleEchoError(c, err)
		}

		c.Set("habit", habit)
		return next(c)
	}
}

func GetHabitsValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func CreateHabitLogValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		habitLog := models.HabitLog{}
		var err error

		if habitLog.HabitID, err = validateInt("habit_id", c.FormValue("habit_id")); err != nil {
			return utils.HandleEchoError(c, err)
		}

		if habitLog.ResultCount, err = validateInt("count", c.FormValue("count")); err != nil {
			return utils.HandleEchoError(c, err)
		}

		var stringDate string
		if stringDate, err = validateStringFromForm(c, "result_date"); err != nil {
			return utils.HandleEchoError(c, err)
		}

		if dateTimeVal, err := time.Parse("1/2/2006", stringDate); err != nil {
			return utils.HandleEchoError(c, err)
		} else {
			habitLog.ResultDate = datatypes.Date(dateTimeVal)
		}

		habitLog.Comment, _ = validateStringFromForm(c, "comment")

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
