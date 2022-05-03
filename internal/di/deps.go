package di

import (
	"context"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

// Dependency Injection context for the CLI
type Deps struct {
	Config    *conf.Conf
	Printer   printer.Printer
	Context   context.Context
	Client    StackStateClient
	IsVerBose bool
	IsJson    bool
	NoColor   bool
	Version   string
	Commit    string
	Date      string
	CLIType   string
}

func NewDeps() Deps {
	return Deps{
		Client:  nil,
		Config:  nil,
		Printer: printer.NewPrinter(),
		Context: nil,
	}
}

func (cli *Deps) CmdRunE(runFn func(*Deps, *cobra.Command) common.CLIError) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return runFn(cli, cmd)
	}
}

type CmdWithApiFn = func(
	*cobra.Command,
	*Deps,
	*stackstate_client.APIClient,
	stackstate_client.ServerInfo,
) common.CLIError

func (cli *Deps) CmdRunEWithApi(
	runFn CmdWithApiFn) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {

		// needs to happen before run, but after execute
		// so flag-config bindings can take hold
		if cli.Config == nil {
			cfg, err := conf.ReadConf(cmd)
			if err != nil {
				return err
			}
			cli.Config = &cfg
			log.Printf("Loaded config %+v", cli.Config)
		}

		if cli.Client == nil {
			ctx, client, err := createClient(cli, cmd)
			if err != nil {
				return err
			}
			cli.Context = ctx
			cli.Client = client
		}

		api, serverInfo, err := cli.Client.Connect()
		if err != nil {
			return err
		}

		return runFn(cmd, cli, api, serverInfo)
	}
}

func createClient(cli *Deps, cmd *cobra.Command) (context.Context, StackStateClient, error) {
	configuration := stackstate_client.NewConfiguration()
	configuration.Servers[0] = stackstate_client.ServerConfiguration{
		URL:         cli.Config.ApiURL,
		Description: "",
		Variables:   nil,
	}
	configuration.Debug = cli.IsVerBose
	var client *stackstate_client.APIClient

	stopSpinner := func() {}
	configuration.OnPreCallAPI = func(r *http.Request) {
		stopSpinner = cli.Printer.StartSpinner(printer.AwaitingServer)
	}
	configuration.OnPostCallAPI = func(r *http.Request) {
		stopSpinner()
	}
	client = stackstate_client.NewAPIClient(configuration)

	auth := make(map[string]stackstate_client.APIKey)
	auth["ApiToken"] = stackstate_client.APIKey{
		Key:    cli.Config.ApiToken,
		Prefix: "",
	}
	ctx := context.WithValue(
		cmd.Context(),
		stackstate_client.ContextAPIKeys,
		auth,
	)

	return ctx, NewStackStateClient(client, ctx), nil
}
