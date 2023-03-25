package rules

import "fmt"

func NotNull(fieldName string, fieldValue interface{}, locale string, ve map[string][]string) map[string][]string {
	if fieldValue == nil || fieldValue == "" {
		ve[fieldName] = append(ve[fieldName], fmt.Sprintf("The field '%s' cannot be null", fieldName))
	}
	return ve
}

func MinLength(fieldName string, fieldValue interface{}, length int, locale string, ve map[string][]string) map[string][]string {
	if s, ok := fieldValue.(string); !ok || len(s) < length {
		ve[fieldName] = append(ve[fieldName], fmt.Sprintf("The field '%s' must have a minimum length of %d", fieldName, length))
	}
	return ve
}
