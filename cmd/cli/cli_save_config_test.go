package cli

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func setupSaveConfigCmd(t *testing.T) (*di.MockDeps, *cobra.Command, string, func()) {
	cli := di.NewMockDeps()
	cmd := CliSaveConfigCommand(&cli.Deps)

	oldConfHome := os.Getenv("XDG_CONFIG_HOME")
	tmpConfDir := t.TempDir()
	os.Setenv("XDG_CONFIG_HOME", tmpConfDir)
	expectedFile := tmpConfDir + "/stackstate-cli/config.yaml"

	return &cli, cmd, expectedFile, func() {
		os.Setenv("XDG_CONFIG_HOME", oldConfHome)
	}
}

func TestSaveConfig(t *testing.T) {
	cli, cmd, expectedFile, cleanup := setupSaveConfigCmd(t)
	defer cleanup()

	di.ExecuteCommandWithContextUnsafe(
		&cli.Deps,
		cmd,
		"--api-url",
		"https://test.stackstate.io/api",
		"--api-token",
		"blaat",
	)

	fileExists, _ := util.DoesFileExist(expectedFile)
	assert.Equal(t, "Connection verified to https://test.stackstate.io/api (StackState version: 0.0.0+-)", (*cli.MockPrinter.SuccessCalls)[0])
	assert.Equal(t, "Config saved to: "+expectedFile, (*cli.MockPrinter.SuccessCalls)[1])
	assert.True(t, fileExists)
	assert.Equal(t, "blaat", cli.Config.ApiToken)
	assert.Equal(t, "https://test.stackstate.io/api", cli.Config.ApiURL)
}

func TestSaveConfigToJson(t *testing.T) {
	cli, cmd, expectedFile, cleanup := setupSaveConfigCmd(t)
	defer cleanup()

	di.ExecuteCommandWithContextUnsafe(
		&cli.Deps,
		cmd,
		"--api-url",
		"https://test.stackstate.io/api",
		"--api-token",
		"blaat",
		"--json",
	)
	assert.Equal(t,
		[]map[string]interface{}{{
			"connection-tested": true,
			"config-file":       expectedFile,
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestSaveConfigSkipValidate(t *testing.T) {
	cli, cmd, expectedFile, cleanup := setupSaveConfigCmd(t)
	defer cleanup()

	cli.MockClient.ConnectError = common.NewResponseError(fmt.Errorf("should not have tried to connect, because --skip-validate was used"), nil)
	di.ExecuteCommandWithContextUnsafe(
		&cli.Deps,
		cmd,
		"--api-url",
		"https://test.stackstate.io/api",
		"--api-token",
		"blaat",
		"--skip-validate",
	)

	fileExists, _ := util.DoesFileExist(expectedFile)
	assert.Equal(t, "Config saved to: "+expectedFile, (*cli.MockPrinter.SuccessCalls)[0])
	assert.True(t, fileExists)
	assert.Equal(t, "https://test.stackstate.io/api", cli.Config.ApiURL)
	assert.Equal(t, "blaat", cli.Config.ApiToken)
}

func TestSaveConfigShouldNotSaveWhenFailedConnection(t *testing.T) {
	cli, cmd, expectedFile, cleanup := setupSaveConfigCmd(t)
	defer cleanup()

	cli.MockClient.ConnectError = common.NewResponseError(fmt.Errorf("failed connection"), nil)
	_, err := di.ExecuteCommandWithContext(
		&cli.Deps,
		cmd,
		"--api-url",
		"https://test.stackstate.io/api",
		"--api-token",
		"blaat",
	)

	assert.IsType(t, common.StdCLIError{}, err)
	fileExists, err := util.DoesFileExist(expectedFile)
	assert.False(t, fileExists)
	assert.Nil(t, err)
}
