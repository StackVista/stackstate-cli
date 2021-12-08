package cli

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/cmd/persistent_flags"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func TestWriteConfig(t *testing.T) {
	mockPrinter := printer.NewMockPrinter()
	cli := &di.Deps{
		Printer: &mockPrinter,
	}
	cmd := CliSaveConfigCommand(cli)
	persistent_flags.AddPersistentFlags(cmd)

	oldConfHome := os.Getenv("XDG_CONFIG_HOME")
	defer os.Setenv("XDG_CONFIG_HOME", oldConfHome)
	tmpConfDir := t.TempDir()
	os.Setenv("XDG_CONFIG_HOME", tmpConfDir)

	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "--api-url", "https://test.stackstate.io/api", "--api-token", "blaat")
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := conf.ReadConf(CliSaveConfigCommand(cli))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "https://test.stackstate.io/api", cfg.ApiUrl)
	assert.Equal(t, "blaat", cfg.ApiToken)
}
