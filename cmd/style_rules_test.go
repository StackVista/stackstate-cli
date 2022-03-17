package cmd

import (
	"context"
	"regexp"
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

func TestRuleEachNounCommandHasVerbs(t *testing.T) {
	root := setupCmd()
	for _, nounCmd := range root.Commands() {
		if nounCmd.Use == "version" {
			continue
		}
		if len(nounCmd.Commands()) == 0 {
			assert.Fail(t, nounCmd.Use+" has no verbs")
		}
	}
}

func TestUseStartsWithLowerCaseWord(t *testing.T) {
	root := setupCmd()
	r := regexp.MustCompile(`^[a-z][a-z0-9-]*`)
	forAllCmd(root, func(cmd *cobra.Command) {
		if !r.MatchString(cmd.Use) {
			assert.Fail(t, cmd.Use+" does not match "+r.String())
		}
	})
}

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
