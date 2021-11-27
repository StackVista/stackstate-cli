package printer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alecthomas/chroma/quick"
	color "github.com/logrusorgru/aurora/v3"
	"github.com/pterm/pterm"
	msg "gitlab.com/stackvista/stackstate-cli2/internal/messages"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gopkg.in/yaml.v3"
)

type Printer interface {
	PrintStruct(s interface{})
	PrintErr(err error)
	PrintErrResponse(err error, resp *http.Response)
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

// kills any ongoing processes
func resetStdPrinter(p *StdPrinter) {
	p.StopSpinner()
}

func (p *StdPrinter) PrintStruct(s interface{}) {
	resetStdPrinter(p)
	msg, _ := json.MarshalIndent(s, "", " ")
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
		return // no change
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

func (p *StdPrinter) PrintErrResponse(err error, resp *http.Response) {
	resetStdPrinter(p)

	var status string
	if resp != nil && resp.Status != "" {
		status = resp.Status
	}

	suggestion := ""
	if resp != nil && resp.StatusCode == 401 {
		suggestion = "\nPlease check your configured API token."
	}

	switch v := err.(type) {
	case sts.GenericOpenAPIError:
		var bodyStr string
		if resp.Header.Get("Content-Type") == "application/json" {
			var bodyStruct interface{}
			json.Unmarshal(v.Body(), &bodyStruct)
			yaml, err := yaml.Marshal(bodyStruct)
			if err == nil && yaml != nil && bodyStruct != nil {
				bodyStr = string(yaml)
			}
		}
		if p.useColor {
			fmt.Fprintf(p.stdErr, "❌ %s\n", color.Red(status))
			if bodyStr != "" {
				fmt.Fprintf(p.stdErr, "\n----------------\n"+
					"Server response:\n"+
					"----------------\n")
				quick.Highlight(p.stdErr, bodyStr, "yaml", "terminal", "monokai")
			}
			if suggestion != "" {
				fmt.Fprintf(p.stdErr, "\nSuggestion: %s", suggestion)
			}
		} else {
			fmt.Fprintf(p.stdErr, "%s\n", status)
			if bodyStr != "" {
				fmt.Fprintf(p.stdErr, "\n----------------\n"+
					"Server response:\n"+
					"----------------\n"+
					"%s", bodyStr)
			}
			if suggestion != "" {
				fmt.Fprintf(p.stdErr, "\nSuggestion: %s", suggestion)
			}
		}
	default:
		if p.useColor {
			fmt.Fprintf(p.stdErr, "❌ %v%v%s", color.Red(status), v, suggestion)
		} else {
			fmt.Fprintf(p.stdErr, "%v%v", status, suggestion)
		}
	}
}
