package di

import (
	"context"

	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

type MockDeps struct {
	Deps
	MockClient  *MockStackStateClient
	MockPrinter *printer.MockPrinter
}

func NewMockDeps() MockDeps {
	mockClient := NewMockStackStateClient()
	mockPrinter := printer.NewMockPrinter()
	return MockDeps{
		Deps: Deps{
			Client:  &mockClient,
			Printer: &mockPrinter,
			Context: context.Background(),
			Config:  &conf.Conf{},
		},
		MockClient:  &mockClient,
		MockPrinter: &mockPrinter,
	}
}
