package util

import (
	"boilerplate/internal/common"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type AppValidator struct {
	valiadator *validator.Validate
}

func NewAppValidator() *AppValidator {
	return &AppValidator{
		valiadator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (v *AppValidator) ValidateRequest(ctxModel *common.ContextModel, dt interface{}) (map[string]string, error) {
	dtValue := reflect.ValueOf(dt)

	// Check if dt is a pointer and dereference it
	if dtValue.Kind() == reflect.Ptr {
		dtValue = dtValue.Elem()
	}

	maps := make(map[string]string)

	// Validate the struct
	err := v.valiadator.Struct(dtValue.Interface())
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil, err
		}

		// Get type of the struct (unboxed from pointer if necessary)
		dtType := reflect.TypeOf(dtValue.Interface())

		for _, err := range err.(validator.ValidationErrors) {
			// Field that has the validation error
			fieldName := err.StructField()

			// Find the field in the struct by name
			if field, found := dtType.FieldByName(fieldName); found {
				// Print the JSON tag if it exists
				jsonTag := field.Tag.Get("json")
				if jsonTag == "" {
					jsonTag = fieldName // If no JSON tag, use field name
				}

				codeError, templateData := getErrorCode(jsonTag, err)
				maps[jsonTag] = GetI18nErrorMessage(ctxModel.Locale, codeError, templateData) // atau translatorId tergantung locale
			}
		}

		return maps, err
	}

	return nil, nil
}

func getErrorCode(jsonTag string, err validator.FieldError) (string, map[string]interface{}) {
	var codeError string
	templateData := make(map[string]interface{})

	switch err.ActualTag() {
	case "required":
		codeError = "E-400-VAL-1"
	case "min":
		switch err.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
			codeError = "E-400-VAL-5"
		default:
			codeError = "E-400-VAL-2"
		}
		templateData["min"] = err.Param()
	case "max":
		switch err.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
			codeError = "E-400-VAL-6"
		default:
			codeError = "E-400-VAL-3"
		}
		templateData["max"] = err.Param()
	}

	templateData["field"] = jsonTag

	return codeError, templateData
}
