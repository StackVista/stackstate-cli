package cli

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func TestSaveConfig(t *testing.T) {
	doTestSaveConfig(t, false)
}

func TestSaveConfigSkipValidate(t *testing.T) {
	doTestSaveConfig(t, true)
}

func TestSaveConfigShouldNotSaveWhenFailedConnection(t *testing.T) {
	cli := di.NewMockDeps()
	cmd := CliSaveConfigCommand(&cli.Deps)
	common.AddPersistentFlags(cmd)

	oldConfHome := os.Getenv("XDG_CONFIG_HOME")
	defer os.Setenv("XDG_CONFIG_HOME", oldConfHome)
	tmpConfDir := t.TempDir()
	os.Setenv("XDG_CONFIG_HOME", tmpConfDir)

	cli.MockClient.ConnectError = common.NewResponseError(fmt.Errorf("failed connection"), nil)
	_, err := util.ExecuteCommandWithContext(
		cli.Context,
		cmd,
		"--api-url",
		"https://test.stackstate.io/api",
		"--api-token",
		"blaat",
	)

	// should have had a connection error
	assert.IsType(t, common.StdCLIError{}, err)

	// should not have written config, so reading fails
	_, err = conf.ReadConf(CliSaveConfigCommand(&cli.Deps))
	assert.IsType(t, conf.ReadConfError{}, err)
}

func doTestSaveConfig(t *testing.T, skipValidate bool) {
	cli := di.NewMockDeps()
	cmd := CliSaveConfigCommand(&cli.Deps)
	common.AddPersistentFlags(cmd)

	oldConfHome := os.Getenv("XDG_CONFIG_HOME")
	defer os.Setenv("XDG_CONFIG_HOME", oldConfHome)
	tmpConfDir := t.TempDir()
	os.Setenv("XDG_CONFIG_HOME", tmpConfDir)

	if skipValidate {
		cli.MockClient.ConnectError = common.NewResponseError(fmt.Errorf("should not have tried to connect, becasue --skip-validate was used"), nil)
		util.ExecuteCommandWithContextUnsafe(
			cli.Context,
			cmd,
			"--api-url",
			"https://test.stackstate.io/api",
			"--api-token",
			"blaat",
			"--skip-validate",
		)
	} else {
		util.ExecuteCommandWithContextUnsafe(
			cli.Context,
			cmd,
			"--api-url",
			"https://test.stackstate.io/api",
			"--api-token",
			"blaat",
		)
	}

	cfg, err := conf.ReadConf(CliSaveConfigCommand(&cli.Deps))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "https://test.stackstate.io/api", cfg.ApiURL)
	assert.Equal(t, "blaat", cfg.ApiToken)
}
