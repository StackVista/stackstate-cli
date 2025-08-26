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

// setupStackPackPackageCmd creates a test command with mock dependencies
func setupStackPackPackageCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackPackageCommand(&cli.Deps)
	return &cli, cmd
}

// createTestStackpack creates a valid stackpack directory structure for testing
func createTestStackpack(t *testing.T, dir string, name string, version string) {
	// Create required directories
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "provisioning"), 0755))
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "resources"), 0755))

	// Create stackpack.conf
	stackpackConf := fmt.Sprintf(`# schemaVersion -- Stackpack specification version.
schemaVersion = "2.0"
# name -- Name of the StackPack.
name = "%s"
# displayName -- Display name of the StackPack.
displayName = "Test %s"
# version -- Semantic version of the StackPack.
version = "%s"
`, name, name, version)
	require.NoError(t, os.WriteFile(filepath.Join(dir, "stackpack.conf"), []byte(stackpackConf), 0644))

	// Create README.md
	readme := fmt.Sprintf("# %s\n\nThis is a test stackpack.", name)
	require.NoError(t, os.WriteFile(filepath.Join(dir, "README.md"), []byte(readme), 0644))

	// Create test files in subdirectories
	require.NoError(t, os.WriteFile(filepath.Join(dir, "provisioning", "test.sty"), []byte("test provisioning"), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "resources", "overview.md"), []byte("test overview"), 0644))
}

func TestStackpackPackageCommand_DefaultBehavior(t *testing.T) {
	// Create temporary directory structure
	tempDir, err := os.MkdirTemp("", "stackpack-package-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackDir := filepath.Join(tempDir, "test-stackpack")
	require.NoError(t, os.MkdirAll(stackpackDir, 0755))
	createTestStackpack(t, stackpackDir, "test-stackpack", "1.0.0")

	// Change to temp directory to test default current directory behavior
	originalWd, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		err := os.Chdir(originalWd)
		require.NoError(t, err)
	}()
	err = os.Chdir(tempDir)
	require.NoError(t, err)

	cli, cmd := setupStackPackPackageCmd(t)

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "-d", stackpackDir)
	require.NoError(t, err)

	// Verify zip file was created in current directory
	expectedZipPath := filepath.Join(tempDir, "test-stackpack-1.0.0.zip")
	_, err = os.Stat(expectedZipPath)
	assert.NoError(t, err, "Zip file should be created")

	// Verify text output
	require.NotEmpty(t, *cli.MockPrinter.SuccessCalls)
	successCall := (*cli.MockPrinter.SuccessCalls)[0]
	assert.Contains(t, successCall, "✓ Stackpack packaged successfully!")

	require.NotEmpty(t, *cli.MockPrinter.PrintLnCalls)
	printLnCalls := *cli.MockPrinter.PrintLnCalls
	allOutput := fmt.Sprintf("%v", printLnCalls)
	assert.Contains(t, allOutput, "test-stackpack (v1.0.0)")
	assert.Contains(t, allOutput, expectedZipPath)
}

func TestStackpackPackageCommand_CustomArchiveFile(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "stackpack-package-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackDir := filepath.Join(tempDir, "my-stackpack")
	require.NoError(t, os.MkdirAll(stackpackDir, 0755))
	createTestStackpack(t, stackpackDir, "my-stackpack", "2.1.0")

	customZipPath := filepath.Join(tempDir, "custom-name.zip")
	cli, cmd := setupStackPackPackageCmd(t)

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "-d", stackpackDir, "-f", customZipPath)
	require.NoError(t, err)

	// Verify zip file was created with custom name
	_, err = os.Stat(customZipPath)
	assert.NoError(t, err, "Custom zip file should be created")

	// Verify output mentions custom path
	printLnCalls := *cli.MockPrinter.PrintLnCalls
	allOutput := fmt.Sprintf("%v", printLnCalls)
	assert.Contains(t, allOutput, customZipPath)
}

func TestStackpackPackageCommand_ForceFlag(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "stackpack-package-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackDir := filepath.Join(tempDir, "test-stackpack")
	require.NoError(t, os.MkdirAll(stackpackDir, 0755))
	createTestStackpack(t, stackpackDir, "test-stackpack", "1.0.0")

	zipPath := filepath.Join(tempDir, "test-stackpack-1.0.0.zip")

	// Create existing zip file
	require.NoError(t, os.WriteFile(zipPath, []byte("existing content"), 0644))

	cli, cmd := setupStackPackPackageCmd(t)

	// Test without force flag - should fail
	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "-d", stackpackDir, "-f", zipPath)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "zip file already exists")
	assert.Contains(t, err.Error(), "--force")

	// Test with force flag - should succeed
	cli2, cmd2 := setupStackPackPackageCmd(t)
	_, err = di.ExecuteCommandWithContext(&cli2.Deps, cmd2, "-d", stackpackDir, "-f", zipPath, "--force")
	require.NoError(t, err)

	// Verify success message
	require.NotEmpty(t, *cli2.MockPrinter.SuccessCalls)
	successCall := (*cli2.MockPrinter.SuccessCalls)[0]
	assert.Contains(t, successCall, "✓ Stackpack packaged successfully!")
}

