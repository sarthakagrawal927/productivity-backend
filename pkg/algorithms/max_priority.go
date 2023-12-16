package algorithms

import (
	db "todo/pkg/database"
	validators "todo/pkg/middlewares"
	"todo/pkg/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetMaxPriorityTask(c echo.Context) error {
	var tasks []models.Task
	//status := c.Get("status").(uint)
	var queryResult *gorm.DB
	queryResult = db.DB_CONNECTION.GetDB().Order("Priority desc").First(&tasks)

	return validators.HandleQueryResult(queryResult, c, validators.RequestResponse{Message: "Success", Data: tasks}, true)
}
