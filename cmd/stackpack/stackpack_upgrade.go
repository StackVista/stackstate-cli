package stackpack

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func StackpacUpgradeCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "upgrade StackPack instance",
		Long:  "Upgrade StackPack instance if nextVersion is available",
		RunE:  cli.CmdRunEWithApi(RunStackpackUpgradeCommand),
	}
	cmd.Flags().String(NameFlag, "", "name of the instance")
	cmd.MarkFlagRequired(NameFlag) //nolint:errcheck
	return cmd
}

func RunStackpackUpgradeCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo stackstate_api.ServerInfo,
) common.CLIError {
	name, err := cmd.Flags().GetString(NameFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	stackpackList, resp, err := api.StackpackApi.StackpackList(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}
	stackPack, err := filterStackpack(name, stackpackList)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	if !stackPack.HasNextVersion() {
		return common.NewCLIArgParseError(fmt.Errorf("StackPack %s cannot be upgraded at this moment", name))
	}

	return nil
}

func filterStackpack(name string, stackpackList []stackstate_api.Sstackpack) (*stackstate_api.Sstackpack, error) {
	for _, v := range stackpackList {
		if v.GetName() == name {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("StackPack %s does not exist", name)
}
