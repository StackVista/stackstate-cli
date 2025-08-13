package stackpack

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/pkg/scaffold"
)

const (
	defaultTemplateGitHubRepo = "StackVista/stackpack-templates" // Default GitHub repository for templates
	defaultTemplateGitHubRef  = "main"                           // Default branch for GitHub templates
	defaultTemplateGitHubPath = "templates"                      // Default path in GitHub repo for templates
	defaultTemplateName       = "generic"                        // Default template name to use
)

type ScaffoldArgs struct {
	// Local template source
	TemplateLocalDir string

	// GitHub template source
	TemplateGitHubRepo string // Format: "owner/repo"
	TemplateGitHubRef  string
	TemplateGitHubPath string

	// Common flags
	DestinationDir string
	Name           string
	DisplayName    string
	TemplateName   string
	Force          bool
}

func StackpackScaffoldCommand(cli *di.Deps) *cobra.Command {
	args := &ScaffoldArgs{}
	cmd := &cobra.Command{
		Use:   "scaffold",
		Short: "Create a stackpack skeleton from a template",
		Long: `Create a stackpack skeleton from a template.

This command scaffolds a new stackpack project structure from a template source.
The template can be from a local directory or a GitHub repository.
The template can be customized with the stackpack name and other variables.`,
		Example: `# Create a stackpack using defaults (uses default GitHub repo and template)
sts stackpack scaffold --name my-stackpack

# Create a stackpack from a local template (looks for ./templates/stackpack/ subdirectory)
sts stackpack scaffold --template-local-dir ./templates --name my-awesome-stackpack --template-name stackpack

# Overwrite existing files without prompting
sts stackpack scaffold --name my-awesome-stackpack --force

# Create a stackpack from a specific GitHub repository
sts stackpack scaffold --template-github-repo stackvista/my-templates --name my-awesome-stackpack --template-name generic`,
		RunE: cli.CmdRunE(RunStackpackScaffoldCommand(args)),
	}

	// Template source flags (mutually exclusive, defaults to GitHub repo if none specified)
	cmd.Flags().StringVar(&args.TemplateLocalDir, "template-local-dir", "", "Path to local directory containing template subdirectories")
	cmd.Flags().StringVar(&args.TemplateGitHubRepo, "template-github-repo", "", fmt.Sprintf("GitHub repository in format 'owner/repo' (default: %s)", defaultTemplateGitHubRepo))
	cmd.Flags().StringVar(&args.TemplateGitHubRef, "template-github-ref", "main", fmt.Sprintf("Git reference (branch, tag, or commit SHA) (default: %s)", defaultTemplateGitHubRef))
	cmd.Flags().StringVar(&args.TemplateGitHubPath, "template-github-path", "", fmt.Sprintf("Path within the repository containing template subdirectories (default: %s)", defaultTemplateGitHubPath))

	// Common flags
	cmd.Flags().StringVar(&args.DestinationDir, "destination-dir", "", "Target directory where scaffolded files will be created. If not specified, uses current working directory")
	cmd.Flags().StringVar(&args.Name, "name", "", "Name of the stackpack (required). Must start with [a-z] and contain only lowercase letters, digits, and hyphens")
	cmd.Flags().StringVar(&args.DisplayName, "display-name", "", "Name that's displayed on both the StackPack listing page and on the title of the StackPack page. If not provided, the value of --name will be used")
	cmd.Flags().StringVar(&args.TemplateName, "template-name", defaultTemplateName, fmt.Sprintf("Name of the template subdirectory to use (default: %s)", defaultTemplateName))
	cmd.Flags().BoolVar(&args.Force, "force", false, "Overwrite existing files without prompting")

	// Mark required flags
	cmd.MarkFlagRequired("name") //nolint:errcheck

	// Template sources are mutually exclusive but not required (will use default GitHub repo if none specified)
	stscobra.MarkMutexFlags(cmd, []string{"template-local-dir", "template-github-repo"}, "template-source", false)

	return cmd
}

