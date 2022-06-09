package context

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func TestSetContext(t *testing.T) {
	cli := di.NewMockDeps(t)
	cli.StsConfig.CurrentContext = "foo"
	cli.StsConfig.Contexts = []*config.NamedContext{
		{Name: "foo", Context: newContext("http://foo.com", "apiToken", "", "/api")},
		{Name: "bar", Context: newContext("http://bar.com", "apiToken", "", "/api")},
	}
	cmd := SetCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "bar")
	assert.NoError(t, err)

	assert.Equal(t, "bar", cli.StsConfig.CurrentContext)

	cfg, err := config.ReadConfig(cli.ConfigPath)
	assert.NoError(t, err)
	assert.Equal(t, "bar", cfg.CurrentContext)
}

func TestSetUnknownContext(t *testing.T) {
	cli := di.NewMockDeps(t)
	cli.StsConfig.CurrentContext = "foo"
	cli.StsConfig.Contexts = []*config.NamedContext{
		{Name: "foo", Context: newContext("http://foo.com", "apiToken", "", "/api")},
		{Name: "bar", Context: newContext("http://bar.com", "apiToken", "", "/api")},
	}
	cmd := SetCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "x")
	assert.Errorf(t, err, "context with name 'x' not found")

	assert.Equal(t, "foo", cli.StsConfig.CurrentContext)
}
