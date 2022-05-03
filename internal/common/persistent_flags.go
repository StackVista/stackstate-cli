package common

import (
	"github.com/spf13/cobra"
)

const (
	VerboseFlag  = "verbose"
	URLFlag      = "url"
	APITokenFlag = "api-token"
	NoColorFlag  = "no-color"
	JsonFlag     = "json"
)

func AddPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(URLFlag, "", "specify the URL of the StackState server")
	cmd.PersistentFlags().String(APITokenFlag, "", "specify the API token of the StackState server")
	cmd.PersistentFlags().Bool(VerboseFlag, false, "print verbose logging to the terminal to track what the CLI is doing")
	cmd.PersistentFlags().Bool(NoColorFlag, false, "disable color when printing to the terminal")
	cmd.PersistentFlags().Bool(JsonFlag, false, "print as json")
}
