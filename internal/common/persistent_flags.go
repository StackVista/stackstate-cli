package common

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/pkg/pflags"
)

const (
	VerboseFlag             = "verbose"
	VerboseFlagShort        = "v"
	VersionFlag             = "version"
	VersionFlagUse          = "Prints the minimum StackState version supported by the command"
	URLFlag                 = "url"
	URLFlagUse              = "Specify the URL of the StackState server"
	APITokenFlag            = "api-token"
	APITokenFlagUse         = "Specify the API token of the StackState server" //nolint:gosec
	ServiceTokenFlag        = "service-token"
	ServiceTokenFlagUse     = "Specify the Service token of the StackState server" //nolint:gosec
	K8sSATokenFlag          = "k8s-sa-token"                                       //nolint:gosec
	K8sSATokenFlagUse       = "Specify the Kubernetes Service Account Token"
	K8sSATokenPathFlag      = "k8s-sa-token-path" //nolint:gosec
	K8sSATokenPathFlagUse   = "Specify the path to the Kubernetes Service Account Token"
	NoColorFlag             = "no-color"
	OutputFlag              = "output"
	OutputFlagShort         = "o"
	ConfigFlag              = "config"
	ContextFlag             = "context"
	ContextFlagShort        = "c"
	SkipSSLFlag             = "skip-ssl"
	SkipSSLFlagUse          = "Whether to skip SSL certificate verification when connecting to StackState"
	CaCertPathFlag          = "ca-cert-path"
	CaCertPathFlagUse       = "Path to a private CA or self-signed certificate file. Ignored if skip-ssl is set"
	CaCertBase64DataFlag    = "ca-cert-base64-data"
	CaCertBase64DataFlagUse = "Base64-encoded private CA or self-signed certificate data to use for SSL verification. Ignored if skip-ssl or ca-cert-path is set"
)

var AllowedOutputs = []string{JSONOutput.String(), TextOutput.String()}

func AddPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(URLFlag, "", URLFlagUse)
	cmd.PersistentFlags().String(APITokenFlag, "", APITokenFlagUse)
	cmd.PersistentFlags().String(ServiceTokenFlag, "", ServiceTokenFlagUse)
	cmd.PersistentFlags().String(K8sSATokenFlag, "", K8sSATokenFlagUse)
	cmd.PersistentFlags().String(K8sSATokenPathFlag, "", K8sSATokenPathFlagUse)
	cmd.PersistentFlags().CountP(VerboseFlag, VerboseFlagShort, "Print verbose logging to the terminal to track what the CLI is doing (use multiple times to increase verbosity)")
	cmd.PersistentFlags().Bool(NoColorFlag, false, "Disable color when printing to the terminal")
	cmd.PersistentFlags().String(ConfigFlag, "", "Override the path to the config file")
	cmd.PersistentFlags().StringP(ContextFlag, ContextFlagShort, "", "Override the context to use")
	cmd.PersistentFlags().Bool(SkipSSLFlag, false, SkipSSLFlagUse)
	cmd.PersistentFlags().String(CaCertBase64DataFlag, "", CaCertBase64DataFlagUse)
	cmd.PersistentFlags().String(CaCertPathFlag, "", CaCertPathFlagUse)
	pflags.EnumP(cmd.PersistentFlags(), OutputFlag, OutputFlagShort, "text", AllowedOutputs, fmt.Sprintf("Specify the output format (must be { %s })", strings.Join(AllowedOutputs, " | ")))

	// NOTE Add as a dummy `--version` flag and hides it, so that we omit the auto-generated Cobra flag on each versioned command.
	cmd.PersistentFlags().Bool(VersionFlag, false, VersionFlagUse)
	cmd.PersistentFlags().MarkHidden(VersionFlag) //nolint:errcheck
}