func TestStackpackPackageCommand_JSONOutput(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "stackpack-package-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackDir := filepath.Join(tempDir, "json-test")
	require.NoError(t, os.MkdirAll(stackpackDir, 0755))
	createTestStackpack(t, stackpackDir, "json-test", "3.0.0")

	// Change to temp directory to ensure zip is created there
	originalWd, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		err := os.Chdir(originalWd)
		require.NoError(t, err)
	}()
	err = os.Chdir(tempDir)
	require.NoError(t, err)

	cli, cmd := setupStackPackPackageCmd(t)

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "-d", stackpackDir, "-o", "json")
	require.NoError(t, err)

	// Verify JSON output
	require.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]

	// Check JSON structure
	assert.Equal(t, true, jsonOutput["success"])
	assert.Equal(t, "json-test", jsonOutput["stackpack_name"])
	assert.Equal(t, "3.0.0", jsonOutput["stackpack_version"])
	assert.Contains(t, jsonOutput["zip_file"], "json-test-3.0.0.zip")
	assert.Contains(t, jsonOutput["source_dir"], stackpackDir)

	// Verify no text output for JSON mode
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestStackpackPackageCommand_MissingRequiredFiles(t *testing.T) {
	tests := []struct {
		name          string
		setupFunc     func(dir string)
		expectedError string
	}{
		{
			name: "missing provisioning directory",
			setupFunc: func(dir string) {
				require.NoError(t, os.MkdirAll(filepath.Join(dir, "resources"), 0755))
				require.NoError(t, os.WriteFile(filepath.Join(dir, "README.md"), []byte("readme"), 0644))
				require.NoError(t, os.WriteFile(filepath.Join(dir, "stackpack.conf"), []byte("name = \"test\"\nversion = \"1.0.0\""), 0644))
			},
			expectedError: "required stackpack item not found: provisioning",
		},
		{
			name: "missing README.md file",
			setupFunc: func(dir string) {
				require.NoError(t, os.MkdirAll(filepath.Join(dir, "provisioning"), 0755))
				require.NoError(t, os.MkdirAll(filepath.Join(dir, "resources"), 0755))
				require.NoError(t, os.WriteFile(filepath.Join(dir, "stackpack.conf"), []byte("name = \"test\"\nversion = \"1.0.0\""), 0644))
			},
			expectedError: "required stackpack item not found: README.md",
		},
		{
			name: "missing resources directory",
			setupFunc: func(dir string) {
				require.NoError(t, os.MkdirAll(filepath.Join(dir, "provisioning"), 0755))
				require.NoError(t, os.WriteFile(filepath.Join(dir, "README.md"), []byte("readme"), 0644))
				require.NoError(t, os.WriteFile(filepath.Join(dir, "stackpack.conf"), []byte("name = \"test\"\nversion = \"1.0.0\""), 0644))
			},
			expectedError: "required stackpack item not found: resources",
		},
		{
			name: "missing stackpack.conf file",
			setupFunc: func(dir string) {
				require.NoError(t, os.MkdirAll(filepath.Join(dir, "provisioning"), 0755))
				require.NoError(t, os.MkdirAll(filepath.Join(dir, "resources"), 0755))
				require.NoError(t, os.WriteFile(filepath.Join(dir, "README.md"), []byte("readme"), 0644))
			},
			expectedError: "failed to parse stackpack.conf",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "stackpack-package-validation-test-*")
			require.NoError(t, err)
			defer os.RemoveAll(tempDir)

			stackpackDir := filepath.Join(tempDir, "invalid-stackpack")
			require.NoError(t, os.MkdirAll(stackpackDir, 0755))
			tt.setupFunc(stackpackDir)

			cli, cmd := setupStackPackPackageCmd(t)

			_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "-d", stackpackDir)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedError)
		})
	}
}

