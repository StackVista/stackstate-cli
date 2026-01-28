package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/util"
)

const (
	HTTPStatusUnauthorized = 401
)

func NewWriteFileError(err error, filepath string) CLIError {
	return StdCLIError{
		Err:      fmt.Errorf("cannot write to file %s (%s)", filepath, err.Error()),
		exitCode: WriteFileErrorExitCode,
	}
}

func NewReadFileError(err error, filepath string) CLIError {
	return StdCLIError{
		Err:      fmt.Errorf("cannot read file %s (%s)", filepath, err.Error()),
		exitCode: ReadFileErrorExitCode,
	}
}

func NewResponseError(err error, serverResponse *http.Response) CLIError {
	return StdCLIError{Err: err, ServerResponse: serverResponse}
}

func NewConnectError(err error, apiURL string, serverResponse *http.Response) CLIError {
	var statusCode int

	// is access error?
	if serverResponse != nil {
		statusCode = serverResponse.StatusCode
	}

	if statusCode == HTTPStatusUnauthorized {
		var errMessage string

		if apiErr, ok := err.(*stackstate_api.GenericOpenAPIError); ok {
			if detailedErr, ok := apiErr.Model().(stackstate_api.GenericErrorsResponse); ok {
				// Now you can access your 401 error details
				var messages []string
				for _, apiErr := range detailedErr.Errors {
					messages = append(messages, apiErr.Message)
				}
				errMessage = strings.Join(messages, ", ")
				// Some responses do not yield more than the 401, others give in the body the error message as the RefreshToken failed or so
				if errMessage == "" {
					errMessage = "invalid api-token"
				}
			} else {
				errMessage = err.Error()
			}
		} else {
			errMessage = err.Error()
		}

		return StdCLIError{
			Err: fmt.Errorf("could not connect to %s: %s\n"+
				"For more information: https://l.stackstate.com/cli-invalid-api-token", apiURL, errMessage),
			exitCode: ConnectErrorExitCode,
		}
	} else {
		return StdCLIError{
			Err:      fmt.Errorf("could not connect to %s (%s)", apiURL, err),
			exitCode: ConnectErrorExitCode,
		}
	}
}

func NewCLIArgParseError(err error) CLIError {
	return StdCLIError{
		Err:            err,
		ServerResponse: nil,
		showUsage:      true,
		exitCode:       CommandFailedRequirementExitCode,
	}
}

func CheckFlagIsValidChoice(flagName string, flagValue string, choices []string) CLIError {
	if !util.StringInSlice(flagValue, choices) {
		return NewCLIArgParseError(
			fmt.Errorf("invalid '%s' flag value '%s' (must be { %s })", flagName, flagValue, strings.Join(choices, " | ")),
		)
	} else {
		return nil
	}
}

func NewNotFoundError(err error) CLIError {
	return StdCLIError{
		Err:            err,
		ServerResponse: nil,
		showUsage:      false,
		exitCode:       NotFoundExitCode,
	}
}

func NewExecutionError(err error) CLIError {
	return StdCLIError{
		Err:            err,
		ServerResponse: nil,
		showUsage:      true,
		exitCode:       ExecutionErrorCode,
	}
}

func NewAPIVersionError(err error) CLIError {
	return StdCLIError{
		Err:            err,
		ServerResponse: nil,
		showUsage:      true,
		exitCode:       APIVersionErrorCode,
	}
}

func NewAPIClientCreateError(message string) CLIError {
	return StdCLIError{
		Err:            errors.New(message),
		ServerResponse: nil,
		showUsage:      false,
		exitCode:       APIClientCreateErrorCode,
	}
}

func NewRuntimeError(err error) CLIError {
	return StdCLIError{
		Err:            err,
		ServerResponse: nil,
		showUsage:      false,
		exitCode:       ExecutionErrorCode,
	}
}
