package monitor

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/settings"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/util"
)

type DescribeArgs struct {
	ID       int64
	FilePath string
}

func MonitorDescribeCommand(cli *di.Deps) *cobra.Command {
	args := &DescribeArgs{}
	cmd := &cobra.Command{
		Use:   "describe --id ID",
		Short: "Export a monitor definition in STY format",
		Long:  "Export the monitor definition in STY (SUSE Observability YAML) format. Output can be printed to stdout or saved to a file for backup or migration.",
		Example: `# describe a monitor to stdout
sts monitor describe --id 123456789

# export a monitor to a file
sts monitor describe --id 123456789 --file my-monitor.sty`,
		RunE: cli.CmdRunEWithApi(RunMonitorDescribeCommand(args)),
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
					"format": "sty",
				})
			} else {
				cli.Printer.PrintLn(data)
			}
			return nil
		}
	}
}
