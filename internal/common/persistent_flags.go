package common

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
)

const (
	VerboseFlag          = "verbose"
	VerboseFlagShort     = "v"
	URLFlag              = "url"
	URLFlagUse           = "Specify the URL of the StackState server"
	APITokenFlag         = "api-token"
	APITokenFlagUse      = "Specify the API token of the StackState server" //nolint:gosec
	ServiceTokenFlag     = "service-token"
	ServiceTokenFlagUse  = "Specify the Service token of the StackState server" //nolint:gosec
	ServiceBearerFlag    = "service-bearer"
	ServiceBearerFlagUse = "Specify the Service bearer of the StackState server" //nolint:gosec
	NoColorFlag          = "no-color"
	OutputFlag           = "output"
	OutputFlagShort      = "o"
	ConfigFlag           = "config"
	ContextFlag          = "context"
	ContextFlagShort     = "c"
)

var AllowedOutputs = []string{JSONOutput.String(), TextOutput.String()}

func AddPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(URLFlag, "", URLFlagUse)
	cmd.PersistentFlags().String(APITokenFlag, "", APITokenFlagUse)
	cmd.PersistentFlags().String(ServiceTokenFlag, "", ServiceTokenFlagUse)
	cmd.PersistentFlags().String(ServiceBearerFlag, "", ServiceBearerFlagUse)
	cmd.PersistentFlags().CountP(VerboseFlag, VerboseFlagShort, "Print verbose logging to the terminal to track what the CLI is doing (use multiple times to increase verbosity)")
	cmd.PersistentFlags().Bool(NoColorFlag, false, "Disable color when printing to the terminal")
	cmd.PersistentFlags().String(ConfigFlag, "", "Override the path to the config file")
	cmd.PersistentFlags().StringP(ContextFlag, ContextFlagShort, "", "Override the context to use")
	pflags.EnumP(cmd.PersistentFlags(), OutputFlag, OutputFlagShort, "text", AllowedOutputs, fmt.Sprintf("Specify the output format (must be { %s })", strings.Join(AllowedOutputs, " | ")))
}
