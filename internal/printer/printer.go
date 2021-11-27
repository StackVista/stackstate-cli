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
	stdOut   io.Writer
	stdErr   io.Writer
	useColor bool
	spinner  *pterm.SpinnerPrinter
}

func NewStdPrinter() Printer {
	spinner := pterm.DefaultSpinner.WithRemoveWhenDone()
	return &StdPrinter{
		useColor: false, // IMPORTANT: use progressive enhancement!
		spinner:  spinner,
		stdOut:   os.Stdout,
		stdErr:   os.Stderr,
	}
}

func (p *StdPrinter) PrintStruct(s interface{}) {
	p.StopSpinner()
	msg, _ := json.Marshal(s)
	fmt.Fprintf(p.stdOut, "%s\n", string(msg))
}

func (p *StdPrinter) PrintErr(err error) {
	p.StopSpinner()
	if p.useColor {
		fmt.Fprintf(p.stdErr, "%s\n", color.Red(err.Error()))
	} else {
		fmt.Fprintf(p.stdErr, "%s\n", err.Error())
	}
}

func (p *StdPrinter) StartSpinner(loadingMsg msg.LoadingMsg) {
	p.StopSpinner()
	if p.useColor {
		p.spinner.Text = loadingMsg.String()
		p.spinner.Start()
	}
}

func (p *StdPrinter) StopSpinner() {
	p.spinner.Stop()
}

func (p *StdPrinter) SetUseColor(useColor bool) {
	p.useColor = useColor
	if !useColor {
		p.StopSpinner()
	}
}

func (p *StdPrinter) GetUseColor() bool {
	return p.useColor
}
