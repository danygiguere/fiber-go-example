package utils

import (
	r "go-fiber-example/m/v2/rules"
	"strconv"
	"strings"
)

type ValidationErrors map[string][]string

const (
	NotNull   = "NotNull"
	MinLength = "MinLength"
)

func splitRule(rule string) (string, int) {
	parts := strings.Split(rule, ":")
	if len(parts) == 1 {
		return parts[0], 0
	} else {
		value, _ := strconv.Atoi(parts[1])
		return parts[0], value
	}
}

func Check(fieldName string, fieldValue interface{}, rules []string, locale string) ValidationErrors {
	ve := make(ValidationErrors)
	for _, rule := range rules {
		ruleKey, ruleValue := splitRule(rule)

		switch ruleKey {
		case NotNull:
			ve = r.NotNull(fieldName, fieldValue, locale, ve)
		case MinLength:
			ve = r.MinLength(fieldName, fieldValue, ruleValue, locale, ve)
		default:
			// Unknown rule, log an error or return an error message
		}
	}
	return ve
}
