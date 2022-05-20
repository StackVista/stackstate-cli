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
	"gitlab.com/stackvista/stackstate-cli2/cmd"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
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
		Clock:     pflags.NewFixedTimeClock(),
	}
	if cli.CLIType == "" {
		cli.CLIType = "local"
	}

	cmd := cmd.STSCommand(cli)

	os.Exit(execute(ctx, cli, cmd))
}

func execute(ctx context.Context, cli *di.Deps, sts *cobra.Command) common.ExitCode {
	common.AddPersistentFlags(sts)
	common.AddRequiredFlagsToCmd(sts)
	setUsageTemplates(sts)
	throwErrorOnUnknownSubCommand(sts, cli)
	sts.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		return common.NewCLIArgParseError(err)
	})

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
				cli.Printer.PrintLn(strings.TrimRight(cmd.UsageString(), "\n"))
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
	case 0:
		// Nothing to do
	case 1:
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

/**
For commands that have sub-commands we can show a simpler usage template.
*/
func setUsageTemplates(sts *cobra.Command) {
	cobraCommand := cobra.Command{}
	fullTemplate := cobraCommand.UsageTemplate()
	subCommandTemplate := `Usage:{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

	util.ForAllCmd(sts, func(c *cobra.Command) {
		c.SetUsageTemplate(fullTemplate)
	})
	sts.SetUsageTemplate(subCommandTemplate)

	for _, c := range sts.Commands() {
		if c.HasSubCommands() {
			c.SetUsageTemplate(subCommandTemplate)

			// we need to set verb command to the full template,
			// because otherwise they will inherit the simple template from their noun command parent
			for _, verbCommands := range c.Commands() {
				verbCommands.SetUsageTemplate(fullTemplate)
			}
		}
	}
}

/**
By default Cobra does not provide an error when a wrong sub-command is entered.
We got some customer feedback that this was missing.
*/
func throwErrorOnUnknownSubCommand(sts *cobra.Command, cli *di.Deps) {
	for _, c := range sts.Commands() {
		if c.HasSubCommands() {
			c.RunE = func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					// this behaves just like the root. Not calling a noun command with a verb sub-command
					// simply prints the full usage string prefixed with the long description
					cli.Printer.PrintLn(cmd.Long + "\n\n" + strings.TrimRight(cmd.UsageString(), "\n"))
					return nil
				}
				for _, sub := range cmd.Commands() {
					if args[0] == sub.Name() {
						return nil
					}
				}
				return common.NewCLIArgParseError(fmt.Errorf("unknown sub-command \"%s\" for \"%s\"", args[0], cmd.Name()))
			}
		}
	}
}
