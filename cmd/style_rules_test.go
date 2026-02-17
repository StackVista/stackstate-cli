package cmd

import (
	"regexp"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

var (
	startWithUpperCaseWord = regexp.MustCompile(`^[A-Z0-9]`)
	endsWithFullStop       = regexp.MustCompile(`\s*\.\s*$`)
)

func setupCmd(t *testing.T) *cobra.Command {
	cli := di.NewMockDeps(t)
	return STSCommand(&cli.Deps)
}

// --- cmd.Use ---

func TestEachNounCommandHasVerbsAndEachVerbHasNoChildren(t *testing.T) {
	root := setupCmd(t)
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

func TestUseShouldOnlyContainCommandName(t *testing.T) {
	root := setupCmd(t)
	stscobra.ForAllCmd(root, func(cmd *cobra.Command) {
		if cmd.Use != cmd.Name() {
			assert.Fail(t, cmd.CommandPath()+" Use field should only contain the command name: "+cmd.Use)
		}
	})
}

// --- cmd.Short ---

func TestShortShouldExist(t *testing.T) {
	root := setupCmd(t)
	stscobra.ForAllCmd(root, func(cmd *cobra.Command) {
		if cmd.Short == "" {
			assert.Fail(t, cmd.Use+" does not have a short")
		}
	})
}

func TestShortShouldStartUpperCase(t *testing.T) {
	root := setupCmd(t)
	stscobra.ForAllCmd(root, func(cmd *cobra.Command) {
		if !startWithUpperCaseWord.MatchString(cmd.Short) {
			assert.Fail(t, cmd.Use+" short should start with upper-case: "+cmd.Short)
		}
	})
}

func TestShortShouldNotEndWithFullStop(t *testing.T) {
	root := setupCmd(t)
	stscobra.ForAllCmd(root, func(cmd *cobra.Command) {
		if endsWithFullStop.MatchString(cmd.Short) {
			assert.Fail(t, cmd.Use+" short should not end with a fullstop: "+cmd.Short)
		}
	})
}

// --- cmd.Long ---

func TestLongShouldExist(t *testing.T) {
	root := setupCmd(t)
	stscobra.ForAllCmd(root, func(cmd *cobra.Command) {
		if cmd.Long == "" {
			assert.Fail(t, cmd.Name()+" does not have a Long description")
		}
	})
}

func TestLongShouldStartWithUpperCase(t *testing.T) {
	root := setupCmd(t)
	stscobra.ForAllCmd(root, func(cmd *cobra.Command) {
		if !startWithUpperCaseWord.MatchString(cmd.Long) {
			assert.Fail(t, cmd.Name()+" long should start with upper-case: "+cmd.Long)
		}
	})
}

func TestLongShouldEndWithAFullStop(t *testing.T) {
	root := setupCmd(t)
	stscobra.ForAllCmd(root, func(cmd *cobra.Command) {
		if !endsWithFullStop.MatchString(cmd.Long) {
			assert.Fail(t, cmd.Name()+" long should end with a full stop: "+cmd.Long)
		}
	})
}

// --- flag.Usage ---

func TestFlagUsageShouldNotEndWithFullStop(t *testing.T) {
	root := setupCmd(t)
	stscobra.ForAllFlags(root, func(cmd *cobra.Command, flag *pflag.Flag) {
		if endsWithFullStop.MatchString(flag.Usage) {
			assert.Fail(t, flag.Name+" flag of command "+cmd.Use+" should not end with a fullstop: "+cmd.Short)
		}
	})
}

func TestFlagUsageShouldStartWithUpperCase(t *testing.T) {
	root := setupCmd(t)
	stscobra.ForAllFlags(root, func(cmd *cobra.Command, flag *pflag.Flag) {
		if !startWithUpperCaseWord.MatchString(flag.Usage) {
			assert.Fail(t, flag.Name+" flag of command "+cmd.Use+" should start with uppercase: "+cmd.Short)
		}
	})
}

// --- flag.Shorthand ---

func TestFlagShortHandMustBeConsistentAmongstCommands(t *testing.T) {
	root := setupCmd(t)
	shorthands := make(map[string]string, 0)
	stscobra.ForAllFlags(root, func(cmd *cobra.Command, flag *pflag.Flag) {
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
