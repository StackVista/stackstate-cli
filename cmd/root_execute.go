package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	pr "gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

//nolint:funlen
func Execute(ctx context.Context) {
	printer := pr.NewPrinter()

	cli := &di.Deps{
		Printer: printer,
	}

	cmd := RootCommand(cli)
	if CLIType == "" {
		CLIType = "local"
	}
	if CLIType != "saas" && CLIType != "full" && CLIType != "local" {
		panic("Type must either 'full', 'local' or 'saas'.")
	}

	runCmd := cmd
	zerolog.SetGlobalLevel(zerolog.Disabled)
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		runCmd = cmd
		verbose, _ := cmd.Flags().GetBool("verbose")
		cli.IsVerBose = verbose
		if verbose {
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		}
		// needs to happen before run, but after execute
		// so flag-config bindings can take hold
		cfg, err := conf.ReadConf(cmd)
		if err != nil {
			printer.PrintErr(err)
			os.Exit(common.ConfigErrorExitCode)
		}
		cli.Config = &cfg
		log.Printf("Loaded config %+v", cli.Config)

		cli.Printer.SetUseColor(!cfg.NoColor)

		switch strings.ToUpper(cfg.Output) {
		case "JSON":
			cli.Printer.SetOutputType(pr.JSON)
		case "YAML":
			cli.Printer.SetOutputType(pr.YAML)
		case "AUTO":
			cli.Printer.SetOutputType(pr.Auto)
		default:
			return fmt.Errorf("invalid choice for output flag: %s. Must be JSON, YAML or auto", cfg.Output)
		}

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
			stopSpinner = cli.Printer.StartSpinner(common.AwaitingServer)
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

		cli.Context = ctx
		cli.Client = di.NewStackStateClient(client, ctx)

		return nil
	}

	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	if err := cmd.ExecuteContext(ctx); err != nil {
		cli.Printer.PrintErr(err)
		var showUsage bool
		var exitCode int
		switch v := err.(type) {
		case common.CLIError:
			exitCode = v.GetExitCode()
			showUsage = v.ShowUsage()
		default:
			exitCode = common.CommandFailedRequirementExitCode
			showUsage = true
		}
		if showUsage {
			println(runCmd.UsageString())
		}
		os.Exit(exitCode)
	}
}
