package cmd

import (
	"testing"

	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestSTSCommand(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := STSCommand(&cli.Deps)

	assert.Equal(t, "sts", cmd.Use)
	assert.Equal(t, "SUSE Observability: topology-powered observability", cmd.Short)
	assert.Equal(t, "SUSE Observability: topology-powered observability.", cmd.Long)
	assert.Contains(t, cmd.UsageTemplate(), "For more information about this CLI visit https://documentation.suse.com/cloudnative/suse-observability/latest/en/setup/cli/cli-sts.html\n")
}

func TestSTSCommandContainsExpectedSubcommands(t *testing.T) {
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
		"dashboard",
		"topology",
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
}

func TestSTSCommandStructure(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := STSCommand(&cli.Deps)

	assert.Len(t, cmd.Commands(), 17, "Expected 17 subcommands")
}

func TestSTSCommandUsageTemplate(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := STSCommand(&cli.Deps)

	usageTemplate := cmd.UsageTemplate()
	assert.Contains(t, usageTemplate, "For more information about this CLI visit https://documentation.suse.com/cloudnative/suse-observability/latest/en/setup/cli/cli-sts.html\n")

	// Verify it contains the standard usage template plus our addition
	assert.Contains(t, usageTemplate, "Usage:")
	assert.Contains(t, usageTemplate, "Available Commands:")
	assert.Contains(t, usageTemplate, "Flags:")
}
