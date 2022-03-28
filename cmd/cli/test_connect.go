package cli

import (
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func testConect(cli *di.Deps) common.CLIError {
	_, _, err := cli.Client.Connect()
	if err != nil {
		return common.NewConnectError(err)
	}

	cli.Printer.Success("Connection verified to: " + cli.Config.ApiUrl)
	return nil
}
