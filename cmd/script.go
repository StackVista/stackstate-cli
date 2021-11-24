package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/script"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
)

func ScriptCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "script",
		Short: "StackState scripting related commands.",
	}
	cmd.AddCommand(script.ScriptExecuteCommand(cfg))

	return cmd
}
