package printer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	color "github.com/logrusorgru/aurora/v3"
	"github.com/pterm/pterm"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gopkg.in/yaml.v3"
)

type Printer interface {
	// Expects s to be JSON or YAML serializable. You'll get an error otherwise.
	PrintStruct(s interface{}) error
	PrintErr(err error)
	StartSpinner(loadingMsg common.LoadingMsg)
	StopSpinner()
	SetUseColor(useColor bool)
	GetUseColor() bool
	SetOutputType(outputType OutputType)
	GetOutputType() OutputType
}

type OutputType int

const (
	YAML OutputType = iota
	JSON
	// Auto formatting, sometimes prints in table format, sometimes in YAML
	Auto
	ServerErrorColorSymbol  = "❌"
	GenericErrorColorSymbol = "❗"
)

type StdPrinter struct {
	stdOut     io.Writer // for test purposes
	stdErr     io.Writer // for test purposes
	useColor   bool
	spinner    *pterm.SpinnerPrinter
	outputType OutputType
}

func NewPrinter() Printer {
	return &StdPrinter{
		useColor:   false, // IMPORTANT: use progressive enhancement!
		spinner:    nil,
		stdOut:     os.Stdout,
		stdErr:     os.Stderr,
		outputType: Auto,
	}
}

// kills any ongoing processes
func resetStdPrinter(p *StdPrinter) {
	p.StopSpinner()
}

func (p *StdPrinter) PrintStruct(s interface{}) error {
	resetStdPrinter(p)

	msg, err := p.sprintStruct(s)
	if err != nil {
		return err
	}
	fmt.Fprintf(p.stdOut, "%s\n", msg)

	return nil
}

func (p *StdPrinter) sprintStruct(s interface{}) (string, error) {
	var sructStr string
	if p.outputType == JSON {
		msg, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			return "", err
		}

		sructStr = string(msg)
	} else if p.outputType == YAML || p.outputType == Auto {

		var buf bytes.Buffer
		yamlEncoder := yaml.NewEncoder(&buf)
		yamlEncoder.SetIndent(2)
		err := yamlEncoder.Encode(&s)
		if err != nil {
			return "", err
		}
		sructStr = buf.String()
	} else {
		return "", fmt.Errorf("unknown outputType %v", p.outputType)
	}

	if p.useColor {
		return colorizeStruct(sructStr, p.outputType)
	} else {
		return sructStr, nil
	}
}

func colorizeStruct(structStr string, outputType OutputType) (string, error) {
	var lexer chroma.Lexer
	switch outputType {
	case JSON:
		lexer = lexers.Get("json")
	case YAML:
		lexer = lexers.Get("yaml")
	case Auto:
		lexer = lexers.Get("yaml")
	default:
		return "", fmt.Errorf("unknown outputType %v", outputType)
	}
	style := styles.Get("monokai")
	formatter := formatters.Get("terminal")

	it, err := lexer.Tokenise(nil, structStr)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = formatter.Format(&buf, style, it)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (p *StdPrinter) PrintErr(err error) {
	resetStdPrinter(p)

	switch e := err.(type) {
	case common.ResponseError:
		p.printErrResponse(e.Err, e.Resp)
	default:
		if p.useColor {
			fmt.Fprintf(p.stdErr, "%s%s\n", GenericErrorColorSymbol, color.Red(err.Error()))
		} else {
			fmt.Fprintf(p.stdErr, "Error: %s\n", err.Error())
		}
	}
}

func (p *StdPrinter) printErrResponse(rtnErr error, resp *http.Response) {
	var httpStatus, errorStr, suggestion string

	// get HTTP status string
	if resp.Status != "" {
		httpStatus = resp.Status
	}

	// get error string
	switch v := rtnErr.(type) {
	case sts.GenericOpenAPIError:
		// did server repond with JSON? Then show that to the user.
		if resp.Header.Get("Content-Type") == "application/json" {
			var bodyStruct interface{}
			json.Unmarshal(v.Body(), &bodyStruct)
			body, err := p.sprintStruct(bodyStruct)
			if err != nil {
				errorStr = v.Error()
			} else {
				errorStr = body
			}
		}
	default:
		errorStr = v.Error()
	}

	// get suggestion
	if resp.StatusCode == 401 {
		suggestion = "Please check your configured API token."
	}

	// print
	if p.useColor {
		if suggestion != "" {
			suggestion = fmt.Sprintf("%s %s", color.Blue("ⓘ"), suggestion)
		}

		fmt.Fprintf(p.stdErr,
			"%s %v\n%s%s",
			ServerErrorColorSymbol,
			color.Red(httpStatus),
			withNewLine(errorStr),
			withNewLine(suggestion),
		)
	} else {
		fmt.Fprintf(p.stdErr, "Server error: %s\n%s%s", httpStatus, withNewLine(errorStr), withNewLine(suggestion))
	}
}

func (p *StdPrinter) StartSpinner(loadingMsg common.LoadingMsg) {
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

// add a new line to the string, if it exists
func withNewLine(s string) string {
	if s != "" {
		return s + "\n"
	} else {
		return s
	}
}

func (p *StdPrinter) SetOutputType(outputType OutputType) {
	p.outputType = outputType
}

func (p *StdPrinter) GetOutputType() OutputType {
	return p.outputType
}
