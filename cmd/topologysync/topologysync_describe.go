package topologysync

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

const (
	NodeIdType     stackstate_api.IdentifierType = "NodeId"
	IdentifierType stackstate_api.IdentifierType = "Identifier"
)

type DescribeArgs struct {
	ID         int64
	Identifier string
}

func DescribeCommand(deps *di.Deps) *cobra.Command {
	args := &DescribeArgs{}
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Show the details of an active Topology synchronization",
		Long:  "Show the details of an active Topology synchronization.",
		RunE:  deps.CmdRunEWithApi(RunDescribeCommand(args)),
	}
	cmd.Flags().Int64VarP(&args.ID, common.IDFlag, common.IDFlagShort, 0, "The ID of a topology synchronization")
	cmd.Flags().StringVar(&args.Identifier, common.IdentifierFlag, "", "The identifier of a topology synchronization")
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}

func IdOrIdentifier(id int64, identifier string) (stackstate_api.IdentifierType, string) {
	if id != 0 {
		return NodeIdType, fmt.Sprintf("%d", id)
	} else {
		return IdentifierType, identifier
	}
}

func RunDescribeCommand(args *DescribeArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		idType, id := IdOrIdentifier(args.ID, args.Identifier)
		sync, resp, err := api.TopologySynchronizationAPI.
			GetTopologySynchronizationStreamById(cli.Context).
			IdentifierType(idType).
			Identifier(id).
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
