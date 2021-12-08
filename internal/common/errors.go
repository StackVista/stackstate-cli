package common

import (
	"net/http"
)

type CLIError interface {
	error
	GetExitCode() ExitCode
	Error() string
}

func NewCLIError(err error) CLIError {
	return StdCLIError{Err: err}
}

type StdCLIError struct {
	Err error
}

func (p StdCLIError) Error() string {
	return p.Err.Error()
}

func (p StdCLIError) GetExitCode() ExitCode {
	return CommandExecutionExitCode
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

func NewResponseError(err error, resp *http.Response) ResponseError {
	return ResponseError{
		Err:  err,
		Resp: resp,
	}
}
