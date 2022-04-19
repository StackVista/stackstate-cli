package anomaly

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

const (
	StartTimeFlag = "start-time"
	EndTimeFlag = "end-time"
	HistoryFlag = "history"
)

func AnomalyCollect(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collect --start-time=TIMESTAMP [--end-time=TIMESTAMP] [--history=DURATION]",
		Short: "collect anomaly feedback",
		RunE:  cli.CmdRunEWithApi(RunCollectFeedbackCommand),
	}
	cmd.Flags().DurationP(StartTimeFlag, "s", -1, "start time of interval with anomalies")
	cmd.MarkFlagRequired(StartTimeFlag) //nolint:errcheck
	cmd.Flags().DurationP(EndTimeFlag, "e", -1, "end time of interval with anomalies")
	cmd.Flags().DurationP(HistoryFlag, "d", 24 * time.Hour, "length of metric data history preceding the anomaly")

	return cmd
}

func RunCollectFeedbackCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_client.APIClient,
	serverInfo stackstate_client.ServerInfo,
) common.CLIError {
	startTime, err := cmd.Flags().GetDuration(StartTimeFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	stackpack, resp, err := api.StackpackApi.StackpackUpload(cli.Context).StackPack(file).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.Success(fmt.Sprintf("uploaded StackPack: %s", filePath))
	cli.Printer.Table(
		[]string{"name", "display name", "version"},
		[][]interface{}{{stackpack.Name, stackpack.DisplayName, stackpack.Version}},
		stackpack,
	)

	return nil
}
