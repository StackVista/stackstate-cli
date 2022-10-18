package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/script"
	"github.com/stackvista/stackstate-cli/internal/di"
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
