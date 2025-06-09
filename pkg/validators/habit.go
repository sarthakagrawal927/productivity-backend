package validators

import (
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

		if habit.LowerLimit, err = validateInt("lower_limit", c.FormValue("lower_limit")); err != nil {
			return utils.HandleEchoError(c, err)
		}

		if habit.UpperLimit, err = validateInt("upper_limit", c.FormValue("upper_limit")); err != nil {
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

		if habit.Priority, err = validateIntFromArrayFromForm(c, "priority", constants.PriorityTypeList); err != nil {
			return utils.HandleEchoError(c, err)
		}

		// Handle Category field using constants
		if habit.Category, err = validateIntFromArrayFromForm(c, "category", constants.HabitCategoryList, constants.HabitCategoryProductivity); err != nil {
			// Default to Productivity category if not specified
			habit.Category = constants.HabitCategoryProductivity
		}

		// Handle PreferredWeekdaysMask as a bitmask
		weekdaysMask, err := validateInt("preferred_weekdays_mask", c.FormValue("preferred_weekdays_mask"))
		if err == nil {
			// Ensure the mask is valid (between 0 and 127 - all 7 days)
			if weekdaysMask >= 0 && weekdaysMask <= 127 {
				habit.PreferredWeekdaysMask = uint8(weekdaysMask)
			}
		}

		// Handle preferred start time
		if startTimeStr := c.FormValue("preferred_start_time"); startTimeStr != "" {
			timeVal, err := validateTime("preferred_start_time", startTimeStr)
			if err == nil {
				timeData := datatypes.NewTime(timeVal.Hour(), timeVal.Minute(), timeVal.Second(), 0)
				habit.PreferredStartTime = &timeData
			}
		}

		// Handle preferred month date
		monthDate, err := validateInt("preferred_month_date", c.FormValue("preferred_month_date"))
		if err == nil && monthDate > 0 && monthDate <= 31 {
			habit.PreferredMonthDate = &monthDate
		}

		approxTimeNeeded, err := validateInt("approx_time_needed", c.FormValue("approx_time_needed"))
		if err == nil {
			habit.ApproxTimeNeeded = &approxTimeNeeded
		}

		habit.UserId = c.Get("user_id").(uint)
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

		if habitLog.Count, err = validateInt("count", c.FormValue("count")); err != nil {
			return utils.HandleEchoError(c, err)
		}

		if habitLog.LoggedForDate, err = validateDate("result_date", c.FormValue("result_date")); err != nil {
			return utils.HandleEchoError(c, err)
		}

		timeVal, err := validateTime("result_start_time", c.FormValue("result_start_time"))
		if err != nil {
			return utils.HandleEchoError(c, err)
		}
		// Convert time.Time to datatypes.Time
		timeData := datatypes.NewTime(timeVal.Hour(), timeVal.Minute(), timeVal.Second(), 0)
		habitLog.HabitStartTime = &timeData

		// Handle mood rating using constants
		if habitLog.MoodRating, err = validateIntFromArrayFromForm(c, "mood_rating", constants.MoodRatingList, constants.MoodRatingNeutral); err != nil {
			// Default to neutral mood if not specified
			habitLog.MoodRating = constants.MoodRatingNeutral
		}

		commentValue, _ := validateStringFromForm(c, "comment")
		if commentValue != "" {
			habitLog.Comment = &commentValue
		}

		habitLog.UserId = c.Get("user_id").(uint)
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
