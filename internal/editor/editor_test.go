package editor

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEditorYAMLRoundTrip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("The test editor uses a POSIX shell")
	}
	dir := t.TempDir()
	scriptPath := filepath.Join(dir, "editor")
	markerPath := filepath.Join(dir, "edited-file")
	const edited = "url: https://suse-observability.example.com\n"
	const script = `#!/bin/sh
printf 'url: https://suse-observability.example.com\n' > "$1"
printf '%s' "$1" > "$EDITED_FILE_MARKER"
`
	require.NoError(t, os.WriteFile(scriptPath, []byte(script), 0700))
	t.Setenv("VISUAL", scriptPath)
	t.Setenv("EDITOR", "an-editor-that-does-not-exist")
	t.Setenv("EDITED_FILE_MARKER", markerPath)

	content, err := NewEditor().Edit("context", ".yaml", strings.NewReader("url: https://original.example.com\n"))
	require.NoError(t, err)
	assert.Equal(t, edited, string(content))
	path, err := os.ReadFile(markerPath)
	require.NoError(t, err)
	_, err = os.Stat(string(path))
	assert.ErrorIs(t, err, os.ErrNotExist)
}
