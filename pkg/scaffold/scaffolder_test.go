package scaffold

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockTemplateSource for testing
type MockTemplateSource struct {
	fetchResult   string
	fetchError    error
	validateError error
	cleanupError  error
	stringResult  string
	cleanupCalled bool
	fetchCalled   bool
}

func (m *MockTemplateSource) Fetch(ctx context.Context) (string, error) {
	m.fetchCalled = true
	return m.fetchResult, m.fetchError
}

func (m *MockTemplateSource) Validate() error {
	return m.validateError
}

func (m *MockTemplateSource) Cleanup() error {
	m.cleanupCalled = true
	return m.cleanupError
}

func (m *MockTemplateSource) String() string {
	return m.stringResult
}

// MockPrinter for testing
type MockPrinter struct {
	printedLines    []string
	printedWarnings []string
}

func (m *MockPrinter) PrintLn(message string) {
	m.printedLines = append(m.printedLines, message)
}

func (m *MockPrinter) PrintWarn(message string) {
	m.printedWarnings = append(m.printedWarnings, message)
}

func TestNewScaffolder(t *testing.T) {
	mockSource := &MockTemplateSource{stringResult: "mock-source"}
	context := TemplateContext{Name: "test-name", DisplayName: "test-name", TemplateName: "test-template"}
	mockPrinter := &MockPrinter{}

	scaffolder := NewScaffolder(mockSource, "/test/dest", context, true, mockPrinter, false)

	if scaffolder.source != mockSource {
		t.Errorf("Expected source to be mockSource")
	}
	if scaffolder.destination != "/test/dest" {
		t.Errorf("Expected destination to be '/test/dest', got '%s'", scaffolder.destination)
	}
	if scaffolder.context.Name != "test-name" {
		t.Errorf("Expected context.Name to be 'test-name', got '%s'", scaffolder.context.Name)
	}
	if scaffolder.force != true {
		t.Errorf("Expected force to be true, got %v", scaffolder.force)
	}
	if scaffolder.printer != mockPrinter {
		t.Errorf("Expected printer to be mockPrinter")
	}
}

func TestScaffolder_validateArgs(t *testing.T) {
	tests := []struct {
		name        string
		context     TemplateContext
		sourceErr   error
		wantErr     bool
		expectedErr string
	}{
		{
			name:      "valid args",
			context:   TemplateContext{Name: "test-name", DisplayName: "test-name", TemplateName: "test-template"},
			sourceErr: nil,
			wantErr:   false,
		},
		{
			name:        "source validation error",
			context:     TemplateContext{Name: "test-name", TemplateName: "test-template"},
			sourceErr:   os.ErrInvalid,
			wantErr:     true,
			expectedErr: "invalid argument",
		},
		{
			name:        "empty name",
			context:     TemplateContext{Name: "", TemplateName: "test-template"},
			sourceErr:   nil,
			wantErr:     true,
			expectedErr: "stackpack name is required",
		},
		{
			name:        "empty template name",
			context:     TemplateContext{Name: "test-name", TemplateName: ""},
			sourceErr:   nil,
			wantErr:     true,
			expectedErr: "template name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSource := &MockTemplateSource{
				validateError: tt.sourceErr,
				stringResult:  "mock-source",
			}
			mockPrinter := &MockPrinter{}

			scaffolder := NewScaffolder(mockSource, "/test/dest", tt.context, false, mockPrinter, false)
			err := scaffolder.validateArgs()

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if tt.expectedErr != "" && !strings.Contains(err.Error(), tt.expectedErr) {
					t.Errorf("Expected error to contain '%s', got '%s'", tt.expectedErr, err.Error())
				}
			} else if err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestScaffolder_checkDestinationDirectory(t *testing.T) {
	tests := []struct {
		name        string
		destination string
		setupFunc   func() (string, func())
		wantErr     bool
	}{
		{
			name:        "valid destination directory",
			destination: "",
			setupFunc: func() (string, func()) {
				tempDir, _ := os.MkdirTemp("", "test-dest-*")
				return tempDir, func() { os.RemoveAll(tempDir) }
			},
			wantErr: false,
		},
		{
			name:        "empty destination defaults to current dir",
			destination: "",
			setupFunc: func() (string, func()) {
				return "", func() {}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			destination, cleanup := tt.setupFunc()
			defer cleanup()

			if destination == "" && tt.name != "empty destination defaults to current dir" {
				tt.destination = destination
			}

			mockSource := &MockTemplateSource{stringResult: "mock-source"}
			mockPrinter := &MockPrinter{}
			context := TemplateContext{Name: "test-name", TemplateName: "test-template"}

			scaffolder := NewScaffolder(mockSource, tt.destination, context, false, mockPrinter, false)
			err := scaffolder.checkDestinationDirectory()

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}

				// Check that destination is set correctly
				expectedDest := tt.destination
				if expectedDest == "" {
					expectedDest = "."
				}
				if scaffolder.destination != expectedDest {
					t.Errorf("Expected destination to be '%s', got '%s'", expectedDest, scaffolder.destination)
				}
			}
		})
	}
}

