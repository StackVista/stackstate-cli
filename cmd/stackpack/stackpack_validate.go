package stackpack

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

// ValidateArgs contains arguments for stackpack validate command
type ValidateArgs struct {
	StackpackDir  string
	StackpackFile string
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
		Long: `Validate a stackpack against a SUSE Observability server.

This command validates a stackpack by uploading it to the server.
- If a directory is provided, it is automatically packaged into a .sts file before uploading
- If a .sts file is provided, it is uploaded directly

Exactly one of --stackpack-directory or --stackpack-file must be specified.

This command is experimental and requires STS_EXPERIMENTAL_STACKPACK environment variable to be set.`,
		Example: `# Validate a stackpack directory (automatically packaged)
sts stackpack validate --stackpack-directory ./my-stackpack

# Validate a pre-packaged .sts file
sts stackpack validate --stackpack-file ./my-stackpack.sts`,
		RunE: cli.CmdRunEWithApi(RunStackpackValidateCommand(args)),
	}

	cmd.Flags().StringVarP(&args.StackpackDir, "stackpack-directory", "d", "", "Path to stackpack directory")
	cmd.Flags().StringVarP(&args.StackpackFile, "stackpack-file", "f", "", "Path to .sts file")

	return cmd
}

// RunStackpackValidateCommand executes the validate command
func RunStackpackValidateCommand(args *ValidateArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		// Validate exactly one of directory or file is set
		if (args.StackpackDir == "" && args.StackpackFile == "") ||
			(args.StackpackDir != "" && args.StackpackFile != "") {
			return common.NewCLIArgParseError(fmt.Errorf("exactly one of --stackpack-directory or --stackpack-file must be specified"))
		}

		// Prepare file to validate - if directory is provided, package it first
		fileToValidate, cleanup, err := prepareStackpackFile(args)
		if err != nil {
			return err
		}
		defer cleanup()

		// Open the file
		file, openErr := os.Open(fileToValidate)
		if openErr != nil {
			return common.NewRuntimeError(fmt.Errorf("failed to open stackpack file: %w", openErr))
		}
		defer file.Close()

		// Call validate endpoint
		result, resp, validateErr := api.StackpackApi.StackPackValidate(cli.Context).StackPack(file).Execute()
		if validateErr != nil {
			return common.NewResponseError(validateErr, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"success": true,
				"result":  result,
			})
		} else {
			cli.Printer.Success("Stackpack validation successful!")
			if result != "" {
				fmt.Println(result)
			}
		}

		return nil
	}
}

// prepareStackpackFile returns the path to the stackpack file to validate.
// If a directory is provided, it packages it into a temporary .sts file.
// Returns the file path and a cleanup function that should be deferred.
func prepareStackpackFile(args *ValidateArgs) (string, func(), common.CLIError) {
	if args.StackpackFile != "" {
		// Use provided .sts file directly
		if _, err := os.Stat(args.StackpackFile); err != nil {
			return "", func() {}, common.NewRuntimeError(fmt.Errorf("failed to access stackpack file: %w", err))
		}
		return args.StackpackFile, func() {}, nil
	}

	// Package the directory
	absDir, err := filepath.Abs(args.StackpackDir)
	if err != nil {
		return "", func() {}, common.NewRuntimeError(fmt.Errorf("failed to resolve stackpack directory: %w", err))
	}

	// Validate stackpack directory
	if err := validateStackpackDirectory(absDir); err != nil {
		return "", func() {}, common.NewCLIArgParseError(err)
	}

	// Parse stackpack info
	parser := &YamlParser{}
	stackpackInfo, err := parser.Parse(filepath.Join(absDir, "stackpack.yaml"))
	if err != nil {
		return "", func() {}, common.NewRuntimeError(fmt.Errorf("failed to parse stackpack.yaml: %w", err))
	}

	// Create temporary .sts file
	tmpFile, err := os.CreateTemp("", fmt.Sprintf("%s-*.sts", stackpackInfo.Name))
	if err != nil {
		return "", func() {}, common.NewRuntimeError(fmt.Errorf("failed to create temporary file: %w", err))
	}
	tmpFile.Close()
	tmpPath := tmpFile.Name()

	// Package stackpack into temporary file
	if err := createStackpackZip(absDir, tmpPath); err != nil {
		os.Remove(tmpPath) // Clean up on error
		return "", func() {}, common.NewRuntimeError(fmt.Errorf("failed to package stackpack: %w", err))
	}

	// Return cleanup function that removes the temporary file
	cleanup := func() {
		os.Remove(tmpPath)
	}

	return tmpPath, cleanup, nil
}
