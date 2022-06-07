package di

import (
	"context"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/client"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
)

// Dependency Injection context for the CLI
type Deps struct {
	// Config    *conf.Conf
	StsConfig      *config.Config
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
	CLIType        string
}

func (cli *Deps) CmdRunE(runFn func(*Deps, *cobra.Command) common.CLIError) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return runFn(cli, cmd)
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
		if cli.Client == nil {
			err := cli.LoadClient(cmd, cli.CurrentContext.URL, cli.CurrentContext.APIPath, cli.CurrentContext.APIToken, cli.CurrentContext.ServiceToken)
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

// Called from the PersistentPreRunE
func (cli *Deps) LoadConfig() common.CLIError {
	cfg, err := config.ReadConfig(cli.ConfigPath)
	if err != nil {
		return err
	}

	cli.StsConfig = cfg
	cli.CurrentContext = cfg.GetCurrentContext()

	return nil
}

func (cli *Deps) LoadClient(cmd *cobra.Command, apiURL string, apiPath string, apiToken, serviceToken string) common.CLIError {
	cli.Client, cli.Context = client.NewStackStateClient(cmd.Context(), cli.IsVerBose, cli.Printer, apiURL, apiPath, apiToken, serviceToken)
	return nil
}

func (cli *Deps) IsJson() bool {
	return cli.Output == common.JSONOutput
}
