package common

import "github.com/spf13/cobra"

const (
	FileFlag      = "file"
	FileFlagShort = "f"
	IDFlag        = "id"
	IDFlagShort   = "i"
	NameFlag      = "name"
)

func AddFileFlag(cmd *cobra.Command, use string) {
	cmd.Flags().StringP(FileFlag, FileFlagShort, "", use)
}

func AddRequiredFileFlag(cmd *cobra.Command, use string) {
	AddFileFlag(cmd, use)
	cmd.MarkFlagRequired(FileFlag) //nolint:errcheck
}

func AddIDFlag(cmd *cobra.Command, use string) {
	cmd.Flags().StringP(IDFlag, IDFlagShort, "", use)
}

func AddRequiredIDFlag(cmd *cobra.Command, use string) {
	AddIDFlag(cmd, use)
	cmd.MarkFlagRequired(IDFlag) //nolint:errcheck
}

func AddNameFlag(cmd *cobra.Command, use string) {
	cmd.Flags().String(NameFlag, "", use)
}

func AddRequiredNameFlag(cmd *cobra.Command, use string) {
	AddNameFlag(cmd, use)
	cmd.MarkFlagRequired(NameFlag) //nolint:errcheck
}
