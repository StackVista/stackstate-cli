package stackpack

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

// ValidateArgs contains arguments for stackpack validate command
type ValidateArgs struct {
	Name          string
	StackpackDir  string
	StackpackFile string
	DockerImage   string

	dockerRunner func([]string) error
}

// StackpackValidateCommand creates the validate subcommand
func StackpackValidateCommand(cli *di.Deps) *cobra.Command {
	return stackpackValidateCommandWithArgs(cli, &ValidateArgs{})
}

// stackpackValidateCommandWithArgs creates the validate command with injected args (for testing)
func stackpackValidateCommandWithArgs(cli *di.Deps, args *ValidateArgs) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate a stackpack",
		Long: `Validate a stackpack using either the API or Docker mode.

In API mode (when a configured backend context is active), this command calls POST /stackpack/{name}/validate
against the live instance.

In Docker mode (when --image is specified), it spins up quay.io/stackstate/stackstate-server:<tag>
with stack-pack-validator as the entrypoint.

This command is experimental and requires STS_EXPERIMENTAL_STACKPACK environment variable to be set.`,
		Example: `# Validate using API
sts stackpack validate --name my-stackpack

# Validate using Docker with a directory
sts stackpack validate --image quay.io/stackstate/stackstate-server:latest --stackpack-directory ./my-stackpack

# Validate using Docker with a file
sts stackpack validate --image quay.io/stackstate/stackstate-server:latest --stackpack-file ./my-stackpack.sts`,
		RunE: cli.CmdRunE(RunStackpackValidateCommand(args)),
	}

	cmd.Flags().StringVarP(&args.Name, "name", "n", "", "Stackpack name (required for API mode)")
	cmd.Flags().StringVarP(&args.StackpackDir, "stackpack-directory", "d", "", "Path to stackpack directory (Docker mode)")
	cmd.Flags().StringVarP(&args.StackpackFile, "stackpack-file", "f", "", "Path to .sts file (Docker mode)")
	cmd.Flags().StringVar(&args.DockerImage, "image", "", "Docker image reference (triggers Docker mode)")

	// Set default docker runner if not already set
	if args.dockerRunner == nil {
		args.dockerRunner = defaultDockerRunner
	}

	return cmd
}

// RunStackpackValidateCommand executes the validate command
func RunStackpackValidateCommand(args *ValidateArgs) func(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	return func(cli *di.Deps, cmd *cobra.Command) common.CLIError {
		// Determine mode: use Docker if image is provided, otherwise check if context is available
		useDocker := args.DockerImage != ""
		if !useDocker {
			// Try to load context if not already loaded
			if cli.CurrentContext == nil {
				_ = cli.LoadContext(cmd) // Silently ignore error, context is optional
			}
			// Use docker mode if no context or no URL
			useDocker = cli.CurrentContext == nil || cli.CurrentContext.URL == ""
		}

		if useDocker {
			return runDockerValidation(args)
		}
		return runAPIValidation(cli, cmd, args)
	}
}

// runAPIValidation validates stackpack via API
func runAPIValidation(cli *di.Deps, cmd *cobra.Command, args *ValidateArgs) common.CLIError {
	if args.Name == "" {
		return common.NewCLIArgParseError(fmt.Errorf("stackpack name is required (use --name)"))
	}

	// Ensure client is loaded
	if cli.Client == nil {
		err := cli.LoadClient(cmd, cli.CurrentContext)
		if err != nil {
			return err
		}
	}

	// Connect to API
	api, _, connectErr := cli.Client.Connect()
	if connectErr != nil {
		return common.NewRuntimeError(fmt.Errorf("failed to connect to API: %w", connectErr))
	}

	// Call validate endpoint
	_, resp, validateErr := api.StackpackApi.ValidateStackPack(cli.Context, args.Name).Execute()
	if validateErr != nil {
		return common.NewResponseError(validateErr, resp)
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"success": true,
		})
	} else {
		cli.Printer.Success("Stackpack validation successful!")
	}

	return nil
}

// runDockerValidation validates stackpack via Docker
func runDockerValidation(args *ValidateArgs) common.CLIError {
	// Validate required flags
	if args.DockerImage == "" {
		return common.NewCLIArgParseError(fmt.Errorf("--image is required for Docker mode"))
	}

	// Validate exactly one of directory or file is set
	if (args.StackpackDir == "" && args.StackpackFile == "") ||
		(args.StackpackDir != "" && args.StackpackFile != "") {
		return common.NewCLIArgParseError(fmt.Errorf("exactly one of --stackpack-directory or --stackpack-file must be specified"))
	}

	// Check docker is available
	if _, err := exec.LookPath("docker"); err != nil {
		return common.NewRuntimeError(fmt.Errorf("docker is not available: %w", err))
	}

	// Build docker command arguments
	dockerArgs := []string{"run", "--rm", "--entrypoint", "/opt/docker/bin/stack-pack-validator"}

	if args.StackpackDir != "" {
		// Convert to absolute path
		absDir, err := filepath.Abs(args.StackpackDir)
		if err != nil {
			return common.NewRuntimeError(fmt.Errorf("failed to resolve stackpack directory: %w", err))
		}
		dockerArgs = append(dockerArgs, "-v", fmt.Sprintf("%s:/stackpack", absDir), args.DockerImage, "-directory", "/stackpack")
	} else {
		// Convert to absolute path
		absFile, err := filepath.Abs(args.StackpackFile)
		if err != nil {
			return common.NewRuntimeError(fmt.Errorf("failed to resolve stackpack file: %w", err))
		}
		dockerArgs = append(dockerArgs, "-v", fmt.Sprintf("%s:/stackpack.sts", absFile), args.DockerImage, "-file", "/stackpack.sts")
	}

	// Execute docker command
	if err := args.dockerRunner(dockerArgs); err != nil {
		return common.NewRuntimeError(fmt.Errorf("docker validation failed: %w", err))
	}

	return nil
}

// defaultDockerRunner executes docker command with streaming output
func defaultDockerRunner(dockerArgs []string) error {
	cmd := exec.CommandContext(context.Background(), "docker", dockerArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
