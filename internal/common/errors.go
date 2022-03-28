package common

import (
	"fmt"
	"net/http"
)

type CLIError interface {
	error
	GetExitCode() ExitCode
	Error() string
	ShowUsage() bool
}

func NewCLIError(err error) CLIError {
	return StdCLIError{Err: err}
}

func NewConnectError(err error) CLIError {
	var statusCode int
	switch v := err.(type) {
	case ResponseError:
		// is access error?
		if v.Resp != nil {
			statusCode = v.Resp.StatusCode
		}
	}

	if statusCode == 401 {
		return StdCLIError{
			Err: fmt.Errorf("could not connect to StackState: invalid api-token\n" +
				"For more information: https://l.stackstate.com/cli-invalid-api-token"),
			exitCode: ConnectErrorExitCode,
		}
	} else {
		return StdCLIError{
			Err: fmt.Errorf("could not connect to StackState: invalid api-url (%s)\n"+
				"Get the URL of your StackState instance and add \"/api\" at the end\n"+
				"Be aware that this CLI can only connect with StackState version 5 or greater", err),
			exitCode: ConnectErrorExitCode,
		}
	}
}

func NewCLIArgParseError(err error) CLIError {
	return StdCLIError{
		Err:       err,
		showUsage: true,
		exitCode:  CommandExecutionExitCode,
	}
}

type StdCLIError struct {
	Err       error
	showUsage bool
	exitCode  int
}

func (p StdCLIError) Error() string {
	return p.Err.Error()
}

func (p StdCLIError) GetExitCode() ExitCode {
	return p.exitCode
}

func (p StdCLIError) ShowUsage() bool {
	return p.showUsage
}

// --------------
// ResponseError is for errors that involve the server
// --------------

func NewResponseError(err error, resp *http.Response) ResponseError {
	return ResponseError{
		Err:  err,
		Resp: resp,
	}
}

type ResponseError struct {
	Err  error
	Resp *http.Response
}

func (p ResponseError) Error() string {
	return p.Err.Error()
}

func (p ResponseError) GetExitCode() ExitCode {
	return ReponseErrorExitCode
}

func (p ResponseError) ShowUsage() bool {
	return false
}
