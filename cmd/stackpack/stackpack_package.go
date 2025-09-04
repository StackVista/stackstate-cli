package stackpack

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gurkankaymak/hocon"
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

const (
	defaultDirMode = 0755 // Default directory permissions
)

// PackageArgs contains arguments for stackpack package command
type PackageArgs struct {
	StackpackDir string
	ArchiveFile  string
	Force        bool
}

// StackpackInfo contains parsed stackpack metadata
type StackpackInfo struct {
	Name    string
	Version string
}

// StackpackConfigParser interface for parsing stackpack configuration
type StackpackConfigParser interface {
	Parse(filePath string) (*StackpackInfo, error)
}

// HoconParser implements StackpackConfigParser for HOCON format
type HoconParser struct{}

func (h *HoconParser) Parse(filePath string) (*StackpackInfo, error) {
	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse stackpack.conf content
	conf, err := hocon.ParseString(string(content))
	if err != nil {
		return nil, fmt.Errorf("failed to parse stackpack.conf file: %w", err)
	}

	name := strings.Trim(conf.GetString("name"), `"`)
	version := strings.Trim(conf.GetString("version"), `"`)

	if name == "" {
		return nil, fmt.Errorf("name not found in stackpack.conf")
	}

	if version == "" {
		return nil, fmt.Errorf("version not found in stackpack.conf")
	}

	return &StackpackInfo{
		Name:    name,
		Version: version,
	}, nil
}

// YamlParser implements StackpackConfigParser for YAML format (future)
type YamlParser struct{}

func (y *YamlParser) Parse(filePath string) (*StackpackInfo, error) {
	// TODO: Implement YAML parsing when format changes
	return nil, fmt.Errorf("YAML format not yet implemented")
}

// Required files and directories for a valid stackpack
var requiredStackpackItems = []string{
	"provisioning",
	"README.md",
	"resources",
	"stackpack.conf",
}

// StackpackPackageCommand creates the package subcommand
func StackpackPackageCommand(cli *di.Deps) *cobra.Command {
	args := &PackageArgs{}
	cmd := &cobra.Command{
		Use:   "package",
		Short: "Package a stackpack into a zip file",
		Long: `Package a stackpack into a zip file.

Creates a zip file containing all required stackpack files and directories:
- provisioning/ (directory)
- README.md (file)
- resources/ (directory)
- stackpack.conf (file)

The zip file is named <stackpack_name>-<version>.zip where the name and
version are extracted from stackpack.conf and created in the current directory.`,
		Example: `# Package stackpack in current directory
sts stackpack package

# Package specific stackpack directory
sts stackpack package -d ./my-stackpack

# Package with custom archive filename
sts stackpack package -f my-custom-archive.zip

# Force overwrite existing zip file
sts stackpack package --force`,
		RunE: cli.CmdRunE(RunStackpackPackageCommand(args)),
	}

	cmd.Flags().StringVarP(&args.StackpackDir, "stackpack-directory", "d", "", "Path to stackpack directory (defaults to current directory)")
	cmd.Flags().StringVarP(&args.ArchiveFile, "archive-file", "f", "", "Path to the zip file to create (defaults to <stackpack_name>-<version>.zip in current directory)")
	cmd.Flags().BoolVar(&args.Force, "force", false, "Overwrite existing zip file without prompting")

	return cmd
}

