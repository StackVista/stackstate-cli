package cobra

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stackvista/stackstate-cli/internal/util"
)

type ValidationErrors []error

func (v *ValidationErrors) Add(err error) {
	*v = append(*v, err)
}

func (v *ValidationErrors) HasErrors() bool {
	return len(*v) > 0
}

func (v *ValidationErrors) Error() string {
	if len(*v) == 1 {
		return (*v)[0].Error()
	}

	err := "Validation failed:\n"
	for i, e := range *v {
		err += fmt.Sprintf("\tError %d: %s\n", i, e.Error())
	}

	return fmt.Sprintf("%v", err)
}

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

func GetAllMutexNames(cmd *cobra.Command) map[bool]util.StringSet {
	mutexNames := map[bool]util.StringSet{
		true:  {},
		false: {},
	}
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if len(flag.Annotations["rmutex"]) > 0 {
			mutexNames[true].Add(flag.Annotations["rmutex"][0])
		}
		if len(flag.Annotations["mutex"]) > 0 {
			mutexNames[false].Add(flag.Annotations["rmutex"][0])
		}
	})
	return mutexNames
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

func CheckMutuallyExclusiveFlags(cmd *cobra.Command, flagNames []string, isRequired bool) error {
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

func ValidateMutexFlags(cmd *cobra.Command) error {
	errors := &ValidationErrors{}
	mutexNames := GetAllMutexNames(cmd)
	for required, mutexes := range mutexNames {
		for mutex := range mutexes {
			err := CheckMutuallyExclusiveFlags(cmd, GetAllFlagNamesOfMutex(cmd, mutex), required)
			if err != nil {
				errors.Add(err)
			}
		}
	}

	if errors.HasErrors() {
		return errors
	}

	return nil
}

func NewMutuallyExclusiveFlagsRequiredError(flagNames []string) error {
	sort.Strings(flagNames)
	return fmt.Errorf("one of the required flags {%s} not set", strings.Join(flagNames, " | "))
}

func NewMutuallyExclusiveFlagsMultipleError(flagNames []string, flagsUsed []string) error {
	sort.Strings(flagNames)
	sort.Strings(flagsUsed)
	return fmt.Errorf("only one of the required flags {%s} must be entered, but got: %s", strings.Join(flagNames, " | "), strings.Join(flagsUsed, " & "))
}
