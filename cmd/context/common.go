package context

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/client"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
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

func ValidateContext(cli *di.Deps, cmd *cobra.Command, cfg *config.StsContext) (*stackstate_api.ServerInfo, common.CLIError) {
	if cli.Client == nil {
		err := cli.LoadClient(cmd, cfg.URL, cfg.APIPath, cfg.APIToken, cfg.ServiceToken)
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
