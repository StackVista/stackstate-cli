package context

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/mutex_flags"
)

const (
	APIPathFlag = "api-path"
)

type SaveArgs struct {
	Name         string
	URL          string
	APIToken     string
	ServiceToken string
	APIPath      string
	SkipValidate bool
}

func SaveCommand(cli *di.Deps) *cobra.Command {
	args := &SaveArgs{}
	cmd := &cobra.Command{
		Use:   "save",
		Short: "save a context",
		Long:  "Save a context to the local config file.",
		RunE:  cli.CmdRunE(RunContextSaveCommand(args)),
	}

	common.AddNameFlagVarVal(cmd, &args.Name, "default", "name of the context")
	cmd.Flags().StringVar(&args.URL, common.URLFlag, "", common.URLFlagUse)
	cmd.Flags().StringVar(&args.APIToken, common.APITokenFlag, "", common.APITokenFlagUse)
	cmd.Flags().StringVar(&args.ServiceToken, common.ServiceTokenFlag, "", common.ServiceTokenFlagUse)
	cmd.Flags().StringVar(&args.APIPath, APIPathFlag, "/api", "specify the path of the API end-point, e.g. the part that comes after the URL")
	cmd.Flags().BoolVar(&args.SkipValidate, "skip-validate", false, "skip validation of the context")

	cmd.MarkFlagRequired(common.URLFlag) //nolint:errcheck
	mutex_flags.MarkMutexFlags(cmd, []string{common.APITokenFlag, common.ServiceTokenFlag}, "tokens", true)

	return cmd
}

func RunContextSaveCommand(args *SaveArgs) func(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	return func(cli *di.Deps, cmd *cobra.Command) common.CLIError {
		cfg := cli.StsConfig

		namedCtx := &config.NamedContext{
			Name: args.Name,
			Context: &config.StsContext{
				URL:          args.URL,
				APIToken:     args.APIToken,
				ServiceToken: args.ServiceToken,
				APIPath:      args.APIPath,
			},
		}

		if !args.SkipValidate {
			if err := ValidateContext(cli, cmd, namedCtx.Context); err != nil {
				return err
			}
		}

		existingCtx := cfg.GetContext(args.Name)
		if existingCtx == nil {
			cfg.Contexts = append(cfg.Contexts, namedCtx)
		} else {
			existingCtx.Context = namedCtx.Context
		}

		cfg.CurrentContext = args.Name

		if err := config.WriteConfig(cli.ConfigPath, cfg); err != nil {
			return common.NewWriteFileError(err, cli.ConfigPath)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"current-context": cfg.CurrentContext,
				"context":         namedCtx.Context,
			})
		} else {
			cli.Printer.Successf("Saved context: '%s'", args.Name)
		}

		return nil
	}
}
