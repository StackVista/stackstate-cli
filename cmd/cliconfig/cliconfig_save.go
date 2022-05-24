package cliconfig

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
	ApiPathFlag          = "api-path"
)

func SaveCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save",
		Short: "save CLI configuration",
		Long:  "Save the configuration of this CLI to disk.",
		Example: "# save a new API token to the config file" +
			`sts cli-config save --url "https://my.stackstate.com --api-token l9x5g14cMcI97IS4785HWgwEpdPr3KJ4"`,
		RunE: cli.CmdRunE(RunSave),
	}
	cmd.Flags().Bool(
		SkipValidateFlagName,
		false,
		"skip validating the connection before saving",
	)
	cmd.Flags().String(ApiPathFlag, "/api", "specify the path of the API end-point, e.g. the part that comes after the URL")

	return cmd
}

func RunSave(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	// get required --url and --api-token
	URL, missingApiURL := cmd.Flags().GetString(common.URLFlag)
	apiToken, missingApiToken := cmd.Flags().GetString(common.APITokenFlag)
	missing := make([]string, 0)
	if URL == "" || missingApiURL != nil {
		missing = append(missing, common.URLFlag)
	}
	if apiToken == "" || missingApiToken != nil {
		missing = append(missing, common.APITokenFlag)
	}
	if len(missing) > 0 {
		return common.NewCLIArgParseError(fmt.Errorf("missing required flag(s): %v", strings.Join(missing, ", ")))
	}

	apiPath, err := cmd.Flags().GetString(ApiPathFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	// set config
	cli.Config = &conf.Conf{
		URL:      URL,
		ApiToken: apiToken,
		ApiPath:  apiPath,
	}

	// get test-connect flag
	skipValidate, err := cmd.Flags().GetBool(SkipValidateFlagName)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	// test connect
	if !skipValidate {
		if cli.Client == nil {
			err := cli.LoadClient(cmd, URL, apiPath, apiToken)
			if err != nil {
				return err
			}
		}

		_, serverInfo, err := cli.Client.Connect()
		if err != nil {
			return err
		}

		if !cli.IsJson() {
			PrintConnectionSuccess(cli.Printer, URL, serverInfo)
		}
	}

	// write config
	filename, err := conf.WriteConf(*cli.Config)
	if err != nil {
		return common.NewReadFileError(err, filename)
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"connection-tested": !skipValidate,
			"config-file":       filename,
		})
	} else {
		cli.Printer.Success("Config saved to: " + filename)
	}

	return nil
}
