package graph

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func RetentionCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "retention",
		Short: "Manage the StackState Graph data retention.",
		Long:  "Manage the StackState Graph data retention.",
		RunE:  deps.CmdRunEWithApi(runRetentionCommand),
	}

	return cmd
}

func runRetentionCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {

	// TODO
	window := 691200000
	epoch := 1663493504140000000

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"window": window,
			"epoch": epoch,
		})
	} else {
		cli.Printer.Success(fmt.Sprintf("Retention window: %d milliseconds", window))
		cli.Printer.Success(fmt.Sprintf("Epoch transactionId: %d", epoch))
	}

	return nil
}
