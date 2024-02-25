package utils

import (
	"log"
	"net/http"
	"sync"
	"todo/pkg/constants"

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

func FetchDataAsync(slice interface{}, dbInstance *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	dbInstance.Find(slice)
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
