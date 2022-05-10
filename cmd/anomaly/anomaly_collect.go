package anomaly

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
		Use:   "collect",
		Short: "collect anomalies that have user feedback",
		Long:  "Fetch anomalies that users have given feedback on, as thumbs-up/-down and/or comments",
		Example: "# Export anomalies with feedback in the last 7 days, include 1 day of metric data for each anomaly" +
			`sts anomaly collect --start-time=-168h --file anomaly-feedback.json"`,
		RunE: cli.CmdRunEWithApi(RunCollectFeedbackCommand),
	}
	cmd.Flags().StringP(StartTimeFlag, "", "-7d", "start time of interval with anomalies.  Format is ISO8601, milliseconds since epoch or relative (-12h)")
	cmd.MarkFlagRequired(StartTimeFlag) //nolint:errcheck
	cmd.Flags().StringP(FileFlag, "f", "", "path to output file")
	cmd.MarkFlagRequired(FileFlag) //nolint:errcheck
	cmd.Flags().StringP(EndTimeFlag, "", "-0h", "end time of interval with anomalies")
	cmd.Flags().DurationP(HistoryFlag, "d", Day, "length of metric data history preceding the anomaly")

	return cmd
}

type Clock interface {
	Now() time.Time
}

type StdClock struct{}

func newClock() Clock {
	return &StdClock{}
}

func (clock *StdClock) Now() time.Time {
	return time.Now()
}

func NewTimeParseError(value string) common.CLIError {
	return common.StdCLIError{
		Err: fmt.Errorf("invalid time format %s - expected ISO8601 (2016-09-13T13:12:04Z), milliseconds since epoch or relative time (-168h) ", value),
	}
}

func parseTime(clock Clock, strStartTime string) (time.Time, common.CLIError) {
	var startTime time.Time
	startTimeInt, err := strconv.ParseInt(strStartTime, 0, 64) //nolint:gomnd
	if err != nil {
		startTime, err = time.Parse(time.RFC3339, strStartTime)
		if err != nil {
			duration, err := time.ParseDuration(strStartTime)
			if err != nil {
				return time.Time{}, NewTimeParseError(strStartTime)
			}
			startTime = clock.Now().Add(duration)
		}
	} else {
		startTime = time.UnixMilli(startTimeInt)
	}
	return startTime, nil
}

func RunCollectFeedbackCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	clock := newClock()

	strStartTime, err := cmd.Flags().GetString(StartTimeFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	startTime, parseErr := parseTime(clock, strStartTime)
	if parseErr != nil {
		return parseErr
	}

	filePath, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	strEndTime, err := cmd.Flags().GetString(EndTimeFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	endTime, parseErr := parseTime(clock, strEndTime)
	if parseErr != nil {
		return parseErr
	}

	history, err := cmd.Flags().GetDuration(HistoryFlag)
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
