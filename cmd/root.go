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
	cmd.AddCommand(VersionCommand())
	cmd.AddCommand(ScriptCommand(cli))

	cmd.PersistentFlags().Bool("verbose", false, "Print more verbose logging.")
	cmd.PersistentFlags().String("api-url", "", "StackState API URL.")
	cmd.PersistentFlags().String("api-token", "", "StackState API Token.")

	return cmd
}
