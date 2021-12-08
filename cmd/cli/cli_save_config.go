package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

const (
	TestConnectFlagName = "test-connect"
)

func CliSaveConfigCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save-config --api-url <api-url> --api-token <api-token>",
		Short: "Save config to file.",
		RunE:  di.CmdRunEWithDeps(cli, RunCliSaveConfig),
	}
	cmd.Flags().BoolP(TestConnectFlagName, "t", false, "Test connection to StackState after the config file has been saved.")

	return cmd
}

func RunCliSaveConfig(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	// get required --api-url and --api-token
	apiUrl, missingApiUrl := cmd.Flags().GetString(common.ApiUrlFlag)
	apiToken, missingApiToken := cmd.Flags().GetString(common.ApiTokenFlag)
	missing := make([]string, 0)
	if apiUrl == "" || missingApiUrl != nil {
		missing = append(missing, common.ApiUrlFlag)
	}
	if apiToken == "" || missingApiToken != nil {
		missing = append(missing, common.ApiTokenFlag)
	}
	if len(missing) > 0 {
		return common.NewCLIArgParseError(fmt.Errorf("missing required flag(s): %v", strings.Join(missing, ", ")))
	}

	// get test-connect flag
	testConnect, err := cmd.Flags().GetBool(TestConnectFlagName)
	if err != nil {
		return common.NewCLIError(err)
	}

	// write config
	filename, err := conf.WriteConf(conf.Conf{
		ApiUrl:   apiUrl,
		ApiToken: apiToken,
	})
	if err != nil {
		return common.NewCLIError(err)
	}
	cli.Printer.Success("Config saved to: " + filename)

	// test connect
	if testConnect {
		return testConect(cli)
	}

	return nil
}
