package common

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
)

const (
	VerboseFlag         = "verbose"
	VerboseFlagShort    = "v"
	URLFlag             = "url"
	URLFlagUse          = "specify the URL of the StackState server"
	APITokenFlag        = "api-token"
	APITokenFlagUse     = "specify the API token of the StackState server"
	ServiceTokenFlag    = "service-token"
	ServiceTokenFlagUse = "specify the Service token of the StackState server"
	NoColorFlag         = "no-color"
	OutputFlag          = "output"
	OutputFlagShort     = "o"
	ConfigFlag          = "config"
	ContextFlag         = "context"
	ContextFlagShort    = "c"
)

var AllowedOutputs = []string{JSONOutput.String(), TextOutput.String()}

func AddPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(URLFlag, "", URLFlagUse)
	cmd.PersistentFlags().String(APITokenFlag, "", APITokenFlagUse)
	cmd.PersistentFlags().String(ServiceTokenFlag, "", ServiceTokenFlagUse)
	cmd.PersistentFlags().CountP(VerboseFlag, VerboseFlagShort, "print verbose logging to the terminal to track what the CLI is doing (use multiple times to increase verbosity)")
	cmd.PersistentFlags().Bool(NoColorFlag, false, "disable color when printing to the terminal")
	cmd.PersistentFlags().String(ConfigFlag, "", "override the path to the config file")
	cmd.PersistentFlags().StringP(ContextFlag, ContextFlagShort, "", "override the context to use")
	pflags.EnumP(cmd.PersistentFlags(), OutputFlag, OutputFlagShort, "text", AllowedOutputs, fmt.Sprintf("specify the output format (must be { %s })", strings.Join(AllowedOutputs, " | ")))
}
