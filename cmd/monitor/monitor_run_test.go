package monitor

import (
	"testing"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

//nolint:deadcode,unused
func setMonitorRunCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := MonitorRunCommand(&cli.Deps)
	return &cli, cmd
}

// PLEASE WRITE TESTS!
