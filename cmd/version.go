package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
)

var (
	Version string
	Commit  string
	Date    string
)

func VersionCommand(_ *config.Config) *cobra.Command {
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
