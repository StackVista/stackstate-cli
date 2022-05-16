package di

import (
	"bytes"
	"context"

	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

type MockDeps struct {
	Deps
	MockClient  *MockStackStateClient
	MockPrinter *printer.MockPrinter
	StdOut      *bytes.Buffer
	StdErr      *bytes.Buffer
}

func NewMockDeps() MockDeps {
	stdOut := new(bytes.Buffer)
	stdErr := new(bytes.Buffer)
	mockClient := NewMockStackStateClient()
	mockPrinter := printer.NewMockPrinter(printer.NewStdPrinter("linux", stdOut, stdErr))
	return MockDeps{
		Deps: Deps{
			Client:    &mockClient,
			Printer:   &mockPrinter,
			Context:   context.Background(),
			Config:    &conf.Conf{},
			Version:   "1.0.0",
			Commit:    "123124",
			BuildDate: "1-1-2022",
			CLIType:   "full",
		},
		MockClient:  &mockClient,
		MockPrinter: &mockPrinter,
		StdOut:      stdOut,
		StdErr:      stdErr,
	}
}
