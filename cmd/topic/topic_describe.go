package topic

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

const (
	Name      = "name"
	Offset    = "offset"
	Number    = "nr"
	PageSize  = "pagesize"
	Partition = "partition"

	NameUsage      = "Topic name"
	OffsetUsage    = "The starting offset"
	NumberUsage    = "The number of messages to show"
	PageSizeUsage  = "The pagination size"
	PartitionUsage = "The Kafka partition to query"
	FileUsage      = "The JSON output file to save the messages to"

	DefaultOffset   = int32(0)
	DefaultNumber   = int32(10)
	DefaultPageSize = int32(10000)
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
	common.AddFileFlagVar(cmd, &args.File, FileUsage)

	return cmd
}

func argValueError(name string, value int32) common.CLIError {
	return common.NewCLIArgParseError(fmt.Errorf("invalid value for argument '%s' specified: %d", name, value))
}

func fetchMessages(request stackstate_api.ApiDescribeRequest, args *DescribeArgs) ([]stackstate_api.Message, common.CLIError) {
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
			return nil, common.NewResponseError(err, resp)
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

	return messages, nil
}

func RunDescribeCommand(args *DescribeArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		if args.Offset < 0 {
			return argValueError(Offset, args.Offset)
		}
		if args.Number < 1 {
			return argValueError(Number, args.Number)
		}
		if args.PageSize < 1 {
			return argValueError(PageSize, args.PageSize)
		}

		request := api.TopicApi.Describe(cli.Context, args.Name)

		if args.Partition != -1 {
			request = request.Partition(args.Partition)
		}

		messages, err := fetchMessages(request, args)
		if err != nil {
			return err
		}

		if args.File != "" {
			json, err := json.Marshal(messages)
			if err != nil {
				return common.NewExecutionError(err)
			}

			err = util.WriteFile(args.File, json)
			if err != nil {
				return common.NewWriteFileError(err, args.File)
			}

			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"filepath": args.File,
				})
			} else {
				cli.Printer.Successf("Messages saved to '%s'", args.File)
			}
		} else {
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
		}
		return nil
	}
}