func TestScaffolder_processTemplate(t *testing.T) {
	tests := []struct {
		name           string
		content        string
		context        TemplateContext
		filename       string
		expectedResult string
		wantErr        bool
		expectedErr    string
	}{
		{
			name:           "no template variables",
			content:        "plain text content",
			context:        TemplateContext{Name: "test-name", TemplateName: "test-template"},
			filename:       "test.txt",
			expectedResult: "plain text content",
			wantErr:        false,
		},
		{
			name:           "template with Name variable",
			content:        "name: <<.Name>>",
			context:        TemplateContext{Name: "my-stackpack", TemplateName: "generic"},
			filename:       "test.conf",
			expectedResult: "name: my-stackpack",
			wantErr:        false,
		},
		{
			name:           "template with TemplateName variable",
			content:        "template: <<.TemplateName>>",
			context:        TemplateContext{Name: "my-stackpack", TemplateName: "webapp"},
			filename:       "test.conf",
			expectedResult: "template: webapp",
			wantErr:        false,
		},
		{
			name:           "template with multiple variables",
			content:        "name=<<.Name>>\ntemplate=<<.TemplateName>>",
			context:        TemplateContext{Name: "awesome-pack", TemplateName: "api"},
			filename:       "config.txt",
			expectedResult: "name=awesome-pack\ntemplate=api",
			wantErr:        false,
		},
		{
			name:        "invalid template syntax",
			content:     "name: <<.InvalidField>>",
			context:     TemplateContext{Name: "test-name", TemplateName: "test-template"},
			filename:    "test.conf",
			wantErr:     true,
			expectedErr: "template execution error",
		},
		{
			name:           "malformed template delimiters",
			content:        "name: <<.Name",
			context:        TemplateContext{Name: "test-name", TemplateName: "test-template"},
			filename:       "test.conf",
			wantErr:        false, // Go templates are more permissive, this actually works as literal text
			expectedResult: "name: <<.Name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSource := &MockTemplateSource{stringResult: "mock-source"}
			mockPrinter := &MockPrinter{}

			scaffolder := NewScaffolder(mockSource, "/test/dest", tt.context, false, mockPrinter, false)
			result, err := scaffolder.processTemplate(tt.content, tt.filename)

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
				if result != tt.expectedResult {
					t.Errorf("Expected result to be '%s', got '%s'", tt.expectedResult, result)
				}
			}
		})
	}
}

