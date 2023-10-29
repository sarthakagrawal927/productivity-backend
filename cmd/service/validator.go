package service

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"todo/pkg/constants"
	"todo/pkg/models"

	"github.com/labstack/echo/v4"
)

func validateAndSanitizeCreateTodo(c echo.Context) (models.Task, error) {
	task := models.Task{}

	// check whether title exists and is string
	title := c.FormValue("title")
	if title == "" {
		return task, errors.New("title is required")
	} else {
		task.Title = title
	}

	// check whether status exists and is number
	if c.FormValue("status") == "" {
		task.Status = int(constants.Todo)
	} else {
		status, err := validateAndGetStatus(c.FormValue("status"))
		if err != nil {
			return task, err
		}
		task.Status = status
	}

	return task, nil
}

func validateAndSanitizeGetTodo(c echo.Context) (int, error) {
	if c.QueryParam("status") == "" {
		return constants.AllStatus, nil
	}
	status, err := validateAndGetStatus(c.QueryParam("status"))
	if err != nil {
		return status, err
	}

	return status, nil
}

func validateAndGetStatus(status string) (int, error) {
	var sanitizedStatus int

	sanitizedStatus, err := strconv.Atoi(status)
	if err != nil {
		return sanitizedStatus, fmt.Errorf("status should be number: %v", err)
	}

	if !slices.Contains(constants.TaskTypeList, sanitizedStatus) {
		return sanitizedStatus, fmt.Errorf("status should be one of %v", constants.TaskTypeList)
	}

	return sanitizedStatus, nil
}

func validateAndGetId(id string) (uint, error) {
	var sanitizedId int

	if id == "" {
		return uint(sanitizedId), errors.New("id is required")
	}

	sanitizedId, err := (strconv.Atoi(id))
	if err != nil {
		return uint(sanitizedId), errors.New("id should be number")
	}

	return uint(sanitizedId), nil
}

func validateAndSanitizeUpdateTodoByID(c echo.Context) (uint, error) {
	var sanitizedID uint

	id, err := validateAndGetId(c.FormValue("id"))
	if err != nil {
		return sanitizedID, err
	}

	updateObj := make(map[string]interface{})

	if c.FormValue("title") != "" {
		updateObj["title"] = c.FormValue("title")
	}
	if c.FormValue("desc") != "" {
		updateObj["desc"] = c.FormValue("desc")
	}

	status := c.FormValue("status")

	if status != "" {
		status, err := validateAndGetStatus(status)
		if err != nil {
			return sanitizedID, err
		}
		updateObj["status"] = status
	}

	if len(updateObj) == 0 {
		return sanitizedID, errors.New("no update field")
	} else {
		c.Set("updateObj", updateObj)
	}

	return id, nil
}
