package common

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func MarkMutexFlags(cmd *cobra.Command, flagNames []string, mutexName string, isRequired bool) {
	for _, flagName := range flagNames {
		var annotationName string
		if isRequired {
			annotationName = "rmutex"
		} else {
			annotationName = "mutex"
		}
		cmd.Flags().SetAnnotation(flagName, annotationName, []string{mutexName}) //nolint:errcheck
	}
}

func CheckMutuallyExclusiveFlags(cmd *cobra.Command, flagNames []string, isRequired bool) CLIError {
	flagUsageCount := 0
	flagsUsed := make([]string, 0)
	for _, flagName := range flagNames {
		f := cmd.Flags().Lookup(flagName)
		if f.Changed {
			flagUsageCount++
			flagsUsed = append(flagsUsed, flagName)
		}
	}

	if flagUsageCount > 1 {
		return NewMutuallyExclusiveFlagsMultipleError(flagNames, flagsUsed)
	}
	if isRequired && flagUsageCount == 0 {
		return NewMutuallyExclusiveFlagsRequiredError(flagNames)
	}

	return nil
}

func NewMutuallyExclusiveFlagsRequiredError(flagNames []string) CLIError {
	return StdCLIError{
		Err:       fmt.Errorf("one of the required flags {%s} not set", strings.Join(flagNames, " | ")),
		showUsage: true,
		exitCode:  CommandFailedRequirementExitCode,
	}
}

func NewMutuallyExclusiveFlagsMultipleError(flagNames []string, flagsUsed []string) CLIError {
	return StdCLIError{
		Err:       fmt.Errorf("only one of the required flags {%s} must be entered, but got: %s", strings.Join(flagNames, " | "), strings.Join(flagsUsed, " & ")),
		showUsage: true,
		exitCode:  CommandFailedRequirementExitCode,
	}
}
