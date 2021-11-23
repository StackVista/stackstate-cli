package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/openapi"
)

func ScriptCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "script",
		Short: "StackState scripting related commands.",
	}
	cmd.AddCommand(ScriptExecuteCommand())

	return cmd
}

func ScriptExecuteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute",
		Short: "Execute a STSL script.",
		Run: func(_ *cobra.Command, _ []string) {
			openapi.NewConfiguration()
		},
	}

	return cmd
}
