package cmd

import (
	"testing"

	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestStackPackCommand_EnabledCommands(t *testing.T) {
	// Ensure that other commands are always present regardless of environment variable
	cli := di.NewMockDeps(t)
	cmd := StackPackCommand(&cli.Deps)

	expectedCommands := []string{
		"upload",
		"list",
		"list-instances",
		"install",
		"list-parameters",
		"uninstall",
		"upgrade",
		"confirm-manual-steps",
		"describe",
		"scaffold",
		"package",
		"test-deploy",
	}

	for _, cmdName := range expectedCommands {
		t.Run("command_"+cmdName+"_enabled", func(t *testing.T) {
			foundCmd, _, err := cmd.Find([]string{cmdName})
			assert.NoError(t, err, "Command %s should always be present", cmdName)
			assert.NotNil(t, foundCmd, "Command %s should always be present", cmdName)
		})
	}
}
