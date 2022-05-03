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
	if cli.IsJson {
		cli.Printer.PrintJson(map[string]interface{}{
			"version":  cli.Version,
			"commit":   cli.Commit,
			"date":     cli.Date,
			"cli-type": cli.CLIType,
		})
	} else {
		cli.Printer.Table(printer.TableData{
			Header: []string{"Version", "Date", "CLI Type", "Commit"},
			Data:   [][]interface{}{{cli.Version, cli.Date, cli.CLIType, cli.Commit}},
		})
	}

	return nil
}
