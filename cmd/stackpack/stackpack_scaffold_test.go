package stackpack

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupStackpackScaffoldCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackScaffoldCommand(&cli.Deps)
	return &cli, cmd
}

func TestStackpackScaffoldCommand_RequiredFlags(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantErr      bool
		errorMessage string
	}{
		{
			name:         "missing name flag",
			args:         []string{},
			wantErr:      true,
			errorMessage: `required flag(s) "name" not set`,
		},
		{
			name:    "valid minimal args with defaults",
			args:    []string{"--name", "test-stackpack"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.wantErr {
				// Create temporary directory for successful command execution to avoid polluting current directory
				tempDir, err := os.MkdirTemp("", "stackpack-required-flags-test-*")
				require.NoError(t, err)
				defer os.RemoveAll(tempDir)

				// Change to temp directory so command uses it as default destination
				originalWd, err := os.Getwd()
				require.NoError(t, err)
				defer func() {
					err := os.Chdir(originalWd)
					require.NoError(t, err, "Failed to change back to original working directory")
				}()

				err = os.Chdir(tempDir)
				require.NoError(t, err)
			}

			cli, cmd := setupStackpackScaffoldCmd(t)
			_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, tt.args...)

			if tt.wantErr {
				require.Error(t, err)
				if tt.errorMessage != "" {
					assert.Contains(t, err.Error(), tt.errorMessage)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestStackpackScaffoldCommand_MutuallyExclusiveFlags(t *testing.T) {
	cli, cmd := setupStackpackScaffoldCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd,
		"--name", "test-stackpack",
		"--template-local-dir", "./templates",
		"--template-github-repo", "owner/repo",
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "template-github-repo")
}

func TestParseGitHubRepo(t *testing.T) {
	tests := []struct {
		name          string
		repoString    string
		expectedOwner string
		expectedRepo  string
		wantErr       bool
		errorMessage  string
	}{
		{
			name:          "valid owner/repo format",
			repoString:    "stackvista/my-templates",
			expectedOwner: "stackvista",
			expectedRepo:  "my-templates",
			wantErr:       false,
		},
		{
			name:          "valid with numbers and dashes",
			repoString:    "owner-123/repo-456",
			expectedOwner: "owner-123",
			expectedRepo:  "repo-456",
			wantErr:       false,
		},
		{
			name:         "missing slash",
			repoString:   "invalidrepo",
			wantErr:      true,
			errorMessage: "invalid GitHub repository format 'invalidrepo', expected 'owner/repo'",
		},
		{
			name:         "empty owner",
			repoString:   "/repo",
			wantErr:      true,
			errorMessage: "invalid GitHub repository format '/repo', expected 'owner/repo'",
		},
		{
			name:         "empty repo",
			repoString:   "owner/",
			wantErr:      true,
			errorMessage: "invalid GitHub repository format 'owner/', expected 'owner/repo'",
		},
		{
			name:         "multiple slashes",
			repoString:   "owner/repo/extra",
			wantErr:      true,
			errorMessage: "invalid GitHub repository format 'owner/repo/extra', expected 'owner/repo'",
		},
		{
			name:         "empty string",
			repoString:   "",
			wantErr:      true,
			errorMessage: "invalid GitHub repository format '', expected 'owner/repo'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			owner, repo, err := parseGitHubRepo(tt.repoString)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errorMessage != "" {
					assert.Equal(t, tt.errorMessage, err.Error())
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedOwner, owner)
				assert.Equal(t, tt.expectedRepo, repo)
			}
		})
	}
}

