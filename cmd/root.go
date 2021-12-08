package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sts",
		Short: "StackState Command Line Interface",
	}
	return cmd
}

func AllCommands(cli *di.Deps) *cobra.Command {
	cmd := RootCommand()
	cmd.AddCommand(VersionCommand(cli))
	cmd.AddCommand(ScriptCommand(cli))
	cmd.AddCommand(CliCommand(cli))

	cmd.PersistentFlags().String(common.ApiUrlFlag, "", "StackState API URL.")
	cmd.PersistentFlags().String(common.ApiTokenFlag, "", "StackState API Token.")
	cmd.PersistentFlags().Bool(common.VerboseFlag, false, "Print verbose logging to see what the CLI is doing.")
	cmd.PersistentFlags().Bool(common.NoColorFlag, false, "Print to terminal without color.")
	cmd.PersistentFlags().StringP(common.OutputFlag, "o", "Auto", "Format output as: JSON, YAML or Auto.")

	return cmd
}
