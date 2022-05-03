package cli

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func setupDescribeConfigCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := CliDescribeConfigCommand(&cli.Deps)
	return &cli, cmd
}

func TestDescribeConfig(t *testing.T) {
	cli, cmd := setupDescribeConfigCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--url", "1", "--api-token", "bla")

	assert.Equal(t, []interface{}{map[string]interface{}{"api-path": "/bla", "api-token": "bla", "url": "1"}}, *cli.MockPrinter.PrintStructCalls)
}

func TestDescribeConfigJson(t *testing.T) {
	cli, cmd := setupDescribeConfigCmd(t)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--url", "1", "--api-token", "bla", "--json")

	assert.Equal(t, []map[string]interface{}{{"api-path": "/bla", "api-token": "bla", "url": "1"}}, *cli.MockPrinter.PrintJsonCalls)
}
