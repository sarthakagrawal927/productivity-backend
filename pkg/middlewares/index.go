package validators

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

/*
ValidateInt validates if the value is a number and returns the sanitized value.
Extra params first field will be optional parameter
*/
func validateInt(key, value string, extraParams ...uint) (uint, error) {
	var sanitizedInt uint

	if value == "" {
		var hasDefaultValue = len(extraParams) > 0
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
