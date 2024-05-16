package topic

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stackvista/stackstate-cli/internal/util"
)

const (
	Name      = "name"
	Offset    = "offset"
	Limit     = "limit"
	PageSize  = "pagesize"
	Partition = "partition"

	NameUsage      = "Topic name"
	OffsetUsage    = "The starting offset"
	LimitUsage     = "The limit of messages to show"
	PartitionUsage = "The Kafka partition to query"
	FileUsage      = "The JSON output file to save the messages to"

	DefaultLimit = int32(10)
)

type DescribeArgs struct {
	Name      string
	Offset    int64
	Limit     int32
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

	cmd.Flags().Int64Var(&args.Offset, Offset, -1, OffsetUsage)
	cmd.Flags().Int32Var(&args.Limit, Limit, DefaultLimit, LimitUsage)
	cmd.Flags().Int32Var(&args.Partition, Partition, -1, PartitionUsage)
	common.AddFileFlagVar(cmd, &args.File, FileUsage)

	return cmd
}

func argValueError(name string, value int32) common.CLIError {
	return common.NewCLIArgParseError(fmt.Errorf("invalid value for argument '%s' specified: %d", name, value))
}
func arg64ValueError(name string, value int64) common.CLIError {
	return common.NewCLIArgParseError(fmt.Errorf("invalid value for argument '%s' specified: %d", name, value))
}

func fetchMessages(request stackstate_api.ApiDescribeRequest, args *DescribeArgs) ([]stackstate_api.Message, common.CLIError) {
	if args.Offset != -1 {
		request = request.Offset(args.Offset)
	}

	result, resp, err := request.Limit(args.Limit).Execute()

	if err != nil {
		return nil, common.NewResponseError(err, resp)
	}

	return result.Messages, nil
}

func RunDescribeCommand(args *DescribeArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		if args.Offset < -1 {
			return arg64ValueError(Offset, args.Offset)
		}
		if args.Limit < 1 {
			return argValueError(Limit, args.Limit)
		}

		request := api.TopicAPI.Describe(cli.Context, args.Name)

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
