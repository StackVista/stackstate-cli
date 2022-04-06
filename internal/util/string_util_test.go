package util

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUcFirst(t *testing.T) {
	assert.Equal(t, UcFirst(""), "")
	assert.Equal(t, UcFirst("test"), "Test")
	assert.Equal(t, UcFirst("ést"), "Ést")
}

func TestToStringPrintNaN(t *testing.T) {
	assert.Equal(t, ToString(math.NaN()), "NaN")
}
func TestToStringPrintInf(t *testing.T) {
	assert.Equal(t, ToString(math.Inf(1)), "+Inf")
}

func TestToStringPrintLargeNum(t *testing.T) {
	assert.Equal(t, ToString(107915343054010.0), "107915343054010")
}