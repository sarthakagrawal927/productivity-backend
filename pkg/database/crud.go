package db

import (
	"todo/pkg/utils"

	"github.com/labstack/echo/v4"
)

func InsertIntoDB(c echo.Context, item interface{}) error {
	queryResult := DB_CONNECTION.GetDB().Create(item)
	return utils.HandleQueryResult(queryResult, c, utils.RequestResponse{Message: "Created Successfully", Data: item}, false)
}
