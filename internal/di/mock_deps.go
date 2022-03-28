package di

import "gitlab.com/stackvista/stackstate-cli2/internal/printer"

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
			Client: mockClient,
		},
		MockClient:  &mockClient,
		MockPrinter: &mockPrinter,
	}
}
