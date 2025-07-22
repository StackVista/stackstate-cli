package di

import (
	"context"
	"encoding/base64"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	stackstate_admin_api "github.com/stackvista/stackstate-cli/generated/stackstate_admin_api"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/client"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/editor"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stackvista/stackstate-cli/pkg/pflags"
)

// Dependency Injection context for the CLI
type Deps struct {
	CurrentContext *config.StsContext
	ConfigPath     string
	Printer        printer.Printer
	Editor         editor.Editor
	Context        context.Context
	Client         client.StackStateClient
	Clock          pflags.Clock
	IsVerBose      bool
	Output         common.Output
	NoColor        bool
	Version        string
	Commit         string
	BuildDate      string
}

func (cli *Deps) CmdRunE(runFn func(*Deps, *cobra.Command) common.CLIError) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return runFn(cli, cmd)
	}
}

func (cli *Deps) CmdRunEWithConfig(runFn func(*Deps, *cobra.Command, *config.Config) common.CLIError) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfig(cli.ConfigPath)
		if err != nil {
			return err
		}

		return runFn(cli, cmd, cfg)
	}
}

type CmdWithApiFn = func(
	*cobra.Command,
	*Deps,
	*stackstate_api.APIClient,
	*stackstate_api.ServerInfo,
) common.CLIError

func (cli *Deps) CmdRunEWithApi(runFn CmdWithApiFn) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if cli.CurrentContext == nil {
			err := cli.LoadContext(cmd)
			if err != nil {
				return err
			}
		}

		log.Debug().Interface("context", cli.CurrentContext).Msg("Using context")

		if cli.Client == nil {
			err := cli.LoadClient(cmd, cli.CurrentContext)
			if err != nil {
				return err
			}
		}

		api, serverInfo, err := cli.Client.Connect()
		if err != nil {
			return err
		}

		err = client.CheckVersionCompatibility(serverInfo.Version, cmd.Version)
		if err != nil {
			return err
		}

		return runFn(cmd, cli, api, serverInfo)
	}
}

func (cli *Deps) LoadContext(cmd *cobra.Command) common.CLIError {
	currCtx, err := config.LoadCurrentContext(cmd.Context(), cmd, viper.GetViper(), cli.ConfigPath)
	if err != nil {
		return err
	}

	cli.CurrentContext = currCtx
	return nil
}

func (cli *Deps) LoadClient(cmd *cobra.Command, context *config.StsContext) common.CLIError {
	logger := log.Ctx(cmd.Context())
	caCertData, err := getCaCertificate(context, logger)
	if err != nil {
		return err
	}
	cli.Client, cli.Context, err = client.NewStackStateClient(
		cmd.Context(),
		cli.IsVerBose,
		cli.Printer,
		cli.Version,
		context.URL,
		context.APIPath,
		context.AdminAPIPath,
		context.APIToken,
		context.ServiceToken,
		context.K8sSAToken,
		context.SkipSSL,
		caCertData,
	)
	return err
}

type CmdWithAdminApiFn = func(
	*cobra.Command,
	*Deps,
	*stackstate_admin_api.APIClient,
) common.CLIError

func (cli *Deps) CmdRunEWithAdminApi(runFn CmdWithAdminApiFn) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if cli.CurrentContext == nil {
			err := cli.LoadContext(cmd)
			if err != nil {
				return err
			}
		}

		log.Debug().Interface("context", cli.CurrentContext).Msg("Using context")

		if cli.Client == nil {
			err := cli.LoadClient(cmd, cli.CurrentContext)
			if err != nil {
				return err
			}
		}

		api, err := cli.Client.AdminConnect()
		if err != nil {
			return err
		}

		// NOTE Can't validate whether the server version is compatible.

		return runFn(cmd, cli, api)
	}
}

func (cli *Deps) IsJson() bool {
	return cli.Output == common.JSONOutput
}

func getCaCertificate(context *config.StsContext, logger *zerolog.Logger) ([]byte, common.CLIError) {
	if context.SkipSSL {
		if context.HasCaCertificateSet() {
			logger.Warn().Msg("Both skip-ssl and one of ca-cert-path or ca-cert-base64-data are set. ca-cert-path and/or ca-cert-base64-data will be ignored.")
		}
		return []byte{}, nil
	}
	switch {
	case context.HasCaCertificateFromFileSet() && context.HasCaCertificateFromArgSet():
		logger.Warn().Msg("Both ca-cert-path and ca-cert-base64-data specified, ca-cert-path will be used.")
		return readFile(context.CaCertPath)
	case context.HasCaCertificateFromFileSet():
		return readFile(context.CaCertPath)
	case context.HasCaCertificateFromArgSet():
		caCertData, err := base64.StdEncoding.DecodeString(context.CaCertBase64Data)
		if err != nil {
			return nil, common.NewAPIClientCreateError("CaCertBase64Data is not a valid base64 encoded string")
		}
		return caCertData, nil
	default:
		return []byte{}, nil
	}
}

func readFile(path string) ([]byte, common.CLIError) {
	caCertData, err := os.ReadFile(path)
	if err != nil {
		return nil, common.NewReadFileError(err, path)
	}
	return caCertData, nil
}
