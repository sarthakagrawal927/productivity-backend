package utils

import (
	"fmt"
	"log"
	"net/http"
	"todo/pkg/constants"

	"github.com/golang-jwt/jwt"
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

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("simple"), nil
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
