package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"todo/pkg/constants"
	"todo/pkg/types"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
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

// ConvertToScheduleEntryFromTime converts a datatypes.Time to a ScheduleEntry
func ConvertToScheduleEntryFromTime(t datatypes.Time) (types.ScheduleEntry, error) {
	// Extract hours and minutes from the time string representation
	// datatypes.Time is stored as a string in format "15:04:05"
	var hour, minute int
	_, err := fmt.Sscanf(t.String(), "%d:%d:", &hour, &minute)
	if err != nil {
		return types.ScheduleEntry{}, err
	}

	// Create a schedule entry with 1 hour duration
	timeStr := fmt.Sprintf("%02d:%02d", hour, minute)
	endTimeStr := fmt.Sprintf("%02d:%02d", (hour+1)%24, minute)

	return ConvertToScheduleEntry(timeStr + "-" + endTimeStr)
}

func InsertElementsInSliceAfterIdx[T any](slice []T, elements []T, idx int) []T {
	return append(slice[:idx+1], append(elements, slice[idx+1:]...)...)
}

func IsWeekendToday() bool {
	return time.Now().Weekday() == time.Saturday || time.Now().Weekday() == time.Sunday
}

func ReadAllCookies(c echo.Context) {
	cookies := c.Cookies()
	fmt.Printf("Number of cookies: %d\n", len(cookies))
	for _, cookie := range cookies {
		fmt.Printf("Cookie Name: %s, Value: %s\n", cookie.Name, cookie.Value)
	}
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("signature invalid")
		}
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	// Extract and verify email from claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
