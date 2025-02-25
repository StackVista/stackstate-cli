package di

import (
	"bytes"
	"context"
	"path/filepath"
	"testing"

	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/editor"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stackvista/stackstate-cli/pkg/pflags"
)

type MockDeps struct {
	Deps
	MockClient  *MockStackStateClient
	MockPrinter *printer.MockPrinter
	StdOut      *bytes.Buffer
	StdErr      *bytes.Buffer
}

func NewMockDeps(t *testing.T) MockDeps {
	stdOut := new(bytes.Buffer)
	stdErr := new(bytes.Buffer)
	mockClient := NewMockStackStateClient()
	mockPrinter := printer.NewMockPrinter(printer.NewStdPrinter("linux", stdOut, stdErr))
	return MockDeps{
		Deps: Deps{
			Client:         &mockClient,
			Printer:        &mockPrinter,
			Editor:         &editor.ReverseEditor{},
			Clock:          pflags.NewTestClock(1652108645000), //nolint:mnd
			Context:        context.Background(),
			CurrentContext: &config.StsContext{},
			ConfigPath:     filepath.Join(t.TempDir(), config.ConfigFileName),
			Version:        "1.0.0",
			Commit:         "123124",
			BuildDate:      "1-1-2022",
		},
		MockClient:  &mockClient,
		MockPrinter: &mockPrinter,
		StdOut:      stdOut,
		StdErr:      stdErr,
	}
}
