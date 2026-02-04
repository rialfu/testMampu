package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func TranslateValidationError(err error, validationMessages map[string]map[string]string) map[string]string {
	result := make(map[string]string)

	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return result
	}

	for _, fe := range ve {
		field := fe.Field()
		tag := fe.Tag()

		if fieldMsg, ok := validationMessages[field]; ok {
			if msg, ok := fieldMsg[tag]; ok && msg != "" {
				if strings.Contains(msg, "%s") {
					result[field] = fmt.Sprintf(msg, fe.Param())
				} else {
					result[field] = msg
				}
				continue
			}
		}

		// 2️⃣ fallback ke default validator
		result[field] = defaultValidatorMessage(fe)
	}

	return result
}
func defaultValidatorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "field is required"
	case "email":
		return "invalid email format"
	case "min":
		kind := fe.Kind().String()
		if kind == "string" {
			return fmt.Sprintf("minimum %s characters", fe.Param())
		}
		return fmt.Sprintf("minimum value %s", fe.Param())

	case "max":
		kind := fe.Kind().String()
		if kind == "string" {
			return fmt.Sprintf("maximum %s characters", fe.Param())
		}
		return fmt.Sprintf("maximum value %s", fe.Param())
	case "numeric":
		return "field must numeric"
	default:
		return "invalid value"
	}
}
