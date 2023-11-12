package validators

import (
	"fmt"
	"slices"
	"todo/pkg/constants"
	"todo/pkg/models"

	"github.com/labstack/echo/v4"
)

func validateAndGetStatus(status string, extraParams ...uint) (uint, error) {
	sanitizedStatus, err := validateInt("status", status, extraParams...)
	if err != nil {
		return sanitizedStatus, err
	}

	if !slices.Contains(constants.TaskTypeList, uint(sanitizedStatus)) {
		return sanitizedStatus, fmt.Errorf("status should be one of %v", constants.TaskTypeList)
	}

	return sanitizedStatus, nil
}

func validateAndGetId(id string) (uint, error) {
	sanitizedId, err := validateInt("id", id)
	if err != nil {
		return uint(sanitizedId), err
	}

	return uint(sanitizedId), nil
}

func CreateTaskValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		task := models.Task{}
		var err error

		if task.Title, err = validateStringFromForm(c, "title"); err != nil {
			return HandleEchoError(c, err)
		}

		if task.Status, err = validateAndGetStatus(c.FormValue("status"), constants.Todo); err != nil {
			return HandleEchoError(c, err)
		}

		task.Desc = c.FormValue("desc")
		c.Set("task", task)
		return next(c)
	}
}

func GetTasksValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		status, err := validateInt("status", c.QueryParam("status"), constants.AllStatus)
		if err != nil {
			return HandleEchoError(c, err)
		}
		c.Set("status", status)
		return next(c)
	}
}

func DeleteTaskValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := validateAndGetId(c.FormValue("id"))
		if err != nil {
			return HandleEchoError(c, err)
		}
		c.Set("id", id)
		return next(c)
	}
}

func UpdateTaskValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		updateObj := make(map[string]interface{})
		var err error

		if updateObj["id"], err = validateAndGetId(c.FormValue("id")); err != nil {
			return HandleEchoError(c, err)
		}

		if c.FormValue("title") != "" {
			updateObj["title"] = c.FormValue("title")
		}

		if c.FormValue("desc") != "" {
			updateObj["desc"] = c.FormValue("desc")
		}

		if c.FormValue("status") != "" {
			if updateObj["status"], err = validateAndGetStatus(c.FormValue("status")); err != nil {
				return HandleEchoError(c, err)
			}
		}

		if len(updateObj) == 0 {
			return HandleEchoError(c, fmt.Errorf("no update field"))
		} else {
			c.Set("updateObj", updateObj)
		}

		return next(c)
	}
}
