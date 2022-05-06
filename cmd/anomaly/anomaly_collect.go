package anomaly

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

const (
	StartTimeFlag = "start-time"
	EndTimeFlag   = "end-time"
	HistoryFlag   = "history"
	FileFlag      = "file"

	Day = 24 * time.Hour
)

func AnomalyCollect(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collect --start-time START-TIME -f FILE",
		Short: "collect anomaly feedback",
		RunE:  cli.CmdRunEWithApi(RunCollectFeedbackCommand),
	}
	cmd.Flags().DurationP(StartTimeFlag, "", 0, "start time of interval with anomalies")
	cmd.MarkFlagRequired(StartTimeFlag) //nolint:errcheck
	cmd.Flags().StringP(FileFlag, "f", "", "path to output file")
	cmd.MarkFlagRequired(FileFlag) //nolint:errcheck
	cmd.Flags().DurationP(EndTimeFlag, "", 0, "end time of interval with anomalies")
	cmd.Flags().DurationP(HistoryFlag, "d", Day, "length of metric data history preceding the anomaly")

	return cmd
}

func RunCollectFeedbackCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	startTime, err := cmd.Flags().GetDuration(StartTimeFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	filePath, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	endTime, err := cmd.Flags().GetDuration(EndTimeFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	history, err := cmd.Flags().GetDuration(HistoryFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	now := time.Now().UnixMilli()
	feedback, resp, err := api.AnomalyFeedbackApi.CollectAnomalyFeedback(cli.Context).
		StartTime(now + startTime.Milliseconds()).
		EndTime(now + endTime.Milliseconds()).
		History(history.Milliseconds()).
		Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	const fileMode = 0644
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.FileMode(fileMode))
	if err != nil {
		return common.NewWriteFileError(err, filePath)
	}
	defer file.Close()

	data, err := json.Marshal(feedback)
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if _, err = file.Write(data); err != nil {
		return common.NewWriteFileError(err, filePath)
	}

	cli.Printer.Success(fmt.Sprintf("Downloaded %d anomalies with feedback.", len(feedback)))

	return nil
}
