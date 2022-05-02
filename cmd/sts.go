package cmd

import (
	"context"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func STSCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sts",
		Short: "StackState: topology-powered observability", // never actually visible
		Long:  "StackState: topology-powered observability.",
	}
	cmd.SetUsageTemplate(cmd.UsageTemplate() +
		"For more information about this CLI visit https://l.stackstate.com/cli\n",
	)

	cmd.AddCommand(VersionCommand(cli))
	cmd.AddCommand(CliCommand(cli))
	if CLIType != "saas" {
		cmd.AddCommand(ScriptCommand(cli))
		cmd.AddCommand(SettingsCommand(cli))
		cmd.AddCommand(StackPackCommand(cli))
	}
	cmd.AddCommand(MonitorCommand(cli))

	return cmd
}

func Execute(ctx context.Context) {
	cli := &di.Deps{}
	cmd := STSCommand(cli)
	execute(ctx, cli, cmd)
}

func execute(ctx context.Context, cli *di.Deps, sts *cobra.Command) common.ExitCode {
	common.AddPersistentFlags(sts)

	if cli.Printer == nil {
		cli.Printer = printer.NewPrinter()
	}

	decapitalizeHelpCommand(sts)
	if CLIType == "" {
		CLIType = "local"
	}
	if CLIType != "saas" && CLIType != "full" && CLIType != "local" {
		panic("Type must either 'full', 'local' or 'saas'.")
	}

	zerolog.SetGlobalLevel(zerolog.Disabled)
	sts.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		return PreRunCommand(cli, cmd, args)
	}

	sts.SilenceErrors = true
	sts.SilenceUsage = true
	if err := sts.ExecuteContext(ctx); err != nil {
		var showUsage bool
		var exitCode int
		switch v := err.(type) {
		case common.CLIError:
			exitCode = v.ExitCode()
			showUsage = v.ShowUsage()
		default:
			showUsage = true
		}
		if exitCode == common.OkExitCode {
			exitCode = common.CommandFailedRequirementExitCode
		}

		if cli.IsJson {
			cli.Printer.PrintErrJson(err)
		} else {
			cli.Printer.PrintErr(err)
			if showUsage {
				cli.Printer.PrintLn(sts.UsageString())
			}
		}
		return exitCode
	} else {
		return common.OkExitCode
	}
}

func PreRunCommand(cli *di.Deps, cmd *cobra.Command, args []string) error {
	json, _ := cmd.Flags().GetBool(common.JsonFlag)
	cli.IsJson = json

	verbose, _ := cmd.Flags().GetBool(common.VerboseFlag)
	cli.IsVerBose = verbose
	if verbose {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}
	// needs to happen before run, but after execute
	// so flag-config bindings can take hold
	cfg, err := conf.ReadConf(cmd)
	if err != nil {
		cli.Printer.PrintErr(err)
		os.Exit(common.ConfigErrorExitCode)
	}
	cli.Config = &cfg
	log.Printf("Loaded config %+v", cli.Config)

	cli.Printer.SetUseColor(!cfg.NoColor)

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

	cli.Context = ctx
	cli.Client = di.NewStackStateClient(client, ctx)

	return nil
}

// we do this to be consistent with the styling rules
func decapitalizeHelpCommand(root *cobra.Command) {
	root.InitDefaultHelpCmd()
	help, _, _ := root.Find([]string{"help"})
	if help != nil {
		help.Short = "help about any command"
	}
}
