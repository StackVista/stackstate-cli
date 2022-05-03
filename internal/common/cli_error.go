package common

import "net/http"

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
