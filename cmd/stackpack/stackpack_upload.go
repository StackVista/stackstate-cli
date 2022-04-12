package stackpack

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

const (
	FileFlag = "file"
)

func StackpackUploadCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload -f FILE",
		Short: "upload a stackpack",
		RunE:  cli.CmdRunEWithApi(RunStackpackUploadCommand),
	}
	cmd.Flags().StringP(FileFlag, "f", "", "stackpack to upload (.sts file")
	cmd.MarkFlagRequired(FileFlag) //nolint:errcheck

	return cmd
}

func RunStackpackUploadCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_client.APIClient,
	serverInfo stackstate_client.ServerInfo,
) common.CLIError {
	filePath, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return common.NewCLIError(err)
	}
	defer file.Close()

	stackpack, resp, err := api.StackpackApi.StackpackUpload(cli.Context).StackPack(file).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.Success(fmt.Sprintf("uploaded StackPack: %s", filePath))
	cli.Printer.Table(
		[]string{"name", "display name", "version"},
		[][]interface{}{{stackpack.Name, stackpack.DisplayName, stackpack.Version}},
		stackpack,
	)

	return nil
}
