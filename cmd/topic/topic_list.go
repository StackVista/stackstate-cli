package topic

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

func ListCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all available Kafka topics",
		Long:  "List all Kafka topics available in SUSE Observability.",
		Example: `# list all topics
sts topic list`,
		RunE: deps.CmdRunEWithApi(RunListCommand),
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
