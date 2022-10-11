package topic

import (
	"encoding/json"
	"sort"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

const (
	Name      = "name"
	Offset    = "offset"
	Number    = "nr"
	PageSize  = "pagesize"
	Partition = "partition"
	File      = "file"
	FileShort = "f"

	NameUsage      = "Topic name"
	OffsetUsage    = "The starting offset"
	NumberUsage    = "The number of messages to show"
	PageSizeUsage  = "The pagination size"
	PartitionUsage = "The Kafka partition to query"
	FileUsage      = "The JSON output file to save the messages to"

	DefaultOffset   = 0
	DefaultNumber   = 10
	DefaultPageSize = 10000
)

type DescribeArgs struct {
	Name      string
	Offset    int32
	Number    int32
	PageSize  int32
	Partition int32
	File      string
}

func DescribeCommand(deps *di.Deps) *cobra.Command {
	args := &DescribeArgs{}
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe a given topic",
		Long:  "Describe a given topic.",
		RunE:  deps.CmdRunEWithApi(RunDescribeCommand(args)),
	}

	cmd.Flags().StringVar(&args.Name, Name, "", NameUsage)
	cmd.MarkFlagRequired(Name) //nolint:errcheck

	cmd.Flags().Int32Var(&args.Offset, Offset, DefaultOffset, OffsetUsage)
	cmd.Flags().Int32Var(&args.Number, Number, DefaultNumber, NumberUsage)
	cmd.Flags().Int32Var(&args.PageSize, PageSize, DefaultPageSize, PageSizeUsage)
	cmd.Flags().Int32Var(&args.Partition, Partition, -1, PartitionUsage)
	cmd.Flags().StringVarP(&args.File, File, FileShort, "", FileUsage)

	return cmd
}

func RunDescribeCommand(args *DescribeArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		request := api.TopicApi.Describe(cli.Context, args.Name)

		if args.Partition != -1 {
			request = request.Partition(args.Partition)
		}

		// NOTE Fetch up to args.Number messages from the topic, at most args.PageSize at a time starting at args.Offset.
		remaining := args.Number
		offset := args.Offset
		messages := make([]stackstate_api.Message, 0)

		for remaining > 0 {
			pageSize := args.PageSize

			if remaining < args.PageSize {
				pageSize = remaining
			}

			result, resp, err := request.Offset(offset).Limit(pageSize).Execute()

			if err != nil {
				return common.NewResponseError(err, resp)
			}

			messages = append(messages, result.Messages...)

			if int32(len(result.Messages)) < pageSize {
				// Exhausted the topic.
				remaining = 0
			} else {
				// Prepare to fetch the next page.
				offset += pageSize
				remaining -= pageSize
			}
		}

		sort.SliceStable(messages, func(i, j int) bool {
			return messages[i].Offset < messages[j].Offset
		})

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"messages": messages,
			})
		} else {
			data := make([][]interface{}, len(messages))
			for i, message := range messages {
				json, err := json.Marshal(message.Message)
				if err != nil {
					return common.NewExecutionError(err)
				}
				data[i] = []interface{}{message.Key, message.Partition, message.Offset, string(json)}
			}
			cli.Printer.Table(printer.TableData{
				Header:              []string{"Key", "Partition", "Offset", "Message"},
				Data:                data,
				MissingTableDataMsg: printer.NotFoundMsg{Types: "messages"},
			})
		}

		return nil
	}
}
