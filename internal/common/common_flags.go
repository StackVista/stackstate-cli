package common

import "github.com/spf13/cobra"

const (
	FileFlag      = "file"
	FileFlagShort = "f"
	IDFlag        = "id"
	IDFlagShort   = "i"
	NameFlag      = "name"
	NameFlagShort = "n"
)

func AddFileFlag(cmd *cobra.Command, use string) *string {
	return cmd.Flags().StringP(FileFlag, FileFlagShort, "", use)
}

func AddFileFlagVar(cmd *cobra.Command, v *string, use string) {
	cmd.Flags().StringVarP(v, FileFlag, FileFlagShort, "", use)
}

func AddRequiredFileFlag(cmd *cobra.Command, use string) *string {
	s := AddFileFlag(cmd, use)
	cmd.MarkFlagRequired(FileFlag) //nolint:errcheck

	return s
}

func AddRequiredFileFlagVar(cmd *cobra.Command, v *string, use string) {
	AddFileFlagVar(cmd, v, use)
	cmd.MarkFlagRequired(FileFlag) //nolint:errcheck
}

func AddIDFlag(cmd *cobra.Command, use string) *string {
	return cmd.Flags().StringP(IDFlag, IDFlagShort, "", use)
}

func AddIDFlagVar(cmd *cobra.Command, v *string, use string) {
	cmd.Flags().StringVarP(v, IDFlag, IDFlagShort, "", use)
}

func AddRequiredIDFlag(cmd *cobra.Command, use string) *string {
	s := AddIDFlag(cmd, use)
	cmd.MarkFlagRequired(IDFlag) //nolint:errcheck

	return s
}

func AddRequiredIDFlagVar(cmd *cobra.Command, v *string, use string) {
	AddIDFlagVar(cmd, v, use)
	cmd.MarkFlagRequired(IDFlag) //nolint:errcheck
}

func AddNameFlag(cmd *cobra.Command, use string) *string {
	return cmd.Flags().StringP(NameFlag, NameFlagShort, "", use)
}

func AddNameFlagVar(cmd *cobra.Command, v *string, use string) {
	cmd.Flags().StringVar(v, NameFlag, "", use)
}

func AddRequiredNameFlag(cmd *cobra.Command, use string) *string {
	s := AddNameFlag(cmd, use)
	cmd.MarkFlagRequired(NameFlag) //nolint:errcheck

	return s
}

func AddRequiredNameFlagVar(cmd *cobra.Command, v *string, use string) {
	AddNameFlagVar(cmd, v, use)
	cmd.MarkFlagRequired(NameFlag) //nolint:errcheck
}
