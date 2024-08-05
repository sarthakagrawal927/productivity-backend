package validators

import (
	"fmt"
	"todo/pkg/constants"
	"todo/pkg/models"

	utils "todo/pkg/utils"

	"github.com/labstack/echo/v4"
)

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
			return utils.HandleEchoError(c, err)
		}

		if task.Status, err = validateIntFromArrayFromForm(c, "status", constants.TaskTypeList, constants.Todo); err != nil {
			return utils.HandleEchoError(c, err)
		}

		if task.TimeToSpend, err = validateInt("time_to_spend", c.FormValue("time_to_spend"), 0); err != nil {
			return utils.HandleEchoError(c, err)
		}

		if task.Priority, err = validateIntFromArrayFromForm(c, "priority", constants.PriorityTypeList, constants.NoPriority); err != nil {
			return utils.HandleEchoError(c, err)
		}

		taskDeadline, err := validateDate("deadline", c.FormValue("deadline"))
		if err == nil {
			task.Deadline = &taskDeadline
		}

		task.Desc = c.FormValue("desc")
		task.UserId = c.Get("user_id").(uint)
		c.Set("task", task)
		return next(c)
	}
}

func GetTasksValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		status, err := validateInt("status", c.QueryParam("status"), constants.AllStatus)
		if err != nil {
			return utils.HandleEchoError(c, err)
		}
		c.Set("status", status)
		return next(c)
	}
}

func DeleteTaskValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := validateAndGetId(c.FormValue("id"))
		if err != nil {
			return utils.HandleEchoError(c, err)
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
			return utils.HandleEchoError(c, err)
		}

		if c.FormValue("title") != "" {
			updateObj["title"] = c.FormValue("title")
		}

		if c.FormValue("desc") != "" {
			updateObj["desc"] = c.FormValue("desc")
		}

		if c.FormValue("status") != "" {
			if updateObj["status"], err = validateIntFromArrayFromForm(c, "status", constants.TaskTypeList); err != nil {
				return utils.HandleEchoError(c, err)
			}
		}

		if c.FormValue("priority") != "" {
			if updateObj["priority"], err = validateIntFromArrayFromForm(c, "priority", constants.PriorityTypeList); err != nil {
				return utils.HandleEchoError(c, err)
			}
		}

		if len(updateObj) == 0 {
			return utils.HandleEchoError(c, fmt.Errorf("no update field"))
		} else {
			updateObj["user_id"] = c.Get("user_id").(uint)
			c.Set("updateObj", updateObj)
		}

		return next(c)
	}
}
