package monitor

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/settings"
	sts "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

type DescribeArgs struct {
	ID       int64
	FilePath string
}

func MonitorDescribeCommand(cli *di.Deps) *cobra.Command {
	args := &DescribeArgs{}
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe a monitor in STJ format",
		Long:  "Describe a monitor in StackState Templated JSON.",
		RunE:  cli.CmdRunEWithApi(RunMonitorDescribeCommand(args)),
	}

	common.AddRequiredIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddFileFlagVar(cmd, &args.FilePath, "Path to the output file")

	return cmd
}

func RunMonitorDescribeCommand(args *DescribeArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *sts.APIClient, serverInfo *sts.ServerInfo) common.CLIError {
		data, resp, err := settings.DoExport(cli.Context, api, []int64{args.ID}, "", []string{}, []string{})
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if args.FilePath != "" {
			if err := util.WriteFile(args.FilePath, []byte(data)); err != nil {
				return common.NewWriteFileError(err, args.FilePath)
			}

			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"filepath": args.FilePath,
				})
			} else {
				cli.Printer.Success(fmt.Sprintf("settings exported to: %s", args.FilePath))
			}

			return nil
		} else {
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"data":   data,
					"format": "stj",
				})
			} else {
				cli.Printer.PrintLn(data)
			}
			return nil
		}
	}
}