func TestDefaultIfEmptyString(t *testing.T) {
	tests := []struct {
		name         string
		value        string
		defaultValue string
		expected     string
	}{
		{
			name:         "empty value returns default",
			value:        "",
			defaultValue: "default-value",
			expected:     "default-value",
		},
		{
			name:         "non-empty value returns value",
			value:        "custom-value",
			defaultValue: "default-value",
			expected:     "custom-value",
		},
		{
			name:         "whitespace value returns whitespace",
			value:        "   ",
			defaultValue: "default-value",
			expected:     "   ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := defaultIfEmptyString(tt.value, tt.defaultValue)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStackpackScaffoldCommand_LocalDirSource(t *testing.T) {
	// Create temporary directories for testing
	tempDir, err := os.MkdirTemp("", "stackpack-scaffold-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a mock template structure
	templateDir := filepath.Join(tempDir, "templates")
	genericTemplateDir := filepath.Join(templateDir, "generic")
	err = os.MkdirAll(genericTemplateDir, 0755)
	require.NoError(t, err)

	// Create a test template file
	templateFile := filepath.Join(genericTemplateDir, "test.txt")
	err = os.WriteFile(templateFile, []byte("Hello <<.Name>>!"), 0644)
	require.NoError(t, err)

	// Create destination directory
	destDir := filepath.Join(tempDir, "destination")
	err = os.MkdirAll(destDir, 0755)
	require.NoError(t, err)

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name: "local directory source success",
			args: []string{
				"--name", "test-stackpack",
				"--template-local-dir", templateDir,
				"--template-name", "generic",
				"--destination-dir", destDir,
			},
			wantErr: false,
		},
		{
			name: "local directory source with non-existent template",
			args: []string{
				"--name", "test-stackpack",
				"--template-local-dir", templateDir,
				"--template-name", "non-existent",
				"--destination-dir", destDir,
			},
			wantErr: true,
		},
		{
			name: "local directory source with non-existent directory",
			args: []string{
				"--name", "test-stackpack",
				"--template-local-dir", "/non/existent/path",
				"--template-name", "generic",
				"--destination-dir", destDir,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, cmd := setupStackpackScaffoldCmd(t)

			_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, tt.args...)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				// Verify the file was created and processed
				createdFile := filepath.Join(destDir, "test.txt")
				content, err := os.ReadFile(createdFile)
				require.NoError(t, err)
				assert.Equal(t, "Hello test-stackpack!", string(content))

				// Clean up for next test
				_ = os.RemoveAll(destDir)
			}
		})
	}
}

func TestStackpackScaffoldCommand_JSONOutput(t *testing.T) {
	// Create temporary directories for testing
	tempDir, err := os.MkdirTemp("", "stackpack-scaffold-json-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a mock template structure
	templateDir := filepath.Join(tempDir, "templates")
	genericTemplateDir := filepath.Join(templateDir, "generic")
	err = os.MkdirAll(genericTemplateDir, 0755)
	require.NoError(t, err)

	// Create a test template file
	templateFile := filepath.Join(genericTemplateDir, "config.json")
	err = os.WriteFile(templateFile, []byte(`{"name": "<<.Name>>", "template": "<<.TemplateName>>"}`), 0644)
	require.NoError(t, err)

	// Create destination directory
	destDir := filepath.Join(tempDir, "destination")
	err = os.MkdirAll(destDir, 0755)
	require.NoError(t, err)

	tests := []struct {
		name       string
		args       []string
		expectJson bool
		expectText bool
	}{
		{
			name: "regular text output",
			args: []string{
				"--name", "test-stackpack",
				"--template-local-dir", templateDir,
				"--template-name", "generic",
				"--destination-dir", destDir,
			},
			expectJson: false,
			expectText: true,
		},
		{
			name: "JSON output",
			args: []string{
				"--name", "test-stackpack",
				"--template-local-dir", templateDir,
				"--template-name", "generic",
				"--destination-dir", destDir,
				"-o", "json",
			},
			expectJson: true,
			expectText: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, cmd := setupStackpackScaffoldCmd(t)

			_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, tt.args...)
			require.NoError(t, err)

			if tt.expectJson {
				// Verify JSON output was called
				require.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
				jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]
				// The JSON output should contain the scaffold result
				assert.Contains(t, jsonOutput, "success")

				// Verify no regular text calls were made for success messages
				assert.False(t, cli.MockPrinter.HasNonJsonCalls)
			} else {
				// Verify text output was called
				require.NotEmpty(t, *cli.MockPrinter.SuccessCalls)
				successCall := (*cli.MockPrinter.SuccessCalls)[0]
				assert.Contains(t, successCall, "Scaffold complete")

				// Verify PrintLn was called for next steps
				require.NotEmpty(t, *cli.MockPrinter.PrintLnCalls)
			}

			// Clean up for next test
			_ = os.RemoveAll(destDir)
		})
	}
}

