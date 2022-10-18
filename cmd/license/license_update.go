package license

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type UpdateArgs struct {
	Key string
}

func UpdateCommand(deps *di.Deps) *cobra.Command {
	args := &UpdateArgs{}
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update the license",
		Long:  "Update the license.",
		RunE:  deps.CmdRunEWithApi(runLicenseUpdateCommand(args)),
	}

	cmd.Flags().StringVar(&args.Key, "key", "", "The license key")
	cmd.MarkFlagRequired("key") //nolint:errcheck

	return cmd
}

func runLicenseUpdateCommand(args *UpdateArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		subscriptionAPI := api.SubscriptionApi.PostSubscription(cli.Context)
		l := stackstate_api.NewNewLicense(args.Key)
		subscriptionAPI = subscriptionAPI.NewLicense(*l)

		license, resp, err := subscriptionAPI.Execute()
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
}
