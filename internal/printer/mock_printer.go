package printer

type MockPrinter struct {
	PrintStructCalls *[]interface{}
	PrintErrCalls    *[]error
}

func NewMockPrinter() MockPrinter {
	printStructCalls := make([]interface{}, 0)
	printErrCalls := make([]error, 0)
	return MockPrinter{
		PrintStructCalls: &printStructCalls,
		PrintErrCalls:    &printErrCalls,
	}
}

func (p MockPrinter) PrintStruct(s interface{}) {
	*p.PrintStructCalls = append(*p.PrintStructCalls, s)
}

func (p MockPrinter) PrintErr(err error) {
	*p.PrintErrCalls = append(*p.PrintErrCalls, err)
}
