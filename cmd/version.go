package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func VersionCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "display version info",
		Long:  "Display the version of this StackState CLI.",
		RunE:  cli.CmdRunE(RunVersionCommand),
	}
	return cmd
}

func RunVersionCommand(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"version":    cli.Version,
			"commit":     cli.Commit,
			"build-date": cli.BuildDate,
			"cli-type":   cli.CLIType,
		})
	} else {
		cli.Printer.Table(printer.TableData{
			Header: []string{"Version", "Build Date", "Commit", "CLI Type"},
			Data:   [][]interface{}{{cli.Version, cli.BuildDate, cli.Commit, cli.CLIType}},
		})
	}

	return nil
}
