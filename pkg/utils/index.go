package utils

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"todo/pkg/constants"
	"todo/pkg/types"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RequestResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleQueryResult(queryResult *gorm.DB, c echo.Context, successResponse RequestResponse, isRead bool) error {
	if queryResult.Error != nil {
		return HandleEchoError(c, queryResult.Error)
	}
	if queryResult.RowsAffected == 0 && !isRead {
		return HandleEchoError(c, gorm.ErrRecordNotFound)
	}
	return c.JSON(http.StatusOK, successResponse)
}

func HandleEchoError(c echo.Context, err error) error {
	log.Println(err)
	return c.JSON(http.StatusBadRequest, RequestResponse{Message: err.Error()})
}

// maybe usable in future
func getDaysCountForHabitFreq(habitFreq uint) int {
	switch habitFreq {
	case constants.HabitDailyFreq:
		return 1
	case constants.HabitWeeklyFreq:
		return 7
	case constants.HabitMonthlyFreq:
		return 30
	default:
		return 1
	}
}

// ParseTime converts a string in "HH:MM" format to an HourMinute struct
func ParseTime(timeStr string) (types.HourMinute, error) {
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return types.HourMinute{}, err
	}
	return types.HourMinute{Hour: t.Hour(), Minute: t.Minute()}, nil
}

// ConvertToScheduleEntry takes a string like "12:00-14:00" and converts it to a ScheduleEntry
func ConvertToScheduleEntry(timeRange string) (types.ScheduleEntry, error) {
	times := strings.Split(timeRange, "-")
	if len(times) != 2 {
		return types.ScheduleEntry{}, fmt.Errorf("invalid time range format")
	}

	startTime, err := ParseTime(times[0])
	if err != nil {
		return types.ScheduleEntry{}, err
	}

	endTime, err := ParseTime(times[1])
	if err != nil {
		return types.ScheduleEntry{}, err
	}

	return types.ScheduleEntry{
		StartTime: startTime,
		EndTime:   endTime,
	}, nil
}

// todo try to make generic
func InsertElementsInSliceAfterIdx(slice []types.ScheduleEntry, elements []types.ScheduleEntry, idx int) []types.ScheduleEntry {
	return append(slice[:idx+1], append(elements, slice[idx+1:]...)...)
}

func IsWeekendToday() bool {
	return time.Now().Weekday() == time.Saturday || time.Now().Weekday() == time.Sunday
}
