package stackpack

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/blang/semver/v4"
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

const (
	stackpackConfigFile = "stackpack.conf"
)

// TestArgs contains arguments for stackpack test command
type TestArgs struct {
	StackpackDir     string
	Params           map[string]string
	Yes              bool
	UnlockedStrategy string
}

// StackpackTestCommand creates the test subcommand
func StackpackTestCommand(cli *di.Deps) *cobra.Command {
	args := &TestArgs{
		Params: make(map[string]string),
	}
	cmd := &cobra.Command{
		Use:   "test",
		Short: "Test a stackpack by packaging, uploading, and installing/upgrading",
		Long: `Test a stackpack by running package, upload, and install/upgrade commands in sequence.

This command will:
1. Read the stackpack name and version from ` + stackpackConfigFile + `
2. Create a temporary copy of the stackpack directory
3. Update the version in the temporary copy with -cli-test.N suffix
4. Package the stackpack from the temporary copy into a zip file
5. Upload the zip file to the server (with confirmation)
6. Install the stackpack (if not installed) or upgrade it (if already installed)

The original stackpack directory is left untouched. The version is automatically incremented for each test run.
The stackpack name is read from ` + stackpackConfigFile + `, so no --name flag is required.`,
		Example: `# Test stackpack with confirmation
sts stackpack test -p "param1=value1"

# Skip confirmation prompt
sts stackpack test --yes

# Test stackpack in specific directory with unlocked strategy
sts stackpack test -d ./my-stackpack --yes --unlocked-strategy force

# Test with custom unlocked strategy
sts stackpack test --unlocked-strategy skip --yes`,
		RunE: cli.CmdRunEWithApi(RunStackpackTestCommand(args)),
	}

	cmd.Flags().StringVarP(&args.StackpackDir, "stackpack-directory", "d", "", "Path to stackpack directory (defaults to current directory)")
	cmd.Flags().StringToStringVarP(&args.Params, ParameterFlag, "p", args.Params, "List of parameters of the form \"key=value\"")
	cmd.Flags().BoolVarP(&args.Yes, "yes", "y", false, "Skip confirmation prompt before upload")
	cmd.Flags().StringVar(&args.UnlockedStrategy, UnlockedStrategyFlag, "fail", "Strategy for dealing with unlocked StackPacks. Valid options are: fail, force, skip")

	return cmd
}