func TestStackpackPackageCommand_InvalidStackpackConf(t *testing.T) {
	tests := []struct {
		name          string
		confContent   string
		expectedError string
	}{
		{
			name:          "invalid HOCON syntax",
			confContent:   `name = "test" invalid syntax {`,
			expectedError: "failed to parse stackpack.conf file",
		},
		{
			name:          "missing name field",
			confContent:   `version = "1.0.0"`,
			expectedError: "name not found in stackpack.conf",
		},
		{
			name:          "missing version field",
			confContent:   `name = "test"`,
			expectedError: "version not found in stackpack.conf",
		},
		{
			name: "empty name field",
			confContent: `name = ""
version = "1.0.0"`,
			expectedError: "name not found in stackpack.conf",
		},
		{
			name: "empty version field",
			confContent: `name = "test"
version = ""`,
			expectedError: "version not found in stackpack.conf",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "stackpack-package-hocon-test-*")
			require.NoError(t, err)
			defer os.RemoveAll(tempDir)

			stackpackDir := filepath.Join(tempDir, "invalid-stackpack")
			require.NoError(t, os.MkdirAll(stackpackDir, 0755))

			// Create required directories and files except stackpack.conf
			require.NoError(t, os.MkdirAll(filepath.Join(stackpackDir, "provisioning"), 0755))
			require.NoError(t, os.MkdirAll(filepath.Join(stackpackDir, "resources"), 0755))
			require.NoError(t, os.WriteFile(filepath.Join(stackpackDir, "README.md"), []byte("readme"), 0644))

			// Create invalid stackpack.conf
			require.NoError(t, os.WriteFile(filepath.Join(stackpackDir, "stackpack.conf"), []byte(tt.confContent), 0644))

			cli, cmd := setupStackPackPackageCmd(t)

			_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "-d", stackpackDir)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedError)
		})
	}
}

func TestStackpackPackageCommand_NonExistentDirectory(t *testing.T) {
	cli, cmd := setupStackPackPackageCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "-d", "/non/existent/directory")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse stackpack.conf")
	assert.Contains(t, err.Error(), "no such file or directory")
}

func TestStackpackPackageCommand_CreateOutputDirectory(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "stackpack-package-output-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackDir := filepath.Join(tempDir, "test-stackpack")
	require.NoError(t, os.MkdirAll(stackpackDir, 0755))
	createTestStackpack(t, stackpackDir, "test-stackpack", "1.0.0")

	// Create nested output directory path that doesn't exist
	outputDir := filepath.Join(tempDir, "output", "nested", "path")
	zipPath := filepath.Join(outputDir, "custom.zip")

	cli, cmd := setupStackPackPackageCmd(t)

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "-d", stackpackDir, "-f", zipPath)
	require.NoError(t, err)

	// Verify output directory was created
	_, err = os.Stat(outputDir)
	assert.NoError(t, err, "Output directory should be created")

	// Verify zip file was created
	_, err = os.Stat(zipPath)
	assert.NoError(t, err, "Zip file should be created in nested directory")
}

func TestHoconParser_Parse(t *testing.T) {
	tests := []struct {
		name          string
		content       string
		expectedName  string
		expectedVer   string
		expectError   bool
		errorContains string
	}{
		{
			name: "valid HOCON with quotes",
			content: `name = "my-stackpack"
version = "1.2.3"`,
			expectedName: "my-stackpack",
			expectedVer:  "1.2.3",
			expectError:  false,
		},
		{
			name: "valid HOCON without quotes",
			content: `name = my-stackpack
version = "1.2.3"`,
			expectedName: "my-stackpack",
			expectedVer:  "1.2.3",
			expectError:  false,
		},
		{
			name: "HOCON with comments",
			content: `# This is a comment
name = "test-app"
# Another comment
version = "2.0.0"`,
			expectedName: "test-app",
			expectedVer:  "2.0.0",
			expectError:  false,
		},
		{
			name:          "missing name",
			content:       `version = "1.0.0"`,
			expectError:   true,
			errorContains: "name not found in stackpack.conf",
		},
		{
			name:          "missing version",
			content:       `name = "test"`,
			expectError:   true,
			errorContains: "version not found in stackpack.conf",
		},
		{
			name:          "invalid HOCON syntax",
			content:       `name = "test" { invalid`,
			expectError:   true,
			errorContains: "failed to parse stackpack.conf file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "hocon-parser-test-*")
			require.NoError(t, err)
			defer os.RemoveAll(tempDir)

			confPath := filepath.Join(tempDir, "test.conf")
			require.NoError(t, os.WriteFile(confPath, []byte(tt.content), 0644))

			parser := &HoconParser{}
			result, err := parser.Parse(confPath)

			if tt.expectError {
				require.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				require.NotNil(t, result)
				assert.Equal(t, tt.expectedName, result.Name)
				assert.Equal(t, tt.expectedVer, result.Version)
			}
		})
	}
}

