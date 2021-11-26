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

type ValidateConfError struct {
	ValidationErrors []error
}

func (s ValidateConfError) Error() string {
	strs := make([]string, len(s.ValidationErrors))
	for _, e := range s.ValidationErrors {
		strs = append(strs, e.Error())
	}
	return strings.Join(strs, "\n")
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
