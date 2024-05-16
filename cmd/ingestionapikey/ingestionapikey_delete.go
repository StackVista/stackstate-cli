package ingestionapikey

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type DeleteArgs struct {
	ID int64
}

func DeleteCommand(deps *di.Deps) *cobra.Command {
	args := &DeleteArgs{}
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an Ingestion Api Key",
		Long:  "Deleted key can't be used by sources, so all ingestion pipelines for that key will fail.",
		RunE:  deps.CmdRunEWithApi(RunIngestionApiKeyDeleteCommand(args)),
	}

	common.AddRequiredIDFlagVar(cmd, &args.ID, "ID of the Ingestion Api Key to delete")

	return cmd
}

func RunIngestionApiKeyDeleteCommand(args *DeleteArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		resp, err := api.IngestionApiKeyAPI.DeleteIngestionApiKey(cli.Context, args.ID).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"deleted-ingestion-api-key": args.ID,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Ingestion Api Key deleted: %d", args.ID))
		}

		return nil
	}
}
