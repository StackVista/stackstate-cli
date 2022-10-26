package topologysync

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

type DescribeArgs struct {
	Id int64
}

func DescribeCommand(deps *di.Deps) *cobra.Command {
	args := &DescribeArgs{}
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Shows the configuration of a topology synchronization",
		Long:  "Shows the configuration of a topology synchronization.",
		RunE:  deps.CmdRunEWithApi(RunDescribeCommand(args)),
	}
	cmd.Flags().Int64VarP(&args.Id, common.IDFlag, common.IDFlagShort, 0, "The ID of topology synchronization")
	cmd.MarkFlagRequired(common.IDFlag) //nolint:errcheck

	return cmd
}

func RunDescribeCommand(args *DescribeArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		sync, resp, err := api.TopologySynchronizationApi.
			GetTopologySynchronizationStreamById(cli.Context).
			IdentifierType(DefaultIdentifierType).
			Identifier(fmt.Sprintf("%d", args.Id)).
			Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"synchronization": sync,
			})
		} else {
			cli.Printer.PrintLn("Topology Synchronization:")
			cli.Printer.Table(FormatSyncTable([]stackstate_api.TopologyStreamListItem{
				sync.Item,
			}))

			data := make([][]interface{}, len(sync.ErrorDetails))
			for i, error := range sync.ErrorDetails {
				id := "-"
				if error.ExternalId != nil {
					id = *error.ExternalId
				}
				data[i] = []interface{}{id, error.Level, error.Message}
			}

			cli.Printer.PrintLn("\nTopology Synchronization Errors:")
			cli.Printer.Table(printer.TableData{
				Header:              []string{"External ID", "Level", "Message"},
				Data:                data,
				MissingTableDataMsg: printer.NotFoundMsg{Types: "synchronization errors"},
			})
		}
		return nil
	}
}
