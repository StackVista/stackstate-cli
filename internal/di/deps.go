package di

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

// Depedency Injection context for the CLI
type Deps struct {
	Config  *conf.Conf
	Client  *sts.APIClient
	Printer printer.Printer
	Context context.Context
}

func NewDeps() Deps {
	return Deps{
		Config:  nil,
		Client:  nil,
		Printer: printer.NewPrinter(),
		Context: nil,
	}
}

func CmdRunEWithDeps(cli *Deps, runFn func(*Deps, *cobra.Command, []string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {

		log.Ctx(cmd.Context()).Info().Msg(fmt.Sprintf("Loaded config %+v", cli.Config))

		if cli.Client == nil {
			configuration := sts.NewConfiguration()
			configuration.Servers[0] = sts.ServerConfiguration{
				URL:         cli.Config.ApiUrl,
				Description: "",
				Variables:   nil,
			}
			verbose, _ := cmd.Flags().GetBool("verbose")
			if verbose {
				configuration.Debug = true
			}

			client := sts.NewAPIClient(configuration)
			cli.Client = client

			auth := make(map[string]sts.APIKey)
			auth["ApiToken"] = sts.APIKey{
				Key:    cli.Config.ApiToken,
				Prefix: "",
			}
			cli.Context = context.WithValue(
				cmd.Context(),
				sts.ContextAPIKeys,
				auth,
			)
		}

		return runFn(cli, cmd, args)
	}
}
