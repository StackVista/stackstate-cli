package cmd

import (
	"github.com/spf13/cobra"
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

	cmd.PersistentFlags().Bool("verbose", false, "Print verbose logging to see what the CLI is doing.")
	cmd.PersistentFlags().String("api-url", "", "StackState API URL.")
	cmd.PersistentFlags().String("api-token", "", "StackState API Token.")
	cmd.PersistentFlags().Bool("no-color", false, "Print to terminal without color.")
	cmd.PersistentFlags().StringP("output", "o", "Auto", "Format output as: JSON, YAML or Auto.")

	return cmd
}
