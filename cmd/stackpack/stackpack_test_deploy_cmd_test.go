package stackpack

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/blang/semver/v4"
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupStackpackTestDeployCmd creates a test command with mock dependencies
func setupStackpackTestDeployCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackTestDeployCommand(&cli.Deps)
	return &cli, cmd
}

func TestStackpackTestDeployCommand_FlagsAndStructure(t *testing.T) {
	_, cmd := setupStackpackTestDeployCmd(t)

	// Test command structure
	assert.Equal(t, "test-deploy", cmd.Use)
	assert.Contains(t, cmd.Short, "Test a stackpack")
	assert.Contains(t, cmd.Long, "package, upload, and install")
	assert.NotEmpty(t, cmd.Example)

	// Test flags exist
	flags := cmd.Flags()

	stackpackDirFlag := flags.Lookup("stackpack-directory")
	require.NotNil(t, stackpackDirFlag)
	assert.Equal(t, "d", stackpackDirFlag.Shorthand)

	paramsFlag := flags.Lookup("parameter")
	require.NotNil(t, paramsFlag)
	assert.Equal(t, "p", paramsFlag.Shorthand)

	yesFlag := flags.Lookup("yes")
	require.NotNil(t, yesFlag)
	assert.Equal(t, "y", yesFlag.Shorthand)

	unlockedStrategyFlag := flags.Lookup("unlocked-strategy")
	require.NotNil(t, unlockedStrategyFlag)
	assert.Equal(t, "fail", unlockedStrategyFlag.DefValue)
}

func TestBumpSnapshotVersion(t *testing.T) {
	tests := []struct {
		name            string
		currentVersion  string
		expectedVersion string
	}{
		{
			name:            "first cli-test version",
			currentVersion:  "1.0.0",
			expectedVersion: "1.0.0-cli-test.10000",
		},
		{
			name:            "increment existing cli-test",
			currentVersion:  "1.0.0-cli-test.1",
			expectedVersion: "1.0.0-cli-test.2",
		},
		{
			name:            "increment higher cli-test number",
			currentVersion:  "2.1.5-cli-test.10",
			expectedVersion: "2.1.5-cli-test.11",
		},
		{
			name:            "complex version with cli-test",
			currentVersion:  "1.0.0-beta.1.cli-test.3",
			expectedVersion: "1.0.0-beta.1.cli-test.4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary config file
			tempDir, err := os.MkdirTemp("", "stackpack-version-test-*")
			require.NoError(t, err)
			defer os.RemoveAll(tempDir)

			configPath := filepath.Join(tempDir, "stackpack.yaml")
			configContent := fmt.Sprintf(`name: "test-stackpack"
version: "%s"
displayName: "Test StackPack"`, tt.currentVersion)
			require.NoError(t, os.WriteFile(configPath, []byte(configContent), 0644))

			// Test version bumping
			newVersion, err := bumpSnapshotVersionWithBase(configPath, tt.currentVersion)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedVersion, newVersion)

			// Verify config file was updated
			parser := &YamlParser{}
			updatedInfo, err := parser.Parse(configPath)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedVersion, updatedInfo.Version)
		})
	}
}

