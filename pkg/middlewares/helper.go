package validators

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RequestResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleQueryResult(queryResult *gorm.DB, c echo.Context, successResponse RequestResponse) error {
	if queryResult.Error != nil {
		return HandleEchoError(c, queryResult.Error)
	}
	if queryResult.RowsAffected == 0 {
		return HandleEchoError(c, gorm.ErrRecordNotFound)
	}
	return c.JSON(http.StatusOK, successResponse)
}

func HandleEchoError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, RequestResponse{Message: err.Error()})
}
