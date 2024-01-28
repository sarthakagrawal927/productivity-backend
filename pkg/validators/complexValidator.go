package validators

import (
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

func handleValidationArray(validationArray ValidationArray, c echo.Context) (map[string]interface{}, error) {
	tempInterface := make(map[string]interface{})
	var err error
	for _, validationObj := range validationArray {
		value := getValueFromSource(validationObj.Source, validationObj.Field, c)
		switch validationObj.Kind {
		case KIND_INT:
			// TODO: Rewrite validateIntFromArray & validateInt to better utilize this function
			if validationObj.ShouldBeFrom != nil && len(validationObj.ShouldBeFrom) > 0 {
				if tempInterface[validationObj.Field],
					err = validateIntFromArray(validationObj.Field, value, validationObj.ShouldBeFrom, validationObj.Default); err != nil {
					return nil, err
				}
			} else {
				if validationObj.Required {
					if tempInterface[validationObj.Field], err = validateInt(validationObj.Field, value); err != nil {
						return nil, err
					}
				} else {
					if tempInterface[validationObj.Field], err = validateInt(validationObj.Field, value, validationObj.Default); err != nil {
						return nil, err
					}
				}
			}

		case KIND_STRING:
			if tempInterface[validationObj.Field], err = validateString(validationObj.Field, value); err != nil && validationObj.Required {
				return nil, err
			}

		case KIND_BOOL:
			if tempInterface[validationObj.Field], err = validateBool(validationObj.Field, value); err != nil {
				return nil, err
			}
		}
	}
	return tempInterface, nil
}

var ValidationArrayForMeta = ValidationArray{
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
}
