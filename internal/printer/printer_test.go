package printer

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func setupPrinter() (*StdPrinter, *bytes.Buffer, *bytes.Buffer) {
	var stdOut, stdErr bytes.Buffer
	p := NewStdPrinter("linux", &stdOut, &stdErr)
	return p, &stdOut, &stdErr
}

func TestNewPrinter(t *testing.T) {
	p := NewPrinter().(*StdPrinter)
	assert.Equal(t, os.Stdout, p.stdOut)
	assert.Equal(t, os.Stderr, p.stdErr)
}

func TestPrintErrWithoutColor(t *testing.T) {
	p, _, stdErr := setupPrinter()
	p.SetUseColor(false)
	assert.Equal(t, false, p.GetUseColor())

	p.PrintErr(fmt.Errorf("test"))
	assert.Equal(t, "[ERROR] Test\n", stdErr.String())
}

func TestPrintErrWithColorIsDefault(t *testing.T) {
	p, _, stdErr := setupPrinter()
	assert.Equal(t, true, p.GetUseColor())

	p.PrintErr(fmt.Errorf("test"))
	assert.Equal(t, "\u274C Test\n", stdErr.String())
}

func TestPrintWithoutSymbolButWithColorOnWindows(t *testing.T) {
	var stdOut, stdErr bytes.Buffer
	p := NewStdPrinter("windows", &stdOut, &stdErr)
	p.PrintErr(fmt.Errorf("test"))
	assert.Equal(t, "\x1b[31m[ERROR]\x1b[0m Test\n", stdErr.String())
}

func TestPrintStructAsJsonWithoutColor(t *testing.T) {
	testStruct := map[string]interface{}{
		"foo": 1,
		"bar": map[string]interface{}{
			"baz": "foobarbaz",
			"in": map[string]interface{}{
				"out": 1,
			},
		},
	}
	const expectedJSON = `{
  "bar": {
    "baz": "foobarbaz",
    "in": {
      "out": 1
    }
  },
  "foo": 1
}
`
	p, _, _ := setupPrinter()
	p.SetUseColor(true) // should not matter

	var buf bytes.Buffer
	p.stdOut = &buf
	p.PrintJson(testStruct)
	assert.Equal(t, expectedJSON, buf.String())
}

func TestPrintErrJson(t *testing.T) {
	testStruct := fmt.Errorf("test error")
	const expectedJSON = `{
  "error": true,
  "error-message": "test error"
}
`
	p, _, _ := setupPrinter()
	p.SetUseColor(true) // should not matter

	var buf bytes.Buffer
	p.stdErr = &buf
	p.PrintErrJson(testStruct)
	assert.Equal(t, expectedJSON, buf.String())
}

func TestPrintStructAsYamlWithoutColor(t *testing.T) {
	testStruct := map[string]interface{}{
		"foo": 1,
		"bar": map[string]interface{}{
			"baz": "foobarbaz",
		},
	}
	const expectedYaml = `bar:
  baz: foobarbaz
foo: 1
`
	p, _, _ := setupPrinter()
	p.SetUseColor(false)
	testPrintStruct(t, p, testStruct, expectedYaml)
}

func TestPrintStructAsYamlWithColor(t *testing.T) {
	testStruct := map[string]interface{}{
		"foo": 1,
		"bar": map[string]interface{}{
			"baz": "foobarbaz",
		},
	}
	//nolint:lll
	const expectedYaml = "\x1b[1m\x1b[34mbar\x1b[0m:\x1b[37m\n\x1b[0m\x1b[37m  \x1b[0m\x1b[1m\x1b[34mbaz\x1b[0m:\x1b[37m \x1b[0mfoobarbaz\x1b[37m\n\x1b[0m\x1b[1m\x1b[34mfoo\x1b[0m:\x1b[37m \x1b[0m\x1b[36m1\x1b[0m\n"
	p, _, _ := setupPrinter()
	p.SetUseColor(true)
	testPrintStruct(t, p, testStruct, expectedYaml)
}

func testPrintStruct(t *testing.T, p *StdPrinter, testStruct interface{}, expectedOutput string) {
	var buf bytes.Buffer
	p.stdOut = &buf
	p.PrintStruct(testStruct)
	assert.Equal(t, expectedOutput, buf.String())
}

func TestPrintCLIErrorWith503WithColor(t *testing.T) {
	p, _, stdErr := setupPrinter()
	p.SetUseColor(true)

	resp := http.Response{
		Status: "503 Service Unavailable",
		Header: map[string][]string{"Content-Type": {"application/json"}},
		Body:   io.NopCloser(strings.NewReader(`{ "hello": "world" }`)),
	}
	p.PrintErr(common.NewResponseError(fmt.Errorf(""), &resp))

	expected := "\u274C 503 Service Unavailable\n\x1b[1m\x1b[31mhello\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m \x1b[0m\x1b[33mworld\x1b[0m\n"
	assert.Equal(t, expected, stdErr.String())
}

func TestPrintCLIError101WithoutColor(t *testing.T) {
	p, _, _ := setupPrinter()
	p.SetUseColor(false)
	var buf bytes.Buffer
	p.stdErr = &buf
	resp := http.Response{Status: "101 Switching Protocols", StatusCode: 101}
	p.PrintErr(common.NewResponseError(fmt.Errorf(""), &resp))
	expected := fmt.Sprintf("[ERROR] %s\n", "101 Switching Protocols")
	assert.Equal(t, expected, buf.String())
}

