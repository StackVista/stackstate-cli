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
		name                       string
		envVarValue                string
		expectExperimentalCommands bool
		description                string
	}{
		{
			name:                       "experimental commands hidden by default",
			envVarValue:                "",
			expectExperimentalCommands: false,
			description:                "When environment variable is not set, experimental commands should be hidden",
		},
		{
			name:                       "experimental commands visible when env var is set to 1",
			envVarValue:                "1",
			expectExperimentalCommands: true,
			description:                "When environment variable is set to '1', experimental commands should be visible",
		},
		{
			name:                       "experimental commands visible when env var is set to any value",
			envVarValue:                "true",
			expectExperimentalCommands: true,
			description:                "When environment variable is set to any non-empty value, experimental commands should be visible",
		},
		{
			name:                       "experimental commands visible when env var is set to enabled",
			envVarValue:                "enabled",
			expectExperimentalCommands: true,
			description:                "When environment variable is set to 'enabled', experimental commands should be visible",
		},
	}

	experimentalCommands := []string{"scaffold", "package", "test"}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Store original environment value to restore later
			originalValue := os.Getenv(experimentalStackpackEnvVar)
			defer func() {
				if originalValue == "" {
					os.Unsetenv(experimentalStackpackEnvVar)
				} else {
					os.Setenv(experimentalStackpackEnvVar, originalValue)
				}
			}()

			// Set the environment variable for this test
			if tt.envVarValue == "" {
				os.Unsetenv(experimentalStackpackEnvVar)
			} else {
				err := os.Setenv(experimentalStackpackEnvVar, tt.envVarValue)
				require.NoError(t, err)
			}

			// Create the command
			cli := di.NewMockDeps(t)
			cmd := StackPackCommand(&cli.Deps)

			// Check each experimental command
			for _, cmdName := range experimentalCommands {
				foundCmd, _, err := cmd.Find([]string{cmdName})

				if tt.expectExperimentalCommands {
					assert.NoError(t, err, tt.description+" (command: %s)", cmdName)
					assert.NotNil(t, foundCmd, tt.description+" (command: %s)", cmdName)
					assert.Equal(t, cmdName, foundCmd.Use, tt.description+" (command: %s)", cmdName)
				} else {
					assert.Error(t, err, tt.description+" (command: %s)", cmdName)
					assert.Contains(t, err.Error(), "unknown command", tt.description+" (command: %s)", cmdName)
				}
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
