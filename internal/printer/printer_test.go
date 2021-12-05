package printer

import (
	"bytes"
	"fmt"
	"testing"

	color "github.com/logrusorgru/aurora/v3"
	"github.com/stretchr/testify/assert"
)

func TestStdPrinterNoColorByDefault(t *testing.T) {
	p := NewStdPrinter().(*StdPrinter)
	assert.Equal(t, false, p.GetUseColor())
	assert.Nil(t, p.spinner)

	var buf bytes.Buffer
	p.stdErr = &buf
	p.PrintErr(fmt.Errorf("test"))
	assert.Equal(t, "test\n", buf.String())
}

func TestStdPrinterUseColor(t *testing.T) {
	p := NewStdPrinter().(*StdPrinter)
	p.SetUseColor(true)
	assert.NotNil(t, p.spinner)

	var buf bytes.Buffer
	p.stdErr = &buf
	p.PrintErr(fmt.Errorf("test"))
	assert.Equal(t, color.Red("test").String()+"\n", buf.String())
}

func TestPrintStructAsJsonNoColor(t *testing.T) {
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
	p := NewStdPrinter().(*StdPrinter)
	p.SetStructFormat(JSON)
	testPrintStruct(t, p, testStruct, expectedJson)
}

func TestPrintStructAsYamlNoColorIsDefault(t *testing.T) {
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
	p := NewStdPrinter().(*StdPrinter)
	testPrintStruct(t, p, testStruct, expectedYaml)
}

func TestPrintStructAsYamlWithColor(t *testing.T) {
	testStruct := map[string]interface{}{
		"foo": 1,
		"bar": map[string]interface{}{
			"baz": "foobarbaz",
		},
	}
	const expectedYaml = "\x1b[1m\x1b[31mbar\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m\x1b[1m\x1b[37m  \x1b[0m\x1b[1m\x1b[31mbaz\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m \x1b[0m\x1b[33mfoobarbaz\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m\x1b[1m\x1b[31mfoo\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m \x1b[0m\x1b[33m1\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m"
	p := NewStdPrinter().(*StdPrinter)
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
	const expectedJson = "\x1b[1m\x1b[37m{\x1b[0m\x1b[1m\x1b[37m\n  \x1b[0m\x1b[1m\x1b[31m\"bar\"\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m \x1b[0m\x1b[1m\x1b[37m{\x1b[0m\x1b[1m\x1b[37m\n    \x1b[0m\x1b[1m\x1b[31m\"baz\"\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m \x1b[0m\x1b[37m\"foobarbaz\"\x1b[0m\x1b[1m\x1b[37m\n  \x1b[0m\x1b[1m\x1b[37m}\x1b[0m\x1b[1m\x1b[37m,\x1b[0m\x1b[1m\x1b[37m\n  \x1b[0m\x1b[1m\x1b[31m\"foo\"\x1b[0m\x1b[1m\x1b[37m:\x1b[0m\x1b[1m\x1b[37m \x1b[0m\x1b[33m1\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m\x1b[1m\x1b[37m}\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m"
	p := NewStdPrinter().(*StdPrinter)
	p.SetStructFormat(JSON)
	p.SetUseColor(true)
	testPrintStruct(t, p, testStruct, expectedJson)
}

func testPrintStruct(t *testing.T, p *StdPrinter, testStruct interface{}, expectedOutput string) {
	var buf bytes.Buffer
	p.stdOut = &buf
	p.PrintStruct(testStruct)
	assert.Equal(t, expectedOutput, buf.String())
}
