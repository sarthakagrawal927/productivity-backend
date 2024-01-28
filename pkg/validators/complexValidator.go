package validators

import (
	"todo/pkg/constants"

	"github.com/labstack/echo/v4"
)

const (
	FROM_BODY  = "body"
	FROM_QUERY = "query"
	FROM_PARAM = "param"
	FROM_FORM  = "form"
)

const (
	KIND_INT    = "int"
	KIND_STRING = "string"
	KIND_BOOL   = "bool"
)

type ValidationRules struct {
}

type ValidationStruct struct {
	Field        string `json:"field"`
	Source       string `json:"source"`
	Kind         string `json:"kind"`
	Required     bool   `json:"required"`
	ShouldBeFrom []uint `json:"shouldInclude"`
	Default      uint   `json:"default"`
}

type ValidationArray []ValidationStruct

var ValidationArrayFor = ValidationArray{
	{
		Field:        "status",
		Source:       "form",
		Kind:         "int",
		Required:     true,
		ShouldBeFrom: []uint{},
		Default:      0,
	},
}

var defaultValidationObj = ValidationStruct{
	Source:       FROM_FORM,
	Kind:         KIND_INT,
	Required:     false,
	ShouldBeFrom: []uint{},
	Default:      0,
}

func getSingleValidationObj(initialObj ValidationStruct) ValidationStruct {
	if initialObj.Source == "" {
		initialObj.Source = defaultValidationObj.Source
	}
	if initialObj.Kind == "" {
		initialObj.Kind = defaultValidationObj.Kind
	}
	// bool is default false
	// default values for interface are also fine
	return initialObj
}

func getValueFromSource(source string, key string, c echo.Context) string {
	switch source {
	case FROM_FORM:
		return c.FormValue(key)
	case FROM_QUERY:
		return c.QueryParam(key)
	case FROM_PARAM:
		return c.Param(key)
	case FROM_BODY:
		return c.FormValue(key)
	default:
		return ""
	}
}

func handleValidationArray(validationArray ValidationArray, c echo.Context) error {
	for _, validationObj := range validationArray {
		value := getValueFromSource(validationObj.Source, validationObj.Field, c)
		switch validationObj.Kind {
		case KIND_INT:
			if validationObj.ShouldBeFrom != nil && len(validationObj.ShouldBeFrom) > 0 {
				if _, err := validateIntFromArray(validationObj.Field, value, validationObj.ShouldBeFrom, validationObj.Default); err != nil {
					return err
				}
			} else {
				if _, err := validateInt(validationObj.Field, value, validationObj.Default); err != nil {
					return err
				}
			}

		case KIND_STRING:
			if _, err := validateString(validationObj.Field, value); err != nil && validationObj.Required {
				return err
			}

		case KIND_BOOL:
			if _, err := validateBool(validationObj.Field, value); err != nil {
				return err
			}
		}
	}
	return nil
}

// Sample validation array
var ValidationArrayForCreateHabit = ValidationArray{
	getSingleValidationObj(ValidationStruct{
		Field:    "title",
		Kind:     KIND_STRING,
		Required: true,
	}),
	getSingleValidationObj(ValidationStruct{
		Field:    "desc",
		Kind:     KIND_STRING,
		Required: true,
	}),
	getSingleValidationObj(ValidationStruct{
		Field:    "target",
		Required: true,
	}),
	getSingleValidationObj(ValidationStruct{
		Field:        "frequency_type",
		Required:     true,
		ShouldBeFrom: constants.HabitFreqTypeList,
	}),
	getSingleValidationObj(ValidationStruct{
		Field:        "mode",
		Required:     true,
		ShouldBeFrom: constants.HabitModeList,
	}),
	getSingleValidationObj(ValidationStruct{
		Field:        "status",
		Required:     true,
		ShouldBeFrom: constants.HabitStatusList,
	}),
	getSingleValidationObj(ValidationStruct{
		Field:    "anti",
		Required: true,
		Kind:     KIND_BOOL,
	}),
}
