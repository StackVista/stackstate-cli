package health

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

type StatusArgs struct {
	Urn       string
	SubStream string
	Topology  string
}

func HealthStatusCommand(cli *di.Deps) *cobra.Command {
	args := &StatusArgs{}
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Status of the stream",
		Long:  "Status of an active health synchronization stream.",
		RunE:  cli.CmdRunEWithApi(RunHealthStatusCommand(args)),
	}
	common.AddRequiredUrnFlagVar(cmd, &args.Urn, "Urn of the stream")
	cmd.Flags().StringVar(&args.SubStream, SubStreamFlag, "", SubStreamFlagUsage)
	cmd.Flags().StringVar(&args.Topology, TopologyFlag, "", TopologyFlagUsage)

	return cmd
}

func RunHealthStatusCommand(args *StatusArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		var resp *http.Response
		var err error
		if args.Topology != "" {
			data := make([][]interface{}, 0)
			var topologies *stackstate_api.TopologyMatchResult
			if args.SubStream == "" {
				topologies, resp, err = api.HealthSynchronizationApi.GetHealthSynchronizationStreamTopologyMatches(cli.Context, args.Urn).Execute()
			} else {
				topologies, resp, err = api.HealthSynchronizationApi.GetHealthSynchronizationSubStreamTopologyMatches(cli.Context, args.Urn, args.SubStream).Execute()
			}
			if err != nil {
				return common.NewResponseError(err, resp)
			}
			for _, v := range topologies.GetMultipleMatchesCheckStates() {
				data = append(data, []interface{}{
					v.GetCheckStateId(),
					v.GetTopologyElementIdentifier(),
					v.GetMatchCount(),
				})
			}

			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"status": data,
				})
			} else {
				cli.Printer.Table(printer.TableData{
					Header:              []string{"Check state id", "Topology element identifier", "Number of matched topology elements"},
					Data:                data,
					MissingTableDataMsg: printer.NotFoundMsg{Types: "health status"},
				})
			}
		} else {
			if args.SubStream == "" {
				streamStatus, resp, err := api.HealthSynchronizationApi.GetHealthSynchronizationStreamStatus(cli.Context, args.Urn).Execute()
				if err != nil {
					return common.NewResponseError(err, resp)
				}
				if streamStatus.GetRecoverMessage() != "" {
					cli.Printer.Successf("This stream is in recovery mode.\n"+
						"This means stackstate is reconstructing the state of the health streams. "+
						"In this period no errors will be reported for the stream,\n"+
						"incoming data will be processed as usual.\n"+
						"The reason recovery mode was entered was because: %s", streamStatus.GetRecoverMessage())
				}
				cli.Printer.Successf("Consistency model for the stream and all substreams: %s", streamStatus.GetConsistencyModel())
				if mainStream, ok := streamStatus.GetMainStreamStatusOk(); ok {
					cli.Printer.PrintLn(fmt.Sprintf("Synchronized check state count: %d", mainStream.CheckStateCount))

					printSubStream(cli.Printer, &mainStream.SubStreamState)

					cli.Printer.PrintLn(fmt.Sprintf("\nSynchronization errors:\n %v", mainStream.GetErrors()))
					cli.Printer.PrintLn(fmt.Sprintf("\nSynchronization metrics::\n %v", mainStream.GetMetrics()))
				} else {
					cli.Printer.PrintLn("Synchronized check state count: -")
				}
				cli.Printer.PrintLn(fmt.Sprintf("\nAggregate metrics for the stream and all substreams:\n %v", streamStatus.GetAggregateMetrics()))
				cli.Printer.PrintLn(fmt.Sprintf("\nErrors for non-existing sub streams:\n %v", streamStatus.GetGlobalErrors()))
			} else {
				subSteamStatus, resp, err := api.HealthSynchronizationApi.GetHealthSynchronizationSubStreamStatus(cli.Context, args.Urn, args.SubStream).Execute()
				if err != nil {
					return common.NewResponseError(err, resp)
				}
				if count, ok := subSteamStatus.GetCheckStateCountOk(); ok {
					cli.Printer.PrintLn(fmt.Sprintf("Synchronized check state count: %d", count))
					printSubStream(cli.Printer, &subSteamStatus.SubStreamState)
				} else {
					cli.Printer.PrintLn("Synchronized check state count: -")
				}
			}
		}
		return nil
	}
}

func printSubStream(ptr printer.Printer, subStream *stackstate_api.HealthSubStreamConsistencyState) {
	if subStream == nil {
		ptr.PrintLn("Invalid Sub Stream State")
		return
	}
	const divideBy = int32(1000)
	if subStreamStateType := subStream.HealthSubStreamExpiry; subStreamStateType != nil {
		ptr.PrintLn(fmt.Sprintf("Repeat interval (Seconds): %.0f", float64(subStreamStateType.RepeatIntervalMs/divideBy)))
		ptr.PrintLn(fmt.Sprintf("Expiry (Seconds): %.0f", float64(subStreamStateType.ExpiryIntervalMs/divideBy)))
	} else if statusType := subStream.HealthSubStreamSnapshot; statusType != nil {
		ptr.PrintLn(fmt.Sprintf("Repeat interval (Seconds): %.0f", float64(statusType.RepeatIntervalMs/divideBy)))
		if expSecond := statusType.ExpiryIntervalMs; expSecond != nil {
			ptr.PrintLn(fmt.Sprintf("Expiry (Seconds): %.0f", float64(*expSecond/divideBy)))
		}
	} else if statusType := subStream.HealthSubStreamTransactionalIncrements; statusType != nil {
		ptr.PrintLn(fmt.Sprintf("Checkpoint offset: %d", statusType.GetOffset()))
		if idx := statusType.BatchIndex; idx != nil {
			ptr.PrintLn(fmt.Sprintf("Checkpoint batch index: %d", idx))
		}
	}
}
