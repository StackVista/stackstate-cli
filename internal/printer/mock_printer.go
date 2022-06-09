package printer

import (
	"fmt"
	"net/http"
)

type MockPrinter struct {
	PrintJsonCalls    *[]map[string]interface{}
	PrintErrJsonCalls *[]error
	PrintStructCalls  *[]interface{}
	PrintErrCalls     *[]error
	StartSpinnerCalls *[]LoadingMsg
	StopSpinnerCalls  *int
	UseColor          bool
	SuccessCalls      *[]string
	PrintWarnCalls    *[]string
	TableCalls        *[]TableData
	PrintLnCalls      *[]string
	HasNonJsonCalls   bool
	DelegatePrinter   Printer
}

type PrintErrResponseCall struct {
	Err  error
	Resp *http.Response
}

func NewMockPrinter(pr Printer) MockPrinter {
	printJsonCalls := make([]map[string]interface{}, 0)
	printErrJsonCalls := make([]error, 0)
	printStructCalls := make([]interface{}, 0)
	printErrCalls := make([]error, 0)
	startSpinnerCalls := make([]LoadingMsg, 0)
	successCalls := make([]string, 0)
	printWarnCalls := make([]string, 0)
	printTableCalls := make([]TableData, 0)
	printLnCalls := make([]string, 0)
	return MockPrinter{
		PrintJsonCalls:    &printJsonCalls,
		PrintErrJsonCalls: &printErrJsonCalls,
		PrintStructCalls:  &printStructCalls,
		PrintErrCalls:     &printErrCalls,
		StartSpinnerCalls: &startSpinnerCalls,
		SuccessCalls:      &successCalls,
		PrintWarnCalls:    &printWarnCalls,
		TableCalls:        &printTableCalls,
		PrintLnCalls:      &printLnCalls,
		DelegatePrinter:   pr,
	}
}

func (p *MockPrinter) PrintStruct(s interface{}) {
	*p.PrintStructCalls = append(*p.PrintStructCalls, s)
	p.HasNonJsonCalls = true
	if p.DelegatePrinter != nil {
		p.DelegatePrinter.PrintStruct(s)
	}
}

func (p *MockPrinter) PrintJson(s map[string]interface{}) {
	*p.PrintJsonCalls = append(*p.PrintJsonCalls, s)
	if p.DelegatePrinter != nil {
		p.DelegatePrinter.PrintJson(s)
	}
}

func (p *MockPrinter) PrintErrJson(err error) {
	*p.PrintErrJsonCalls = append(*p.PrintErrJsonCalls, err)
	if p.DelegatePrinter != nil {
		p.DelegatePrinter.PrintErrJson(err)
	}
}

func (p *MockPrinter) PrintErr(err error) {
	*p.PrintErrCalls = append(*p.PrintErrCalls, err)
	p.HasNonJsonCalls = true
	if p.DelegatePrinter != nil {
		p.DelegatePrinter.PrintErr(err)
	}
}

func (p *MockPrinter) StartSpinner(loadingMsg LoadingMsg) StopPrinterFn {
	*p.StartSpinnerCalls = append(*p.StartSpinnerCalls, loadingMsg)
	p.HasNonJsonCalls = true

	if p.DelegatePrinter != nil {
		return p.DelegatePrinter.StartSpinner(loadingMsg)
	} else {
		return func() {}
	}
}

func (p *MockPrinter) SetUseColor(useColor bool) {
	p.UseColor = useColor
	if p.DelegatePrinter != nil {
		p.DelegatePrinter.SetUseColor(useColor)
	}
}

func (p *MockPrinter) GetUseColor() bool {
	return p.UseColor
}

func (p *MockPrinter) Success(msg string) {
	p.HasNonJsonCalls = true
	*p.SuccessCalls = append(*p.SuccessCalls, msg)
	if p.DelegatePrinter != nil {
		p.DelegatePrinter.Success(msg)
	}
}

func (p *MockPrinter) Successf(format string, args ...interface{}) {
	p.Success(fmt.Sprintf(format, args...))
}

func (p *MockPrinter) PrintWarn(msg string) {
	p.HasNonJsonCalls = true
	*p.PrintWarnCalls = append(*p.PrintWarnCalls, msg)
	if p.DelegatePrinter != nil {
		p.DelegatePrinter.PrintWarn(msg)
	}
}

func (p *MockPrinter) Table(t TableData) {
	p.HasNonJsonCalls = true
	*p.TableCalls = append(*p.TableCalls, t)
	if p.DelegatePrinter != nil {
		p.DelegatePrinter.Table(t)
	}
}

func (p *MockPrinter) PrintLn(text string) {
	p.HasNonJsonCalls = true
	*p.PrintLnCalls = append(*p.PrintLnCalls, text)
	if p.DelegatePrinter != nil {
		p.DelegatePrinter.PrintLn(text)
	}
}
