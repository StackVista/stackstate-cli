package scaffold

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGitHubSource_NewGitHubSource(t *testing.T) {
	tests := []struct {
		name         string
		owner        string
		repo         string
		ref          string
		path         string
		templateName string
		expectedRef  string
	}{
		{
			name:         "all parameters provided",
			owner:        "testowner",
			repo:         "testrepo",
			ref:          "v1.0.0",
			path:         "templates",
			templateName: "generic",
			expectedRef:  "v1.0.0",
		},
		{
			name:         "empty ref defaults to main",
			owner:        "testowner",
			repo:         "testrepo",
			ref:          "",
			path:         "templates",
			templateName: "generic",
			expectedRef:  "main",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			source := NewGitHubSource(tt.owner, tt.repo, tt.ref, tt.path, tt.templateName)

			if source.Owner != tt.owner {
				t.Errorf("Expected Owner to be '%s', got '%s'", tt.owner, source.Owner)
			}
			if source.Repo != tt.repo {
				t.Errorf("Expected Repo to be '%s', got '%s'", tt.repo, source.Repo)
			}
			if source.Ref != tt.expectedRef {
				t.Errorf("Expected Ref to be '%s', got '%s'", tt.expectedRef, source.Ref)
			}
			if source.Path != tt.path {
				t.Errorf("Expected Path to be '%s', got '%s'", tt.path, source.Path)
			}
			if source.TemplateName != tt.templateName {
				t.Errorf("Expected TemplateName to be '%s', got '%s'", tt.templateName, source.TemplateName)
			}
		})
	}
}

func TestGitHubSource_String(t *testing.T) {
	tests := []struct {
		name        string
		source      *GitHubSource
		expectedStr string
	}{
		{
			name:        "with path",
			source:      NewGitHubSource("owner", "repo", "main", "templates", "generic"),
			expectedStr: "github:owner/repo@main:templates",
		},
		{
			name:        "without path",
			source:      NewGitHubSource("owner", "repo", "v1.0", "", "generic"),
			expectedStr: "github:owner/repo@v1.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.source.String()
			if result != tt.expectedStr {
				t.Errorf("Expected String() to return '%s', got '%s'", tt.expectedStr, result)
			}
		})
	}
}

