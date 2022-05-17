package cliconfig

import (
	"fmt"

	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/client"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func PrintConnectionSuccess(pr printer.Printer, apiUrl string, serverInfo *stackstate_api.ServerInfo) {
	pr.Success(
		fmt.Sprintf("Connection verified to %s (StackState version: %s)",
			apiUrl,
			client.VersionToString(serverInfo.Version),
		),
	)
}
