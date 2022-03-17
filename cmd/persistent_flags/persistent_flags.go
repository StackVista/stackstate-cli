package persistent_flags

import (
	"github.com/spf13/cobra"
)

const (
	VerboseFlag  = "verbose"
	ApiUrlFlag   = "api-url"
	ApiTokenFlag = "api-token"
	NoColorFlag  = "no-color"
	OutputFlag   = "output"
)

func AddPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(ApiUrlFlag, "", "specify the StackState API URL")
	cmd.PersistentFlags().String(ApiTokenFlag, "", "specify the StackState API token")
	cmd.PersistentFlags().Bool(VerboseFlag, false, "print verbose logging to the terminal to track what the CLI is doing")
	cmd.PersistentFlags().Bool(NoColorFlag, false, "disable color when printing to the terminal")
	cmd.PersistentFlags().StringP(OutputFlag, "o", "Auto", "format output as: JSON, YAML or Auto")
}
