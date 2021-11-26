package conf

import (
	"fmt"
	"strings"
)

type MissingFieldError struct {
	FieldName string
}

func (s MissingFieldError) Error() string {
	return fmt.Sprintf("Missing field: %v", s.FieldName)
}

type ValidationError struct {
	ValidationErrors []error
}

func (s ValidationError) Error() string {
	strs := make([]string, len(s.ValidationErrors))
	for _, e := range s.ValidationErrors {
		strs = append(strs, e.Error())
	}
	return strings.Join(strs, "\n")
}

func ValidateConf(conf Conf) error {
	var errors []error
	if conf.ApiUrl == "" {
		errors = append(errors, MissingFieldError{FieldName: "api-url"})
	}
	if conf.ApiToken == "" {
		errors = append(errors, MissingFieldError{FieldName: "api-token"})
	}

	if len(errors) != 0 {
		return ValidationError{ValidationErrors: errors}
	} else {
		return nil
	}
}
