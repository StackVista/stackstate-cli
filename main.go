package main

import (
	"context"
	"os"
	"runtime"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/static_info"
)

func main() {
	ctx := log.Logger.WithContext(context.Background())
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cli := &di.Deps{
		Version:   static_info.Version,
		Commit:    static_info.Commit,
		CLIType:   static_info.CLIType,
		BuildDate: static_info.BuildDate,
	}
	if cli.CLIType == "" {
		cli.CLIType = "local"
	}

	cmd := cmd.STSCommand(cli)
	execute(ctx, cli, cmd)
}

func execute(ctx context.Context, cli *di.Deps, sts *cobra.Command) common.ExitCode {
	common.AddPersistentFlags(sts)
	common.AddRequiredFlagsToCmd(sts)

	if cli.Printer == nil {
		cli.Printer = printer.NewPrinter()
	}

	decapitalizeHelpCommand(sts)
	if cli.CLIType != "saas" && cli.CLIType != "full" && cli.CLIType != "local" {
		panic("Type must either 'full', 'local' or 'saas'.")
	}

	zerolog.SetGlobalLevel(zerolog.Disabled)
	sts.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		return PreRunCommand(cli, cmd)
	}

	// shush, cobra... we take care of this ourselves
	sts.SilenceErrors = true
	sts.SilenceUsage = true

	if cmd, err := sts.ExecuteContextC(ctx); err != nil {
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
				cli.Printer.PrintLn(cmd.UsageString())
			}
		}
		return exitCode
	} else {
		return common.OkExitCode
	}
}

func PreRunCommand(cli *di.Deps, cmd *cobra.Command) error {
	if strings.ToLower(os.Getenv("TERM")) == "dumb" {
		cli.NoColor = true
	}
	// Implementation of https://no-color.org/
	if _, noColorEnvExists := os.LookupEnv("NO_COLOR"); noColorEnvExists {
		cli.NoColor = true
	}
	noColorFlag, _ := cmd.Flags().GetBool(common.NoColorFlag)
	if noColorFlag {
		cli.NoColor = true
	}

	json, _ := cmd.Flags().GetBool(common.JsonFlag)
	cli.IsJson = json
	if json {
		cli.NoColor = true
	}

	verbosity, _ := cmd.Flags().GetCount(common.VerboseFlag)
	cli.IsVerBose = verbosity > 0
	switch verbosity {
	case 0: //nolint:gomnd
		// Nothing to do
	case 1: //nolint:gomnd
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case 2: //nolint:gomnd
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	log.Info().
		Strs("args", os.Args).
		Str("os", runtime.GOOS).
		Str("version", cli.Version).
		Str("commit", cli.Commit).
		Str("date", cli.BuildDate).
		Str("CLIType", cli.CLIType).
		Msg("starting CLI")

	cli.Printer.SetUseColor(!cli.NoColor)

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
