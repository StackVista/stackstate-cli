package topic

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func ListCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List available topics",
		Long:  "List available topics.",
		RunE:  deps.CmdRunEWithApi(RunListCommand),
	}

	return cmd
}

func RunListCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
	topics, resp, err := api.TopicApi.List(cli.Context).Execute()

	if err != nil {
		return common.NewResponseError(err, resp)
	}

	names := make([]string, len(topics))
	for i, topic := range topics {
		names[i] = topic.Name
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"topics": names,
		})
	} else {
		data := make([][]interface{}, len(names))
		for i, name := range names {
			data[i] = []interface{}{name}
		}
		cli.Printer.Table(printer.TableData{
			Header:              []string{"Topic Name"},
			Data:                data,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "topics"},
		})
	}

	return nil
}
