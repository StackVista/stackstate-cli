package cli

import (
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func testConect(cli *di.Deps) common.CLIError {
	cli.Printer.StartSpinner(common.AwaitingServer)
	defer cli.Printer.StopSpinner()
	client, ctx := cli.NewStsClient()
	_, resp, err := client.UserProfileApi.GetCurrentUserProfile(ctx).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}
	cli.Printer.Success("Connection verified to: " + cli.Config.ApiUrl)
	return nil
}
