package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/persistent_flags"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

const (
	TestConnectFlagName = "test-connect"
)

func CliSaveConfigCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save-config --api-url API-URL --api-token API-TOKEN",
		Short: "save CLI configuration to file",
		Example: "# save a new API token to the config file and test the connection - WARNING this will overwrite the saved config\n" +
			`cli save-config --api-token l9x5g14cMcI97IS4785HWgwEpdPr3KJ4 --api-url "https://my.stackstate.com/api" --test-connect`,
		RunE: cli.CmdRunE(RunCliSaveConfig),
	}
	cmd.Flags().Bool(TestConnectFlagName, false, "test the connection to StackState after configuration has been saved to file")

	return cmd
}

func RunCliSaveConfig(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	// get required --api-url and --api-token
	apiUrl, missingApiUrl := cmd.Flags().GetString(persistent_flags.ApiUrlFlag)
	apiToken, missingApiToken := cmd.Flags().GetString(persistent_flags.ApiTokenFlag)
	missing := make([]string, 0)
	if apiUrl == "" || missingApiUrl != nil {
		missing = append(missing, persistent_flags.ApiUrlFlag)
	}
	if apiToken == "" || missingApiToken != nil {
		missing = append(missing, persistent_flags.ApiTokenFlag)
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
