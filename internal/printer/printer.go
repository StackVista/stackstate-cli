package printer

import (
	"encoding/json"
	"fmt"
	"os"

	color "github.com/logrusorgru/aurora/v3"
	"github.com/pterm/pterm"
	msg "gitlab.com/stackvista/stackstate-cli2/internal/messages"
)

type Printer interface {
	PrintStruct(s interface{})
	PrintErr(err error)
	StartSpinner(loadingMsg msg.LoadingMsg)
	StopSpinner()
}

type StdPrinter struct {
	spinner *pterm.SpinnerPrinter
}

func NewStdPrinter() Printer {
	spinner := pterm.DefaultSpinner.WithRemoveWhenDone()
	return StdPrinter{
		spinner: spinner,
	}
}

func (p StdPrinter) PrintStruct(s interface{}) {
	p.StopSpinner()
	x, _ := json.Marshal(s)
	println(string(x))
}

func (p StdPrinter) PrintErr(err error) {
	p.StopSpinner()
	fmt.Fprintf(os.Stderr, "%s\n", color.Red(err.Error()))
}

func (p StdPrinter) StartSpinner(loadingMsg msg.LoadingMsg) {
	p.StopSpinner()
	p.spinner.Text = loadingMsg.String()
	p.spinner.Start()
}

func (p StdPrinter) StopSpinner() {
	p.spinner.Stop()
}
