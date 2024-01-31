package validators

import (
	"time"
	"todo/pkg/constants"
	middleware "todo/pkg/middlewares"
	"todo/pkg/models"

	"github.com/labstack/echo/v4"
)

func CreateHabitValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		habit := models.Habit{}
		var err error

		if habit.Title, err = validateStringFromForm(c, "title"); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if habit.Desc, err = validateStringFromForm(c, "desc"); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if habit.Target, err = validateInt("target", c.FormValue("target")); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if habit.FrequencyType, err = validateIntFromArrayFromForm(c, "frequency_type", constants.HabitFreqTypeList); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if habit.Mode, err = validateIntFromArrayFromForm(c, "mode", constants.HabitModeList); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if habit.Status, err = validateIntFromArrayFromForm(c, "status", constants.HabitStatusList, constants.HabitActive); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		// if habit.Anti, err = validateBool("anti", c.FormValue("anti")); err != nil {
		// 	return middleware.HandleEchoError(c, err)
		// }

		if habit.ApproxTimeNeeded, err = validateInt("approx_time_needed", c.FormValue("approx_time_needed")); err != nil {
			return middleware.HandleEchoError(c, err)
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
			return middleware.HandleEchoError(c, err)
		}

		if habitLog.ResultCount, err = validateInt("count", c.FormValue("count")); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		var stringDate string
		if stringDate, err = validateStringFromForm(c, "date"); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if habitLog.Date, err = time.Parse("YYYY-MM-DD", stringDate); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		c.Set("habit_log", habitLog)
		return next(c)
	}
}

func GetSingleHabitValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := validateAndGetId(c.Param("id"))
		if err != nil {
			return middleware.HandleEchoError(c, err)
		}

		c.Set("id", int(id))
		return next(c)
	}
}
