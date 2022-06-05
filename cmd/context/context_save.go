package context

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

type SaveArgs struct {
	Name string
}

func SaveCommand(deps *di.Deps) *cobra.Command {
	args := &SaveArgs{}
	cmd := &cobra.Command{
		Use:   "save",
		Short: "save a context",
		Long:  "Save a context to the local config file.",
		RunE:  deps.CmdRunEWithApi(RunContextSaveCommand(args)),
	}

	common.AddNameFlagVarVal(cmd, &args.Name, "default", "name of the context")

	return cmd
}

func RunContextSaveCommand(args *SaveArgs) di.CmdWithApiFn {
	return func(c *cobra.Command, d *di.Deps, a *stackstate_api.APIClient, si *stackstate_api.ServerInfo) common.CLIError {
		return nil
	}
}
