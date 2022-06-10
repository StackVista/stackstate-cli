package di

import (
	"bytes"
	"context"
	"path/filepath"
	"testing"

	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
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
			Clock:          pflags.NewTestClock(1652108645000), //nolint:gomnd
			Context:        context.Background(),
			StsConfig:      &config.Config{},
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
