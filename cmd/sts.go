package cmd

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
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

	// shush, cobra... we take care of this ourselves
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
