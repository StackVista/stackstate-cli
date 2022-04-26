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
	common.AddJsonFlag(cmd)
	return cmd
}

func RunVersionCommand(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	json, err := cmd.Flags().GetBool(common.JsonFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	if json {
		info := map[string]interface{}{
			"version":  Version,
			"commit":   Commit,
			"date":     Date,
			"cli-type": CLIType,
		}

		cli.Printer.PrintJson(info)
	} else {
		cli.Printer.Table(printer.TableData{
			Header: []string{"Version", "Date", "CLI Type", "Commit"},
			Data:   [][]interface{}{{Version, Date, CLIType, Commit}},
		})
	}

	return nil
}
