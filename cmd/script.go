package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/script"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func ScriptCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "script",
		Short: "Execute STSL scripts on the server",
		Long:  "Execute STSL (SUSE Observability Scripting Language) scripts on the server. Scripts can query topology, metrics, and other data.",
	}
	cmd.AddCommand(script.ScriptRunCommand(cli))

	return cmd
}
