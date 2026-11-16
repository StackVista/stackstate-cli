package context

import (
	"testing"

	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

const (
	testBar     = "bar"
	nameFlagArg = "--name"
)

func TestDeleteCurrentContext(t *testing.T) {
	cli := di.NewMockDeps(t)
	setupConfig(t, &cli)
	cmd := DeleteCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, nameFlagArg, "foo")
	assert.Errorf(t, err, "cannot delete the current context (foo)")
}

func TestDeleteNonExistingContext(t *testing.T) {
	cli := di.NewMockDeps(t)
	setupConfig(t, &cli)
	cmd := DeleteCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, nameFlagArg, "x")
	assert.Errorf(t, err, "context with name 'x' not found")
}

func TestDeleteContext(t *testing.T) {
	cli := di.NewMockDeps(t)
	setupConfig(t, &cli)
	cmd := DeleteCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, nameFlagArg, testBar)
	assert.NoError(t, err)

	cfg, err := config.ReadConfig(cli.ConfigPath)
	assert.NoError(t, err)

	assert.Equal(t, 5, len(cfg.Contexts))
}