func TestPrintCLIErrorWith202Response(t *testing.T) {
	p, _, _ := setupPrinter()
	p.SetUseColor(false)
	var buf bytes.Buffer
	p.stdErr = &buf
	resp := http.Response{Status: "202 Accepted", StatusCode: 202}
	p.PrintErr(common.NewResponseError(fmt.Errorf("hello world"), &resp))
	expected := "[ERROR] Hello world\n"
	assert.Equal(t, expected, buf.String())
}

func TestPrintCLIErrorWithNilResponseWithColor(t *testing.T) {
	p, _, _ := setupPrinter()
	p.SetUseColor(true)
	var buf bytes.Buffer
	p.stdErr = &buf
	p.PrintErr(common.NewResponseError(fmt.Errorf("hello world"), nil))
	expected := "\u274C Hello world\n"
	assert.Equal(t, expected, buf.String())
}

func TestPrintTableWithoutColor(t *testing.T) {
	p, stdOut, _ := setupPrinter()
	p.SetUseColor(false)
	p.Table(TableData{
		Header: []string{"A", "B"},
		Data:   [][]interface{}{{"1", "2"}},
	})
	assert.Equal(t, "A | B\n1 | 2\n", stdOut.String())
}

func TestPrintTableWithColor(t *testing.T) {
	p, stdOut, _ := setupPrinter()
	p.SetUseColor(true)
	p.Table(TableData{
		Header: []string{"A", "B"},
		Data:   [][]interface{}{{"1", "2"}},
	})
	assert.Equal(t, "\x1b[36mA\x1b[0m\x1b[36m | \x1b[0m\x1b[36mB\x1b[0m\n1 | 2\n", stdOut.String())
}

func TestPrintTableWrapEqualColumnTreatment(t *testing.T) {
	p, stdOut, _ := setupPrinter()
	p.SetUseColor(false)
	p.MaxWidth = 25
	p.Table(TableData{
		Header: []string{"Hello", "Foo", "World"},
		Data: [][]interface{}{
			{"1", "2", "3"},
			{"hello darkness my old friend", "I've come to talk to you again", "Because a vision softly creeping"},
		},
	})

	expected := `HELLO  | FOO    | WORLD 
1      | 2      | 3     
hello  | I've c | Becaus
darkne | ome to | e a vi
ss my  |  talk  | sion s
old fr | to you | oftly 
iend   |  again | creepi
       |        | ng    
`
	assert.Equal(t, expected, stdOut.String())
}

func TestPrintTableWrapUnequalColumnTreatment(t *testing.T) {
	p, stdOut, _ := setupPrinter()
	p.SetUseColor(false)
	p.MaxWidth = 30
	p.Table(TableData{
		Header: []string{"Hello", "Foo", "World"},
		Data: [][]interface{}{
			{"hello", "darkness my old friend", "Because a vision softly creeping"},
		},
	})

	expected := `HELLO | FOO       | WORLD    
hello | darkness  | Because a
      | my old fr |  vision s
      | iend      | oftly cre
      |           | eping    
`
	assert.Equal(t, expected, stdOut.String())
}

func TestTableDoesNotRenderBiggerThanMaxWidth(t *testing.T) {
	p, stdOut, _ := setupPrinter()
	p.SetUseColor(false)

	data := []string{"hello", "fooba", "world"}
	header1 := []string{"Hello", "Foo", "World"}
	header2 := []string{"x", "y", "z"}
	header3 := []string{"Wooot", "Woot", "It's the sound Of da Police"}
	for i := 0; i < 100; i++ {
		d := [][]interface{}{util.StringSliceToInterfaceSlice(data)}
		if i%2 == 0 {
			p.MaxWidth = 31
			p.Table(TableData{Header: header1, Data: d})
		} else {
			if i%3 == 0 {
				p.MaxWidth = 32
				p.Table(TableData{Header: header2, Data: d})
			} else {
				p.MaxWidth = 33
				p.Table(TableData{Header: header3, Data: d})
			}
		}
		data[0] += "w"
		data[1] += "h "
		data[2] += "h h"
	}

	lines := strings.Split(stdOut.String(), "\n")
	maxLineWidth := 0
	for _, line := range lines {
		if maxLineWidth < len(line) {
			maxLineWidth = len(line)
		}
	}

	assert.LessOrEqual(t, maxLineWidth, 33)
}

func TestPrintTableNoData(t *testing.T) {
	p, stdOut, _ := setupPrinter()
	p.SetUseColor(false)
	p.Table(TableData{
		Header: []string{"A", "B"},
		Data:   [][]interface{}{},
	})
	assert.Equal(t, NoTableDataDefaultMsg+"\n", stdOut.String())
}

func TestPrintTableNoDataCustomMessage(t *testing.T) {
	p, stdOut, _ := setupPrinter()
	p.SetUseColor(false)
	p.Table(TableData{
		Header:              []string{"A", "B"},
		Data:                [][]interface{}{},
		MissingTableDataMsg: NotFoundMsg{Types: "cats"},
	})
	assert.Equal(t, "No cats found.\n", stdOut.String())
}