// RunStackpackTestCommand executes the test command
//
//nolint:funlen
func RunStackpackTestCommand(args *TestArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		// Warn if JSON output is requested - not meaningful for test command
		if cli.IsJson() {
			cli.Printer.PrintLn("Warning: JSON output format is not meaningful for the test command, proceeding with text output")
		}

		// Set default stackpack directory
		if args.StackpackDir == "" {
			currentDir, err := os.Getwd()
			if err != nil {
				return common.NewRuntimeError(fmt.Errorf("failed to get current directory: %w", err))
			}
			args.StackpackDir = currentDir
		}

		// Parse stackpack configuration to get original version
		parser := &HoconParser{}
		stackpackConf := filepath.Join(args.StackpackDir, stackpackConfigFile)
		originalInfo, err := parser.Parse(stackpackConf)
		if err != nil {
			return common.NewRuntimeError(fmt.Errorf("failed to parse %s: %w", stackpackConfigFile, err))
		}

		cli.Printer.Success("Starting stackpack test sequence...")
		cli.Printer.PrintLn(fmt.Sprintf("  Stackpack: %s (current version: %s)", originalInfo.Name, originalInfo.Version))
		cli.Printer.PrintLn("")

		// Step 1: Check installed version and determine base version for snapshot
		cli.Printer.PrintLn("Step 1/5: Checking installed version...")
		installedVersion, err := getInstalledStackpackVersion(cli, api, originalInfo.Name)
		if err != nil {
			return common.NewRuntimeError(fmt.Errorf("failed to check installed stackpack version: %w", err))
		}

		baseVersionForSnapshot := originalInfo.Version
		if installedVersion != "" {
			cli.Printer.PrintLn(fmt.Sprintf("  Found installed version: %s", installedVersion))

			// Compare base versions (strip cli-test suffix from installed version for comparison)
			installedBaseVersion := installedVersion
			if strings.Contains(installedVersion, "-cli-test.") {
				// Extract base version from cli-test version (e.g., "1.0.0-cli-test.5" -> "1.0.0")
				if parts := strings.Split(installedVersion, "-cli-test."); len(parts) > 0 {
					installedBaseVersion = parts[0]
				}
			}

			// Base version selection logic:
			// 1. If installed base version is higher: use installed version
			// 2. If base versions are equal AND installed has cli-test: use installed version (continue test sequence)
			// 3. If local version is higher: use local version (prevents downgrade)
			baseComparison := compareVersions(installedBaseVersion, originalInfo.Version)
			if baseComparison > 0 || (baseComparison == 0 && strings.Contains(installedVersion, "-cli-test.")) {
				baseVersionForSnapshot = installedVersion
				cli.Printer.PrintLn(fmt.Sprintf("  Using installed version as base: %s", baseVersionForSnapshot))
			} else if baseComparison < 0 {
				cli.Printer.PrintLn(fmt.Sprintf("  Using local version as base (higher than installed): %s", baseVersionForSnapshot))
			}
		}

		// Step 2: Create temporary directory and copy stackpack
		cli.Printer.PrintLn("")
		cli.Printer.PrintLn("Step 2/5: Creating temporary copy for testing...")

		tempDir, err := os.MkdirTemp("", "stackpack-test-*")
		if err != nil {
			return common.NewRuntimeError(fmt.Errorf("failed to create temporary directory: %w", err))
		}

		// Ensure cleanup of temporary directory
		defer func() {
			if removeErr := os.RemoveAll(tempDir); removeErr != nil {
				cli.Printer.PrintLn(fmt.Sprintf("Warning: Failed to cleanup temporary directory: %v", removeErr))
			}
		}()

		tempStackpackDir := filepath.Join(tempDir, "stackpack")
		if err := copyDirectory(args.StackpackDir, tempStackpackDir); err != nil {
			return common.NewRuntimeError(fmt.Errorf("failed to copy stackpack to temporary directory: %w", err))
		}
		cli.Printer.Success("Temporary copy created")

		// Step 3: Update version in temporary copy
		cli.Printer.PrintLn("")
		cli.Printer.PrintLn("Step 3/5: Bumping version for testing...")

		tempConfigPath := filepath.Join(tempStackpackDir, stackpackConfigFile)
		newVersion, err := bumpSnapshotVersionWithBase(tempConfigPath, baseVersionForSnapshot)
		if err != nil {
			return common.NewRuntimeError(fmt.Errorf("failed to bump version: %w", err))
		}
		cli.Printer.Success(fmt.Sprintf("Version bumped to: %s", newVersion))

		// Step 4: Package stackpack from temporary directory
		cli.Printer.PrintLn("")
		cli.Printer.PrintLn("Step 4/5: Packaging stackpack...")
		packageArgs := &PackageArgs{
			StackpackDir: tempStackpackDir, // Use temporary directory
			Force:        true,             // Always overwrite for testing
		}

		// Set archive file in current directory
		currentDir, _ := os.Getwd()
		packageArgs.ArchiveFile = filepath.Join(currentDir, fmt.Sprintf("%s-%s.zip", originalInfo.Name, newVersion))

		if err := runPackageStep(cli, packageArgs); err != nil {
			return err
		}
		cli.Printer.Success("Stackpack packaged successfully")

		// Step 5: Confirm upload (if needed) and execute upload/install workflow
		if !args.Yes {
			cli.Printer.PrintLn("")
			if !confirmUpload(cli, packageArgs.ArchiveFile) {
				return common.NewRuntimeError(fmt.Errorf("upload cancelled by user"))
			}
		}

		// Upload stackpack
		cli.Printer.PrintLn("")
		cli.Printer.PrintLn("Step 5/5: Uploading and installing/upgrading stackpack...")
		uploadArgs := &UploadArgs{
			FilePath: packageArgs.ArchiveFile,
		}

		if err := runUploadStep(cli, api, serverInfo, uploadArgs); err != nil {
			return err
		}
		cli.Printer.Success("Stackpack uploaded successfully")

		// Install or upgrade stackpack based on installation status
		if installedVersion != "" {
			upgradeArgs := &UpgradeArgs{
				TypeName:         originalInfo.Name,
				UnlockedStrategy: args.UnlockedStrategy,
				Wait:             true,
				Timeout:          DefaultTimeout,
			}

			if err := runUpgradeStep(cli, api, serverInfo, upgradeArgs); err != nil {
				return err
			}
			cli.Printer.Success("Stackpack upgraded successfully")
		} else {
			installArgs := &InstallArgs{
				Name:             originalInfo.Name,
				UnlockedStrategy: args.UnlockedStrategy,
				Params:           args.Params,
				Wait:             true,
				Timeout:          DefaultTimeout,
			}

			if err := runInstallStep(cli, api, serverInfo, installArgs); err != nil {
				return err
			}
			cli.Printer.Success("Stackpack installed successfully")
		}

		cli.Printer.PrintLn("")
		cli.Printer.Success("🎉 Test sequence completed successfully!")

		// Clean up zip file
		if err := os.Remove(packageArgs.ArchiveFile); err != nil {
			cli.Printer.PrintLn(fmt.Sprintf("Note: Could not clean up zip file %s: %v", packageArgs.ArchiveFile, err))
		}

		return nil
	}
}

