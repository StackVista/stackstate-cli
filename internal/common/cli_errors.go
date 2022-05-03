package common

import (
	"fmt"
	"net/http"
	"strings"

	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

const (
	HTTPStatusUnauthorized = 401
)

type CLIError interface {
	error
	ExitCode() ExitCode
	Error() string
	ShowUsage() bool
	GetServerResponse() *http.Response
}

type StdCLIError struct {
	Err            error
	ServerResponse *http.Response
	showUsage      bool
	exitCode       int
}

func (p StdCLIError) Error() string {
	return p.Err.Error()
}

func (p StdCLIError) ExitCode() ExitCode {
	return p.exitCode
}

func (p StdCLIError) ShowUsage() bool {
	return p.showUsage
}

func (p StdCLIError) GetServerResponse() *http.Response {
	return p.ServerResponse
}

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

func NewConnectError(err error, serverResponse *http.Response) CLIError {
	var statusCode int

	// is access error?
	if serverResponse != nil {
		statusCode = serverResponse.StatusCode
	}

	if statusCode == HTTPStatusUnauthorized {
		return StdCLIError{
			Err: fmt.Errorf("could not connect to StackState: invalid api-token\n" +
				"For more information: https://l.stackstate.com/cli-invalid-api-token"),
			exitCode: ConnectErrorExitCode,
		}
	} else {
		return StdCLIError{
			Err: fmt.Errorf("could not connect to StackState (%s)\n"+
				"Please check your api-url and network connection", err),
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
