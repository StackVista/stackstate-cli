package di

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
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

		if cli.Config == nil {
			err := cli.LoadConfig(cmd)
			if err != nil {
				return err
			}
		}

		if cli.Client == nil {
			err := cli.LoadClient(cmd, cli.Config.ApiURL, cli.Config.ApiToken)
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

// needs to happen before command run, but after cobra command execute
// so flag-config bindings can take hold
func (cli *Deps) LoadConfig(cmd *cobra.Command) common.CLIError {
	cfg, err := conf.ReadConf(cmd)
	if err != nil {
		return err
	}
	log.Printf("Loaded config %+v", cli.Config)
	cli.Config = &cfg
	return nil
}

func (cli *Deps) LoadClient(cmd *cobra.Command, apiURL string, apiToken string) common.CLIError {
	configuration := stackstate_client.NewConfiguration()
	configuration.Servers[0] = stackstate_client.ServerConfiguration{
		URL:         apiURL,
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
		Key:    apiToken,
		Prefix: "",
	}
	ctx := context.WithValue(
		cmd.Context(),
		stackstate_client.ContextAPIKeys,
		auth,
	)

	cli.Context = ctx
	cli.Client = NewStackStateClient(client, ctx)
	return nil
}
