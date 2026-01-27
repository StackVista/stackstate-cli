package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/otelrelationmapping"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func OtelRelationtMappingCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "otel-relation-mapping",
		Short: "Manage OpenTelemetry relation mappings",
		Long:  "Manage OpenTelemetry relation mappings. Maps OTel resources to topology relations.",
	}

	cmd.AddCommand(otelrelationmapping.OtelRelationMappingListCommand(deps))
	cmd.AddCommand(otelrelationmapping.OtelRelationMappingStatusCommand(deps))

	return cmd
}