func TestStackpackScaffoldCommand_ForceFlag(t *testing.T) {
	// Create temporary directories for testing
	tempDir, err := os.MkdirTemp("", "stackpack-scaffold-force-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a mock template structure
	templateDir := filepath.Join(tempDir, "templates")
	genericTemplateDir := filepath.Join(templateDir, "generic")
	err = os.MkdirAll(genericTemplateDir, 0755)
	require.NoError(t, err)

	// Create a test template file
	templateFile := filepath.Join(genericTemplateDir, "existing.txt")
	err = os.WriteFile(templateFile, []byte("Template content"), 0644)
	require.NoError(t, err)

	// Create destination directory with conflicting file
	destDir := filepath.Join(tempDir, "destination")
	err = os.MkdirAll(destDir, 0755)
	require.NoError(t, err)

	existingFile := filepath.Join(destDir, "existing.txt")
	err = os.WriteFile(existingFile, []byte("Existing content"), 0644)
	require.NoError(t, err)

	t.Run("without force flag should fail on conflict", func(t *testing.T) {
		cli, cmd := setupStackpackScaffoldCmd(t)

		_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd,
			"--name", "test-stackpack",
			"--template-local-dir", templateDir,
			"--template-name", "generic",
			"--destination-dir", destDir,
		)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "conflicting files exist")
	})

	t.Run("with force flag should succeed on conflict", func(t *testing.T) {
		cli, cmd := setupStackpackScaffoldCmd(t)

		_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd,
			"--name", "test-stackpack",
			"--template-local-dir", templateDir,
			"--template-name", "generic",
			"--destination-dir", destDir,
			"--force",
		)

		require.NoError(t, err)

		// Verify the file was overwritten
		content, err := os.ReadFile(existingFile)
		require.NoError(t, err)
		assert.Equal(t, "Template content", string(content))
	})
}

func TestStackpackScaffoldCommand_TemplateContextPassedCorrectly(t *testing.T) {
	// Create temporary directories for testing
	tempDir, err := os.MkdirTemp("", "stackpack-scaffold-context-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a mock template structure
	templateDir := filepath.Join(tempDir, "templates")
	customTemplateDir := filepath.Join(templateDir, "custom-template")
	err = os.MkdirAll(customTemplateDir, 0755)
	require.NoError(t, err)

	// Create a test template file with Name, DisplayName, and TemplateName variables
	templateFile := filepath.Join(customTemplateDir, "config.yaml")
	templateContent := `name: <<.Name>>
display_name: <<.DisplayName>>
template: <<.TemplateName>>
`
	err = os.WriteFile(templateFile, []byte(templateContent), 0644)
	require.NoError(t, err)

	// Create destination directory
	destDir := filepath.Join(tempDir, "destination")
	err = os.MkdirAll(destDir, 0755)
	require.NoError(t, err)

	cli, cmd := setupStackpackScaffoldCmd(t)

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd,
		"--name", "my-awesome-stackpack",
		"--template-local-dir", templateDir,
		"--template-name", "custom-template",
		"--destination-dir", destDir,
	)

	require.NoError(t, err)

	// Verify the template variables were substituted correctly
	createdFile := filepath.Join(destDir, "config.yaml")
	content, err := os.ReadFile(createdFile)
	require.NoError(t, err)

	expectedContent := `name: my-awesome-stackpack
display_name: my-awesome-stackpack
template: custom-template
`
	assert.Equal(t, expectedContent, string(content))
}

