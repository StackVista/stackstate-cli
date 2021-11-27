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
	SetUseColor(useColor bool)
	GetUseColor() bool
}

type StdPrinter struct {
	useColor bool
	spinner  *pterm.SpinnerPrinter
}

func NewStdPrinter() Printer {
	spinner := pterm.DefaultSpinner.WithRemoveWhenDone()
	return &StdPrinter{
		useColor: false, // use progressive enhancement
		spinner:  spinner,
	}
}

func (p *StdPrinter) PrintStruct(s interface{}) {
	p.StopSpinner()
	x, _ := json.Marshal(s)
	fmt.Fprintf(os.Stdout, "%s\n", string(x))
}

func (p *StdPrinter) PrintErr(err error) {
	p.StopSpinner()
	if p.useColor {
		fmt.Fprintf(os.Stderr, "%s\n", color.Red(err.Error()))
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
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
