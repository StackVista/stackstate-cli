package license

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func ShowCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show the license information",
		Long:  "Show the license information.",
		RunE:  deps.CmdRunEWithApi(runLicenseShowCommand),
	}
	return cmd
}

func runLicenseShowCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
	licenseAPI := api.SubscriptionApi.GetSubscription(cli.Context)
	license, resp, err := licenseAPI.Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"subscription": license,
		})
	} else {
		cli.Printer.Table(subscriptionAsTable(license))
	}

	return nil
}

func expiration(ts *int64) string {
	if ts == nil {
		return ""
	}
	return time.UnixMilli(*ts).Format("2006-01-02")
}
