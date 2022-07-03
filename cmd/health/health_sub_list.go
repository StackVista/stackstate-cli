package health

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

type SubListArgs struct {
	Urn string
}

func HealthSubListCommand(cli *di.Deps) *cobra.Command {
	args := &SubListArgs{}
	cmd := &cobra.Command{
		Use:   "sub-list",
		Short: "List of sub stream",
		Long:  "List active health synchronization sub streams of a stream.",
		RunE:  cli.CmdRunEWithApi(RunHealthSubListCommand(args)),
	}
	cmd.Flags().StringVar(&args.Urn, UrnFlag, "", UrnFlagUsage)
	cmd.MarkFlagRequired(UrnFlag) //nolint:errcheck

	return cmd
}

func RunHealthSubListCommand(args *SubListArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		subList, resp, err := api.HealthSynchronizationApi.GetHealthSynchronizationSubStreamOverview(cli.Context, args.Urn).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}
		data := make([][]interface{}, 0)
		for _, v := range subList.GetSubStreams() {
			data = append(data, []interface{}{v.GetSubStreamId(), v.GetCheckStateCount()})
		}
		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"sub-stream": data,
			})
		} else {
			cli.Printer.Table(printer.TableData{
				Header:              []string{"Sub stream id", "Check state count"},
				Data:                data,
				MissingTableDataMsg: printer.NotFoundMsg{Types: "health sub-stream"},
			})
		}
		return nil
	}
}
