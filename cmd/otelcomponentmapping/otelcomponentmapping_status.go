package otelcomponentmapping

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/otelmapping"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func OtelComponentMappingStatusCommand(deps *di.Deps) *cobra.Command {
	args := &otelmapping.StatusArgs{}
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get the status of an OTel Component Mappings",
		Long:  "Get the status of an OTel Component Mappings.",
		RunE:  deps.CmdRunEWithApi(otelmapping.RunStatus(args, "component", otelmapping.FetchComponentStatus)),
	}

	common.AddRequiredIdentifierFlagVar(cmd, &args.Identifier, "Identifier of the OTel Component Mapping")

	return cmd
}
