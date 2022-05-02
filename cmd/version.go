package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func VersionCommand(cli *di.Deps) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "display version info",
		Long:  "Display the version of this StackState CLI.",
		RunE:  cli.CmdRunE(RunVersionCommand),
	}
}

func RunVersionCommand(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	info := map[string]interface{}{
		"version":  Version,
		"commit":   Commit,
		"date":     Date,
		"cli-type": CLIType,
	}

	cli.Printer.Table(printer.TableData{
		Header:     []string{"Version", "Date", "CLI Type", "Commit"},
		Data:       [][]interface{}{{Version, Date, CLIType, Commit}},
		StructData: info,
	})

	return nil
}
