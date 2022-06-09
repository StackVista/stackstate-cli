package common

import "github.com/spf13/cobra"

const (
	FileFlag       = "file"
	FileFlagShort  = "f"
	IDFlag         = "id"
	IdentifierFlag = "identifier"
	IDFlagShort    = "i"
	NameFlag       = "name"
	NameFlagShort  = "n"
)

func AddFileFlagVar(cmd *cobra.Command, v *string, use string) {
	cmd.Flags().StringVarP(v, FileFlag, FileFlagShort, "", use)
}

func AddRequiredFileFlagVar(cmd *cobra.Command, v *string, use string) {
	AddFileFlagVar(cmd, v, use)
	cmd.MarkFlagRequired(FileFlag) //nolint:errcheck
}

func AddIDFlagVar(cmd *cobra.Command, v *int64, use string) {
	cmd.Flags().Int64VarP(v, IDFlag, IDFlagShort, 0, use)
}

func AddRequiredIDFlagVar(cmd *cobra.Command, v *int64, use string) {
	AddIDFlagVar(cmd, v, use)
	cmd.MarkFlagRequired(IDFlag) //nolint:errcheck
}

func AddIdentifierFlagVar(cmd *cobra.Command, v *string, use string) {
	cmd.Flags().StringVar(v, IdentifierFlag, "", use)
}

func AddRequiredIdentifierFlagVar(cmd *cobra.Command, v *string, use string) {
	AddIdentifierFlagVar(cmd, v, use)
	cmd.MarkFlagRequired(IDFlag) //nolint:errcheck
}

func AddNameFlagVar(cmd *cobra.Command, v *string, use string) {
	cmd.Flags().StringVarP(v, NameFlag, NameFlagShort, "", use)
}

func AddNameFlagVarVal(cmd *cobra.Command, v *string, val string, use string) {
	cmd.Flags().StringVarP(v, NameFlag, NameFlagShort, val, use)
}

func AddRequiredNameFlagVar(cmd *cobra.Command, v *string, use string) {
	AddNameFlagVar(cmd, v, use)
	cmd.MarkFlagRequired(NameFlag) //nolint:errcheck
}