//nolint:funlen
func TestScaffolder_renderTemplates(t *testing.T) {
	tests := []struct {
		name        string
		setupFiles  map[string]string
		context     TemplateContext
		wantErr     bool
		expectedErr string
	}{
		{
			name: "render text files with templates",
			setupFiles: map[string]string{
				"config.conf": "name=<<.Name>>",
				"readme.md":   "# <<.Name>> StackPack\nTemplate: <<.TemplateName>>",
				"data.json":   `{"name": "<<.Name>>", "template": "<<.TemplateName>>"}`,
			},
			context: TemplateContext{Name: "my-pack", TemplateName: "webapp"},
			wantErr: false,
		},
		{
			name: "copy files without template variables",
			setupFiles: map[string]string{
				"plain.txt": "no template variables here",
				"script.sh": "#!/bin/bash\necho 'hello world'",
			},
			context: TemplateContext{Name: "my-pack", TemplateName: "webapp"},
			wantErr: false,
		},
		{
			name: "mixed template and non-template files",
			setupFiles: map[string]string{
				"config.conf":        "name=<<.Name>>",
				"plain.txt":          "static content",
				"subdir/nested.conf": "template=<<.TemplateName>>",
			},
			context: TemplateContext{Name: "test-pack", TemplateName: "generic"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary source directory
			sourceDir, err := os.MkdirTemp("", "test-render-source-*")
			if err != nil {
				t.Fatalf("Failed to create temp source dir: %v", err)
			}
			defer os.RemoveAll(sourceDir)

			// Create test files
			for filePath, content := range tt.setupFiles {
				fullPath := filepath.Join(sourceDir, filePath)
				if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
					t.Fatalf("Failed to create directory structure: %v", err)
				}
				if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
					t.Fatalf("Failed to create test file: %v", err)
				}
			}

			mockSource := &MockTemplateSource{stringResult: "mock-source"}
			mockPrinter := &MockPrinter{}

			scaffolder := NewScaffolder(mockSource, "/test/dest", tt.context, false, mockPrinter, false)
			renderedDir, err := scaffolder.renderTemplates(sourceDir)

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

				// Clean up rendered directory
				defer os.RemoveAll(renderedDir)

				// Verify rendered files exist and have correct content
				for filePath, originalContent := range tt.setupFiles {
					renderedPath := filepath.Join(renderedDir, filePath)
					if _, err := os.Stat(renderedPath); os.IsNotExist(err) {
						t.Errorf("Expected rendered file %s to exist", filePath)
						continue
					}

					renderedContent, err := os.ReadFile(renderedPath)
					if err != nil {
						t.Errorf("Failed to read rendered file %s: %v", filePath, err)
						continue
					}

					// Check if template processing occurred
					contentStr := string(renderedContent)
					if strings.Contains(originalContent, "<<") && strings.Contains(originalContent, ">>") {
						// Should be processed
						if strings.Contains(contentStr, "<<") || strings.Contains(contentStr, ">>") {
							t.Errorf("Template variables not processed in file %s: %s", filePath, contentStr)
						}
						// Check for either Name or TemplateName in content depending on which was used
						if strings.Contains(originalContent, "<<.Name>>") && !strings.Contains(contentStr, tt.context.Name) {
							t.Errorf("Expected rendered content to contain name '%s' in file %s", tt.context.Name, filePath)
						}
						if strings.Contains(originalContent, "<<.TemplateName>>") && !strings.Contains(contentStr, tt.context.TemplateName) {
							t.Errorf("Expected rendered content to contain template name '%s' in file %s", tt.context.TemplateName, filePath)
						}
					} else if contentStr != originalContent {
						t.Errorf("Non-template file %s was modified: expected '%s', got '%s'",
							filePath, originalContent, contentStr)
					}
				}
			}
		})
	}
}

func TestScaffolder_checkForConflicts(t *testing.T) {
	tests := []struct {
		name          string
		sourceFiles   []string
		destFiles     []string
		expectedCount int
		wantErr       bool
	}{
		{
			name:          "no conflicts",
			sourceFiles:   []string{"file1.txt", "file2.txt"},
			destFiles:     []string{"other.txt"},
			expectedCount: 0,
			wantErr:       false,
		},
		{
			name:          "some conflicts",
			sourceFiles:   []string{"file1.txt", "file2.txt", "file3.txt"},
			destFiles:     []string{"file1.txt", "file3.txt"},
			expectedCount: 2,
			wantErr:       false,
		},
		{
			name:          "all conflicts",
			sourceFiles:   []string{"file1.txt", "file2.txt"},
			destFiles:     []string{"file1.txt", "file2.txt"},
			expectedCount: 2,
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary source directory
			sourceDir, err := os.MkdirTemp("", "test-conflicts-source-*")
			if err != nil {
				t.Fatalf("Failed to create temp source dir: %v", err)
			}
			defer os.RemoveAll(sourceDir)

			// Create temporary destination directory
			destDir, err := os.MkdirTemp("", "test-conflicts-dest-*")
			if err != nil {
				t.Fatalf("Failed to create temp dest dir: %v", err)
			}
			defer os.RemoveAll(destDir)

			// Create source files
			for _, fileName := range tt.sourceFiles {
				filePath := filepath.Join(sourceDir, fileName)
				if err := os.WriteFile(filePath, []byte("source content"), 0644); err != nil {
					t.Fatalf("Failed to create source file: %v", err)
				}
			}

			// Create destination files
			for _, fileName := range tt.destFiles {
				filePath := filepath.Join(destDir, fileName)
				if err := os.WriteFile(filePath, []byte("dest content"), 0644); err != nil {
					t.Fatalf("Failed to create dest file: %v", err)
				}
			}

			mockSource := &MockTemplateSource{stringResult: "mock-source"}
			mockPrinter := &MockPrinter{}
			context := TemplateContext{Name: "test-name", TemplateName: "test-template"}

			scaffolder := NewScaffolder(mockSource, destDir, context, false, mockPrinter, false)
			conflicts, err := scaffolder.checkForConflicts(sourceDir, destDir)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
					return
				}
				if len(conflicts) != tt.expectedCount {
					t.Errorf("Expected %d conflicts, got %d: %v", tt.expectedCount, len(conflicts), conflicts)
				}
			}
		})
	}
}

