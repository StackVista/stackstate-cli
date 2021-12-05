package printer

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	color "github.com/logrusorgru/aurora/v3"
	"github.com/stretchr/testify/assert"
)

func TestPrintErr(t *testing.T) {
	p := NewPrinter().(*StdPrinter)
	assert.Equal(t, false, p.GetUseColor())
	assert.Nil(t, p.spinner)

	var buf bytes.Buffer
	p.stdErr = &buf
	p.PrintErr(fmt.Errorf("test"))
	assert.Equal(t, "Error: test\n", buf.String())
}

func TestPrintErrWithColor(t *testing.T) {
	p := NewPrinter().(*StdPrinter)
	p.SetUseColor(true)
	assert.NotNil(t, p.spinner)

	var buf bytes.Buffer
	p.stdErr = &buf
	p.PrintErr(fmt.Errorf("test"))
	assert.Equal(t, "❗"+color.Red("test").String()+"\n", buf.String())
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
	const expectedYaml = "\x1b[1m\x1b[31mbar\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m\x1b[1m\x1b[37m  \x1b[0m\x1b[1m\x1b[31mbaz\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m \x1b[0m\x1b[33mfoobarbaz\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m\x1b[1m\x1b[31mfoo\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m \x1b[0m\x1b[33m1\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m\n"
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

func TestPrintErrResponseWithoutResponseIsSameAsPrintErr(t *testing.T) {
	p := NewPrinter().(*StdPrinter)
	var buf bytes.Buffer
	p.stdErr = &buf
	p.PrintErrResponse(fmt.Errorf("test"), nil)
	assert.Equal(t, "Error: test\n", buf.String())
}

func TestPrintErrResponseWithoutResponseIsSameAsPrintErrWithColor(t *testing.T) {
	p := NewPrinter().(*StdPrinter)
	p.SetUseColor(true)
	var buf bytes.Buffer
	p.stdErr = &buf
	p.PrintErrResponse(fmt.Errorf("test"), nil)
	expected := fmt.Sprintf("❗%s\n", color.Red("test"))
	assert.Equal(t, expected, buf.String())
}

func TestPrintErrResponse503WithColor(t *testing.T) {
	p := NewPrinter().(*StdPrinter)
	p.SetUseColor(true)
	var buf bytes.Buffer
	p.stdErr = &buf
	resp := http.Response{Status: "503 Service Unavailable"}
	p.PrintErrResponse(fmt.Errorf(""), &resp)
	expected := fmt.Sprintf("❌ %s\n", color.Red("503 Service Unavailable"))
	assert.Equal(t, expected, buf.String())
}

func TestPrintErrResponse503(t *testing.T) {
	p := NewPrinter().(*StdPrinter)
	var buf bytes.Buffer
	p.stdErr = &buf
	resp := http.Response{Status: "503 Service Unavailable"}
	p.PrintErrResponse(fmt.Errorf(""), &resp)
	expected := fmt.Sprintf("Server error: %s\n", "503 Service Unavailable")
	assert.Equal(t, expected, buf.String())
}
