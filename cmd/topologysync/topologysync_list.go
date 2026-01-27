package topologysync

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

func ListCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all active topology synchronizations",
		Long:  "List all active topology synchronizations. Shows status, component and relation counts, and error counts.",
		Example: `# list all topology synchronizations
sts topology-sync list`,
		RunE: deps.CmdRunEWithApi(RunListCommand),
	}

	return cmd
}

func FormatSyncTable(streams []stackstate_api.TopologyStreamListItem) printer.TableData {
	data := make([][]interface{}, len(streams))
	for i, stream := range streams {
		identifier := "-"
		if stream.SyncIdentifier.IsSet() {
			identifier = *stream.SyncIdentifier.Get()
		}
		data[i] = []interface{}{
			stream.NodeId,
			stream.Name,
			identifier,
			stream.Status,
			fmt.Sprintf("+%-4d %5s", stream.CreatedComponents, fmt.Sprintf("-%d", stream.DeletedComponents)),
			fmt.Sprintf("+%-4d %5s", stream.CreatedRelations, fmt.Sprintf("-%d", stream.DeletedRelations)),
			stream.Errors,
		}
	}
	return printer.TableData{
		Header:              []string{"Id", "Name", "Identifier", "Status", "Components", "Relations", "Errors"},
		Data:                data,
		MissingTableDataMsg: printer.NotFoundMsg{Types: "synchronizations"},
	}
}

func RunListCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
	streamList, resp, err := api.TopologySynchronizationApi.GetTopologySynchronizationStreams(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	streams := streamList.Streams

	sort.SliceStable(streams, func(i, j int) bool {
		return streams[i].Name < streams[j].Name
	})

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"synchronizations": streams,
		})
	} else {
		cli.Printer.Table(FormatSyncTable(streams))
	}

	return nil
}