func TestScaffolder_JSONOutput(t *testing.T) {
	// Create test files map
	testFiles := map[string]string{
		"config.conf": "name=<<.Name>>",
		"readme.md":   "# <<.Name>> StackPack",
		"static.txt":  "no template variables here",
	}

	tests := []struct {
		name       string
		jsonOutput bool
		expectLogs bool
	}{
		{
			name:       "normal output mode",
			jsonOutput: false,
			expectLogs: true,
		},
		{
			name:       "JSON output mode",
			jsonOutput: true,
			expectLogs: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary source directory for each test iteration
			sourceDir, err := os.MkdirTemp("", "test-json-source-*")
			if err != nil {
				t.Fatalf("Failed to create temp source dir: %v", err)
			}
			defer os.RemoveAll(sourceDir)

			// Create temporary destination directory for each test iteration
			destDir, err := os.MkdirTemp("", "test-json-dest-*")
			if err != nil {
				t.Fatalf("Failed to create temp dest dir: %v", err)
			}
			defer os.RemoveAll(destDir)

			// Create test files in source directory
			for filePath, content := range testFiles {
				fullPath := filepath.Join(sourceDir, filePath)
				if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
					t.Fatalf("Failed to create test file: %v", err)
				}
			}

			mockSource := &MockTemplateSource{
				fetchResult:  sourceDir,
				stringResult: "mock-source",
			}
			mockPrinter := &MockPrinter{}
			templateCtx := TemplateContext{Name: "test-pack", TemplateName: "test-template"}

			scaffolder := NewScaffolder(mockSource, destDir, templateCtx, false, mockPrinter, tt.jsonOutput)
			result, cleanup, err := scaffolder.Scaffold(context.TODO())

			if err != nil {
				t.Errorf("Expected no error but got: %v", err)
				return
			}

			// Check result structure
			if result.Success != true {
				t.Errorf("Expected success=true, got %v", result.Success)
			}
			if result.Name != "test-pack" {
				t.Errorf("Expected name='test-pack', got %s", result.Name)
			}
			if result.Template != "test-template" {
				t.Errorf("Expected template='test-template', got %s", result.Template)
			}
			if result.Source != "mock-source" {
				t.Errorf("Expected source='mock-source', got %s", result.Source)
			}
			if result.FilesCount != len(testFiles) {
				t.Errorf("Expected files_count=%d, got %d", len(testFiles), result.FilesCount)
			}

			// Check progress output behavior
			if tt.expectLogs {
				if len(mockPrinter.printedLines) == 0 {
					t.Errorf("Expected progress messages in normal mode, but got none")
				}
			} else {
				if len(mockPrinter.printedLines) > 0 {
					t.Errorf("Expected no progress messages in JSON mode, but got: %v", mockPrinter.printedLines)
				}
			}

			// Verify files were actually created
			for fileName := range testFiles {
				destPath := filepath.Join(destDir, fileName)
				if _, err := os.Stat(destPath); os.IsNotExist(err) {
					t.Errorf("Expected file %s to be copied to destination", fileName)
				}
			}

			err = cleanup()
			assert.NoError(t, err)
		})
	}
}
