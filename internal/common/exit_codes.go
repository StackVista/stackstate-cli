package common

type ExitCode = int

const (
	OkExitCode ExitCode = iota
	ConfigErrorExitCode
	ReponseErrorExitCode
	CommandFailedRequirementExitCode
	ConnectErrorExitCode
)