// RunStackpackPackageCommand executes the package command
func RunStackpackPackageCommand(args *PackageArgs) func(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	return func(cli *di.Deps, cmd *cobra.Command) common.CLIError {
		// Set default stackpack directory
		if args.StackpackDir == "" {
			currentDir, err := os.Getwd()
			if err != nil {
				return common.NewRuntimeError(fmt.Errorf("failed to get current working directory: %w", err))
			}
			args.StackpackDir = currentDir
		}

		// Convert to absolute path
		absStackpackDir, err := filepath.Abs(args.StackpackDir)
		if err != nil {
			return common.NewRuntimeError(fmt.Errorf("failed to get absolute path for stackpack directory: %w", err))
		}
		args.StackpackDir = absStackpackDir

		// Parse stackpack.conf using HOCON parser to get name and version
		parser := &HoconParser{}
		stackpackInfo, err := parser.Parse(filepath.Join(args.StackpackDir, "stackpack.conf"))
		if err != nil {
			return common.NewRuntimeError(fmt.Errorf("failed to parse stackpack.conf: %w", err))
		}

		// Set default archive file path if not specified
		if args.ArchiveFile == "" {
			currentDir, err := os.Getwd()
			if err != nil {
				return common.NewRuntimeError(fmt.Errorf("failed to get current working directory: %w", err))
			}
			zipFileName := fmt.Sprintf("%s-%s.zip", stackpackInfo.Name, stackpackInfo.Version)
			args.ArchiveFile = filepath.Join(currentDir, zipFileName)
		} else {
			// Convert to absolute path
			absArchiveFile, err := filepath.Abs(args.ArchiveFile)
			if err != nil {
				return common.NewRuntimeError(fmt.Errorf("failed to get absolute path for archive file: %w", err))
			}
			args.ArchiveFile = absArchiveFile
		}

		// Validate stackpack directory
		if err := validateStackpackDirectory(args.StackpackDir); err != nil {
			return common.NewCLIArgParseError(err)
		}

		// Check if zip file exists and handle force flag
		if _, err := os.Stat(args.ArchiveFile); err == nil && !args.Force {
			return common.NewRuntimeError(fmt.Errorf("zip file already exists: %s (use --force to overwrite)", args.ArchiveFile))
		}

		// Create output directory if it doesn't exist
		outputDir := filepath.Dir(args.ArchiveFile)
		if err := os.MkdirAll(outputDir, os.FileMode(defaultDirMode)); err != nil {
			return common.NewRuntimeError(fmt.Errorf("failed to create output directory: %w", err))
		}

		// Create zip file
		if err := createStackpackZip(args.StackpackDir, args.ArchiveFile); err != nil {
			return common.NewRuntimeError(fmt.Errorf("failed to create zip file: %w", err))
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"success":           true,
				"stackpack_name":    stackpackInfo.Name,
				"stackpack_version": stackpackInfo.Version,
				"zip_file":          args.ArchiveFile,
				"source_dir":        args.StackpackDir,
			})
		} else {
			cli.Printer.Successf("Stackpack packaged successfully!")
			cli.Printer.PrintLn("")
			cli.Printer.PrintLn(fmt.Sprintf("Stackpack: %s (v%s)", stackpackInfo.Name, stackpackInfo.Version))
			cli.Printer.PrintLn(fmt.Sprintf("Zip file: %s", args.ArchiveFile))
		}

		return nil
	}
}

func validateStackpackDirectory(dir string) error {
	for _, item := range requiredStackpackItems {
		itemPath := filepath.Join(dir, item)
		if _, err := os.Stat(itemPath); err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("required stackpack item not found: %s", item)
			}
			return fmt.Errorf("failed to check stackpack item %s: %w", item, err)
		}
	}
	return nil
}

func createStackpackZip(sourceDir, zipPath string) error {
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Add each required item to the zip
	for _, item := range requiredStackpackItems {
		itemPath := filepath.Join(sourceDir, item)
		if err := addToZip(zipWriter, itemPath, item); err != nil {
			return fmt.Errorf("failed to add %s to zip: %w", item, err)
		}
	}

	return nil
}

func addToZip(zipWriter *zip.Writer, sourcePath, zipPath string) error {
	fileInfo, err := os.Stat(sourcePath)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return addDirToZip(zipWriter, sourcePath, zipPath)
	}

	return addFileToZip(zipWriter, sourcePath, zipPath)
}

func addFileToZip(zipWriter *zip.Writer, sourcePath, zipPath string) error {
	file, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer file.Close()

	zipFileWriter, err := zipWriter.Create(zipPath)
	if err != nil {
		return err
	}

	_, err = io.Copy(zipFileWriter, file)
	return err
}

func addDirToZip(zipWriter *zip.Writer, sourceDir, zipDir string) error {
	return filepath.Walk(sourceDir, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get relative path from source directory
		relPath, err := filepath.Rel(sourceDir, filePath)
		if err != nil {
			return err
		}

		// Skip the root directory itself
		if relPath == "." {
			return nil
		}

		// Create zip path with forward slashes for cross-platform compatibility
		zipPath := filepath.ToSlash(filepath.Join(zipDir, relPath))

		if fileInfo.IsDir() {
			// Create directory entry in zip
			_, err := zipWriter.Create(zipPath + "/")
			return err
		}

		// Add file to zip
		return addFileToZip(zipWriter, filePath, zipPath)
	})
}
