package anomaly

import (
	"fmt"
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
	cmd.Flags().DurationP(StartTimeFlag, "s", 0, "start time of interval with anomalies")
	cmd.MarkFlagRequired(StartTimeFlag) //nolint:errcheck
	cmd.Flags().DurationP(EndTimeFlag, "e", 0, "end time of interval with anomalies")
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
	endTime, err := cmd.Flags().GetDuration(EndTimeFlag)
	if err != nil {
		return common.NewCLIError(err)
	}
	now := time.Now().UnixMilli()

	feedback, resp, err := api.AnomalyFeedbackApi.AnomalyFeedbackGet(cli.Context).
		StartTime(now + startTime.Milliseconds()).
		EndTime(now + endTime.Milliseconds()).
		History(24 * 60 * 60 * 1000).
		Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.Success(fmt.Sprintf("Downloaded %d anomalies with feedback.", len(feedback)))

	return nil
}
