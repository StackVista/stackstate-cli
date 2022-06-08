package context

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func setupSaveCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cli.StsConfig.CurrentContext = "foo"
	cli.StsConfig.Contexts = []*config.NamedContext{
		{Name: "foo", Context: newContext("http://foo.com", "apiToken", "", "/api")},
		{Name: "bar", Context: newContext("http://bar.com", "apiToken", "", "/api")},
	}
	cmd := SaveCommand(&cli.Deps)

	return &cli, cmd
}

func TestSaveNewContext(t *testing.T) {
	cli, cmd := setupSaveCmd(t)
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "baz", "--url", "http://baz.com", "--api-token", "my-token")
	assert.NoError(t, err)

	assert.Equal(t, "baz", cli.StsConfig.CurrentContext)
	assert.Len(t, cli.StsConfig.Contexts, 3)

	cfg, err := config.ReadConfig(cli.ConfigPath)
	assert.NoError(t, err)
	assert.Equal(t, "baz", cfg.CurrentContext)
	assert.Len(t, cfg.Contexts, 3)
	curr, err := cfg.GetContext(cfg.CurrentContext)
	assert.NoError(t, err)
	assert.Equal(t, "my-token", curr.Context.APIToken)
	assert.Equal(t, "http://baz.com", curr.Context.URL)
	assert.Equal(t, "/api", curr.Context.APIPath)
	assert.Empty(t, curr.Context.ServiceToken)
}

func TestSaveExistingContext(t *testing.T) {
	cli, cmd := setupSaveCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "bar", "--url", "http://bar.com", "--service-token", "my-token")
	assert.NoError(t, err)

	assert.Equal(t, "bar", cli.StsConfig.CurrentContext)
	assert.Len(t, cli.StsConfig.Contexts, 2)

	cfg, err := config.ReadConfig(cli.ConfigPath)
	assert.NoError(t, err)
	assert.Equal(t, "bar", cfg.CurrentContext)
	assert.Len(t, cfg.Contexts, 2)
	curr, err := cfg.GetContext(cfg.CurrentContext)
	assert.NoError(t, err)
	assert.Equal(t, "my-token", curr.Context.ServiceToken)
	assert.Equal(t, "http://bar.com", curr.Context.URL)
	assert.Equal(t, "/api", curr.Context.APIPath)
	assert.Empty(t, curr.Context.APIToken)
}

func TestNoSaveOnMissingTokens(t *testing.T) {
	cli, cmd := setupSaveCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "bar", "--url", "http://my-bar.com")
	assert.Errorf(t, err, "missing required argument: --api-token")

	assert.Equal(t, "foo", cli.StsConfig.CurrentContext)
	assert.Len(t, cli.StsConfig.Contexts, 2)

	// Should not have written config file
	assert.NoFileExists(t, cli.ConfigPath)
}
