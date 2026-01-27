package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/otelcomponentmapping"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func OtelComponentMappingCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "otel-component-mapping",
		Short: "Manage OpenTelemetry component mappings",
		Long:  "Manage OpenTelemetry component mappings. Maps OTel resources to topology components.",
	}

	cmd.AddCommand(otelcomponentmapping.OtelComponentMappingListCommand(deps))
	cmd.AddCommand(otelcomponentmapping.OtelComponentMappingStatusCommand(deps))

	return cmd
}
