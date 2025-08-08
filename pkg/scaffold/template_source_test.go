package scaffold

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLocalDirSource_NewLocalDirSource(t *testing.T) {
	source := NewLocalDirSource("/test/path", "template-name")

	if source.Path != "/test/path" {
		t.Errorf("Expected Path to be '/test/path', got '%s'", source.Path)
	}

	if source.TemplateName != "template-name" {
		t.Errorf("Expected TemplateName to be 'template-name', got '%s'", source.TemplateName)
	}
}

func TestLocalDirSource_String(t *testing.T) {
	source := NewLocalDirSource("/test/path", "template-name")
	expected := "localdir:/test/path"

	if source.String() != expected {
		t.Errorf("Expected String() to return '%s', got '%s'", expected, source.String())
	}
}

func TestLocalDirSource_Validate(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		templateName string
		setupFunc    func(string) error
		wantErr      bool
		expectedErr  string
	}{
		{
			name:         "valid directory with template",
			path:         "",
			templateName: "test-template",
			setupFunc: func(tempDir string) error {
				return os.MkdirAll(filepath.Join(tempDir, "test-template"), 0755)
			},
			wantErr: false,
		},
		{
			name:         "empty path",
			path:         "",
			templateName: "test-template",
			setupFunc:    func(tempDir string) error { return nil },
			wantErr:      true,
			expectedErr:  "template path is required",
		},
		{
			name:         "empty template name",
			path:         "/test/path",
			templateName: "",
			setupFunc:    func(tempDir string) error { return nil },
			wantErr:      true,
			expectedErr:  "template name is required",
		},
		{
			name:         "non-existent directory",
			path:         "/non/existent/path",
			templateName: "test-template",
			setupFunc:    func(tempDir string) error { return nil },
			wantErr:      true,
			expectedErr:  "template directory does not exist: /non/existent/path",
		},
		{
			name:         "path is a file not directory",
			path:         "",
			templateName: "test-template",
			setupFunc: func(tempDir string) error {
				return os.WriteFile(tempDir+"/file.txt", []byte("test"), 0644)
			},
			wantErr:     true,
			expectedErr: "template path is not a directory:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "test-local-source-*")
			if err != nil {
				t.Fatalf("Failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tempDir)

			// Setup test scenario
			if err := tt.setupFunc(tempDir); err != nil {
				t.Fatalf("Failed to setup test: %v", err)
			}

			// Use temp dir as path if not specified
			path := tt.path
			if path == "" && tt.name != "empty path" {
				path = tempDir
			}
			if tt.name == "path is a file not directory" {
				path = tempDir + "/file.txt"
			}

			source := NewLocalDirSource(path, tt.templateName)
			err = source.Validate()

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if tt.expectedErr != "" && !containsString(err.Error(), tt.expectedErr) {
					t.Errorf("Expected error to contain '%s', got '%s'", tt.expectedErr, err.Error())
				}
			} else if err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestLocalDirSource_Fetch(t *testing.T) {
	tests := []struct {
		name         string
		templateName string
		setupFunc    func(string) error
		wantErr      bool
		expectedErr  string
	}{
		{
			name:         "successful fetch",
			templateName: "test-template",
			setupFunc: func(tempDir string) error {
				return os.MkdirAll(filepath.Join(tempDir, "test-template"), 0755)
			},
			wantErr: false,
		},
		{
			name:         "template not found",
			templateName: "non-existent-template",
			setupFunc:    func(tempDir string) error { return nil },
			wantErr:      true,
			expectedErr:  "template 'non-existent-template' not found in directory",
		},
		{
			name:         "template is file not directory",
			templateName: "file-template",
			setupFunc: func(tempDir string) error {
				return os.WriteFile(filepath.Join(tempDir, "file-template"), []byte("test"), 0644)
			},
			wantErr:     true,
			expectedErr: "template 'file-template' is not a directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "test-local-fetch-*")
			if err != nil {
				t.Fatalf("Failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tempDir)

			// Setup test scenario
			if err := tt.setupFunc(tempDir); err != nil {
				t.Fatalf("Failed to setup test: %v", err)
			}

			source := NewLocalDirSource(tempDir, tt.templateName)
			result, err := source.Fetch(context.Background())

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if tt.expectedErr != "" && !containsString(err.Error(), tt.expectedErr) {
					t.Errorf("Expected error to contain '%s', got '%s'", tt.expectedErr, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
					return
				}

				expectedPath := filepath.Join(tempDir, tt.templateName)
				if result != expectedPath {
					t.Errorf("Expected result to be '%s', got '%s'", expectedPath, result)
				}

				// Verify the returned path exists
				if _, err := os.Stat(result); os.IsNotExist(err) {
					t.Errorf("Returned path does not exist: %s", result)
				}
			}
		})
	}
}

func TestLocalDirSource_Cleanup(t *testing.T) {
	source := NewLocalDirSource("/test/path", "template-name")

	// Cleanup should not return error for local directories
	err := source.Cleanup()
	if err != nil {
		t.Errorf("Expected no error from Cleanup, got: %v", err)
	}
}

// Helper function to check if a string contains a substring
func containsString(s, substr string) bool {
	return strings.Contains(s, substr)
}
