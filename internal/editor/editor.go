package editor

import (
	"io"
	"os"

	"k8s.io/kubectl/pkg/cmd/util/editor"
)

type Editor interface {
	Edit(prefix, suffix string, contents io.Reader) ([]byte, error)
}

type defaultEditor struct {
	editor editor.Editor
}

func NewEditor() Editor {
	return &defaultEditor{editor: editor.NewDefaultEditor([]string{"VISUAL", "EDITOR"})}
}

func (e *defaultEditor) Edit(prefix, suffix string, contents io.Reader) ([]byte, error) {
	c, f, err := e.editor.LaunchTempFile(prefix, suffix, contents)
	if f != "" {
		defer os.Remove(f)
	}

	return c, err
}
