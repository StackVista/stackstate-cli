package dev

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/devimagepatcher"
)

func PatchDeploymentImageCommand(cfg *config.Dev) *cobra.Command {
	servicesList, err := devimagepatcher.Services()
	if err != nil {
		panic(err)
	}
	services := strings.Join(servicesList, ", ")

	cmd := &cobra.Command{
		Use:   "patch-deployment-image",
		Short: fmt.Sprintf("Patch a single deployment in a running StackState with a new development image (locally built). Provide a service name to patch, supported: %s", services),

		RunE: func(cmd *cobra.Command, args []string) error {
			patcher, err := devimagepatcher.New(cmd.Context(), cfg)
			if err != nil {
				return err
			}

			noPush, _ := cmd.Flags().GetBool(NoPushFlag)
			namespace, _ := cmd.Flags().GetString(NamespaceFlag)
			tag, _ := cmd.Flags().GetString(TagFlag)
			service := &args[0]
			if !noPush {
				err := patcher.PushNewDockerImages(service, nil, tag)
				if err != nil {
					return err
				}
			}

			err = patcher.UpdateDeployments(service, nil, tag, namespace)
			if err != nil {
				return err
			}

			return nil
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("dev-image patch-deployment needs a service name as argument. Supported service names are: %s", services)
			}
			return nil
		},
	}

	return cmd
}
