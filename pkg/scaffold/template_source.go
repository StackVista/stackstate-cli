package scaffold

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

// TemplateSource interface for different template sources
type TemplateSource interface {
	// Fetch retrieves templates to a temporary directory
	Fetch(ctx context.Context) (string, error)
	// Validate checks if the source is valid
	Validate() error
	// Cleanup removes temporary files
	Cleanup() error
	// String returns a user-friendly description of the template source
	String() string
}

// LocalDirSource implements TemplateSource for local directories
type LocalDirSource struct {
	Path         string
	TemplateName string
	tempDir      string
}

// NewLocalDirSource creates a new LocalDirSource
func NewLocalDirSource(path, templateName string) *LocalDirSource {
	return &LocalDirSource{
		Path:         path,
		TemplateName: templateName,
	}
}

// Validate checks if the source is valid
func (l *LocalDirSource) Validate() error {
	if l.Path == "" {
		return fmt.Errorf("template path is required")
	}

	if l.TemplateName == "" {
		return fmt.Errorf("template name is required")
	}

	// Check if the path exists
	if _, err := os.Stat(l.Path); os.IsNotExist(err) {
		return fmt.Errorf("template directory does not exist: %s", l.Path)
	}

	// Check if it's a directory
	info, err := os.Stat(l.Path)
	if err != nil {
		return fmt.Errorf("failed to check template directory: %w", err)
	}

	if !info.IsDir() {
		return fmt.Errorf("template path is not a directory: %s", l.Path)
	}

	return nil
}

// Fetch retrieves templates to a temporary directory
func (l *LocalDirSource) Fetch(ctx context.Context) (string, error) {
	// Resolve the base template directory
	absPath, err := filepath.Abs(l.Path)
	if err != nil {
		return "", fmt.Errorf("failed to resolve absolute path: %w", err)
	}

	// Look for a subdirectory with the template name
	templatePath := filepath.Join(absPath, l.TemplateName)
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return "", fmt.Errorf("template '%s' not found in directory '%s'", l.TemplateName, absPath)
	}

	// Check if it's a directory
	info, err := os.Stat(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to check template directory: %w", err)
	}

	if !info.IsDir() {
		return "", fmt.Errorf("template '%s' is not a directory in '%s'", l.TemplateName, absPath)
	}

	l.tempDir = templatePath
	return templatePath, nil
}

// Cleanup removes temporary files
func (l *LocalDirSource) Cleanup() error {
	// For local directories, we don't need to cleanup anything
	// In the future implementations, this might delete temp directories
	return nil
}

// String returns a user-friendly description of the local directory source
func (l *LocalDirSource) String() string {
	return fmt.Sprintf("localdir:%s", l.Path)
}
