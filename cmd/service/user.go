package service

import (
	"net/http"
	"strings"
	"time"
	db "todo/pkg/database"
	"todo/pkg/models"
	utils "todo/pkg/utils"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type HabitLogWithHabitType struct {
	HabitID     uint           `json:"habit_id"`
	ResultCount uint           `json:"result_count"`
	ResultDate  datatypes.Date `json:"result_time"`
	Comment     string         `json:"comment"`

	Title         string `json:"title"`
	Mode          uint   `json:"mode"`
	FrequencyType uint   `json:"frequency_type"`
	Anti          bool   `json:"anti"`
	gorm.Model
}

func GetDailyLogs(c echo.Context) error {
	var habitLog []HabitLogWithHabitType
	currentTime := time.Now()
	formattedTodayDate := currentTime.Format("2006-01-02 15:04:05")
	formattedYesterdayDate := currentTime.AddDate(0, 0, -1).Format("2006-01-02 15:04:05")

	queryResult := db.DB_CONNECTION.GetDB().Table("habit_logs").Select("habit_logs.*, habits.mode, habits.frequency_type, habits.title").Joins("LEFT JOIN habits on habits.id = habit_logs.habit_id").Where("habit_logs.result_date IN (?, ?)", formattedTodayDate, formattedYesterdayDate).Scan(&habitLog)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Success", Data: habitLog}, true)
}

func GetTodaySchedule(c echo.Context) error {
	formattedSchedule, schedule, taskEntries := getFormattedSchedule()
	return c.JSON(http.StatusOK, utils.RequestResponse{Message: "Success", Data: map[string]interface{}{
		"formatted_schedule": formattedSchedule,
		"schedule":           schedule,
		"task_entries":       taskEntries,
	}})
}

func ValidateGoogleJWT(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")

	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, "Missing authorization header")
	}

	tokenString := strings.SplitN(authHeader, " ", 2)[1] // Extract token from "Bearer <token>" format

	claims, err := utils.ParseJWT(tokenString)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	if claims["email"] == nil || claims["email"].(string) == "" {
		return c.JSON(http.StatusUnauthorized, "Missing email claim")
	}

	// Access authenticated user's email from claims
	email := claims["email"].(string)

	// once we have the email, create a user and return a jwt token
	userId, err := CreateUserWithEmailIfNotExists(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error creating user")
	}

	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims = token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["id"] = userId

	// Sign the token with a secret
	tokenString, err = token.SignedString([]byte("simple"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error signing token")
	}
	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

func CreateUserWithEmailIfNotExists(email string) (uint, error) {
	user := &models.User{Email: email}
	queryResult := db.DB_CONNECTION.GetDB().First(&user, "email = ?", email) // Modify the query to use the email field
	if queryResult.RowsAffected == 0 {
		user.Email = email
		queryResult = db.DB_CONNECTION.GetDB().Create(user)
	}
	if queryResult.Error != nil {
		return 0, queryResult.Error
	}
	return user.ID, nil
}
