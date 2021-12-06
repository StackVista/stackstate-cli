package script

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func ScriptExecuteCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute <script>",
		Args:  cobra.ExactArgs(1),
		Short: "Execute a STSL script.",
		RunE:  di.CmdRunEWithDeps(cli, RunScriptExecuteCommand),
	}

	cmd.Flags().StringP("arguments-script", "a", "", "A script that returns a java.util.Map with arguments that can be used as variables within the actual script.")
	cmd.Flags().IntP("timeout", "t", 0, "Timeout in milli-seconds.")

	return cmd
}

func RunScriptExecuteCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	var argumentsScript *string
	a, _ := cmd.Flags().GetString("arguments-script")
	if a != "" {
		argumentsScript = &a
	}
	var timeoutMs *int32
	t, _ := cmd.Flags().GetInt("timeout")
	if t != 0 {
		t32 := int32(t)
		timeoutMs = &t32
	}

	scriptRequest := sts.ExecuteScriptRequest{
		TimeoutMs:       timeoutMs,
		Script:          args[0],
		ArgumentsScript: argumentsScript,
	}

	cli.Printer.StartSpinner(common.AwaitingServer)
	defer cli.Printer.StopSpinner()
	scriptResponse, resp, err := cli.Client.ScriptingApi.
		ScriptExecute(cli.Context).
		ExecuteScriptRequest(scriptRequest).
		Execute()

	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.PrintStruct(scriptResponse["result"])

	return nil
}
