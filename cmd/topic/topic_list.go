package topic

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func TopicListCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list topic",
		Long:  "List available Topic.",
		RunE:  cli.CmdRunEWithApi(RunTopicListCommand),
	}
	return cmd
}

func RunTopicListCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	topicList, resp, err := api.TopicApi.ListTopics(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}
	data := make([][]interface{}, 0)
	for _, v := range topicList {
		row := []interface{}{v.GetName()}
		data = append(data, row)
	}
	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"topics": data,
		})
	} else {
		cli.Printer.Table(printer.TableData{
			Header:              []string{"name"},
			Data:                data,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "Topics"},
		})
	}
	return nil
}
