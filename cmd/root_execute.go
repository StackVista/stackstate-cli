package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	pr "gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func Execute(ctx context.Context) {
	printer := pr.NewPrinter()

	cli := &di.Deps{
		Printer: printer,
	}

	cmd := AllCommands(cli)

	runCmd := cmd
	zerolog.SetGlobalLevel(zerolog.Disabled)
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		runCmd = cmd
		verbose, _ := cmd.Flags().GetBool("verbose")
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
			return fmt.Errorf("invalid choice for output flag: %s. Must be JSON, YAML or Auto", cfg.Output)
		}

		configuration := sts.NewConfiguration()
		configuration.Servers[0] = sts.ServerConfiguration{
			URL:         cli.Config.ApiUrl,
			Description: "",
			Variables:   nil,
		}
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
			exitCode = common.CommandExecutionExitCode
			showUsage = true
		}
		if showUsage {
			println(runCmd.UsageString())
		}
		os.Exit(exitCode)
	}
}
