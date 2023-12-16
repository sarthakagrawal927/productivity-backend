package validators

import (
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

		if habit.Goal, err = validateInt("goal", c.FormValue("goal")); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if habit.FrequencyType, err = validateIntFromArray(c.FormValue("frequency_type"), constants.HabitFreqTypeList); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if habit.Mode, err = validateIntFromArray(c.FormValue("mode"), constants.HabitModeList); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if habit.Status, err = validateIntFromArray(c.FormValue("status"), constants.HabitStatusList, constants.HabitActive); err != nil {
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
