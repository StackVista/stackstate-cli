package cmd

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

var (
	startWithLowerCaseWordExceptions = `StackState`
	startWithLowerCaseWord           = regexp.MustCompile(`^([a-z][a-z0-9-]*|` + startWithLowerCaseWordExceptions + `)`)
	endsWithFullStop                 = regexp.MustCompile(`\s*\.\s*$`)
	startsLowerCaseExceptions        = `\.stj`
	startsLowerCase                  = regexp.MustCompile(`^([a-z]|` + startsLowerCaseExceptions + `)`)
)

func setupCmd() *cobra.Command {
	mockPrinter := printer.NewMockPrinter()
	cli := di.Deps{
		Config:  &conf.Conf{},
		Printer: &mockPrinter,
		Context: context.Background(),
		Client:  nil,
	}
	return RootCommand(&cli)
}

func forAllCmd(parent *cobra.Command, fn func(*cobra.Command)) {
	fn(parent)
	for _, child := range parent.Commands() {
		fn(child)
		forAllCmd(child, fn)
	}
}

func forAllFlags(parent *cobra.Command, fn func(*cobra.Command, *pflag.Flag)) {
	forAllCmd(parent, func(cmd *cobra.Command) {
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			fn(cmd, flag)
		})
	})
}

//--- cmd ---

func TestEachNounCommandHasVerbsAndEachVerbHasNoChildren(t *testing.T) {
	root := setupCmd()
	for _, nounCmd := range root.Commands() {
		if nounCmd.Use == "version" { // exception to the rule
			continue
		}
		if len(nounCmd.Commands()) == 0 {
			assert.Fail(t, nounCmd.Use+" has no verbs")
		}
		for _, verbCmd := range nounCmd.Commands() {
			if len(verbCmd.Commands()) > 0 {
				assert.Fail(t, verbCmd.Use+" should have no children")
			}
		}
	}
}

//--- cmd.Use USAGE ---

func TestUseStartsWithLowerCaseWord(t *testing.T) {
	root := setupCmd()
	r := regexp.MustCompile(`^[a-z][a-z0-9-]*`)
	forAllCmd(root, func(cmd *cobra.Command) {
		if !r.MatchString(cmd.Use) {
			assert.Fail(t, cmd.Use+" does not match "+r.String())
		}
	})
}

func isFlagRequired(flag *pflag.Flag) bool {
	if len(flag.Annotations[cobra.BashCompOneRequiredFlag]) > 0 {
		return true
	}
	if len(flag.Annotations["rmutex"]) > 0 {
		return true
	}
	return false
}

func TestUseShouldMentionRequiredFlags(t *testing.T) {
	root := setupCmd()
	forAllFlags(root, func(cmd *cobra.Command, flag *pflag.Flag) {
		isRequiredFlag := isFlagRequired(flag)
		if isRequiredFlag {
			var requiredFlagInUse string
			if len(flag.Shorthand) == 1 {
				requiredFlagInUse = fmt.Sprintf("-%s %s", flag.Shorthand, strings.ToUpper(flag.Name))
			} else {
				requiredFlagInUse = fmt.Sprintf("--%s %s", flag.Name, strings.ToUpper(flag.Name))
			}
			if !strings.Contains(cmd.Use, requiredFlagInUse) {
				assert.Fail(t, cmd.Use+" does not contain: "+requiredFlagInUse)
			}
		} else {
			if len(flag.Shorthand) == 1 {
				if strings.Contains(cmd.Use, fmt.Sprintf("-%s", flag.Shorthand)) {
					assert.Fail(t, cmd.Use+" should not contain optional flag: "+flag.Name)
				}
			}
			if strings.Contains(cmd.Use, fmt.Sprintf("--%s", flag.Name)) {
				assert.Fail(t, cmd.Use+" should not contain optional flag: "+flag.Name)
			}
		}
	})
}

//--- cmd.Short USAGE ---

func TestShortShouldExist(t *testing.T) {
	root := setupCmd()
	forAllCmd(root, func(cmd *cobra.Command) {
		if cmd.Short == "" {
			assert.Fail(t, cmd.Use+" does not have a short")
		}
	})
}

func TestShortShouldStartLowerCse(t *testing.T) {
	root := setupCmd()
	forAllCmd(root, func(cmd *cobra.Command) {
		if !startWithLowerCaseWord.MatchString(cmd.Short) {
			assert.Fail(t, cmd.Use+" short should start with lower-case: "+cmd.Short)
		}
	})
}

func TestShortShouldNotEndWithFullStop(t *testing.T) {
	root := setupCmd()
	forAllCmd(root, func(cmd *cobra.Command) {
		if endsWithFullStop.MatchString(cmd.Short) {
			assert.Fail(t, cmd.Use+" short should not end with a fullstop: "+cmd.Short)
		}
	})
}

//--- flag.Usage ---

func TestFlagUsageShouldNotEndWithFullStop(t *testing.T) {
	root := setupCmd()
	forAllFlags(root, func(cmd *cobra.Command, flag *pflag.Flag) {
		if endsWithFullStop.MatchString(flag.Usage) {
			assert.Fail(t, flag.Name+" flag of command "+cmd.Use+" should not end with a fullstop: "+cmd.Short)
		}
	})
}

func TestFlagUsageShouldStartWithLowerCase(t *testing.T) {
	root := setupCmd()
	forAllFlags(root, func(cmd *cobra.Command, flag *pflag.Flag) {
		if !startsLowerCase.MatchString(flag.Usage) {
			assert.Fail(t, flag.Name+" flag of command "+cmd.Use+" should start with lowercase: "+cmd.Short)
		}
	})
}

//--- flag.Shorthand ---

func TestFlagShortHandMustBeConsistentAmongstCommands(t *testing.T) {
	root := setupCmd()
	shorthands := make(map[string]string, 0)
	forAllFlags(root, func(cmd *cobra.Command, flag *pflag.Flag) {
		if len(flag.Shorthand) == 1 {
			flagName := shorthands[flag.Shorthand]
			if flagName != "" {
				if flag.Name != flagName {
					assert.Fail(t, flag.Name+" flag of command "+cmd.Use+" does not match the long hand: "+flagName)
				}
			}
			shorthands[flag.Shorthand] = flag.Name
		}
	})
}
