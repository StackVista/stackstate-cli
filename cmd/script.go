package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/script"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func ScriptCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "script",
		Short: "Run scripts",
		Long:  "Run, test and develop scripts.",
	}
	cmd.AddCommand(script.ScriptRunCommand(cli))

	return cmd
}
