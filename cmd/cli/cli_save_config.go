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
	SkipValidateFlagName = "skip-validate"
)

func CliSaveConfigCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save-config --api-url API-URL --api-token API-TOKEN",
		Short: "save CLI configuration",
		Long:  "Save the configuration of this CLI to disk.",
		Example: "# save a new API token to the config file" +
			`sts cli save-config --api-token l9x5g14cMcI97IS4785HWgwEpdPr3KJ4 --api-url "https://my.stackstate.com/api"`,
		RunE: cli.CmdRunE(RunCliSaveConfig),
	}
	cmd.Flags().Bool(
		SkipValidateFlagName,
		false,
		"skip validating the connection before saving",
	)

	return cmd
}

func RunCliSaveConfig(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	// get required --api-url and --api-token
	apiURL, missingApiURL := cmd.Flags().GetString(common.APIURLFlag)
	apiToken, missingApiToken := cmd.Flags().GetString(common.APITokenFlag)
	missing := make([]string, 0)
	if apiURL == "" || missingApiURL != nil {
		missing = append(missing, common.APIURLFlag)
	}
	if apiToken == "" || missingApiToken != nil {
		missing = append(missing, common.APITokenFlag)
	}
	if len(missing) > 0 {
		return common.NewCLIArgParseError(fmt.Errorf("missing required flag(s): %v", strings.Join(missing, ", ")))
	}

	// this is really only necessary for testing
	// in production this should've already been handled by Viper
	cli.Config.ApiURL = apiURL
	cli.Config.ApiToken = apiToken

	// get test-connect flag
	skipValidate, err := cmd.Flags().GetBool(SkipValidateFlagName)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	// test connect
	if !skipValidate {
		err := testConect(cli)
		if err != nil {
			return err
		}
	}

	// write config
	filename, err := conf.WriteConf(*cli.Config)
	if err != nil {
		return common.NewCLIError(err, nil)
	}
	cli.Printer.Success("Config saved to: " + filename)

	return nil
}
