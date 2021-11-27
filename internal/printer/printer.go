package printer

import (
	"encoding/json"
	"fmt"
	"io"
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
	SetUseColor(useColor bool)
	GetUseColor() bool
}

type StdPrinter struct {
	stdOut   io.Writer // for test purposes
	stdErr   io.Writer // for test purposes
	useColor bool
	spinner  *pterm.SpinnerPrinter
}

func NewStdPrinter() Printer {
	return &StdPrinter{
		useColor: false, // IMPORTANT: use progressive enhancement!
		spinner:  nil,
		stdOut:   os.Stdout,
		stdErr:   os.Stderr,
	}
}

func (p *StdPrinter) PrintStruct(s interface{}) {
	resetStdPrinter(p)
	msg, _ := json.Marshal(s)
	fmt.Fprintf(p.stdOut, "%s\n", string(msg))
}

func (p *StdPrinter) PrintErr(err error) {
	resetStdPrinter(p)
	if p.useColor {
		fmt.Fprintf(p.stdErr, "%s\n", color.Red(err.Error()))
	} else {
		fmt.Fprintf(p.stdErr, "%s\n", err.Error())
	}
}

func (p *StdPrinter) StartSpinner(loadingMsg msg.LoadingMsg) {
	resetStdPrinter(p)
	if p.spinner != nil {
		p.spinner.Text = loadingMsg.String()
		p.spinner.Start()
	}
}

func (p *StdPrinter) StopSpinner() {
	if p.spinner != nil {
		p.spinner.Stop()
	}
}

func (p *StdPrinter) SetUseColor(useColor bool) {
	if p.useColor == useColor {
		return
	}

	resetStdPrinter(p)
	p.useColor = useColor
	if useColor {
		p.spinner = pterm.DefaultSpinner.WithRemoveWhenDone()
	} else {
		p.spinner = nil
	}
}

func (p *StdPrinter) GetUseColor() bool {
	return p.useColor
}

func resetStdPrinter(p *StdPrinter) {
	p.StopSpinner()
}
