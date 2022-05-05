package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gitlab.com/stackvista/stackstate-cli2/cmd"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/mutex_flags"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
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
	addRequiredFlagsToCmd(sts)

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

	verbose, _ := cmd.Flags().GetBool(common.VerboseFlag)
	cli.IsVerBose = verbose
	if verbose {
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

func addRequiredFlagsToCmd(root *cobra.Command) {
	util.ForAllCmd(root, func(cmd *cobra.Command) {
		required := ""
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			if util.IsSingleRequiredFlag(flag) {
				required += fmt.Sprintf("--%s %s", flag.Name, strings.ToUpper(flag.Value.Type()))
			}
		})

		requiredMutexFlags := mutex_flags.GetAllMutexNames(cmd, true)
		for _, mutex := range requiredMutexFlags {
			mutexFlagUses := make([]string, 0)
			for _, mutexFlag := range mutex_flags.GetAllFlagsOfMutex(cmd, mutex) {
				mutexFlagUses = append(mutexFlagUses, fmt.Sprintf("--%s %s", mutexFlag.Name, strings.ToUpper(mutexFlag.Value.Type())))
			}
			required += "{ " + strings.Join(mutexFlagUses, " | ") + " }"
		}

		if required != "" {
			cmd.Use += " " + required
		}
	})
}
