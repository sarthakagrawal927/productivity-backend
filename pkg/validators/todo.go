package validators

import (
	"fmt"
	"slices"
	"todo/pkg/constants"
	"todo/pkg/models"

	middleware "todo/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

func validateIntFromArray(status string, options []uint, extraParams ...uint) (uint, error) {
	sanitizedStatus, err := ValidateInt("status", status, extraParams...)
	if err != nil {
		return sanitizedStatus, err
	}

	if !slices.Contains(options, uint(sanitizedStatus)) {
		return sanitizedStatus, fmt.Errorf("status should be one of %v", options)
	}

	return sanitizedStatus, nil
}

func validateAndGetId(id string) (uint, error) {
	sanitizedId, err := ValidateInt("id", id)
	if err != nil {
		return uint(sanitizedId), err
	}

	return uint(sanitizedId), nil
}

func CreateTaskValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		task := models.Task{}
		var err error

		if task.Title, err = ValidateStringFromForm(c, "title"); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if task.Status, err = validateIntFromArray(c.FormValue("status"), constants.TaskTypeList, constants.Todo); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if task.Complexity, err = validateIntFromArray(c.FormValue("complexity"), constants.ComplexityTypeList, constants.NoComplexity); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		if task.Priority, err = validateIntFromArray(c.FormValue("priority"), constants.PriorityTypeList, constants.NoPriority); err != nil {
			return middleware.HandleEchoError(c, err)
		}

		task.Desc = c.FormValue("desc")
		c.Set("task", task)
		return next(c)
	}
}

func GetTasksValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		status, err := ValidateInt("status", c.QueryParam("status"), constants.AllStatus)
		if err != nil {
			return middleware.HandleEchoError(c, err)
		}
		c.Set("status", status)
		return next(c)
	}
}

func DeleteTaskValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := validateAndGetId(c.FormValue("id"))
		if err != nil {
			return middleware.HandleEchoError(c, err)
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
			return middleware.HandleEchoError(c, err)
		}

		if c.FormValue("title") != "" {
			updateObj["title"] = c.FormValue("title")
		}

		if c.FormValue("desc") != "" {
			updateObj["desc"] = c.FormValue("desc")
		}

		if c.FormValue("status") != "" {
			if updateObj["status"], err = validateIntFromArray(c.FormValue("status"), constants.TaskTypeList); err != nil {
				return middleware.HandleEchoError(c, err)
			}
		}

		if c.FormValue("complexity") != "" {
			if updateObj["complexity"], err = validateIntFromArray(c.FormValue("complexity"), constants.ComplexityTypeList); err != nil {
				return middleware.HandleEchoError(c, err)
			}
		}

		if c.FormValue("priority") != "" {
			if updateObj["priority"], err = validateIntFromArray(c.FormValue("priority"), constants.PriorityTypeList); err != nil {
				return middleware.HandleEchoError(c, err)
			}
		}

		if len(updateObj) == 0 {
			return middleware.HandleEchoError(c, fmt.Errorf("no update field"))
		} else {
			c.Set("updateObj", updateObj)
		}

		return next(c)
	}
}