// bumpSnapshotVersionWithBase increments cli-test version using a different base version for computation
// Logic: If baseVersion contains "-cli-test.N", increments N to N+1
// If no cli-test suffix exists, adds "-cli-test.10000" to the base version (high number for alphanumeric ordering)
// Examples: "1.0.0" -> "1.0.0-cli-test.10000", "1.0.0-cli-test.10005" -> "1.0.0-cli-test.10006"
func bumpSnapshotVersionWithBase(configPath, baseVersion string) (string, error) {
	// Parse the base version using semver
	version, err := semver.Parse(baseVersion)
	if err != nil {
		return "", fmt.Errorf("failed to parse base version %s: %w", baseVersion, err)
	}

	var newVersion string
	// Search for existing "cli-test" pre-release identifier
	cliTestIndex := -1
	for i, pre := range version.Pre {
		if pre.VersionStr == "cli-test" && i+1 < len(version.Pre) {
			cliTestIndex = i
			break
		}
	}

	if cliTestIndex >= 0 && cliTestIndex+1 < len(version.Pre) {
		// Existing cli-test found: increment its number
		nextPart := version.Pre[cliTestIndex+1]
		var currentNum uint64

		// Extract current cli-test number (handles both numeric and string formats)
		if nextPart.IsNumeric() {
			// Numeric format: use VersionNum directly
			currentNum = nextPart.VersionNum
		} else if num, err := strconv.ParseUint(nextPart.VersionStr, 10, 64); err == nil {
			// String format: parse to number
			currentNum = num
		} else {
			// Invalid format: reset to cli-test.10000
			newVersion = fmt.Sprintf("%d.%d.%d-cli-test.10000", version.Major, version.Minor, version.Patch)
			return newVersion, updateVersionInHocon(configPath, newVersion)
		}

		// Rebuild pre-release parts with incremented cli-test number
		newPreParts := make([]string, len(version.Pre))
		for i, pre := range version.Pre {
			switch {
			case i == cliTestIndex+1:
				// Increment the cli-test number
				newPreParts[i] = fmt.Sprintf("%d", currentNum+1)
			case pre.IsNumeric():
				// Convert numeric pre-release parts to string
				newPreParts[i] = fmt.Sprintf("%d", pre.VersionNum)
			default:
				// Keep string pre-release parts as-is
				newPreParts[i] = pre.VersionStr
			}
		}
		newVersion = fmt.Sprintf("%d.%d.%d-%s", version.Major, version.Minor, version.Patch, strings.Join(newPreParts, "."))
	} else {
		// No cli-test found: add initial cli-test.10000 suffix
		newVersion = fmt.Sprintf("%d.%d.%d-cli-test.10000", version.Major, version.Minor, version.Patch)
	}

	return newVersion, updateVersionInHocon(configPath, newVersion)
}

// confirmUpload prompts user for confirmation before upload
func confirmUpload(cli *di.Deps, zipFile string) bool {
	serverURL := "unknown"
	if cli.CurrentContext != nil {
		serverURL = cli.CurrentContext.URL
	}

	cli.Printer.PrintLn(fmt.Sprintf("⚠️  This will upload '%s' to SUSE Observability server:", filepath.Base(zipFile)))
	cli.Printer.PrintLn(fmt.Sprintf("   Server: %s", serverURL))
	fmt.Print("   Continue? (y/N): ")

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y" || response == "yes"
}

