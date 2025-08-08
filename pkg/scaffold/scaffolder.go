package scaffold

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	defaultDirMode = 0755 // Default directory permissions
)

// TemplateContext holds the variables for template rendering
type TemplateContext struct {
	Name         string // Stackpack name (e.g. "my-stackpack")
	TemplateName string // Template name used for scaffolding
}

// Scaffolder handles the scaffolding process
type Scaffolder struct {
	source      TemplateSource
	destination string
	context     TemplateContext
	force       bool
	printer     Printer
	jsonOutput  bool
}

// Printer interface for output (to avoid importing internal packages)
type Printer interface {
	PrintLn(message string)
	PrintWarn(message string)
}

// NewScaffolder creates a new Scaffolder instance
func NewScaffolder(source TemplateSource, destination string, context TemplateContext, force bool, printer Printer, jsonOutput bool) *Scaffolder {
	return &Scaffolder{
		source:      source,
		destination: destination,
		context:     context,
		force:       force,
		printer:     printer,
		jsonOutput:  jsonOutput,
	}
}

// Scaffold executes the complete scaffolding workflow
func (s *Scaffolder) Scaffold(ctx context.Context) (*ScaffoldResult, func() error, error) {
	result := &ScaffoldResult{
		Success:     false,
		Source:      s.source.String(),
		Destination: s.destination,
		Name:        s.context.Name,
		Template:    s.context.TemplateName,
	}

	cleanUpFn := s.source.Cleanup

	// Validate arguments
	s.printProgress("✓ Validating arguments...")
	if err := s.validateArgs(); err != nil {
		return result, cleanUpFn, fmt.Errorf("validation failed: %w", err)
	}

	// Check destination directory
	s.printProgress("✓ Checking destination directory...")
	if err := s.checkDestinationDirectory(); err != nil {
		return result, cleanUpFn, fmt.Errorf("destination directory check failed: %w", err)
	}
	result.Destination = s.destination // Update with resolved destination

	// Fetch template to temp directory
	s.printProgress("✓ Fetching template from " + s.source.String())
	tempDir, err := s.source.Fetch(ctx)
	if err != nil {
		return result, cleanUpFn, fmt.Errorf("failed to fetch template: %w", err)
	}

	// Render templates
	s.printProgress("✓ Rendering templates...")
	renderedDir, err := s.renderTemplates(tempDir)
	if err != nil {
		return result, cleanUpFn, fmt.Errorf("failed to render templates: %w", err)
	}

	// Validate rendered templates
	s.printProgress("✓ Validating rendered templates...")
	if err := s.validateRenderedTemplates(renderedDir); err != nil {
		return result, cleanUpFn, fmt.Errorf("rendered template validation failed: %w", err)
	}

	// Copy to destination
	s.printProgress("✓ Copying files to destination...")
	copyResult, err := s.copyToDestinationWithResult(renderedDir)
	if err != nil {
		return result, cleanUpFn, fmt.Errorf("failed to copy to destination: %w", err)
	}

	// Update result with copied files info
	result.Success = true
	result.FilesCount = len(copyResult.CopiedFiles)
	result.Files = copyResult.CopiedFiles

	return result, cleanUpFn, nil
}

// validateArgs validates the scaffolding arguments
func (s *Scaffolder) validateArgs() error {
	if err := s.source.Validate(); err != nil {
		return err
	}

	if s.context.Name == "" {
		return fmt.Errorf("stackpack name is required")
	}

	if s.context.TemplateName == "" {
		return fmt.Errorf("template name is required")
	}

	return nil
}

// checkDestinationDirectory validates and prepares the destination directory
func (s *Scaffolder) checkDestinationDirectory() error {
	if s.destination == "" {
		s.destination = "."
	}

	// Create destination directory if it doesn't exist
	if err := os.MkdirAll(s.destination, os.FileMode(defaultDirMode)); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Check if destination is empty (in the future, we might want to enforce this)
	// For now, we'll allow non-empty destinations

	return nil
}

// renderTemplates processes templates with context variables using Go templates with << >> delimiters
func (s *Scaffolder) renderTemplates(sourceDir string) (string, error) {
	// Create a temporary directory for rendered templates
	renderedDir, err := os.MkdirTemp("", "stackpack-rendered-*")
	if err != nil {
		return "", fmt.Errorf("failed to create rendered templates directory: %w", err)
	}

	// Process all files in the source directory
	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate relative path and destination path
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(renderedDir, relPath)

		// Handle directories
		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		// Process files - render templates or copy as-is
		return s.renderFile(path, destPath, info.Mode())
	})

	if err != nil {
		os.RemoveAll(renderedDir) // Cleanup on error
		return "", err
	}

	return renderedDir, nil
}

