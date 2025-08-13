package cmd

import (
	"os"
	"testing"

	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStackPackCommand_FeatureGating(t *testing.T) {
	tests := []struct {
		name                  string
		envVarValue           string
		expectScaffoldCommand bool
		description           string
	}{
		{
			name:                  "scaffold command hidden by default",
			envVarValue:           "",
			expectScaffoldCommand: false,
			description:           "When environment variable is not set, scaffold command should be hidden",
		},
		{
			name:                  "scaffold command visible when env var is set to 1",
			envVarValue:           "1",
			expectScaffoldCommand: true,
			description:           "When environment variable is set to '1', scaffold command should be visible",
		},
		{
			name:                  "scaffold command visible when env var is set to any value",
			envVarValue:           "true",
			expectScaffoldCommand: true,
			description:           "When environment variable is set to any non-empty value, scaffold command should be visible",
		},
		{
			name:                  "scaffold command visible when env var is set to enabled",
			envVarValue:           "enabled",
			expectScaffoldCommand: true,
			description:           "When environment variable is set to 'enabled', scaffold command should be visible",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Store original environment value to restore later
			originalValue := os.Getenv("STS_EXPERIMENTAL_STACKPACK_SCAFFOLD")
			defer func() {
				if originalValue == "" {
					os.Unsetenv("STS_EXPERIMENTAL_STACKPACK_SCAFFOLD")
				} else {
					os.Setenv("STS_EXPERIMENTAL_STACKPACK_SCAFFOLD", originalValue)
				}
			}()

			// Set the environment variable for this test
			if tt.envVarValue == "" {
				os.Unsetenv("STS_EXPERIMENTAL_STACKPACK_SCAFFOLD")
			} else {
				err := os.Setenv("STS_EXPERIMENTAL_STACKPACK_SCAFFOLD", tt.envVarValue)
				require.NoError(t, err)
			}

			// Create the command
			cli := di.NewMockDeps(t)
			cmd := StackPackCommand(&cli.Deps)

			// Check if scaffold command exists
			scaffoldCmd, _, err := cmd.Find([]string{"scaffold"})

			if tt.expectScaffoldCommand {
				assert.NoError(t, err, tt.description)
				assert.NotNil(t, scaffoldCmd, tt.description)
				assert.Equal(t, "scaffold", scaffoldCmd.Use, tt.description)
			} else {
				assert.Error(t, err, tt.description)
				assert.Contains(t, err.Error(), "unknown command", tt.description)
			}
		})
	}
}

func TestStackPackCommand_AlwaysPresentCommands(t *testing.T) {
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
	}

	for _, cmdName := range expectedCommands {
		t.Run("command_"+cmdName+"_always_present", func(t *testing.T) {
			foundCmd, _, err := cmd.Find([]string{cmdName})
			assert.NoError(t, err, "Command %s should always be present", cmdName)
			assert.NotNil(t, foundCmd, "Command %s should always be present", cmdName)
		})
	}
}
