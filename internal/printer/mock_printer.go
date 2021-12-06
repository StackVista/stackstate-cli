package printer

import (
	"net/http"

	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

type MockPrinter struct {
	PrintStructCalls  *[]interface{}
	PrintErrCalls     *[]error
	StartSpinnerCalls *[]common.LoadingMsg
	StopSpinnerCalls  *int
	UseColor          bool
	outputType        OutputType
}

type PrintErrResponseCall struct {
	Err  error
	Resp *http.Response
}

func NewMockPrinter() MockPrinter {
	printStructCalls := make([]interface{}, 0)
	printErrCalls := make([]error, 0)
	startSpinnerCalls := make([]common.LoadingMsg, 0)
	var stopSpinnerCalls int
	return MockPrinter{
		PrintStructCalls:  &printStructCalls,
		PrintErrCalls:     &printErrCalls,
		StartSpinnerCalls: &startSpinnerCalls,
		StopSpinnerCalls:  &stopSpinnerCalls,
	}
}

func (p *MockPrinter) PrintStruct(s interface{}) error {
	*p.PrintStructCalls = append(*p.PrintStructCalls, s)
	return nil
}

func (p *MockPrinter) PrintErr(err error) {
	*p.PrintErrCalls = append(*p.PrintErrCalls, err)
}

func (p *MockPrinter) StartSpinner(loadingMsg common.LoadingMsg) {
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

func (p *MockPrinter) SetOutputType(outputType OutputType) {
	p.outputType = outputType
}

func (p *MockPrinter) GetOutputType() OutputType {
	return p.outputType
}
