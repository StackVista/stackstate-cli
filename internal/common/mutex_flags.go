package common

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
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

func GetAllMutexNames(cmd *cobra.Command, isRequired bool) []string {
	mutexNames := make([]string, 0)
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if isRequired && len(flag.Annotations["rmutex"]) > 0 {
			mutexNames = append(mutexNames, flag.Annotations["rmutex"][0])
		}
		if !isRequired && len(flag.Annotations["mutex"]) > 0 {
			mutexNames = append(mutexNames, flag.Annotations["mutex"][0])
		}
	})
	return util.UniqueStrings(mutexNames)
}

func GetAllFlagsOfMutex(cmd *cobra.Command, mutexName string) []*pflag.Flag {
	flags := make([]*pflag.Flag, 0)
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if len(flag.Annotations["rmutex"]) > 0 && flag.Annotations["rmutex"][0] == mutexName {
			flags = append(flags, flag)
		}
		if len(flag.Annotations["mutex"]) > 0 && flag.Annotations["mutex"][0] == mutexName {
			flags = append(flags, flag)
		}
	})
	return flags
}

func GetAllFlagNamesOfMutex(cmd *cobra.Command, mutexName string) []string {
	mutexFlagNames := make([]string, 0)
	for _, mutexFlag := range GetAllFlagsOfMutex(cmd, mutexName) {
		mutexFlagNames = append(mutexFlagNames, mutexFlag.Name)
	}
	return mutexFlagNames
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
