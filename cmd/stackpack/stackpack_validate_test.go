package stackpack

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupValidateCmd creates a test command with API context
func setupValidateCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cfg := &config.Config{
		CurrentContext: "test-context",
		Contexts: []*config.NamedContext{
			{
				Name: "test-context",
				Context: &config.StsContext{
					URL:     "https://test-server.example.com",
					APIPath: "/api",
				},
			},
		},
	}
	cli.ConfigPath = filepath.Join(t.TempDir(), "config.yaml")
	err := config.WriteConfig(cli.ConfigPath, cfg)
	require.NoError(t, err)

	cmd := StackpackValidateCommand(&cli.Deps)
	return &cli, cmd
}

// createTestStackpackDir creates a minimal stackpack directory with required items
func createTestStackpackDir(t *testing.T, dir string, name string, version string) {
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "settings"), 0755))
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "resources"), 0755))

	stackpackConf := fmt.Sprintf(`name: "%s"
version: "%s"
`, name, version)
	require.NoError(t, os.WriteFile(filepath.Join(dir, "stackpack.yaml"), []byte(stackpackConf), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "README.md"), []byte("# Test Stackpack"), 0644))
}

// ===== Tests =====

func TestValidate_WithDirectory_AutoPackages(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	tempDir, err := os.MkdirTemp("", "stackpack-validate-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackDir := filepath.Join(tempDir, "test-stackpack")
	require.NoError(t, os.MkdirAll(stackpackDir, 0755))
	createTestStackpackDir(t, stackpackDir, "test-stackpack", "1.0.0")

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--stackpack-directory", stackpackDir)
	require.NoError(t, err)

	// Verify success message
	require.NotEmpty(t, *cli.MockPrinter.SuccessCalls)
	successCall := (*cli.MockPrinter.SuccessCalls)[0]
	assert.Contains(t, successCall, "validation successful")
}

func TestValidate_WithDirectory_InvalidStackpack(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	tempDir, err := os.MkdirTemp("", "stackpack-validate-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create directory with missing required items
	stackpackDir := filepath.Join(tempDir, "invalid-stackpack")
	require.NoError(t, os.MkdirAll(stackpackDir, 0755))

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--stackpack-directory", stackpackDir)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "required stackpack item not found")
}

func TestValidate_WithDirectory_MissingStackpackYaml(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	tempDir, err := os.MkdirTemp("", "stackpack-validate-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackDir := filepath.Join(tempDir, "test-stackpack")
	require.NoError(t, os.MkdirAll(filepath.Join(stackpackDir, "settings"), 0755))
	require.NoError(t, os.MkdirAll(filepath.Join(stackpackDir, "resources"), 0755))
	require.NoError(t, os.WriteFile(filepath.Join(stackpackDir, "README.md"), []byte("test"), 0644))

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--stackpack-directory", stackpackDir)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "required stackpack item not found")
}

func TestValidate_WithPrePackagedFile(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	tempDir, err := os.MkdirTemp("", "stackpack-validate-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a pre-packaged .sts file
	stackpackFile := filepath.Join(tempDir, "test.sts")
	require.NoError(t, os.WriteFile(stackpackFile, []byte("test stackpack content"), 0644))

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--stackpack-file", stackpackFile)
	require.NoError(t, err)

	// Verify success message
	require.NotEmpty(t, *cli.MockPrinter.SuccessCalls)
	successCall := (*cli.MockPrinter.SuccessCalls)[0]
	assert.Contains(t, successCall, "validation successful")
}

func TestValidate_JSONOutput(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	tempDir, err := os.MkdirTemp("", "stackpack-validate-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackFile := filepath.Join(tempDir, "test.sts")
	require.NoError(t, os.WriteFile(stackpackFile, []byte("test content"), 0644))

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--stackpack-file", stackpackFile, "-o", "json")
	require.NoError(t, err)

	// Verify JSON was called
	require.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]

	assert.Equal(t, true, jsonOutput["success"])
}

func TestValidate_MissingPath(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "exactly one of")
}

func TestValidate_MutuallyExclusive(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	tempDir, err := os.MkdirTemp("", "stackpack-validate-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackDir := filepath.Join(tempDir, "stackpack")
	require.NoError(t, os.MkdirAll(stackpackDir, 0755))

	stackpackFile := filepath.Join(tempDir, "test.sts")
	require.NoError(t, os.WriteFile(stackpackFile, []byte("test"), 0644))

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd,
		"-d", stackpackDir,
		"-f", stackpackFile)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "exactly one of")
}

func TestValidate_NonexistentFile(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--stackpack-file", "/nonexistent/path/file.sts")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to access stackpack file")
}

func TestValidate_WithDirectory_IncludingOptionalItems(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	tempDir, err := os.MkdirTemp("", "stackpack-validate-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackDir := filepath.Join(tempDir, "test-stackpack")
	require.NoError(t, os.MkdirAll(stackpackDir, 0755))
	createTestStackpackDir(t, stackpackDir, "test-stackpack", "1.0.0")

	// Add optional items
	require.NoError(t, os.MkdirAll(filepath.Join(stackpackDir, "icons"), 0755))
	require.NoError(t, os.WriteFile(filepath.Join(stackpackDir, "icons", "icon.png"), []byte("fake png"), 0644))
	require.NoError(t, os.MkdirAll(filepath.Join(stackpackDir, "includes"), 0755))
	require.NoError(t, os.WriteFile(filepath.Join(stackpackDir, "includes", "include.txt"), []byte("include data"), 0644))

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--stackpack-directory", stackpackDir)
	require.NoError(t, err)

	// Verify success message
	require.NotEmpty(t, *cli.MockPrinter.SuccessCalls)
	successCall := (*cli.MockPrinter.SuccessCalls)[0]
	assert.Contains(t, successCall, "validation successful")
}

func TestValidate_NonexistentDirectory(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--stackpack-directory", "/nonexistent/stackpack/dir")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "required stackpack item not found")
}
