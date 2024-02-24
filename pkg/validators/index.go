package validators

import (
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
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

// TODO: have ability to make string optional and check for length
func validateStringFromForm(c echo.Context, key string) (string, error) {
	return validateString(key, c.FormValue(key))
}

func validateIntFromArrayFromForm(c echo.Context, key string, options []uint, extraParams ...uint) (uint, error) {
	return validateIntFromArray(key, c.FormValue(key), options, extraParams...)
}

func validateIntFromArray(key, status string, options []uint, extraParams ...uint) (uint, error) {
	sanitizedStatus, err := validateInt(key, status, extraParams...)
	if err != nil {
		return sanitizedStatus, err
	}

	if sanitizedStatus == 0 {
		return sanitizedStatus, nil
	}

	if !slices.Contains(options, uint(sanitizedStatus)) {
		return sanitizedStatus, fmt.Errorf("%s should be one of %v", key, options)
	}

	return sanitizedStatus, nil
}

func validateBool(key, value string) (bool, error) {
	if value == "" {
		return false, fmt.Errorf("%s is required", key)
	}

	sanitizedBool, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("%s should be boolean: %v", key, err)
	}

	return sanitizedBool, nil
}

func validateDate(key, value string) (datatypes.Date, error) {
	var stringDate string
	var err error
	if stringDate, err = validateString(key, value); err != nil {
		return datatypes.Date{}, err
	}

	if dateTimeVal, err := time.Parse("1/2/2006", stringDate); err != nil {
		return datatypes.Date{}, err
	} else {
		return datatypes.Date(dateTimeVal), nil
	}
}
