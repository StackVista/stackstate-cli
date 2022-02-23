package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/persistent_flags"
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
	cmd.AddCommand(MonitorCommand(cli))
	persistent_flags.AddPersistentFlags(cmd)

	return cmd
}
