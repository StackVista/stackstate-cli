package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func VersionCommand(cli *di.Deps) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "display the CLI version number",
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

	cli.Printer.Table(
		[]string{"Version", "Date", "CLI Type", "Commit"},
		[][]interface{}{{Version, Date, CLIType, Commit}},
		info,
	)

	return nil
}
