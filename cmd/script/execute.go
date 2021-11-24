package script

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func ScriptExecuteCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute <script>",
		Args:  cobra.ExactArgs(1),
		Short: "Execute a STSL script.",
		RunE:  RunCmdWithError(cfg, RunScriptExecuteCommand),
	}

	cmd.Flags().StringP("arguments-script", "a", "", "A script that returns a java.util.Map with arguments that can be used as variables within the actual script.")
	cmd.Flags().IntP("timeout", "t", 0, "Timeout in milli-seconds.")

	return cmd
}

func RunCmdWithError(cfg *config.Config, runFn func(*config.Config, *cobra.Command, []string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return runFn(cfg, cmd, args)
	}
}

func RunScriptExecuteCommand(cfg *config.Config, cmd *cobra.Command, args []string) error {
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
	verbose, _ := cmd.Flags().GetCount("verbose")

	if verbose > 0 {
		fmt.Printf("Config: %+v\n", cfg)
	}

	configuration := stackstate_client.NewConfiguration()
	configuration.Servers[0] = stackstate_client.ServerConfiguration{
		URL:         cfg.URL,
		Description: "",
		Variables:   nil,
	}

	auth := make(map[string]stackstate_client.APIKey)
	auth["ApiToken"] = stackstate_client.APIKey{
		Key:    cfg.ApiToken,
		Prefix: "",
	}
	ctx := context.WithValue(
		cmd.Context(),
		stackstate_client.ContextAPIKeys,
		auth,
	)

	client := stackstate_client.NewAPIClient(configuration)
	scriptExecute := client.ScriptingApi.ScriptExecute(ctx)
	scriptRequest := stackstate_client.ExecuteScriptRequest{
		TimeoutMs:       timeoutMs,
		Script:          args[0],
		ArgumentsScript: argumentsScript,
	}

	if verbose > 0 {
		scriptRequestStr, _ := json.Marshal(scriptRequest)
		fmt.Println("Executing script request:", string(scriptRequestStr))
	}

	scriptResponse, resp, err := scriptExecute.ExecuteScriptRequest(scriptRequest).Execute()

	if err != nil {
		var status string
		if resp != nil {
			status = resp.Status + ". "
		}

		switch v := err.(type) {
		case stackstate_client.GenericOpenAPIError:
			return fmt.Errorf("%vError response: %+v", status, string(v.Body()))
		default:
			return fmt.Errorf("%v%+v", status, v)
		}
	}

	fmt.Println(scriptResponse["result"])
	return nil
}