func TestGitHubSource_Validate(t *testing.T) {
	tests := []struct {
		name         string
		owner        string
		repo         string
		templateName string
		wantErr      bool
		expectedErr  string
	}{
		{
			name:         "valid source",
			owner:        "testowner",
			repo:         "testrepo",
			templateName: "generic",
			wantErr:      false,
		},
		{
			name:         "empty owner",
			owner:        "",
			repo:         "testrepo",
			templateName: "generic",
			wantErr:      true,
			expectedErr:  "GitHub repository owner is required",
		},
		{
			name:         "empty repo",
			owner:        "testowner",
			repo:         "",
			templateName: "generic",
			wantErr:      true,
			expectedErr:  "GitHub repository name is required",
		},
		{
			name:         "empty template name",
			owner:        "testowner",
			repo:         "testrepo",
			templateName: "",
			wantErr:      true,
			expectedErr:  "template name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			source := NewGitHubSource(tt.owner, tt.repo, "main", "templates", tt.templateName)
			err := source.Validate()

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if err.Error() != tt.expectedErr {
					t.Errorf("Expected error '%s', got '%s'", tt.expectedErr, err.Error())
				}
			} else if err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestGitHubSource_Cleanup(t *testing.T) {
	// Test cleanup when no temp directory is set
	source := NewGitHubSource("owner", "repo", "main", "", "generic")
	err := source.Cleanup()
	if err != nil {
		t.Errorf("Expected no error when tempDir is empty, got: %v", err)
	}

	// Test cleanup when temp directory is set
	tempDir, err := os.MkdirTemp("", "test-cleanup-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Create a test file in temp directory
	testFile := filepath.Join(tempDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	source.tempDir = tempDir
	err = source.Cleanup()
	if err != nil {
		t.Errorf("Expected no error during cleanup, got: %v", err)
	}

	// Verify temp directory was removed
	if _, err := os.Stat(tempDir); !os.IsNotExist(err) {
		t.Errorf("Expected temp directory to be removed, but it still exists")
	}
}

func TestGitHubSource_extractZip(t *testing.T) {
	// Create a test ZIP file
	zipData := createTestZipFile(t, map[string]string{
		"test-repo-main/file1.txt":     "content1",
		"test-repo-main/dir/file2.txt": "content2",
	})

	// Write ZIP to temporary file
	tempDir, err := os.MkdirTemp("", "test-extract-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	zipFile := filepath.Join(tempDir, "test.zip")
	if err := os.WriteFile(zipFile, zipData, 0644); err != nil {
		t.Fatalf("Failed to write ZIP file: %v", err)
	}

	extractDir := filepath.Join(tempDir, "extracted")
	source := NewGitHubSource("owner", "repo", "main", "", "generic")

	err = source.extractZip(zipFile, extractDir)
	if err != nil {
		t.Fatalf("Expected no error during extraction, got: %v", err)
	}

	// Verify extracted files
	expectedFiles := []string{
		"test-repo-main/file1.txt",
		"test-repo-main/dir/file2.txt",
	}

	for _, expectedFile := range expectedFiles {
		fullPath := filepath.Join(extractDir, expectedFile)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			t.Errorf("Expected file %s to exist after extraction", expectedFile)
		}
	}
}

//nolint:funlen
func TestGitHubSource_findTemplateDirectory(t *testing.T) {
	tests := []struct {
		name         string
		repo         string
		ref          string
		path         string
		templateName string
		setupFiles   map[string]string
		wantErr      bool
		expectedErr  string
	}{
		{
			name:         "find template with path",
			repo:         "test-repo",
			ref:          "main",
			path:         "templates",
			templateName: "generic",
			setupFiles: map[string]string{
				"test-repo-main/templates/generic/file.txt": "content",
			},
			wantErr: false,
		},
		{
			name:         "find template without path",
			repo:         "test-repo",
			ref:          "main",
			path:         "",
			templateName: "generic",
			setupFiles: map[string]string{
				"test-repo-main/generic/file.txt": "content",
			},
			wantErr: false,
		},
		{
			name:         "repo directory not found",
			repo:         "test-repo",
			ref:          "main",
			path:         "",
			templateName: "generic",
			setupFiles: map[string]string{
				"wrong-repo-main/generic/file.txt": "content",
			},
			wantErr:     true,
			expectedErr: "could not find repository directory in extracted archive",
		},
		{
			name:         "template base directory not found",
			repo:         "test-repo",
			ref:          "main",
			path:         "nonexistent",
			templateName: "generic",
			setupFiles: map[string]string{
				"test-repo-main/templates/generic/file.txt": "content",
			},
			wantErr:     true,
			expectedErr: "template base directory not found at path 'nonexistent'",
		},
		{
			name:         "template not found",
			repo:         "test-repo",
			ref:          "main",
			path:         "templates",
			templateName: "nonexistent",
			setupFiles: map[string]string{
				"test-repo-main/templates/generic/file.txt": "content",
			},
			wantErr:     true,
			expectedErr: "template 'nonexistent' not found in directory",
		},
		{
			name:         "template is file not directory",
			repo:         "test-repo",
			ref:          "main",
			path:         "templates",
			templateName: "generic",
			setupFiles: map[string]string{
				"test-repo-main/templates/generic": "file content",
			},
			wantErr:     true,
			expectedErr: "template 'generic' is not a directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "test-find-template-*")
			if err != nil {
				t.Fatalf("Failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tempDir)

			// Create test file structure
			for filePath, content := range tt.setupFiles {
				fullPath := filepath.Join(tempDir, filePath)
				if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
					t.Fatalf("Failed to create directory structure: %v", err)
				}
				if content == "" && strings.HasSuffix(filePath, "/") {
					// Directory
					if err := os.MkdirAll(fullPath, 0755); err != nil {
						t.Fatalf("Failed to create directory: %v", err)
					}
				} else {
					// File
					if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
						t.Fatalf("Failed to create file: %v", err)
					}
				}
			}

			source := NewGitHubSource("owner", tt.repo, tt.ref, tt.path, tt.templateName)
			result, err := source.findTemplateDirectory(tempDir)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if tt.expectedErr != "" && !strings.Contains(err.Error(), tt.expectedErr) {
					t.Errorf("Expected error to contain '%s', got '%s'", tt.expectedErr, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
					return
				}

				// Verify the result is the correct path
				var expectedPath string
				if tt.path != "" {
					expectedPath = filepath.Join(tempDir, fmt.Sprintf("%s-%s", tt.repo, tt.ref), tt.path, tt.templateName)
				} else {
					expectedPath = filepath.Join(tempDir, fmt.Sprintf("%s-%s", tt.repo, tt.ref), tt.templateName)
				}

				if result != expectedPath {
					t.Errorf("Expected result to be '%s', got '%s'", expectedPath, result)
				}
			}
		})
	}
}

func TestGitHubSource_Fetch_Integration(t *testing.T) {
	// Create a mock HTTP server that serves a ZIP file
	zipData := createTestZipFile(t, map[string]string{
		"test-repo-main/":                           "",
		"test-repo-main/templates/":                 "",
		"test-repo-main/templates/generic/":         "",
		"test-repo-main/templates/generic/file.txt": "template content",
	})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedURL := "/testowner/testrepo/archive/main.zip"
		if r.URL.Path != expectedURL {
			t.Errorf("Expected request to '%s', got '%s'", expectedURL, r.URL.Path)
			http.Error(w, "Not found", 404)
			return
		}

		w.Header().Set("Content-Type", "application/zip")
		w.WriteHeader(200)
		_, _ = w.Write(zipData)
	}))
	defer server.Close()

	// We need to modify the Fetch method to use our test server, but since we can't easily do that,
	// let's just test the zip extraction and directory finding parts separately.
	// For a real integration test, we'd need dependency injection for the HTTP client.

	t.Skip("Integration test requires HTTP client dependency injection")
}

// Helper function to create a ZIP file with given files
func createTestZipFile(t *testing.T, files map[string]string) []byte {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	for filePath, content := range files {
		if strings.HasSuffix(filePath, "/") {
			// Directory entry
			_, err := zipWriter.Create(filePath)
			if err != nil {
				t.Fatalf("Failed to create directory entry in ZIP: %v", err)
			}
		} else {
			// File entry
			fileWriter, err := zipWriter.Create(filePath)
			if err != nil {
				t.Fatalf("Failed to create file entry in ZIP: %v", err)
			}
			_, err = io.WriteString(fileWriter, content)
			if err != nil {
				t.Fatalf("Failed to write file content to ZIP: %v", err)
			}
		}
	}

	if err := zipWriter.Close(); err != nil {
		t.Fatalf("Failed to close ZIP writer: %v", err)
	}

	return buf.Bytes()
}
