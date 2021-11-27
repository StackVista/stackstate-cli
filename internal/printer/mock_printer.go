package printer

import (
	msg "gitlab.com/stackvista/stackstate-cli2/internal/messages"
)

type MockPrinter struct {
	PrintStructCalls  *[]interface{}
	PrintErrCalls     *[]error
	StartSpinnerCalls *[]msg.LoadingMsg
	StopSpinnerCalls  *int
	UseColor          bool
}

func NewMockPrinter() MockPrinter {
	printStructCalls := make([]interface{}, 0)
	printErrCalls := make([]error, 0)
	startSpinnerCalls := make([]msg.LoadingMsg, 0)
	var stopSpinnerCalls int
	return MockPrinter{
		PrintStructCalls:  &printStructCalls,
		PrintErrCalls:     &printErrCalls,
		StartSpinnerCalls: &startSpinnerCalls,
		StopSpinnerCalls:  &stopSpinnerCalls,
	}
}

func (p *MockPrinter) PrintStruct(s interface{}) {
	*p.PrintStructCalls = append(*p.PrintStructCalls, s)
}

func (p *MockPrinter) PrintErr(err error) {
	*p.PrintErrCalls = append(*p.PrintErrCalls, err)
}

func (p *MockPrinter) StartSpinner(loadingMsg msg.LoadingMsg) {
	*p.StartSpinnerCalls = append(*p.StartSpinnerCalls, loadingMsg)
}

func (p *MockPrinter) StopSpinner() {
	*p.StopSpinnerCalls++
}

func (p *MockPrinter) SetUseColor(useColor bool) {
	p.UseColor = useColor
}

func (p *MockPrinter) GetUseColor() bool {
	return p.UseColor
}
