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

	var buf bytes.Buffer
	p.stdErr = &buf
	p.PrintErr(fmt.Errorf("test"))
	assert.Equal(t, "test\n", buf.String())
}

func TestStdPrinterUseColor(t *testing.T) {
	p := NewStdPrinter().(*StdPrinter)
	p.SetUseColor(true)

	var buf bytes.Buffer
	p.stdErr = &buf
	p.PrintErr(fmt.Errorf("test"))
	assert.Equal(t, color.Red("test").String()+"\n", buf.String())
}