// renderFile processes a single file through Go template engine
func (s *Scaffolder) renderFile(srcPath, destPath string, mode os.FileMode) error {
	// Read source file content
	content, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("failed to read template file %s: %w", srcPath, err)
	}

	// Skip binary files (simple heuristic: contains null bytes)
	if bytes.Contains(content, []byte{0}) {
		// Copy binary files as-is
		return s.copyBinaryFile(srcPath, destPath, mode)
	}

	// Process text files through Go template engine
	renderedContent, err := s.processTemplate(string(content), srcPath)
	if err != nil {
		return fmt.Errorf("failed to process template %s: %w", srcPath, err)
	}

	// Write rendered content to destination
	if err := os.WriteFile(destPath, []byte(renderedContent), mode); err != nil {
		return fmt.Errorf("failed to write rendered file %s: %w", destPath, err)
	}

	return nil
}

// processTemplate processes template content with Go templates using << >> delimiters
func (s *Scaffolder) processTemplate(content, filename string) (string, error) {
	// Check if content contains our template delimiters
	if !strings.Contains(content, "<<") || !strings.Contains(content, ">>") {
		// No template variables, return as-is
		return content, nil
	}

	// Create Go template with custom delimiters
	tmpl, err := template.New(filepath.Base(filename)).Delims("<<", ">>").Parse(content)
	if err != nil {
		return "", fmt.Errorf("template parsing error: %w", err)
	}

	// Execute template with context
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, s.context); err != nil {
		return "", fmt.Errorf("template execution error: %w", err)
	}

	return buf.String(), nil
}

// copyBinaryFile copies binary files without template processing
func (s *Scaffolder) copyBinaryFile(srcPath, destPath string, mode os.FileMode) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}

	return os.Chmod(destPath, mode)
}

// validateRenderedTemplates checks if the rendered templates are valid
// This is a stub implementation
func (s *Scaffolder) validateRenderedTemplates(renderedDir string) error {
	// Check if the directory exists
	if _, err := os.Stat(renderedDir); os.IsNotExist(err) {
		return fmt.Errorf("rendered template directory does not exist: %s", renderedDir)
	}

	// In the future, this could validate:
	// - Template syntax

	return nil
}

// CopyResult holds information about copied files
type CopyResult struct {
	CopiedFiles []string
}

// ScaffoldResult holds the result of scaffolding operation for JSON output
type ScaffoldResult struct {
	Success     bool     `json:"success"`
	Source      string   `json:"source"`
	Destination string   `json:"destination"`
	Name        string   `json:"name"`
	Template    string   `json:"template"`
	FilesCount  int      `json:"files_count"`
	Files       []string `json:"files,omitempty"`
}

// checkForConflicts scans for existing files that would be overwritten
func (s *Scaffolder) checkForConflicts(src, dst string) ([]string, error) {
	var conflicts []string

	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Calculate the destination path
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(dst, relPath)

		// Check if destination file already exists
		if _, err := os.Stat(destPath); err == nil {
			conflicts = append(conflicts, relPath)
		}

		return nil
	})

	return conflicts, err
}

// copyDir recursively copies a directory (assumes no conflicts when called)
func (s *Scaffolder) copyDir(src, dst string) (*CopyResult, error) {
	result := &CopyResult{
		CopiedFiles: []string{},
	}

	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate the destination path
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		// Copy the file (conflicts already handled in pre-flight check)
		if err := s.copyFile(path, destPath); err != nil {
			return err
		}

		result.CopiedFiles = append(result.CopiedFiles, relPath)
		return nil
	})

	return result, err
}

// copyFile copies a single file
func (s *Scaffolder) copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// Copy file permissions
	sourceInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	return os.Chmod(dst, sourceInfo.Mode())
}

// printProgress conditionally prints progress messages (suppressed in JSON mode)
func (s *Scaffolder) printProgress(message string) {
	if !s.jsonOutput && s.printer != nil {
		s.printer.PrintLn(message)
	}
}

// copyToDestinationWithResult copies files and returns detailed results
func (s *Scaffolder) copyToDestinationWithResult(renderedDir string) (*CopyResult, error) {
	// Pre-flight check: scan for conflicts before making any changes
	if !s.force {
		conflicts, err := s.checkForConflicts(renderedDir, s.destination)
		if err != nil {
			return nil, fmt.Errorf("failed to check for file conflicts: %w", err)
		}

		if len(conflicts) > 0 {
			if !s.jsonOutput && s.printer != nil {
				s.printer.PrintWarn("The following files already exist and would be overwritten:")
				for _, file := range conflicts {
					s.printer.PrintLn("  " + file)
				}
				s.printer.PrintLn("")
				s.printer.PrintLn("Use --force flag to overwrite existing files, or remove/rename the conflicting files.")
			}
			return nil, fmt.Errorf("conflicting files exist, use --force to overwrite or resolve conflicts manually")
		}
	}

	// No conflicts (or force flag used), proceed with copying
	result, err := s.copyDir(renderedDir, s.destination)
	if err != nil {
		return nil, err
	}

	// Print the list of copied files (only in non-JSON mode)
	if !s.jsonOutput && s.printer != nil && len(result.CopiedFiles) > 0 {
		s.printer.PrintLn("Files copied:")
		for _, file := range result.CopiedFiles {
			s.printer.PrintLn("  " + file)
		}
	}

	return result, nil
}
