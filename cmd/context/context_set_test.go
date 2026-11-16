package context

import (
	"testing"

	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestSetContext(t *testing.T) {
	cli := di.NewMockDeps(t)
	setupConfig(t, &cli)
	cmd := SetCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, nameFlagArg, testBar)
	assert.NoError(t, err)

	cfg, err := config.ReadConfig(cli.ConfigPath)
	assert.NoError(t, err)
	assert.Equal(t, testBar, cfg.CurrentContext)
}

func TestSetUnknownContext(t *testing.T) {
	cli := di.NewMockDeps(t)
	setupConfig(t, &cli)
	cmd := SetCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, nameFlagArg, "x")
	assert.Errorf(t, err, "context with name 'x' not found")

	cfg, err := config.ReadConfig(cli.ConfigPath)
	assert.NoError(t, err)
	assert.Equal(t, "foo", cfg.CurrentContext)
}
