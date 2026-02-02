package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/topology"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func TopologyCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "topology",
		Short: "Inspect SUSE Observability topology components",
		Long:  "Inspect SUSE Observability topology components. Query and display topology components using component types, tags, and identifiers.",
	}
	cmd.AddCommand(topology.InspectCommand(cli))

	return cmd
}
