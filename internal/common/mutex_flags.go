package common

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func CheckRequiredMutuallyExclusiveFlags(cmd *cobra.Command, flagNames []string) CLIError {
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
	if flagUsageCount == 0 {
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
