package cli

import (
	"fmt"

	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func testConect(cli *di.Deps) common.CLIError {
	_, serverInfo, err := cli.Client.Connect()
	if err != nil {
		return common.NewConnectError(err)
	}

	dev := ""
	if serverInfo.Version.IsDev {
		dev = "-dev"
	}
	msg := fmt.Sprintf("Connection verified to %s (StackState version: %d.%d.%d+%s-%s%s)",
		cli.Config.ApiURL,
		serverInfo.Version.Major,
		serverInfo.Version.Minor,
		serverInfo.Version.Patch,
		serverInfo.Version.Diff,
		serverInfo.Version.Commit,
		dev)
	cli.Printer.Success(msg)
	return nil
}
