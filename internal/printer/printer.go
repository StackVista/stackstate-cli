package printer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gookit/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/pterm/pterm"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
	"gopkg.in/yaml.v3"
)

type Printer interface {
	// Expects s to be JSON or YAML serializable. Will print an error otherwise
	PrintStruct(s interface{})
	PrintErr(err error)
	StartSpinner(loadingMsg common.LoadingMsg) StopPrinterFn
	SetUseColor(useColor bool)
	GetUseColor() bool
	SetOutputType(outputType OutputType)
	GetOutputType() OutputType
	PrintWarn(msg string)
	Success(msg string)
	Table(header []string, data [][]interface{}, structData interface{})
	PrintLn(text string)
}

const (
	YamlIndent = 2
)

type OutputType int

type StopPrinterFn func()

const (
	YAML OutputType = iota
	JSON
	// Auto formatting, sometimes prints in table format, sometimes in YAML
	Auto
)

type StdPrinter struct {
	stdOut     io.Writer // for test purposes
	stdErr     io.Writer // for test purposes
	useColor   bool
	outputType OutputType
	MaxWidth   int
}

func NewPrinter() Printer {
	x := &StdPrinter{
		useColor:   true,
		stdOut:     os.Stdout,
		stdErr:     os.Stderr,
		outputType: Auto,
		MaxWidth:   pterm.DefaultParagraph.MaxWidth,
	}
	pterm.EnableColor()

	return x
}

func (p *StdPrinter) SetStdPrinterOutput(stdOut io.Writer, stdErr io.Writer) {
	color.SetOutput(stdOut)
	p.stdOut = stdOut
	p.stdErr = stdErr
}

func (p *StdPrinter) PrintStruct(s interface{}) {
	msg, err := p.sprintStruct(s)
	if err != nil {
		p.PrintErr(fmt.Errorf("error (%s) while printing struct: %v", err, s))
	} else {
		fmt.Fprintf(p.stdOut, "%s\n", msg)
	}
}

func (p *StdPrinter) sprintStruct(s interface{}) (string, error) {
	var sructStr string

	switch p.outputType {
	case JSON:
		msg, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			return "", err
		}

		sructStr = string(msg)
	case YAML, Auto:
		var buf bytes.Buffer
		yamlEncoder := yaml.NewEncoder(&buf)
		yamlEncoder.SetIndent(YamlIndent)
		err := yamlEncoder.Encode(&s)
		if err != nil {
			return "", err
		}
		sructStr = strings.TrimRight(buf.String(), "\n")
	default:
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
	switch e := err.(type) {
	case common.CLIError:
		p.printCLIError(e)
	default:
		color.Fprintf(p.stdErr, "%s %s\n", p.sprintSymbol("error"), color.Red.Render(util.UcFirst(err.Error())))
	}
}

func (p *StdPrinter) printCLIError(err common.CLIError) {
	var errorStr, bodyStr string

	resp := err.GetServerResponse()

	//nolint:gomnd
	isErrorResponse := resp != nil && ((resp.StatusCode-200 < 0) || (resp.StatusCode-200 >= 100))

	// get error string with HTTP status
	if isErrorResponse {
		if err.Error() != "" && err.Error() != resp.Status {
			errorStr = fmt.Sprintf("%s (%s)", resp.Status, err.Error())
		} else {
			errorStr = resp.Status
		}
	} else {
		errorStr = util.UcFirst(err.Error())
	}

	// get error string
	// did server repond with JSON? Then show that as an error string!
	if isErrorResponse && resp.Header.Get("Content-Type") == "application/json" {
		var bodyStruct interface{}
		bodyb, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			err := json.Unmarshal(bodyb, &bodyStruct)
			if err == nil {
				body, err := p.sprintStruct(bodyStruct)
				if err == nil {
					bodyStr = body
				}
			}
		}
	}

	color.Fprintf(p.stdErr,
		"%s %v\n%s",
		p.sprintSymbol("error"),
		color.Red.Render(errorStr),
		util.WithNewLine(bodyStr),
	)
}

func (p *StdPrinter) StartSpinner(loadingMsg common.LoadingMsg) StopPrinterFn {
	if p.useColor {
		s, _ := pterm.DefaultSpinner.WithRemoveWhenDone().WithRemoveWhenDone().WithShowTimer(false).WithText(loadingMsg.String()).Start()
		return func() {
			if s != nil {
				//nolint:golint,errcheck
				s.Stop()
			}
		}
	}
	return func() {}
}

