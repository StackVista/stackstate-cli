package di

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/client"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
)

// Dependency Injection context for the CLI
type Deps struct {
	CurrentContext *config.StsContext
	ConfigPath     string
	Printer        printer.Printer
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

func (cli *Deps) CmdRunEWithApi(
	runFn CmdWithApiFn) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if cli.CurrentContext == nil {
			currCtx, err := config.LoadCurrentContext(cmd, viper.GetViper(), cli.ConfigPath)
			if err != nil {
				return err
			}

			cli.CurrentContext = currCtx
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

		return runFn(cmd, cli, api, serverInfo)
	}
}

func (cli *Deps) LoadClient(cmd *cobra.Command, context *config.StsContext) common.CLIError {
	cli.Client, cli.Context = client.NewStackStateClient(cmd.Context(), cli.IsVerBose, cli.Printer, context.URL, context.APIPath, context.APIToken, context.ServiceToken, context.ServiceBearer)
	return nil
}

func (cli *Deps) IsJson() bool {
	return cli.Output == common.JSONOutput
}
