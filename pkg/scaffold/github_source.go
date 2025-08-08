package scaffold

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// GitHubSource implements TemplateSource for GitHub repositories
type GitHubSource struct {
	Owner        string // GitHub repository owner
	Repo         string // GitHub repository name
	Ref          string // Git ref (branch, tag, or commit SHA) - defaults to "main"
	Path         string // Path within the repository to the template directory
	TemplateName string // Name of the template
	tempDir      string // Temporary directory for cleanup
}

// NewGitHubSource creates a new GitHubSource
func NewGitHubSource(owner, repo, ref, path, templateName string) *GitHubSource {
	if ref == "" {
		ref = "main" // Default to main branch
	}

	return &GitHubSource{
		Owner:        owner,
		Repo:         repo,
		Ref:          ref,
		Path:         path,
		TemplateName: templateName,
	}
}

// Validate checks if the source is valid
func (g *GitHubSource) Validate() error {
	if g.Owner == "" {
		return fmt.Errorf("GitHub repository owner is required")
	}

	if g.Repo == "" {
		return fmt.Errorf("GitHub repository name is required")
	}

	if g.TemplateName == "" {
		return fmt.Errorf("template name is required")
	}

	return nil
}

// Fetch retrieves templates from GitHub to a temporary directory
func (g *GitHubSource) Fetch(ctx context.Context) (string, error) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "stackpack-template-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary directory: %w", err)
	}
	g.tempDir = tempDir

	// Download repository archive
	archiveURL := fmt.Sprintf("https://github.com/%s/%s/archive/%s.zip", g.Owner, g.Repo, g.Ref)

	req, err := http.NewRequestWithContext(ctx, "GET", archiveURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to download repository archive from %s: %w", archiveURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download repository archive: HTTP %d from %s", resp.StatusCode, archiveURL)
	}

	// Create temporary zip file
	zipFile := filepath.Join(tempDir, "archive.zip")
	out, err := os.Create(zipFile)
	if err != nil {
		return "", fmt.Errorf("failed to create zip file: %w", err)
	}
	defer out.Close()

	// Download the zip file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to download zip file: %w", err)
	}

	// Extract the zip file
	extractDir := filepath.Join(tempDir, "extracted")
	if err := g.extractZip(zipFile, extractDir); err != nil {
		return "", fmt.Errorf("failed to extract zip file: %w", err)
	}

	// Find the template directory within the extracted repository
	templateDir, err := g.findTemplateDirectory(extractDir)
	if err != nil {
		return "", fmt.Errorf("failed to find template directory: %w", err)
	}

	return templateDir, nil
}

// Cleanup removes temporary files
func (g *GitHubSource) Cleanup() error {
	if g.tempDir != "" {
		return os.RemoveAll(g.tempDir)
	}
	return nil
}

// extractZip extracts a zip file to the specified directory
func (g *GitHubSource) extractZip(src, dest string) error {
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	// Create destination directory
	if err := os.MkdirAll(dest, os.FileMode(defaultDirMode)); err != nil {
		return err
	}

	// Extract files
	for _, file := range reader.File {
		path := filepath.Join(dest, file.Name)

		// Security check: ensure the path is within the destination directory
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("invalid file path in zip: %s", file.Name)
		}

		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.FileInfo().Mode()); err != nil {
				return err
			}
			continue
		}

		// Create parent directories
		if err := os.MkdirAll(filepath.Dir(path), os.FileMode(defaultDirMode)); err != nil {
			return err
		}

		// Extract file
		if err := g.extractFile(file, path); err != nil {
			return err
		}
	}

	return nil
}

// extractFile extracts a single file from the zip
func (g *GitHubSource) extractFile(file *zip.File, destPath string) error {
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	outFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.FileInfo().Mode())
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, rc)
	return err
}

// findTemplateDirectory finds the template directory within the extracted repository
func (g *GitHubSource) findTemplateDirectory(extractedDir string) (string, error) {
	// GitHub archives are extracted with a top-level directory named "{repo}-{ref}"
	// We need to find this directory first
	entries, err := os.ReadDir(extractedDir)
	if err != nil {
		return "", err
	}

	var repoDir string
	expectedPrefix := fmt.Sprintf("%s-%s", g.Repo, g.Ref)

	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), expectedPrefix) {
			repoDir = filepath.Join(extractedDir, entry.Name())
			break
		}
	}

	if repoDir == "" {
		return "", fmt.Errorf("could not find repository directory in extracted archive")
	}

	// Build the template base directory path
	var templateBaseDir string
	if g.Path != "" {
		templateBaseDir = filepath.Join(repoDir, g.Path)
	} else {
		templateBaseDir = repoDir
	}

	// Verify the template base directory exists
	if _, err := os.Stat(templateBaseDir); os.IsNotExist(err) {
		if g.Path != "" {
			return "", fmt.Errorf("template base directory not found at path '%s' in repository %s/%s", g.Path, g.Owner, g.Repo)
		}
		return "", fmt.Errorf("repository directory not found")
	}

	// Look for a subdirectory with the template name
	templateDir := filepath.Join(templateBaseDir, g.TemplateName)
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		return "", fmt.Errorf("template '%s' not found in directory '%s' of repository %s/%s", g.TemplateName, templateBaseDir, g.Owner, g.Repo)
	}

	// Check if it's a directory
	info, err := os.Stat(templateDir)
	if err != nil {
		return "", fmt.Errorf("failed to check template directory: %w", err)
	}

	if !info.IsDir() {
		return "", fmt.Errorf("template '%s' is not a directory in repository %s/%s", g.TemplateName, g.Owner, g.Repo)
	}

	return templateDir, nil
}

// String returns a user-friendly description of the GitHub source
func (g *GitHubSource) String() string {
	if g.Path != "" {
		return fmt.Sprintf("github:%s/%s@%s:%s", g.Owner, g.Repo, g.Ref, g.Path)
	}
	return fmt.Sprintf("github:%s/%s@%s", g.Owner, g.Repo, g.Ref)
}
