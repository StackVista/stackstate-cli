package context

import (
	"testing"

	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCurrentContext(t *testing.T) {
	cli := di.NewMockDeps(t)
	setupConfig(t, &cli)
	cmd := DeleteCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "foo")
	assert.Errorf(t, err, "cannot delete the current context (foo)")
}

func TestDeleteNonExistingContext(t *testing.T) {
	cli := di.NewMockDeps(t)
	setupConfig(t, &cli)
	cmd := DeleteCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "x")
	assert.Errorf(t, err, "context with name 'x' not found")
}

func TestDeleteContext(t *testing.T) {
	cli := di.NewMockDeps(t)
	setupConfig(t, &cli)
	cmd := DeleteCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "bar")
	assert.NoError(t, err)

	cfg, err := config.ReadConfig(cli.ConfigPath)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(cfg.Contexts))
}
