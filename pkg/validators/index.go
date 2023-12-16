package validators

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/labstack/echo/v4"
)

/*
validateInt validates if the value is a number and returns the sanitized value.
Extra params first field will be optional parameter
*/
func validateInt(key, value string, extraParams ...uint) (uint, error) {
	var sanitizedInt uint

	if value == "" {
		hasDefaultValue := len(extraParams) > 0
		if hasDefaultValue {
			return extraParams[0], nil
		}

		return sanitizedInt, fmt.Errorf("%s is required", key)
	}

	sanitizedSignedInt, err := strconv.Atoi(value)
	if err != nil {
		return sanitizedInt, fmt.Errorf("%s should be number: %v", key, err)
	} else {
		sanitizedInt = uint(sanitizedSignedInt)
	}

	return sanitizedInt, nil
}

func validateString(key, value string) (string, error) {
	if value == "" {
		return value, fmt.Errorf("%s is required", key)
	}
	return value, nil
}

func validateStringFromForm(c echo.Context, key string) (string, error) {
	return validateString(key, c.FormValue(key))
}

func validateIntFromArray(status string, options []uint, extraParams ...uint) (uint, error) {
	sanitizedStatus, err := validateInt("status", status, extraParams...)
	if err != nil {
		return sanitizedStatus, err
	}

	if !slices.Contains(options, uint(sanitizedStatus)) {
		return sanitizedStatus, fmt.Errorf("status should be one of %v", options)
	}

	return sanitizedStatus, nil
}

// consider making a function like this
func validateTitleDescInterface(c echo.Context, obj interface{}) (interface{}, error) {
	var err error
	if obj.(map[string]interface{})["title"], err = validateStringFromForm(c, "title"); err != nil {
		return nil, err
	}
	if obj.(map[string]interface{})["desc"], err = validateStringFromForm(c, "desc"); err != nil {
		return nil, err
	}
	return obj, nil
}
