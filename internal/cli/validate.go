package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func Required(s string) error {
	if strings.TrimSpace(s) == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}

func FlagRequired(cmd *cobra.Command, flagName string) error {
	f, err := cmd.Flags().GetString(flagName)
	if err != nil {
		return err
	}

	if err := Required(f); err != nil {
		return fmt.Errorf("flag '%s' is required", flagName)
	}

	return nil
}

func FlagsRequired(flags ...string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		for _, f := range flags {
			if err := FlagRequired(cmd, f); err != nil {
				return err
			}
		}

		return nil
	}
}
