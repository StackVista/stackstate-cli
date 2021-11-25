package printer

import (
	"encoding/json"
	"fmt"
	"os"

	color "github.com/logrusorgru/aurora/v3"
)

type Printer interface {
	PrintStruct(s interface{})
	PrintErr(err error)
}

type StdPrinter struct{}

func NewStdPrinter() Printer {
	return StdPrinter{}
}

func (p StdPrinter) PrintStruct(s interface{}) {
	x, _ := json.Marshal(s)
	println(string(x))
}

func (p StdPrinter) PrintErr(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", color.Red(err.Error()))
}
