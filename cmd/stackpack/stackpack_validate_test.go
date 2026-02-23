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

// setupValidateCmd creates a test command with API mode context (URL set)
func setupValidateCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cli.CurrentContext = &config.StsContext{
		URL: "https://test-server.example.com",
	}
	cmd := StackpackValidateCommand(&cli.Deps)
	return &cli, cmd
}

// setupValidateCmdWithMockDocker creates a test command with Docker mode (no URL) and captures docker args
func setupValidateCmdWithMockDocker(t *testing.T) (*di.MockDeps, *cobra.Command, *[][]string) {
	cli := di.NewMockDeps(t)
	cli.CurrentContext = &config.StsContext{
		URL: "", // No URL = docker mode
	}
	args := &ValidateArgs{}
	dockerArgsCaptured := [][]string{}

	args.dockerRunner = func(dockerArgs []string) error {
		dockerArgsCaptured = append(dockerArgsCaptured, dockerArgs)
		return nil
	}

	cmd := stackpackValidateCommandWithArgs(&cli.Deps, args)
	return &cli, cmd, &dockerArgsCaptured
}

// createTestStackpackDir creates a minimal stackpack directory with stackpack.yaml
func createTestStackpackDir(t *testing.T, dir string, name string, version string) {
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "settings"), 0755))
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "resources"), 0755))

	stackpackConf := fmt.Sprintf(`name: "%s"
version: "%s"
`, name, version)
	require.NoError(t, os.WriteFile(filepath.Join(dir, "stackpack.yaml"), []byte(stackpackConf), 0644))
}

// ===== API Mode Tests =====

func TestValidate_APIMode_ByName(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "my-stackpack")
	require.NoError(t, err)

	// Verify success message
	require.NotEmpty(t, *cli.MockPrinter.SuccessCalls)
	successCall := (*cli.MockPrinter.SuccessCalls)[0]
	assert.Contains(t, successCall, "validation successful")
}

func TestValidate_APIMode_RequiresName(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	tempDir, err := os.MkdirTemp("", "stackpack-validate-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackDir := filepath.Join(tempDir, "test-stackpack")
	require.NoError(t, os.MkdirAll(stackpackDir, 0755))
	createTestStackpackDir(t, stackpackDir, "test-stackpack", "1.0.0")

	// API mode requires --name flag, not --stackpack-directory
	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--stackpack-directory", stackpackDir)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "name is required")
}

func TestValidate_APIMode_JSONOutput(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--name", "test-pack", "-o", "json")
	require.NoError(t, err)

	// Verify JSON was called
	require.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
	jsonOutput := (*cli.MockPrinter.PrintJsonCalls)[0]

	assert.Equal(t, true, jsonOutput["success"])
}

func TestValidate_APIMode_MissingName(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "name is required")
}

// ===== Docker Mode Tests =====

func TestValidate_DockerMode_MissingImage(t *testing.T) {
	cli, cmd, _ := setupValidateCmdWithMockDocker(t)

	tempDir, err := os.MkdirTemp("", "stackpack-validate-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "-d", tempDir)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "--image is required")
}

func TestValidate_DockerMode_MissingPath(t *testing.T) {
	cli, cmd, _ := setupValidateCmdWithMockDocker(t)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--image", "quay.io/stackstate/stackstate-server:latest")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "exactly one of")
}

func TestValidate_DockerMode_MutuallyExclusive(t *testing.T) {
	cli, cmd, _ := setupValidateCmdWithMockDocker(t)

	tempDir, err := os.MkdirTemp("", "stackpack-validate-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackDir := filepath.Join(tempDir, "stackpack")
	require.NoError(t, os.MkdirAll(stackpackDir, 0755))

	stackpackFile := filepath.Join(tempDir, "test.sts")
	require.NoError(t, os.WriteFile(stackpackFile, []byte("test"), 0644))

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd,
		"--image", "quay.io/stackstate/stackstate-server:latest",
		"-d", stackpackDir,
		"-f", stackpackFile)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "exactly one of")
}

func TestValidate_DockerMode_WithDirectory(t *testing.T) {
	cli, cmd, capturedArgs := setupValidateCmdWithMockDocker(t)

	tempDir, err := os.MkdirTemp("", "stackpack-validate-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackDir := filepath.Join(tempDir, "test-stackpack")
	require.NoError(t, os.MkdirAll(stackpackDir, 0755))

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd,
		"--image", "quay.io/stackstate/stackstate-server:latest",
		"-d", stackpackDir)
	require.NoError(t, err)

	// Verify docker args were captured
	require.Len(t, *capturedArgs, 1)
	dockerArgs := (*capturedArgs)[0]

	// Check essential arguments
	assert.Contains(t, dockerArgs, "run")
	assert.Contains(t, dockerArgs, "--rm")
	assert.Contains(t, dockerArgs, "--entrypoint")
	assert.Contains(t, dockerArgs, "/opt/docker/bin/stack-pack-validator")
	assert.Contains(t, dockerArgs, "quay.io/stackstate/stackstate-server:latest")
	assert.Contains(t, dockerArgs, "-directory")
	assert.Contains(t, dockerArgs, "/stackpack")

	// Verify -v flag is present with correct path
	foundVolume := false
	for i, arg := range dockerArgs {
		if arg == "-v" && i+1 < len(dockerArgs) {
			volumeArg := dockerArgs[i+1]
			if volumeArg == stackpackDir+":/stackpack" {
				foundVolume = true
				break
			}
		}
	}
	assert.True(t, foundVolume, "Expected volume mount for stackpack directory")
}

func TestValidate_DockerMode_WithFile(t *testing.T) {
	cli, cmd, capturedArgs := setupValidateCmdWithMockDocker(t)

	tempDir, err := os.MkdirTemp("", "stackpack-validate-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackFile := filepath.Join(tempDir, "test.sts")
	require.NoError(t, os.WriteFile(stackpackFile, []byte("test content"), 0644))

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd,
		"--image", "quay.io/stackstate/stackstate-server:latest",
		"-f", stackpackFile)
	require.NoError(t, err)

	// Verify docker args were captured
	require.Len(t, *capturedArgs, 1)
	dockerArgs := (*capturedArgs)[0]

	// Check essential arguments
	assert.Contains(t, dockerArgs, "run")
	assert.Contains(t, dockerArgs, "--rm")
	assert.Contains(t, dockerArgs, "--entrypoint")
	assert.Contains(t, dockerArgs, "/opt/docker/bin/stack-pack-validator")
	assert.Contains(t, dockerArgs, "quay.io/stackstate/stackstate-server:latest")
	assert.Contains(t, dockerArgs, "-file")
	assert.Contains(t, dockerArgs, "/stackpack.sts")

	// Verify -v flag is present with correct path
	foundVolume := false
	for i, arg := range dockerArgs {
		if arg == "-v" && i+1 < len(dockerArgs) {
			volumeArg := dockerArgs[i+1]
			if volumeArg == stackpackFile+":/stackpack.sts" {
				foundVolume = true
				break
			}
		}
	}
	assert.True(t, foundVolume, "Expected volume mount for stackpack file")
}

// ===== Mock Helper (using the real stackstate_api type) =====

// Note: We use the real stackstate_api.StackPackValidationSuccess type directly
// in tests via the mock API setup
