package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func ScriptCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "script",
		Short: "StackState scripting related commands.",
	}
	cmd.AddCommand(ScriptExecuteCommand())

	return cmd
}

func ScriptExecuteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute",
		Short: "Execute a STSL script.",
		Run:   RunScriptExecuteCommand,
	}

	return cmd
}

func RunScriptExecuteCommand(cmd *cobra.Command, _ []string) {
	configuration := stackstate_client.NewConfiguration()
	configuration.Servers[0] = stackstate_client.ServerConfiguration{
		URL:         "https://master.preprod.stackstate.io/api/",
		Description: "",
		Variables:   nil,
	}
	configuration.DefaultHeader["Authorization"] = "ApiToken h17MZsqcL2Kbi21_FUOd86Moimc8SZwa"
	client := stackstate_client.NewAPIClient(configuration)
	scriptExecute := client.ScriptingApi.ScriptExecute(cmd.Context())
	e := stackstate_client.ExecuteScriptRequest{
		TimeoutMs:       nil,
		Script:          "return 2+2",
		ArgumentsScript: nil,
	}
	resp, _, err := scriptExecute.ExecuteScriptRequest(e).Execute()
	if err != nil {
		fmt.Println("Got err ", err)
	}

	fmt.Println("Success!", resp)
}
