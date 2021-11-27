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
