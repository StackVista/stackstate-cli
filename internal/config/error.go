package config

import (
	"fmt"
	"net/http"
	"strings"

	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

var _ common.CLIError = (*ReadConfError)(nil)        // Compile time check for interface implementation
var _ common.CLIError = (*ValidateContextError)(nil) // Compile time check for interface implementation

type ReadConfError struct {
	RootCause           error
	IsMissingConfigFile bool
}

func (p ReadConfError) Error() string {
	if p.IsMissingConfigFile {
		return fmt.Sprintf("could not load StackState CLI config\n%s\nYou do not have a config file. To create one try running `sts context save --help`", p.RootCause)
	}

	return fmt.Sprintf("could not load StackState CLI config\n%s", p.RootCause)
}

func (p ReadConfError) ExitCode() common.ExitCode {
	return common.ConfigErrorExitCode
}

func (p ReadConfError) ShowUsage() bool {
	return false
}

func (p ReadConfError) GetServerResponse() *http.Response {
	return nil
}

type MissingFieldError struct {
	FieldName string
}

func (s MissingFieldError) Error() string {
	return fmt.Sprintf("Missing field '%v'", s.FieldName)
}

type MustBeOneOfError struct {
	FieldName string
	Value     string
	Choices   []string
}

func (s MustBeOneOfError) Error() string {
	return fmt.Sprintf(
		"Field '%s' cannot be '%s'. Must be one of [%s]",
		s.FieldName,
		s.Value,
		strings.Join(s.Choices, ", "),
	)
}

type ValidateContextError struct {
	ContextName      string
	ValidationErrors []error
}

func (s ValidateContextError) Error() string {
	strs := make([]string, 0)
	for _, e := range s.ValidationErrors {
		strs = append(strs, e.Error())
	}

	contextName := "current"
	if s.ContextName != "" {
		contextName = fmt.Sprintf("'%s'", s.ContextName)
	}

	msg := fmt.Sprintf("Failed to validate the %s context:", contextName)
	for _, str := range strs {
		msg += fmt.Sprintf("\n* %s", str)
	}

	return msg
}

func (p ValidateContextError) ExitCode() common.ExitCode {
	return common.ConfigErrorExitCode
}

func (p ValidateContextError) ShowUsage() bool {
	return false
}

func (p ValidateContextError) GetServerResponse() *http.Response {
	return nil
}
