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

func TestWriteConfig(t *testing.T) {
	cli := di.NewMockDeps()
	cmd := CliSaveConfigCommand(&cli.Deps)
	cli.MockClient.ConnectError = fmt.Errorf("should not have tried to connect")
	common.AddPersistentFlags(cmd)

	oldConfHome := os.Getenv("XDG_CONFIG_HOME")
	defer os.Setenv("XDG_CONFIG_HOME", oldConfHome)
	tmpConfDir := t.TempDir()
	os.Setenv("XDG_CONFIG_HOME", tmpConfDir)

	util.ExecuteCommandWithContextUnsafe(cli.Context, cmd, "--api-url", "https://test.stackstate.io/api", "--api-token", "blaat")

	cfg, err := conf.ReadConf(CliSaveConfigCommand(&cli.Deps))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "https://test.stackstate.io/api", cfg.ApiUrl)
	assert.Equal(t, "blaat", cfg.ApiToken)
}
