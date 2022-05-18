package stackpack

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

type UploadArgs struct {
	FilePath string
}

func StackpackUploadCommand(cli *di.Deps) *cobra.Command {
	args := &UploadArgs{}
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "upload a stackpack",
		Long:  "Upload a StackPack file to StackState.",
		RunE:  cli.CmdRunEWithApi(RunStackpackUploadCommand(args)),
	}
	common.AddRequiredFileFlagVar(cmd, &args.FilePath, "stackpack file to upload (.sts file)")

	return cmd
}

func RunStackpackUploadCommand(args *UploadArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		file, err := os.Open(args.FilePath)
		if err != nil {
			return common.NewCLIArgParseError(err)
		}
		defer file.Close()

		stackpack, resp, err := api.StackpackApi.StackpackUpload(cli.Context).StackPack(file).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson {
			cli.Printer.PrintJson(map[string]interface{}{
				"uploaded-stackpack": stackpack,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("uploaded StackPack: %s", args.FilePath))
			cli.Printer.Table(printer.TableData{
				Header: []string{"name", "display name", "version"},
				Data:   [][]interface{}{{stackpack.Name, stackpack.DisplayName, stackpack.Version}},
			})
		}

		return nil
	}
}