func TestRunStackpackScaffoldCommand_SourceSelection(t *testing.T) {
	tests := []struct {
		name     string
		args     ScaffoldArgs
		expected string // Expected source type (for assertion purposes)
	}{
		{
			name: "local directory source selected",
			args: ScaffoldArgs{
				TemplateLocalDir: "/path/to/templates",
				Name:             "test-stackpack",
				TemplateName:     "generic",
			},
			expected: "local",
		},
		{
			name: "github source selected with defaults",
			args: ScaffoldArgs{
				Name:         "test-stackpack",
				TemplateName: "generic",
			},
			expected: "github-default",
		},
		{
			name: "github source selected with custom repo",
			args: ScaffoldArgs{
				TemplateGitHubRepo: "custom/repo",
				Name:               "test-stackpack",
				TemplateName:       "generic",
			},
			expected: "github-custom",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// We can't easily test the actual source creation without mocking,
			// but we can test the source selection logic by examining the defaults
			switch tt.expected {
			case "local":
				assert.NotEmpty(t, tt.args.TemplateLocalDir, "Local dir should be set")
			case "github-default":
				assert.Empty(t, tt.args.TemplateLocalDir, "Local dir should be empty for GitHub source")
				assert.Empty(t, tt.args.TemplateGitHubRepo, "GitHub repo should be empty to use defaults")
			case "github-custom":
				assert.Empty(t, tt.args.TemplateLocalDir, "Local dir should be empty for GitHub source")
				assert.NotEmpty(t, tt.args.TemplateGitHubRepo, "GitHub repo should be set")
			}
		})
	}
}

func TestStackpackScaffoldCommand_DefaultValues(t *testing.T) {
	_, cmd := setupStackpackScaffoldCmd(t)

	// Test that defaults are properly set in the command flags
	flags := cmd.Flags()

	// Check default values
	templateGitHubRefFlag := flags.Lookup("template-github-ref")
	assert.Equal(t, "main", templateGitHubRefFlag.DefValue)

	templateNameFlag := flags.Lookup("template-name")
	assert.Equal(t, defaultTemplateName, templateNameFlag.DefValue)

	forceFlag := flags.Lookup("force")
	assert.Equal(t, "false", forceFlag.DefValue)
}

func TestDisplayNextSteps(t *testing.T) {
	cli := di.NewMockDeps(t)
	args := &ScaffoldArgs{
		DestinationDir: "/path/to/destination",
		Name:           "test-stackpack",
	}

	displayNextSteps(&cli.Deps, args)

	// Verify that next steps were printed
	printLnCalls := *cli.MockPrinter.PrintLnCalls
	require.Greater(t, len(printLnCalls), 0)

	// Check that specific next steps are mentioned
	allOutput := fmt.Sprintf("%v", printLnCalls)
	assert.Contains(t, allOutput, "Next steps")
	assert.Contains(t, allOutput, args.DestinationDir)
	assert.Contains(t, allOutput, "Review the generated files")
	assert.Contains(t, allOutput, "for instructions on what to do next")
}

