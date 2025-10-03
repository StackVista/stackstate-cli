package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/otelrelationmapping"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func OtelRelationtMappingCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "otel-relation-mapping",
		Short: "Manage the Otel Relation Mapping",
		Long:  "Manage the Otel Relation Mapping.",
	}

	cmd.AddCommand(otelrelationmapping.OtelRelationMappingListCommand(deps))
	cmd.AddCommand(otelrelationmapping.OtelRelationMappingStatusCommand(deps))

	return cmd
}
