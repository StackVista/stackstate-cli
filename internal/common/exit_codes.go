package common

type ExitCode = int

const (
	ConfigErrorExitCode ExitCode = iota + 1
	ReponseErrorExitCode
	CommandFailedRequirementExitCode
	ConnectErrorExitCode
)