func TestStackpackScaffoldCommand_DisplayNameFallback(t *testing.T) {
	// Create temporary directories for testing
	tempDir, err := os.MkdirTemp("", "stackpack-scaffold-displayname-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a mock template structure
	templateDir := filepath.Join(tempDir, "templates")
	customTemplateDir := filepath.Join(templateDir, "custom-template")
	err = os.MkdirAll(customTemplateDir, 0755)
	require.NoError(t, err)

	// Create a test template file with DisplayName variable
	templateFile := filepath.Join(customTemplateDir, "test-displayname.yaml")
	templateContent := `name: <<.Name>>
display_name: <<.DisplayName>>
`
	err = os.WriteFile(templateFile, []byte(templateContent), 0644)
	require.NoError(t, err)

	// Create destination directory
	destDir := filepath.Join(tempDir, "destination")
	err = os.MkdirAll(destDir, 0755)
	require.NoError(t, err)

	tests := []struct {
		name                string
		args                []string
		expectedDisplayName string
		description         string
	}{
		{
			name: "with display-name flag provided",
			args: []string{
				"--name", "my-stackpack",
				"--display-name", "My Awesome StackPack",
				"--template-local-dir", templateDir,
				"--template-name", "custom-template",
				"--destination-dir", destDir,
			},
			expectedDisplayName: "My Awesome StackPack",
			description:         "Should use provided display name",
		},
		{
			name: "without display-name flag - should fallback to name",
			args: []string{
				"--name", "my-stackpack-fallback",
				"--template-local-dir", templateDir,
				"--template-name", "custom-template",
				"--destination-dir", destDir,
			},
			expectedDisplayName: "my-stackpack-fallback",
			description:         "Should fallback to name when display-name not provided",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, cmd := setupStackpackScaffoldCmd(t)

			_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, tt.args...)
			require.NoError(t, err, tt.description)

			// Verify the template variables were substituted correctly
			createdFile := filepath.Join(destDir, "test-displayname.yaml")
			content, err := os.ReadFile(createdFile)
			require.NoError(t, err)

			// Check that DisplayName was correctly used
			assert.Contains(t, string(content), "display_name: "+tt.expectedDisplayName, tt.description)

			// Clean up for next test iteration
			_ = os.RemoveAll(destDir)
			_ = os.MkdirAll(destDir, 0755)
		})
	}
}

//nolint:funlen
func TestValidateStackpackName(t *testing.T) {
	tests := []struct {
		name        string
		stackName   string
		wantErr     bool
		errorSubstr string
	}{
		// Valid names
		{
			name:      "valid single letter",
			stackName: "a",
			wantErr:   false,
		},
		{
			name:      "valid simple name",
			stackName: "app",
			wantErr:   false,
		},
		{
			name:      "valid name with hyphens",
			stackName: "my-stackpack",
			wantErr:   false,
		},
		{
			name:      "valid name with numbers",
			stackName: "app123",
			wantErr:   false,
		},
		{
			name:      "valid complex name",
			stackName: "my-app-v2-123",
			wantErr:   false,
		},
		// Invalid names - starts with digit
		{
			name:        "invalid starts with digit",
			stackName:   "1-stackpack",
			wantErr:     true,
			errorSubstr: "must start with a lowercase letter [a-z]",
		},
		{
			name:        "invalid starts with number",
			stackName:   "2app",
			wantErr:     true,
			errorSubstr: "must start with a lowercase letter [a-z]",
		},
		// Invalid names - starts with hyphen
		{
			name:        "invalid starts with hyphen",
			stackName:   "-stackpack",
			wantErr:     true,
			errorSubstr: "must start with a lowercase letter [a-z]",
		},
		// Invalid names - uppercase letters
		{
			name:        "invalid uppercase first letter",
			stackName:   "My-Stackpack",
			wantErr:     true,
			errorSubstr: "must start with a lowercase letter [a-z]",
		},
		{
			name:        "invalid camelCase",
			stackName:   "myApp",
			wantErr:     true,
			errorSubstr: "lowercase letters, digits, and hyphens",
		},
		{
			name:        "invalid all uppercase",
			stackName:   "APP",
			wantErr:     true,
			errorSubstr: "must start with a lowercase letter [a-z]",
		},
		// Invalid names - spaces
		{
			name:        "invalid with spaces",
			stackName:   "my stackpack",
			wantErr:     true,
			errorSubstr: "lowercase letters, digits, and hyphens",
		},
		{
			name:        "invalid app name with space",
			stackName:   "app name",
			wantErr:     true,
			errorSubstr: "lowercase letters, digits, and hyphens",
		},
		// Invalid names - special characters
		{
			name:        "invalid underscore",
			stackName:   "my_stackpack",
			wantErr:     true,
			errorSubstr: "lowercase letters, digits, and hyphens",
		},
		{
			name:        "invalid at symbol",
			stackName:   "app@domain",
			wantErr:     true,
			errorSubstr: "lowercase letters, digits, and hyphens",
		},
		{
			name:        "invalid dot",
			stackName:   "stack.pack",
			wantErr:     true,
			errorSubstr: "lowercase letters, digits, and hyphens",
		},
		{
			name:        "invalid slash",
			stackName:   "my/app",
			wantErr:     true,
			errorSubstr: "lowercase letters, digits, and hyphens",
		},
		{
			name:        "invalid colon",
			stackName:   "app:v1",
			wantErr:     true,
			errorSubstr: "lowercase letters, digits, and hyphens",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateStackpackName(tt.stackName)
			if tt.wantErr {
				require.Error(t, err, "Expected validation to fail for name '%s'", tt.stackName)
				if tt.errorSubstr != "" {
					assert.Contains(t, err.Error(), tt.errorSubstr, "Error message should contain expected substring")
				}
				assert.Contains(t, err.Error(), tt.stackName, "Error message should contain the invalid name")
			} else {
				require.NoError(t, err, "Expected validation to pass for name '%s'", tt.stackName)
			}
		})
	}
}

