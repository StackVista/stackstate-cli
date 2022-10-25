package common

type ExitCode = int

const (
	OkExitCode ExitCode = iota
	ConfigErrorExitCode
	ReponseErrorExitCode
	CommandFailedRequirementExitCode
	ConnectErrorExitCode
	ReadFileErrorExitCode
	WriteFileErrorExitCode
	UnknownSubCommmandExitCode
	NotFoundExitCode
	ExecutionErrorCode
	APIVersionErrorCode
)
