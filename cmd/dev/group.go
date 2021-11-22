package dev

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/devimagepatcher"
)

func PatchGroupImageCommand(cfg *config.Dev) *cobra.Command {
	groupsList, err := devimagepatcher.Groups()
	if err != nil {
		panic(err)
	}
	groups := strings.Join(groupsList, ", ")

	cmd := &cobra.Command{
		Use:   "patch-group-image",
		Short: fmt.Sprintf("Patch a group of deployments in a running StackState with a new development image (locally built). Provide a group name to patch, supported: %s", groups),

		RunE: func(cmd *cobra.Command, args []string) error {
			patcher, err := devimagepatcher.New(cmd.Context(), cfg)
			if err != nil {
				return err
			}

			noPush, _ := cmd.Flags().GetBool(NoPushFlag)
			namespace, _ := cmd.Flags().GetString(NamespaceFlag)
			tag, _ := cmd.Flags().GetString(TagFlag)
			group := &args[0]
			if !noPush {
				err := patcher.PushNewDockerImages(nil, group, tag)
				if err != nil {
					return err
				}
			}

			err = patcher.UpdateDeployments(nil, group, tag, namespace)
			if err != nil {
				return err
			}

			return nil
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("dev-image patch-group needs a group name as argument. Supported group names are: %s", groups)
			}
			return nil
		},
	}

	return cmd
}
