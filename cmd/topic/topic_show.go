package topic

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

const (
	OffsetFlag    = "offset"
	PartitionFlag = "partition"
	LimitFlag     = "limit"
)

type TopicShowArgs struct {
	Name      string
	Offset    int64
	Partition int64
	Limit     int64
}

func TopicShowCommand(cli *di.Deps) *cobra.Command {
	args := TopicShowArgs{}
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show topic details",
		Long:  "Show details of the Topic.",
		RunE:  cli.CmdRunEWithApi(RunTopicShowCommand(&args)),
	}
	common.AddRequiredNameFlagVar(cmd, &args.Name, "topic name")
	cmd.Flags().Int64VarP(&args.Offset, OffsetFlag, "", 0, "starting offset")
	cmd.Flags().Int64VarP(&args.Partition,
		PartitionFlag,
		"",
		0,
		"partition to pull data from. Defaults to all partitions")
	cmd.Flags().Int64VarP(&args.Limit,
		LimitFlag,
		"",
		0,
		"maximum amount of messages to show",
	)
	return cmd
}

func RunTopicShowCommand(args *TopicShowArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {

		messageList, resp, err := api.TopicApi.GetTopic(cli.Context, args.Name).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}
		countPerMessageType := map[string]int{}
		for _, v := range messageList.GetMessages() {
			for messageType, _ := range v.GetMessage() {
				countPerMessageType[messageType]++
			}
		}
		data := make([][]interface{}, 0)
		for k, v := range countPerMessageType {
			row := []interface{}{k, v}
			data = append(data, row)
		}
		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"topic message": data,
			})
		} else {
			cli.Printer.Table(printer.TableData{
				Header:              []string{"message type", "count"},
				Data:                data,
				MissingTableDataMsg: printer.NotFoundMsg{Types: "topic"},
			})
		}
		return nil
	}
}
