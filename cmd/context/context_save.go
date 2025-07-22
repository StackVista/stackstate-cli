package context

import (
	"encoding/base64"
	"os"

	"github.com/spf13/cobra"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/di"
)

const (
	APIPathFlag = "api-path"
)

type SaveArgs struct {
	Name             string
	URL              string
	APIToken         string
	ServiceToken     string
	APIPath          string
	CaCertPath       string
	CaCertBase64Data string
	SkipValidate     bool
	SkipSSLFlag      bool
}

func SaveCommand(cli *di.Deps) *cobra.Command {
	args := &SaveArgs{}
	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save a context",
		Long:  "Save a context to the local config file.",
		RunE:  cli.CmdRunE(RunContextSaveCommand(args)),
	}

	common.AddNameFlagVarVal(cmd, &args.Name, "default", "Name of the context")
	cmd.Flags().StringVar(&args.URL, common.URLFlag, "", common.URLFlagUse)
	cmd.Flags().StringVar(&args.APIToken, common.APITokenFlag, "", common.APITokenFlagUse)
	cmd.Flags().StringVar(&args.ServiceToken, common.ServiceTokenFlag, "", common.ServiceTokenFlagUse)
	cmd.Flags().StringVar(&args.APIPath, APIPathFlag, "/api", "Specify the path of the API end-point, e.g. the part that comes after the URL")
	cmd.Flags().BoolVar(&args.SkipValidate, "skip-validate", false, "Skip validation of the context")
	cmd.Flags().StringVar(&args.CaCertPath, common.CaCertPathFlag, "", common.CaCertPathFlagUse)
	cmd.Flags().StringVar(&args.CaCertBase64Data, common.CaCertBase64DataFlag, "", common.CaCertBase64DataFlagUse)
	cmd.Flags().BoolVar(&args.SkipSSLFlag, common.SkipSSLFlag, false, common.SkipSSLFlagUse)
	cmd.MarkFlagRequired(common.URLFlag) //nolint:errcheck
	stscobra.MarkMutexFlags(cmd, []string{common.APITokenFlag, common.ServiceTokenFlag, common.K8sSATokenFlag}, "tokens", true)

	return cmd
}

func RunContextSaveCommand(args *SaveArgs) func(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	return func(cli *di.Deps, cmd *cobra.Command) common.CLIError {
		cfg, err := config.ReadConfig(cli.ConfigPath)
		if err != nil {
			cfg = config.EmptyConfig()
		}

		namedCtx := &config.NamedContext{
			Name: args.Name,
			Context: &config.StsContext{
				URL:          args.URL,
				APIToken:     args.APIToken,
				ServiceToken: args.ServiceToken,
				APIPath:      args.APIPath,
				SkipSSL:      args.SkipSSLFlag,
			},
		}
		// Use private CA only if SkipSSL is not enabled
		if !args.SkipSSLFlag {
			// Providing CA certificate from file takes precedence over providing from the command line argument.
			if args.CaCertPath != "" {
				data, serr := os.ReadFile(args.CaCertPath)
				if serr != nil {
					return common.NewReadFileError(serr, args.CaCertPath)
				}
				namedCtx.Context.CaCertBase64Data = base64.StdEncoding.EncodeToString(data)
				namedCtx.Context.CaCertPath = ""
			} else if args.CaCertBase64Data != "" {
				namedCtx.Context.CaCertBase64Data = args.CaCertBase64Data
			}
		}

		if !args.SkipValidate {
			if _, err := ValidateContext(cli, cmd, namedCtx.Context); err != nil {
				return err
			}
		}

		existingCtx, err := cfg.GetContext(args.Name)
		if err != nil {
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
