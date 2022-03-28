package printer

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

func setupPrinter() (StdPrinter, *bytes.Buffer, *bytes.Buffer) {
	p := NewPrinter().(*StdPrinter)
	var stdOut, stdErr bytes.Buffer
	p.SetStdPrinterOutput(&stdOut, &stdErr)
	return *p, &stdOut, &stdErr
}

func TestPrintErr(t *testing.T) {
	p, _, stdErr := setupPrinter()
	assert.Equal(t, false, p.GetUseColor())
	assert.Nil(t, p.spinner)

	p.PrintErr(fmt.Errorf("test"))
	assert.Equal(t, "[ERROR] Test\n", stdErr.String())
}

func TestPrintErrWithColor(t *testing.T) {
	p, _, stdErr := setupPrinter()
	p.SetUseColor(true)
	assert.NotNil(t, p.spinner)

	p.PrintErr(fmt.Errorf("test"))
	assert.Equal(t, "❌ \x1b[31mTest\x1b[0m\n", stdErr.String())
}

func TestPrintStructAsJson(t *testing.T) {
	testStruct := map[string]interface{}{
		"foo": 1,
		"bar": map[string]interface{}{
			"baz": "foobarbaz",
			"in": map[string]interface{}{
				"out": 1,
			},
		},
	}
	const expectedJson = `{
  "bar": {
    "baz": "foobarbaz",
    "in": {
      "out": 1
    }
  },
  "foo": 1
}
`
	p := NewPrinter().(*StdPrinter)
	p.SetOutputType(JSON)
	testPrintStruct(t, p, testStruct, expectedJson)
}

func TestPrintStructAsYaml(t *testing.T) {
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
	p := NewPrinter().(*StdPrinter)
	testPrintStruct(t, p, testStruct, expectedYaml)
}

func TestPrintStructAsYamlWithColor(t *testing.T) {
	testStruct := map[string]interface{}{
		"foo": 1,
		"bar": map[string]interface{}{
			"baz": "foobarbaz",
		},
	}
	const expectedYaml = "\x1b[1m\x1b[31mbar\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m\x1b[1m\x1b[37m  \x1b[0m\x1b[1m\x1b[31mbaz\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m \x1b[0m\x1b[33mfoobarbaz\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m\x1b[1m\x1b[31mfoo\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m \x1b[0m\x1b[33m1\x1b[0m\n"
	p := NewPrinter().(*StdPrinter)
	p.SetUseColor(true)
	testPrintStruct(t, p, testStruct, expectedYaml)
}

func TestPrintStructAsJsonWithColor(t *testing.T) {
	testStruct := map[string]interface{}{
		"foo": 1,
		"bar": map[string]interface{}{
			"baz": "foobarbaz",
		},
	}
	const expectedJson = "\x1b[1m\x1b[37m{\x1b[0m\x1b[1m\x1b[37m\n  \x1b[0m\x1b[1m\x1b[31m\"bar\"\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m \x1b[0m\x1b[1m\x1b[37m{\x1b[0m\x1b[1m\x1b[37m\n    \x1b[0m\x1b[1m\x1b[31m\"baz\"\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m \x1b[0m\x1b[37m\"foobarbaz\"\x1b[0m\x1b[1m\x1b[37m\n  \x1b[0m\x1b[1m\x1b[37m}\x1b[0m\x1b[1m\x1b[37m,\x1b[0m\x1b[1m\x1b[37m\n  \x1b[0m\x1b[1m\x1b[31m\"foo\"\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m \x1b[0m\x1b[33m1\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m\x1b[1m\x1b[37m}\x1b[0m\n"
	p := NewPrinter().(*StdPrinter)
	p.SetOutputType(JSON)
	p.SetUseColor(true)
	testPrintStruct(t, p, testStruct, expectedJson)
}

func testPrintStruct(t *testing.T, p *StdPrinter, testStruct interface{}, expectedOutput string) {
	var buf bytes.Buffer
	p.stdOut = &buf
	p.PrintStruct(testStruct)
	assert.Equal(t, expectedOutput, buf.String())
}

func TestPrintErrResponse503WithColor(t *testing.T) {
	p, _, stdErr := setupPrinter()
	p.SetUseColor(true)

	resp := http.Response{Status: "503 Service Unavailable"}
	p.PrintErr(common.NewResponseError(fmt.Errorf(""), &resp))

	expected := "❌ \x1b[31m503 Service Unavailable\x1b[0m\n"
	assert.Equal(t, expected, stdErr.String())
}

func TestPrintErrResponse503(t *testing.T) {
	p := NewPrinter().(*StdPrinter)
	var buf bytes.Buffer
	p.stdErr = &buf
	resp := http.Response{Status: "503 Service Unavailable"}
	p.PrintErr(common.NewResponseError(fmt.Errorf(""), &resp))
	expected := fmt.Sprintf("[ERROR] %s\n", "503 Service Unavailable")
	assert.Equal(t, expected, buf.String())
}

func TestPrintErrResponse200(t *testing.T) {
	p := NewPrinter().(*StdPrinter)
	var buf bytes.Buffer
	p.stdErr = &buf
	resp := http.Response{Status: "200 Ok", StatusCode: 200}
	p.PrintErr(common.NewResponseError(fmt.Errorf("hello world"), &resp))
	expected := fmt.Sprintf("[ERROR] %s\n%s\n", "Client error", "Hello world")
	assert.Equal(t, expected, buf.String())
}

func TestPrintErrResponseNilResponse(t *testing.T) {
	p := NewPrinter().(*StdPrinter)
	var buf bytes.Buffer
	p.stdErr = &buf
	p.PrintErr(common.NewResponseError(fmt.Errorf("hello world"), nil))
	expected := fmt.Sprintf("[ERROR] %s\n%s\n", "Client error", "Hello world")
	assert.Equal(t, expected, buf.String())
}

func TestPrintTable(t *testing.T) {
	p, stdOut, _ := setupPrinter()
	p.Table([]string{"A", "B"}, [][]string{{"1", "2"}}, nil)
	assert.Equal(t, "A | B\n1 | 2\n", stdOut.String())
}

func TestPrintTableAsStruct(t *testing.T) {
	p, stdOut, _ := setupPrinter()
	p.SetOutputType(YAML)
	p.Table([]string{"A", "B"}, [][]string{{"1", "2"}}, map[string]interface{}{"A": 1, "B": 2})
	assert.Equal(t, "A: 1\nB: 2\n", stdOut.String())
}
