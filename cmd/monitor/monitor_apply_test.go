package monitor

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
)

//nolint:unused
func setMonitorApplyCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := MonitorApplyCommand(&cli.Deps)
	return &cli, cmd
}

// PLEASE WRITE TESTS!