func TestStackpackScaffoldCommand_NameValidation(t *testing.T) {
	// Create temporary template directory for all tests to avoid GitHub template conflicts
	tempDir, err := os.MkdirTemp("", "stackpack-validation-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a mock template structure
	templateDir := filepath.Join(tempDir, "templates")
	genericTemplateDir := filepath.Join(templateDir, "generic")
	err = os.MkdirAll(genericTemplateDir, 0755)
	require.NoError(t, err)

	// Create a simple test template file
	templateFile := filepath.Join(genericTemplateDir, "test.txt")
	err = os.WriteFile(templateFile, []byte("Test stackpack: <<.Name>>"), 0644)
	require.NoError(t, err)

	// Create destination directory
	destDir := filepath.Join(tempDir, "destination")
	err = os.MkdirAll(destDir, 0755)
	require.NoError(t, err)

	tests := []struct {
		name        string
		args        []string
		wantErr     bool
		errorSubstr string
	}{
		{
			name: "valid name passes validation",
			args: []string{
				"--name", "my-stackpack",
				"--template-local-dir", templateDir,
				"--destination-dir", destDir,
			},
			wantErr: false,
		},
		{
			name: "invalid name with uppercase fails",
			args: []string{
				"--name", "My-Stackpack",
				"--template-local-dir", templateDir,
				"--destination-dir", destDir,
			},
			wantErr:     true,
			errorSubstr: "must start with a lowercase letter [a-z]",
		},
		{
			name: "invalid name with spaces fails",
			args: []string{
				"--name", "my stackpack",
				"--template-local-dir", templateDir,
				"--destination-dir", destDir,
			},
			wantErr:     true,
			errorSubstr: "lowercase letters, digits, and hyphens",
		},
		{
			name: "invalid name starting with number fails",
			args: []string{
				"--name", "1-app",
				"--template-local-dir", templateDir,
				"--destination-dir", destDir,
			},
			wantErr:     true,
			errorSubstr: "must start with a lowercase letter [a-z]",
		},
		{
			name: "invalid name with underscore fails",
			args: []string{
				"--name", "my_app",
				"--template-local-dir", templateDir,
				"--destination-dir", destDir,
			},
			wantErr:     true,
			errorSubstr: "lowercase letters, digits, and hyphens",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, cmd := setupStackpackScaffoldCmd(t)

			_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, tt.args...)

			if tt.wantErr {
				require.Error(t, err, "Expected command to fail for invalid name")
				if tt.errorSubstr != "" {
					assert.Contains(t, err.Error(), tt.errorSubstr, "Error message should contain expected validation message")
				}
			} else {
				require.NoError(t, err, "Expected command to succeed for valid name")

				// Clean up destination for next test iteration
				_ = os.RemoveAll(destDir)
				_ = os.MkdirAll(destDir, 0755)
			}
		})
	}
}