func RunStackpackScaffoldCommand(args *ScaffoldArgs) func(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	return func(cli *di.Deps, cmd *cobra.Command) common.CLIError {
		// Create template source based on which source was specified
		var source scaffold.TemplateSource
		var err error

		if args.DestinationDir == "" {
			args.DestinationDir, err = os.Getwd()
			if err != nil {
				return common.NewRuntimeError(fmt.Errorf("failed to get current working directory: %w", err))
			}
		}

		// Validate stackpack name
		if err := validateStackpackName(args.Name); err != nil {
			return common.NewCLIArgParseError(err)
		}

		if args.TemplateLocalDir != "" {
			source = scaffold.NewLocalDirSource(args.TemplateLocalDir, args.TemplateName)
		} else {
			// Use GitHub repository (either specified or default)
			githubRepo := defaultIfEmptyString(args.TemplateGitHubRepo, defaultTemplateGitHubRepo)
			githubRef := defaultIfEmptyString(args.TemplateGitHubRef, defaultTemplateGitHubRef)
			githubPath := defaultIfEmptyString(args.TemplateGitHubPath, defaultTemplateGitHubPath)

			// Parse owner/repo format
			owner, repo, err := parseGitHubRepo(githubRepo)
			if err != nil {
				return common.NewCLIArgParseError(err)
			}
			source = scaffold.NewGitHubSource(owner, repo, githubRef, githubPath, args.TemplateName)
		}

		// Create template context
		displayName := args.DisplayName
		if displayName == "" {
			displayName = args.Name
		}
		context := scaffold.TemplateContext{
			Name:         args.Name,
			DisplayName:  displayName,
			TemplateName: args.TemplateName,
		}

		// Create scaffolder with force flag, printer, and JSON output mode
		scaffolder := scaffold.NewScaffolder(source, args.DestinationDir, context, args.Force, cli.Printer, cli.IsJson())
		// Execute scaffolding
		result, cleanUpFn, err := scaffolder.Scaffold(cmd.Context())
		if err != nil {
			return common.NewRuntimeError(err)
		}

		err = cleanUpFn()
		if err != nil {
			return common.NewRuntimeError(fmt.Errorf("failed to clean up temporary files: %w", err))
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"success":     result.Success,
				"source":      result.Source,
				"destination": result.Destination,
				"name":        result.Name,
				"template":    result.Template,
				"files_count": result.FilesCount,
				"files":       result.Files,
			})
		} else {
			// Display success message and next steps
			cli.Printer.Successf("✓ Scaffold complete!")
			cli.Printer.PrintLn("")
			displayNextSteps(cli, args)
		}

		return nil
	}
}

// parseGitHubRepo parses "owner/repo" format into separate owner and repo
func parseGitHubRepo(repoString string) (string, string, error) {
	parts := strings.Split(repoString, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("invalid GitHub repository format '%s', expected 'owner/repo'", repoString)
	}
	return parts[0], parts[1], nil
}

// validateStackpackName validates the stackpack name according to naming rules
func validateStackpackName(name string) error {
	// Pattern: starts with [a-z], followed by [a-z0-9-]*
	validNamePattern := regexp.MustCompile(`^[a-z][a-z0-9-]*$`)

	if !validNamePattern.MatchString(name) {
		return fmt.Errorf("invalid stackpack name '%s': must start with a lowercase letter [a-z] and contain only lowercase letters, digits, and hyphens", name)
	}

	return nil
}

func displayNextSteps(cli *di.Deps, args *ScaffoldArgs) {
	cli.Printer.PrintLn("Next steps:")
	cli.Printer.PrintLn("1. Review the generated files in: " + args.DestinationDir)
	cli.Printer.PrintLn(fmt.Sprintf("2. Check the %s for instructions on what to do next.", filepath.Join(args.DestinationDir, "README.md")))
}

func defaultIfEmptyString(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
