package printer

import (
	"net/http"

	msg "gitlab.com/stackvista/stackstate-cli2/internal/messages"
)

type MockPrinter struct {
	PrintStructCalls      *[]interface{}
	PrintErrCalls         *[]error
	PrintErrResponseCalls *[]PrintErrResponseCall
	StartSpinnerCalls     *[]msg.LoadingMsg
	StopSpinnerCalls      *int
	UseColor              bool
}

type PrintErrResponseCall struct {
	Err  error
	Resp *http.Response
}

func NewMockPrinter() MockPrinter {
	printStructCalls := make([]interface{}, 0)
	printErrCalls := make([]error, 0)
	printErrResponseCalls := make([]PrintErrResponseCall, 0)
	startSpinnerCalls := make([]msg.LoadingMsg, 0)
	var stopSpinnerCalls int
	return MockPrinter{
		PrintStructCalls:      &printStructCalls,
		PrintErrCalls:         &printErrCalls,
		PrintErrResponseCalls: &printErrResponseCalls,
		StartSpinnerCalls:     &startSpinnerCalls,
		StopSpinnerCalls:      &stopSpinnerCalls,
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

func (p *MockPrinter) PrintErrResponse(err error, resp *http.Response) {
	*p.PrintErrResponseCalls = append(*p.PrintErrResponseCalls, PrintErrResponseCall{
		Err:  err,
		Resp: resp,
	})
}
