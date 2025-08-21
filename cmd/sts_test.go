package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

const dashboardCommand = "dashboard"

// Helper function to manage STS_EXPERIMENTAL_DASHBOARD environment variable
func withDashboardEnv(value string, fn func()) {
	originalValue := os.Getenv("STS_EXPERIMENTAL_DASHBOARD")
	defer func() {
		if originalValue == "" {
			_ = os.Unsetenv("STS_EXPERIMENTAL_DASHBOARD")
		} else {
			_ = os.Setenv("STS_EXPERIMENTAL_DASHBOARD", originalValue)
		}
	}()

	if value == "" {
		_ = os.Unsetenv("STS_EXPERIMENTAL_DASHBOARD")
	} else {
		_ = os.Setenv("STS_EXPERIMENTAL_DASHBOARD", value)
	}

	fn()
}

func TestSTSCommand(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := STSCommand(&cli.Deps)

	assert.Equal(t, "sts", cmd.Use)
	assert.Equal(t, "StackState: topology-powered observability", cmd.Short)
	assert.Equal(t, "StackState: topology-powered observability.", cmd.Long)
	assert.Contains(t, cmd.UsageTemplate(), "For more information about this CLI visit https://l.stackstate.com/cli")
}

func TestSTSCommandContainsExpectedSubcommands(t *testing.T) {
	withDashboardEnv("", func() {
		cli := di.NewMockDeps(t)
		cmd := STSCommand(&cli.Deps)

		expectedCommands := []string{
			"context",
			"version",
			"script",
			"settings",
			"stackpack",
			"monitor",
			"service-token",
			"health",
			"license",
			"graph",
			"rbac",
			"topic",
			"topology-sync",
			"agent",
			"user-session",
		}

		// Verify expected commands are present
		for _, expectedCmd := range expectedCommands {
			found := false
			for _, subCmd := range cmd.Commands() {
				if subCmd.Use == expectedCmd {
					found = true
					break
				}
			}
			assert.True(t, found, "Expected command '%s' not found", expectedCmd)
		}

		// Verify dashboard command is NOT present when env var is not set
		dashboardFound := false
		for _, subCmd := range cmd.Commands() {
			if subCmd.Use == dashboardCommand {
				dashboardFound = true
				break
			}
		}
		assert.False(t, dashboardFound, "Dashboard command should not be present when STS_EXPERIMENTAL_DASHBOARD is not set")
	})
}

func TestSTSCommandDashboardExperimentalFeature(t *testing.T) {
	tests := []struct {
		name                   string
		envVarValue            string
		shouldIncludeDashboard bool
	}{
		{
			name:                   "Dashboard command not included when env var is empty",
			envVarValue:            "",
			shouldIncludeDashboard: false,
		},
		{
			name:                   "Dashboard command included when env var is set to 'true'",
			envVarValue:            "true",
			shouldIncludeDashboard: true,
		},
		{
			name:                   "Dashboard command included when env var is set to any non-empty value",
			envVarValue:            "any-value",
			shouldIncludeDashboard: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			withDashboardEnv(tt.envVarValue, func() {
				cli := di.NewMockDeps(t)
				cmd := STSCommand(&cli.Deps)

				// Check if dashboard command is present
				dashboardFound := false
				for _, subCmd := range cmd.Commands() {
					if subCmd.Use == dashboardCommand {
						dashboardFound = true
						break
					}
				}

				if tt.shouldIncludeDashboard {
					assert.True(t, dashboardFound, "Dashboard command should be present when STS_EXPERIMENTAL_DASHBOARD='%s'", tt.envVarValue)
				} else {
					assert.False(t, dashboardFound, "Dashboard command should not be present when STS_EXPERIMENTAL_DASHBOARD='%s'", tt.envVarValue)
				}
			})
		})
	}
}

func TestSTSCommandStructure(t *testing.T) {
	withDashboardEnv("1", func() {
		cli := di.NewMockDeps(t)
		cmd := STSCommand(&cli.Deps)

		// Verify the command has the expected number of subcommands (15 regular + 1 dashboard)
		assert.Len(t, cmd.Commands(), 16, "Expected 16 subcommands when dashboard is enabled")

		// Verify that dashboard command is included and properly configured
		var dashboardCmd *cobra.Command
		for _, subCmd := range cmd.Commands() {
			if subCmd.Use == dashboardCommand {
				dashboardCmd = subCmd
				break
			}
		}
		assert.NotNil(t, dashboardCmd, "Dashboard command should be present")
		assert.Equal(t, dashboardCommand, dashboardCmd.Use)
		assert.NotEmpty(t, dashboardCmd.Short)
	})
}

func TestSTSCommandWithoutDashboard(t *testing.T) {
	withDashboardEnv("", func() {
		cli := di.NewMockDeps(t)
		cmd := STSCommand(&cli.Deps)

		// Verify the command has the expected number of subcommands (15 regular, no dashboard)
		assert.Len(t, cmd.Commands(), 15, "Expected 15 subcommands when dashboard is disabled")

		// Double-check that no command has "dashboard" as its Use field
		for _, subCmd := range cmd.Commands() {
			assert.NotEqual(t, dashboardCommand, subCmd.Use, "Dashboard command should not be present")
		}
	})
}

func TestSTSCommandUsageTemplate(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := STSCommand(&cli.Deps)

	usageTemplate := cmd.UsageTemplate()
	assert.Contains(t, usageTemplate, "For more information about this CLI visit https://l.stackstate.com/cli")

	// Verify it contains the standard usage template plus our addition
	assert.Contains(t, usageTemplate, "Usage:")
	assert.Contains(t, usageTemplate, "Available Commands:")
	assert.Contains(t, usageTemplate, "Flags:")
}
