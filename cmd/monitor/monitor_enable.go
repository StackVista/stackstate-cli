package monitor

import (
	"github.com/spf13/cobra"
	sts "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	stscobra "gitlab.com/stackvista/stackstate-cli2/internal/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func MonitorEnableCommand(cli *di.Deps) *cobra.Command {
	args := &IdArgs{}
	cmd := &cobra.Command{
		Use:   "disable",
		Short: "Disable a monitor",
		Long:  "Disable a monitor.",
		RunE:  cli.CmdRunEWithApi(RunMonitorEnableDisableCommand(args, sts.MONITORSTATUSVALUE_ENABLED.Ptr(), "enabled")),
	}

	common.AddIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddIdentifierFlagVar(cmd, &args.Identifier, IdentifierFlagUsage)
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}
