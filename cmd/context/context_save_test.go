package context

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

const (
	testToken               = "my-token"
	testContextURL          = "http://bar.com"
	urlFlagArg              = "--url"
	testCACertContext       = "cacertdata"
	testAPIPath             = "/api"
	serviceTokenFlagArg     = "--service-token"
	caCertPathFlagArg       = "--ca-cert-path"
	caCertBase64DataFlagArg = "--ca-cert-base64-data"
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
			args: []string{nameFlagArg, "baz", urlFlagArg, "http://baz.com", "--api-token", testToken},
			expectedContext: config.NamedContext{
				Name: "baz",
				Context: &config.StsContext{
					URL:      "http://baz.com",
					APIToken: testToken,
					APIPath:  testAPIPath,
				},
			},
			totalContextInConfig: 7,
			wantErr:              false,
		},
		{
			name: "existing context",
			args: []string{nameFlagArg, testBar, urlFlagArg, testContextURL, serviceTokenFlagArg, testToken},
			expectedContext: config.NamedContext{
				Name: testBar,
				Context: &config.StsContext{
					URL:          testContextURL,
					ServiceToken: testToken,
					APIPath:      testAPIPath,
				},
			},
			totalContextInConfig: 6,
			wantErr:              false,
		},
		{
			name: "existing context ca-cert is set with ca-cert-path",
			args: []string{nameFlagArg, testBar, urlFlagArg, testContextURL, serviceTokenFlagArg, testToken, caCertPathFlagArg, "testdata/selfSignedCert.crt"},
			expectedContext: config.NamedContext{
				Name: testBar,
				Context: &config.StsContext{
					URL:              testContextURL,
					ServiceToken:     testToken,
					APIPath:          testAPIPath,
					CaCertBase64Data: selfSignedBase64Cert,
				},
			},
			totalContextInConfig: 6,
			wantErr:              false,
		},
		{
			name: "new context ca-cert is set with ca-cert-path",
			args: []string{nameFlagArg, testCACertContext, urlFlagArg, testContextURL, serviceTokenFlagArg, testToken, caCertBase64DataFlagArg, privateCaBase64Cert},
			expectedContext: config.NamedContext{
				Name: testCACertContext,
				Context: &config.StsContext{
					URL:              testContextURL,
					ServiceToken:     testToken,
					APIPath:          testAPIPath,
					CaCertBase64Data: privateCaBase64Cert,
				},
			},
			totalContextInConfig: 7,
			wantErr:              false,
		},
		{
			name: "ca-cert-path takes precedence over ca-cert-base64-data",
			args: []string{nameFlagArg, testCACertContext, urlFlagArg, testContextURL, serviceTokenFlagArg, testToken, caCertPathFlagArg, "testdata/selfSignedCert.crt", caCertBase64DataFlagArg, privateCaBase64Cert},
			expectedContext: config.NamedContext{
				Name: testCACertContext,
				Context: &config.StsContext{
					URL:              testContextURL,
					ServiceToken:     testToken,
					APIPath:          testAPIPath,
					CaCertBase64Data: selfSignedBase64Cert,
				},
			},
			totalContextInConfig: 7,
			wantErr:              false,
		},
		{
			name: "ca-cert ignored if skip-ssl is set",
			args: []string{nameFlagArg, testBar, urlFlagArg, testContextURL, serviceTokenFlagArg, testToken, "--skip-ssl", caCertPathFlagArg, "/path/to/ca.crt", caCertBase64DataFlagArg, "base64-data"},
			expectedContext: config.NamedContext{
				Name: testBar,
				Context: &config.StsContext{
					URL:              testContextURL,
					ServiceToken:     testToken,
					APIPath:          testAPIPath,
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
			args:            []string{nameFlagArg, testBar, urlFlagArg, "http://my-bar.com"},
			expectedContext: config.NamedContext{},
			wantErr:         true,
			errorMessage:    "one of the required flags {api-token | service-token} not set",
		},
		{
			name:            "ca-cert-path is not found",
			args:            []string{nameFlagArg, testBar, urlFlagArg, "http://my-bar.com", serviceTokenFlagArg, testToken, caCertPathFlagArg, "/path/to/ca.crt"},
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
