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
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
)

const (
	StartTimeFlag = "start-time"
	EndTimeFlag   = "end-time"
	HistoryFlag   = "history"
)

type CollectFeedbackArgs struct {
	StartTime time.Time
	EndTime   time.Time
	History   time.Duration
	File      string
}

func AnomalyCollectFeedback(cli *di.Deps) *cobra.Command {
	args := &CollectFeedbackArgs{}
	cmd := &cobra.Command{
		Use:   "collect-feedback",
		Short: "collect anomalies that have user feedback",
		Long:  "Collect anomalies that users have given feedback on, as thumbs-up/-down and/or comments.",
		Example: "# Collect anomalies with feedback in the last 8 days, include 2 days of metric data for each anomaly" +
			`sts anomaly collect --start-time -8d --history 2d --file anomaly-feedback.json`,
		RunE: cli.CmdRunEWithApi(RunCollectFeedbackCommand(args)),
	}
	pflags.RelativeTimeVar(cmd.Flags(), &args.StartTime, StartTimeFlag, "-7d", cli.Clock, "start time of interval with anomalies.  Format is ISO8601, milliseconds since epoch or relative (-12h)")
	cmd.MarkFlagRequired(StartTimeFlag) //nolint:errcheck
	common.AddRequiredFileFlagVar(cmd, &args.File, "path to output file")
	pflags.RelativeTimeVar(cmd.Flags(), &args.EndTime, EndTimeFlag, "-0h", cli.Clock, "end time of interval with anomalies")
	pflags.DurationVar(cmd.Flags(), &args.History, HistoryFlag, pflags.Day, "length of metric data history preceding the anomaly")

	return cmd
}

func RunCollectFeedbackCommand(args *CollectFeedbackArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		anomalies, resp, err := api.ExportAnomalyApi.ExportAnomaly(cli.Context).
			Feedback("present").
			StartTime(args.StartTime.UnixMilli()).
			EndTime(args.EndTime.UnixMilli()).
			History(args.History.Milliseconds()).
			Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		const fileMode = 0644
		file, err := os.OpenFile(args.File, os.O_CREATE|os.O_WRONLY, os.FileMode(fileMode))
		if err != nil {
			return common.NewWriteFileError(err, args.File)
		}
		defer file.Close()

		data, err := json.Marshal(anomalies)
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if _, err = file.Write(data); err != nil {
			return common.NewWriteFileError(err, args.File)
		}

		if cli.IsJson {
			cli.Printer.PrintJson(map[string]interface{}{
				"collected-anomalies": len(anomalies),
				"filepath":            args.File,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Collected %d anomalies to %s.", len(anomalies), args.File))
		}

		return nil
	}
}
