package monitor

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
)

//nolint:unused
func setMonitorRunCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := MonitorRunCommand(&cli.Deps)
	return &cli, cmd
}

// PLEASE WRITE TESTS!
