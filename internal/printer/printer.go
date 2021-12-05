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
	msg "gitlab.com/stackvista/stackstate-cli2/internal/messages"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gopkg.in/yaml.v3"
)

type Printer interface {
	// Expects s to be JSON or YAML serializable. You'll get an error otherwise.
	PrintStruct(s interface{}) error
	PrintErr(err error)
	PrintErrResponse(err error, resp *http.Response)
	StartSpinner(loadingMsg msg.LoadingMsg)
	StopSpinner()
	SetUseColor(useColor bool)
	GetUseColor() bool
	SetStructFormat(structFormat StructFormatType)
	GetStructFormat() StructFormatType
}

type StructFormatType int

const (
	YAML StructFormatType = iota
	JSON
	ServerErrorColorSymbol  = "❌"
	GenericErrorColorSymbol = "❗"
)

type StdPrinter struct {
	stdOut       io.Writer // for test purposes
	stdErr       io.Writer // for test purposes
	useColor     bool
	spinner      *pterm.SpinnerPrinter
	structFormat StructFormatType
}

func NewPrinter() Printer {
	return &StdPrinter{
		useColor:     false, // IMPORTANT: use progressive enhancement!
		spinner:      nil,
		stdOut:       os.Stdout,
		stdErr:       os.Stderr,
		structFormat: YAML,
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
	if p.structFormat == JSON {
		msg, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			return "", err
		}

		sructStr = string(msg)
	} else if p.structFormat == YAML {

		var buf bytes.Buffer
		yamlEncoder := yaml.NewEncoder(&buf)
		yamlEncoder.SetIndent(2)
		err := yamlEncoder.Encode(&s)
		if err != nil {
			return "", err
		}
		sructStr = buf.String()
	} else {
		return "", fmt.Errorf("unknown structFormat %v", p.structFormat)
	}

	if p.useColor {
		return colorizeStruct(sructStr, p.structFormat)
	} else {
		return sructStr, nil
	}
}

func colorizeStruct(structStr string, structFormat StructFormatType) (string, error) {
	var lexer chroma.Lexer
	switch structFormat {
	case JSON:
		lexer = lexers.Get("json")
	case YAML:
		lexer = lexers.Get("yaml")
	default:
		return "", fmt.Errorf("unknown structFormat %v", structFormat)
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
	if p.useColor {
		fmt.Fprintf(p.stdErr, "%s%s\n", GenericErrorColorSymbol, color.Red(err.Error()))
	} else {
		fmt.Fprintf(p.stdErr, "Error: %s\n", err.Error())
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

func (p *StdPrinter) PrintErrResponse(rtnErr error, resp *http.Response) {
	// if reponse is nil then treat this error like any any other CLI error
	// it could be a connection error, timeout, etc.
	if resp == nil {
		p.PrintErr(rtnErr)
		return
	}

	resetStdPrinter(p)
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

// add a new line to the string, if it exists
func withNewLine(s string) string {
	if s != "" {
		return s + "\n"
	} else {
		return s
	}
}

func (p *StdPrinter) SetStructFormat(structFormat StructFormatType) {
	p.structFormat = structFormat
}

func (p *StdPrinter) GetStructFormat() StructFormatType {
	return p.structFormat
}