func TestHoconParser_ParseNonExistentFile(t *testing.T) {
	parser := &HoconParser{}
	result, err := parser.Parse("/non/existent/file.conf")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read file")
	assert.Nil(t, result)
}

func TestYamlParser_Parse(t *testing.T) {
	parser := &YamlParser{}
	result, err := parser.Parse("any-path")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "YAML format not yet implemented")
	assert.Nil(t, result)
}

func TestValidateStackpackDirectory(t *testing.T) {
	tests := []struct {
		name          string
		setupFunc     func(dir string)
		expectError   bool
		errorContains string
	}{
		{
			name: "valid stackpack directory",
			setupFunc: func(dir string) {
				require.NoError(t, os.MkdirAll(filepath.Join(dir, "provisioning"), 0755))
				require.NoError(t, os.MkdirAll(filepath.Join(dir, "resources"), 0755))
				require.NoError(t, os.WriteFile(filepath.Join(dir, "README.md"), []byte("readme"), 0644))
				require.NoError(t, os.WriteFile(filepath.Join(dir, "stackpack.conf"), []byte("conf"), 0644))
			},
			expectError: false,
		},
		{
			name: "missing provisioning directory",
			setupFunc: func(dir string) {
				require.NoError(t, os.MkdirAll(filepath.Join(dir, "resources"), 0755))
				require.NoError(t, os.WriteFile(filepath.Join(dir, "README.md"), []byte("readme"), 0644))
				require.NoError(t, os.WriteFile(filepath.Join(dir, "stackpack.conf"), []byte("conf"), 0644))
			},
			expectError:   true,
			errorContains: "required stackpack item not found: provisioning",
		},
		{
			name: "missing all required items",
			setupFunc: func(dir string) {
				// Create empty directory
			},
			expectError:   true,
			errorContains: "required stackpack item not found: provisioning",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "validate-stackpack-test-*")
			require.NoError(t, err)
			defer os.RemoveAll(tempDir)

			tt.setupFunc(tempDir)

			err = validateStackpackDirectory(tempDir)

			if tt.expectError {
				require.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

//nolint:funlen
func TestStackpackPackageCommand_TextOutput(t *testing.T) {
	tests := []struct {
		name       string
		outputFlag string
		expectText bool
		expectJson bool
	}{
		{
			name:       "default text output",
			outputFlag: "",
			expectText: true,
			expectJson: false,
		},
		{
			name:       "explicit text output",
			outputFlag: "text",
			expectText: true,
			expectJson: false,
		},
		{
			name:       "json output",
			outputFlag: "json",
			expectText: false,
			expectJson: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create unique temp directory for each subtest
			tempDir, err := os.MkdirTemp("", "stackpack-package-text-test-*")
			require.NoError(t, err)
			defer os.RemoveAll(tempDir)

			stackpackDir := filepath.Join(tempDir, "text-test")
			require.NoError(t, os.MkdirAll(stackpackDir, 0755))
			createTestStackpack(t, stackpackDir, "text-test", "4.0.0")

			// Change to temp directory to ensure zip is created there
			originalWd, err := os.Getwd()
			require.NoError(t, err)
			defer func() {
				err := os.Chdir(originalWd)
				require.NoError(t, err)
			}()
			err = os.Chdir(tempDir)
			require.NoError(t, err)

			cli, cmd := setupStackPackPackageCmd(t)

			args := []string{"-d", stackpackDir}
			if tt.outputFlag != "" {
				args = append(args, "-o", tt.outputFlag)
			}

			_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, args...)
			require.NoError(t, err)

			if tt.expectText {
				// Verify success message
				require.NotEmpty(t, *cli.MockPrinter.SuccessCalls)
				successCall := (*cli.MockPrinter.SuccessCalls)[0]
				assert.Contains(t, successCall, "✓ Stackpack packaged successfully!")

				// Verify stackpack info is printed
				printLnCalls := *cli.MockPrinter.PrintLnCalls
				allOutput := fmt.Sprintf("%v", printLnCalls)
				assert.Contains(t, allOutput, "text-test (v4.0.0)")

				// Should not have JSON calls
				assert.Empty(t, *cli.MockPrinter.PrintJsonCalls)
			}

			if tt.expectJson {
				// Verify JSON output
				require.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
				jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]
				assert.Equal(t, "text-test", jsonOutput["stackpack_name"])
				assert.Equal(t, "4.0.0", jsonOutput["stackpack_version"])

				// Should not have text calls in JSON mode
				assert.False(t, cli.MockPrinter.HasNonJsonCalls)
			}
		})
	}
}
