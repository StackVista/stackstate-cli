package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version string
	Commit  string
	Date    string
)

func VersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display the version number",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Version: %s\n", Version)
			fmt.Printf("Commit: %s\n", Commit)
			fmt.Printf("Date built: %s\n", Date)
		},
	}

	return cmd
}
