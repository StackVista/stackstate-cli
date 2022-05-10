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
	FileFlag      = "file"
)

func AnomalyCollect(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collect",
		Short: "collect anomalies that have user feedback",
		Long:  "Fetch anomalies that users have given feedback on, as thumbs-up/-down and/or comments",
		Example: "# Export anomalies with feedback in the last 7 days, include 1 day of metric data for each anomaly" +
			`sts anomaly collect --start-time=-7d --file anomaly-feedback.json"`,
		RunE: cli.CmdRunEWithApi(RunCollectFeedbackCommand),
	}
	cmd.Flags().StringP(StartTimeFlag, "", "-7d", "start time of interval with anomalies.  Format is ISO8601, milliseconds since epoch or relative (-12h)")
	cmd.MarkFlagRequired(StartTimeFlag) //nolint:errcheck
	cmd.Flags().StringP(FileFlag, "f", "", "path to output file")
	cmd.MarkFlagRequired(FileFlag) //nolint:errcheck
	cmd.Flags().StringP(EndTimeFlag, "", "-0h", "end time of interval with anomalies")
	cmd.Flags().StringP(HistoryFlag, "d", "1d", "length of metric data history preceding the anomaly")

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

	filePath, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	feedback, resp, err := api.AnomalyFeedbackApi.CollectAnomalyFeedback(cli.Context).
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

	data, err := json.Marshal(feedback)
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if _, err = file.Write(data); err != nil {
		return common.NewWriteFileError(err, filePath)
	}

	if cli.IsJson {
		cli.Printer.PrintJson(map[string]interface{}{
			"downloaded-anomalies": len(feedback),
		})
	} else {
		cli.Printer.Success(fmt.Sprintf("Downloaded %d anomalies with feedback.", len(feedback)))
	}

	return nil
}