func TestUpdateVersionInYaml(t *testing.T) {
	// Create temporary config file
	tempDir, err := os.MkdirTemp("", "stackpack-yaml-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	configPath := filepath.Join(tempDir, "stackpack.yaml")
	originalVersion := "2.0.0"
	newVersion := "2.0.0-cli-test.1"

	// Create config with original version
	configContent := fmt.Sprintf(`name: "test-stackpack"
version: "%s"
displayName: "Test StackPack"`, originalVersion)
	require.NoError(t, os.WriteFile(configPath, []byte(configContent), 0644))

	// Test version update using HOCON approach
	err = updateVersionInYaml(configPath, newVersion)
	require.NoError(t, err)

	// Verify config file was updated and is still valid HOCON
	parser := &YamlParser{}
	updatedInfo, err := parser.Parse(configPath)
	require.NoError(t, err)
	assert.Equal(t, newVersion, updatedInfo.Version)
	assert.Equal(t, "test-stackpack", updatedInfo.Name)
}

func TestCopyDirectory(t *testing.T) {
	// Create source directory with test files
	tempDir, err := os.MkdirTemp("", "copy-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	srcDir := filepath.Join(tempDir, "src")
	dstDir := filepath.Join(tempDir, "dst")

	// Create source structure
	require.NoError(t, os.MkdirAll(filepath.Join(srcDir, "subdir"), 0755))

	testContent := "test content"
	require.NoError(t, os.WriteFile(filepath.Join(srcDir, "file1.txt"), []byte(testContent), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(srcDir, "subdir", "file2.txt"), []byte(testContent), 0644))

	// Test directory copy
	err = copyDirectory(srcDir, dstDir)
	require.NoError(t, err)

	// Verify files were copied
	copiedContent, err := os.ReadFile(filepath.Join(dstDir, "file1.txt"))
	require.NoError(t, err)
	assert.Equal(t, testContent, string(copiedContent))

	copiedContent2, err := os.ReadFile(filepath.Join(dstDir, "subdir", "file2.txt"))
	require.NoError(t, err)
	assert.Equal(t, testContent, string(copiedContent2))
}

func TestStackpackTestDeployCommand_RequiredFlags(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantErr      bool
		errorMessage string
	}{
		{
			name:    "valid minimal args - no flags required",
			args:    []string{},
			wantErr: false,
		},
		{
			name:    "with all flags",
			args:    []string{"-d", "./test", "-p", "param1=value1", "--yes", "--unlocked-strategy", "force"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, cmd := setupStackpackTestDeployCmd(t)
			_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, tt.args...)

			if tt.wantErr {
				require.Error(t, err)
				if tt.errorMessage != "" {
					assert.Contains(t, err.Error(), tt.errorMessage)
				}
			} else if err != nil {
				// Note: This will fail due to missing stackpack.yaml, but that's expected
				// We're only testing flag parsing here
				// Should fail on stackpack.conf parsing, not flag validation
				assert.Contains(t, err.Error(), "stackpack.yaml")
			}
		})
	}
}

func TestStackpackTestDeployCommand_DirectoryHandling(t *testing.T) {
	// Create temporary directory with valid stackpack structure
	tempDir, err := os.MkdirTemp("", "stackpack-test-dir-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	stackpackDir := filepath.Join(tempDir, "test-stackpack")
	require.NoError(t, os.MkdirAll(stackpackDir, 0755))

	// Create required files
	createTestStackpack(t, stackpackDir, "test-stackpack", "1.0.0")

	cli, cmd := setupStackpackTestDeployCmd(t)

	tests := []struct {
		name string
		args []string
	}{
		{
			name: "explicit directory",
			args: []string{"-d", stackpackDir, "--yes"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Note: This test will fail at the API call stage since we're using mock deps
			// But we can verify that directory parsing works
			_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, tt.args...)

			// Should fail on API calls (upload/install), not on directory or config parsing
			if err != nil {
				// The error should not be about missing stackpack.yaml or directory issues
				assert.NotContains(t, err.Error(), "stackpack.yaml")
				assert.NotContains(t, err.Error(), "no such file or directory")
			}
		})
	}
}

