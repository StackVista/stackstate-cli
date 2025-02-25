package context

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/client"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

func PrintConnectionSuccess(pr printer.Printer, apiUrl string, serverInfo *stackstate_api.ServerInfo) {
	if serverInfo.PlatformVersion != nil {
		pr.Success(
			fmt.Sprintf("Connection verified to %s (Platform version: %s)",
				apiUrl,
				*serverInfo.PlatformVersion,
			),
		)
	} else {
		// Fallback to serverInfo.Version if platformVersion is not present (an updated client could interact with a server not supporting PlatformVersion yet).
		pr.Success(
			fmt.Sprintf("Connection verified to %s (StackState version: %s)",
				apiUrl,
				client.VersionToString(serverInfo.Version),
			),
		)
	}
}

func ValidateContext(cli *di.Deps, cmd *cobra.Command, cfg *config.StsContext) (*stackstate_api.ServerInfo, common.CLIError) {
	if cli.Client == nil {
		err := cli.LoadClient(cmd, cfg)
		if err != nil {
			return nil, err
		}
	}

	_, serverInfo, err := cli.Client.Connect()
	if err != nil {
		return nil, err
	}

	if !cli.IsJson() {
		PrintConnectionSuccess(cli.Printer, cfg.URL, serverInfo)
	}

	return serverInfo, nil
}
