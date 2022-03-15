package cli

import (
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func testConect(cli *di.Deps) common.CLIError {
	cli.Printer.StartSpinner(common.AwaitingServer)
	_, resp, err := cli.Client.UserProfileApi.GetCurrentUserProfile(cli.Context).Execute()
	cli.Printer.StopSpinner()
	if err != nil {
		return common.NewResponseError(err, resp)
	}
	cli.Printer.Success("Connection verified to: " + cli.Config.ApiUrl)
	return nil
}
