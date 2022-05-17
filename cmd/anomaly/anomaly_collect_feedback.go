package anomaly

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/time_flags"
)

const (
	StartTimeFlag = "start-time"
	EndTimeFlag   = "end-time"
	HistoryFlag   = "history"
)

func AnomalyCollectFeedback(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collect-feedback",
		Short: "collect anomalies that have user feedback",
		Long:  "Collect anomalies that users have given feedback on, as thumbs-up/-down and/or comments.",
		Example: "# Collect anomalies with feedback in the last 8 days, include 2 days of metric data for each anomaly" +
			`sts anomaly collect --start-time -8d --history 2d --file anomaly-feedback.json`,
		RunE: cli.CmdRunEWithApi(RunCollectFeedbackCommand),
	}
	cmd.Flags().String(StartTimeFlag, "-7d", "start time of interval with anomalies.  Format is ISO8601, milliseconds since epoch or relative (-12h)")
	cmd.MarkFlagRequired(StartTimeFlag) //nolint:errcheck
	common.AddRequiredFileFlag(cmd, "path to output file")
	cmd.Flags().String(EndTimeFlag, "-0h", "end time of interval with anomalies")
	cmd.Flags().String(HistoryFlag, "1d", "length of metric data history preceding the anomaly")

	return cmd
}

func RunCollectFeedbackCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	clock := time_flags.NewFixedTimeClock()

	startTime, cliErr := time_flags.GetTimeFlag(cmd, StartTimeFlag, clock)
	if cliErr != nil {
		return cliErr
	}

	endTime, cliErr := time_flags.GetTimeFlag(cmd, EndTimeFlag, clock)
	if cliErr != nil {
		return cliErr
	}

	history, cliErr := time_flags.GetDurationFlag(cmd, HistoryFlag)
	if cliErr != nil {
		return cliErr
	}

	filePath, err := cmd.Flags().GetString(common.FileFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	anomalies, resp, err := api.ExportAnomalyApi.ExportAnomaly(cli.Context).
		Feedback("present").
		StartTime(startTime.UnixMilli()).
		EndTime(endTime.UnixMilli()).
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

	data, err := json.Marshal(anomalies)
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if _, err = file.Write(data); err != nil {
		return common.NewWriteFileError(err, filePath)
	}

	if cli.IsJson {
		cli.Printer.PrintJson(map[string]interface{}{
			"collected-anomalies": len(anomalies),
			"filepath":            filePath,
		})
	} else {
		cli.Printer.Success(fmt.Sprintf("Collected %d anomalies to %s.", len(anomalies), filePath))
	}

	return nil
}