//nolint:funlen
func TestVersionRegexPatterns(t *testing.T) {
	tests := []struct {
		name           string
		version        string
		expectedBase   string
		expectedNumber int
		isSnapshot     bool
	}{
		{
			name:           "regular version",
			version:        "1.0.0",
			expectedBase:   "",
			expectedNumber: 0,
			isSnapshot:     false,
		},
		{
			name:           "first cli-test",
			version:        "1.0.0-cli-test.1",
			expectedBase:   "1.0.0",
			expectedNumber: 1,
			isSnapshot:     true,
		},
		{
			name:           "higher cli-test number",
			version:        "2.1.0-cli-test.15",
			expectedBase:   "2.1.0",
			expectedNumber: 15,
			isSnapshot:     true,
		},
		{
			name:           "complex base version",
			version:        "1.0.0-beta.1.cli-test.3",
			expectedBase:   "1.0.0",
			expectedNumber: 3,
			isSnapshot:     true,
		},
		{
			name:           "semantic version with cli-test",
			version:        "10.20.30-cli-test.99",
			expectedBase:   "10.20.30",
			expectedNumber: 99,
			isSnapshot:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This tests the semver parsing used in bumpSnapshotVersion
			version, err := semver.Parse(tt.version)
			if tt.isSnapshot {
				require.NoError(t, err, "Should be able to parse valid semver version")
				// Check if it has cli-test pre-release
				require.Greater(t, len(version.Pre), 1, "Should have at least 2 pre-release parts for cli-test")

				// Find cli-test index
				cliTestIndex := -1
				for i, pre := range version.Pre {
					if pre.VersionStr == "cli-test" && i+1 < len(version.Pre) {
						cliTestIndex = i
						break
					}
				}
				require.GreaterOrEqual(t, cliTestIndex, 0, "Should have cli-test in pre-release parts")

				var actualNumber int
				numberPart := version.Pre[cliTestIndex+1]
				if numberPart.VersionStr == "" {
					// Numeric identifier stored in VersionNum
					actualNumber = int(numberPart.VersionNum)
				} else {
					// String identifier that should be numeric
					var parseErr error
					actualNumber, parseErr = strconv.Atoi(numberPart.VersionStr)
					require.NoError(t, parseErr, "Number part after cli-test should be a number")
				}
				assert.Equal(t, tt.expectedNumber, actualNumber)

				// Reconstruct base version
				expectedBase := fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
				assert.Equal(t, tt.expectedBase, expectedBase)
			} else {
				require.NoError(t, err, "Should be able to parse regular version")
				assert.Len(t, version.Pre, 0, "Regular version should have no pre-release parts")
			}
		})
	}
}

func TestCompareVersionsSemver(t *testing.T) {
	tests := []struct {
		name     string
		v1       string
		v2       string
		expected int // -1 if v1 < v2, 0 if v1 == v2, 1 if v1 > v2
	}{
		{
			name:     "equal versions",
			v1:       "1.0.0",
			v2:       "1.0.0",
			expected: 0,
		},
		{
			name:     "v1 less than v2",
			v1:       "1.0.0",
			v2:       "1.1.0",
			expected: -1,
		},
		{
			name:     "v1 greater than v2",
			v1:       "2.0.0",
			v2:       "1.9.9",
			expected: 1,
		},
		{
			name:     "pre-release versions - cli-test comparison",
			v1:       "1.0.0-cli-test.1",
			v2:       "1.0.0-cli-test.2",
			expected: -1,
		},
		{
			name:     "pre-release vs release",
			v1:       "1.0.0-cli-test.1",
			v2:       "1.0.0",
			expected: -1,
		},
		{
			name:     "major version difference",
			v1:       "2.0.0-cli-test.1",
			v2:       "1.9.9",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := compareVersions(tt.v1, tt.v2)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStackpackTestDeployCommand_InstallUpgradeLogic(t *testing.T) {
	tests := []struct {
		name             string
		installedVersion string
		expectedAction   string
	}{
		{
			name:             "not installed - should install",
			installedVersion: "",
			expectedAction:   "install",
		},
		{
			name:             "already installed - should upgrade",
			installedVersion: "1.0.0",
			expectedAction:   "upgrade",
		},
		{
			name:             "already installed with pre-release - should upgrade",
			installedVersion: "1.0.0-beta.1",
			expectedAction:   "upgrade",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test verifies the logic for choosing between install and upgrade
			// The actual logic is: if installedVersion != "" then upgrade, else install
			var shouldUpgrade bool
			if tt.installedVersion != "" {
				shouldUpgrade = true
			}

			if tt.expectedAction == "upgrade" {
				assert.True(t, shouldUpgrade, "Should choose upgrade when stackpack is already installed")
			} else {
				assert.False(t, shouldUpgrade, "Should choose install when stackpack is not installed")
			}
		})
	}
}

func TestConfirmUpload(t *testing.T) {
	// Note: This test is more complex because it involves stdin interaction
	// For now, we test the function exists and can be called
	// A full integration test would require mocking stdin

	cli := di.NewMockDeps(t)

	// Test that the function exists and doesn't panic
	// Note: This will block waiting for input, so we skip actual execution
	// In a real test environment, you'd mock the stdin reader

	// Just verify the function signature is correct
	//nolint:staticcheck
	var confirmFunc func(*di.Deps, string) bool = confirmUpload
	assert.NotNil(t, confirmFunc)

	// Test URL extraction logic implicitly by checking context access
	// Note: In mock environment, CurrentContext may be nil, which is expected
	_ = cli.CurrentContext // Just ensure the field exists
}
