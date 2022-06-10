package context

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func TestContextListJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	setupConfig(t, &cli)
	cmd := ListCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--output", "json")

	assert.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)

	cfg, err := config.ReadConfig(cli.ConfigPath)
	assert.NoError(t, err)

	assert.Equal(t, cfg.Contexts, (*cli.MockPrinter.PrintJsonCalls)[0]["contexts"])
}
