package context

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func setupSaveCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := SaveCommand(&cli.Deps)

	return &cli, cmd
}

func TestSaveContext(t *testing.T) { //nolint:funlen
	tests := []struct {
		name                 string
		args                 []string
		expectedContext      config.NamedContext
		totalContextInConfig int
		wantErr              bool
		errorMessage         string
	}{
		{
			name: "new context",
			args: []string{"--name", "baz", "--url", "http://baz.com", "--api-token", "my-token"},
			expectedContext: config.NamedContext{
				Name: "baz",
				Context: &config.StsContext{
					URL:      "http://baz.com",
					APIToken: "my-token",
					APIPath:  "/api",
				},
			},
			totalContextInConfig: 7,
			wantErr:              false,
		},
		{
			name: "existing context",
			args: []string{"--name", "bar", "--url", "http://bar.com", "--service-token", "my-token"},
			expectedContext: config.NamedContext{
				Name: "bar",
				Context: &config.StsContext{
					URL:          "http://bar.com",
					ServiceToken: "my-token",
					APIPath:      "/api",
				},
			},
			totalContextInConfig: 6,
			wantErr:              false,
		},
		{
			name: "existing context ca-cert is set with ca-cert-path",
			args: []string{"--name", "bar", "--url", "http://bar.com", "--service-token", "my-token", "--ca-cert-path", "testdata/selfSignedCert.crt"},
			expectedContext: config.NamedContext{
				Name: "bar",
				Context: &config.StsContext{
					URL:              "http://bar.com",
					ServiceToken:     "my-token",
					APIPath:          "/api",
					CaCertBase64Data: selfSignedBase64Cert,
				},
			},
			totalContextInConfig: 6,
			wantErr:              false,
		},
		{
			name: "new context ca-cert is set with ca-cert-path",
			args: []string{"--name", "cacertdata", "--url", "http://bar.com", "--service-token", "my-token", "--ca-cert-base64-data", privateCaBase64Cert},
			expectedContext: config.NamedContext{
				Name: "cacertdata",
				Context: &config.StsContext{
					URL:              "http://bar.com",
					ServiceToken:     "my-token",
					APIPath:          "/api",
					CaCertBase64Data: privateCaBase64Cert,
				},
			},
			totalContextInConfig: 7,
			wantErr:              false,
		},
		{
			name: "ca-cert-path takes precedence over ca-cert-base64-data",
			args: []string{"--name", "cacertdata", "--url", "http://bar.com", "--service-token", "my-token", "--ca-cert-path", "testdata/selfSignedCert.crt", "--ca-cert-base64-data", privateCaBase64Cert},
			expectedContext: config.NamedContext{
				Name: "cacertdata",
				Context: &config.StsContext{
					URL:              "http://bar.com",
					ServiceToken:     "my-token",
					APIPath:          "/api",
					CaCertBase64Data: selfSignedBase64Cert,
				},
			},
			totalContextInConfig: 7,
			wantErr:              false,
		},
		{
			name: "ca-cert ignored if skip-ssl is set",
			args: []string{"--name", "bar", "--url", "http://bar.com", "--service-token", "my-token", "--skip-ssl", "--ca-cert-path", "/path/to/ca.crt", "--ca-cert-base64-data", "base64-data"},
			expectedContext: config.NamedContext{
				Name: "bar",
				Context: &config.StsContext{
					URL:              "http://bar.com",
					ServiceToken:     "my-token",
					APIPath:          "/api",
					SkipSSL:          true,
					CaCertBase64Data: "",
					CaCertPath:       "",
				},
			},
			totalContextInConfig: 6,
			wantErr:              false,
		},
		{
			name:            "no save on missing tokens",
			args:            []string{"--name", "bar", "--url", "http://my-bar.com"},
			expectedContext: config.NamedContext{},
			wantErr:         true,
			errorMessage:    "one of the required flags {api-token | service-token} not set",
		},
		{
			name:            "ca-cert-path is not found",
			args:            []string{"--name", "bar", "--url", "http://my-bar.com", "--service-token", "my-token", "--ca-cert-path", "/path/to/ca.crt"},
			expectedContext: config.NamedContext{},
			wantErr:         true,
			errorMessage:    "no such file or directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, cmd := setupSaveCmd(t)
			setupConfig(t, cli)
			_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, tt.args...)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errorMessage != "" {
					assert.Contains(t, err.Error(), tt.errorMessage)
				}
			} else {
				cfg, err := config.ReadConfig(cli.ConfigPath)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedContext.Name, cfg.CurrentContext)
				assert.Len(t, cfg.Contexts, tt.totalContextInConfig)
				validateContext(t, cfg, tt.expectedContext)
			}
		})
	}
}

func validateContext(t *testing.T, cfg *config.Config, expectedContext config.NamedContext) {
	ctx, err := cfg.GetContext(expectedContext.Name)
	assert.NoError(t, err)
	assert.Equal(t, expectedContext.Context, ctx.Context)
}
