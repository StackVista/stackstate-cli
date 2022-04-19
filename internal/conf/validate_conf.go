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

type MustBeOneOfError struct {
	FieldName string
	Value     string
	Choices   []string
}

func (s MustBeOneOfError) Error() string {
	return fmt.Sprintf(
		"Field %s cannot be '%s'. Must be one of: [%s]",
		s.FieldName,
		s.Value,
		strings.Join(s.Choices, ", "),
	)
}

type ValidateConfError struct {
	ValidationErrors []error
}

func (s ValidateConfError) Error() string {
	strs := make([]string, 0)
	for _, e := range s.ValidationErrors {
		strs = append(strs, e.Error())
	}
	return "Validation errors:\n * " + strings.Join(strs, "\n * ")
}

func ValidateConf(conf Conf) error {
	var errors []error

	validate(conf, &errors)

	if len(errors) != 0 {
		return ValidateConfError{ValidationErrors: errors}
	} else {
		return nil
	}
}