func (p *StdPrinter) SetUseColor(useColor bool) {
	if p.useColor == useColor {
		return // no change
	}

	p.useColor = useColor
	if useColor {
		pterm.EnableColor()
	} else {
		pterm.DisableColor()
	}
}

func (p *StdPrinter) GetUseColor() bool {
	return p.useColor
}

func (p *StdPrinter) SetOutputType(outputType OutputType) {
	p.outputType = outputType
}

func (p *StdPrinter) GetOutputType() OutputType {
	return p.outputType
}

func (p *StdPrinter) Success(msg string) {
	color.Fprintf(p.stdOut, "%s %s\n", p.sprintSymbol("success"), util.UcFirst(msg))
}

func (p *StdPrinter) PrintWarn(msg string) {
	color.Fprintf(p.stdOut, "%s %s\n", p.sprintSymbol("warn"), util.UcFirst(msg))
}

func (p *StdPrinter) Table(header []string, data [][]interface{}, structData interface{}) {
	if p.outputType == Auto {
		tw := table.NewWriter()
		tw.Style().Options.DrawBorder = false
		tw.Style().Options.SeparateHeader = false
		tw.Style().Format.Header = text.FormatUpper
		tw.Style().Box.PaddingLeft = ""
		tw.Style().Box.PaddingRight = ""
		tw.Style().Box.MiddleVertical = " | "
		if p.useColor {
			tw.Style().Color.Header = text.Colors{text.FgCyan}
		}

		tw.AppendHeader(util.StringSliceToInterfaceSlice(header))

		rows := make([]table.Row, 0)
		for _, row := range data {
			columns := make(table.Row, 0)
			for _, v := range row {
				value := util.ToString(v)
				columns = append(columns, value)
			}
			rows = append(rows, columns)
		}

		adjustedColumnWidths := calcColumnWidth(header, data, p.MaxWidth, tw.Style().Box)
		columnConfigs := make([]table.ColumnConfig, len(header))
		for i, h := range header {
			columnConfigs = append(columnConfigs, table.ColumnConfig{
				Name:     h,
				WidthMax: adjustedColumnWidths[i],
			})
		}
		tw.SetColumnConfigs(columnConfigs)

		tw.AppendRows(rows)
		fmt.Fprintf(p.stdOut, "%s\n", tw.Render())
	} else {
		p.PrintStruct(structData)
	}
}

func calcColumnWidth(header []string, data [][]interface{}, maxWidth int, box table.BoxStyle) []int {
	// average column width should be the size of screen minus some table row overhead divided by the number of columns
	tableRowOverheadWidth := len(box.PaddingLeft) + len(box.PaddingRight) +
		((len(header) - 1) * (len(box.MiddleVertical) + len(box.PaddingLeft) + len(box.PaddingRight)))
	avgColumnWidth := (maxWidth - tableRowOverheadWidth) / len(header)

	// minimum column width is the length of the name of the header
	columnWidths := make([]int, len(header))
	for i, h := range header {
		columnWidths[i] = len(h)
	}

	// maximum column width is the length of the longest value
	for _, row := range data {
		for i, v := range row {
			value := util.ToString(v)
			if columnWidths[i] < len(value) {
				columnWidths[i] = len(value)
			}
		}
	}

	// every column smaller than the average column can be added to the bigger columns
	extraRoomFromSmallerColumns := 0
	biggerColumnCount := 0
	for _, cw := range columnWidths {
		if cw < avgColumnWidth {
			extraRoomFromSmallerColumns += (avgColumnWidth - cw)
		} else {
			biggerColumnCount++
		}
	}

	// any column:
	// - smaller than the average stays the same maximum with any bigger
	// - bigger than the average gets the average plus the extra room from smaller columns
	adjustedColumnWidths := make([]int, len(header))
	for i, cw := range columnWidths {
		if cw < avgColumnWidth {
			adjustedColumnWidths[i] = cw
		} else {
			adjustedColumnWidths[i] = avgColumnWidth + (extraRoomFromSmallerColumns / biggerColumnCount)
		}
	}
	return adjustedColumnWidths
}

func (p *StdPrinter) PrintLn(text string) {
	color.Fprintf(p.stdOut, "%s\n", text)
}
