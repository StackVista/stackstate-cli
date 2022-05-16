package cliconfig

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func setupDescribeCmd() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := DescribeCommand(&cli.Deps)
	return &cli, cmd
}

func TestDescribe(t *testing.T) {
	cli, cmd := setupDescribeCmd()

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--url", "https://test", "--api-token", "bla")

	assert.Equal(t, 1, len(*cli.MockPrinter.PrintStructCalls))
	config := (*cli.MockPrinter.PrintStructCalls)[0].(map[string]interface{})
	assert.Equal(t, "bla", config["api-token"])
	assert.Equal(t, "https://test", config["url"])
}

func TestDescribeJson(t *testing.T) {
	cli, cmd := setupDescribeCmd()

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--url", "https://test", "--api-token", "bla", "--json")

	config := (*cli.MockPrinter.PrintJsonCalls)[0]
	assert.Equal(t, "bla", config["api-token"])
	assert.Equal(t, "https://test", config["url"])
}