// runPackageStep executes the package command logic
func runPackageStep(cli *di.Deps, args *PackageArgs) common.CLIError {
	// Reuse the existing package command logic
	packageCmd := &cobra.Command{}
	packageFn := RunStackpackPackageCommand(args)
	return packageFn(cli, packageCmd)
}

// runUploadStep executes the upload command logic
func runUploadStep(cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo, args *UploadArgs) common.CLIError {
	// Reuse the existing upload command logic
	uploadCmd := &cobra.Command{}
	uploadFn := RunStackpackUploadCommand(args)
	return uploadFn(uploadCmd, cli, api, serverInfo)
}

// runInstallStep executes the install command logic
func runInstallStep(cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo, args *InstallArgs) common.CLIError {
	// Reuse the existing install command logic
	installCmd := &cobra.Command{}
	installFn := RunStackpackInstallCommand(args)
	return installFn(installCmd, cli, api, serverInfo)
}

// runUpgradeStep executes the upgrade command logic
func runUpgradeStep(cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo, args *UpgradeArgs) common.CLIError {
	// Reuse the existing upgrade command logic
	upgradeCmd := &cobra.Command{}
	upgradeFn := RunStackpackUpgradeCommand(args)
	return upgradeFn(upgradeCmd, cli, api, serverInfo)
}

// getInstalledStackpackVersion checks if a stackpack is installed and returns its version
func getInstalledStackpackVersion(cli *di.Deps, api *stackstate_api.APIClient, stackpackName string) (string, error) {
	stackPackList, err := fetchAllStackPacks(cli, api)
	if err != nil {
		return "", err
	}

	for _, stackpack := range stackPackList {
		if stackpack.Name == stackpackName && len(stackpack.GetConfigurations()) > 0 {
			// Stackpack is installed, return its version
			return stackpack.Version, nil
		}
	}

	// Stackpack not installed
	return "", nil
}

// compareVersions compares two version strings using semver
// Returns: -1 if v1 < v2, 0 if v1 == v2, 1 if v1 > v2
func compareVersions(v1, v2 string) int {
	ver1, err1 := semver.Parse(v1)
	ver2, err2 := semver.Parse(v2)

	if err1 != nil || err2 != nil {
		// If parsing fails, treat as equal (shouldn't happen with valid semver)
		return 0
	}

	return ver1.Compare(ver2)
}

// copyDirectory recursively copies a directory to a destination
func copyDirectory(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate the destination path
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		return copyFile(path, dstPath, info.Mode())
	})
}

// copyFile copies a single file
func copyFile(src, dst string, mode os.FileMode) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if err := dstFile.Chmod(mode); err != nil {
		return err
	}

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// updateVersionInHocon updates the version field in a HOCON configuration using regex
func updateVersionInHocon(configPath, newVersion string) error {
	// Read the current config
	content, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse as HOCON to validate structure
	parser := &HoconParser{}
	_, err = parser.Parse(configPath)
	if err != nil {
		return fmt.Errorf("failed to parse HOCON config: %w", err)
	}

	// Use regex to replace version while preserving HOCON structure
	// This is more reliable than string replacement as it targets the version field specifically
	oldContent := string(content)
	versionRegex := regexp.MustCompile(`(?m)^(\s*version\s*=\s*)"[^"]*"(.*)$`)
	newContent := versionRegex.ReplaceAllString(oldContent, `${1}"`+newVersion+`"${2}`)

	if oldContent == newContent {
		// Try unquoted version format
		versionRegex = regexp.MustCompile(`(?m)^(\s*version\s*=\s*)[^,\n\r}]*(.*)$`)
		newContent = versionRegex.ReplaceAllString(oldContent, `${1}"`+newVersion+`"${2}`)
	}

	if oldContent == newContent {
		return fmt.Errorf("version field not found in config file")
	}

	// Validate the result can still be parsed
	tempFile := configPath + ".tmp"
	if err := os.WriteFile(tempFile, []byte(newContent), os.FileMode(defaultDirMode)); err != nil {
		return fmt.Errorf("failed to write temporary config: %w", err)
	}

	// Validate the modified config
	_, err = parser.Parse(tempFile)
	if err != nil {
		_ = os.Remove(tempFile)
		return fmt.Errorf("modified config is not valid HOCON: %w", err)
	}

	// Replace the original file
	if err := os.Rename(tempFile, configPath); err != nil {
		_ = os.Remove(tempFile)
		return fmt.Errorf("failed to replace config file: %w", err)
	}

	return nil
}
