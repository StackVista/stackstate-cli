package util

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUcFirst(t *testing.T) {
	assert.Equal(t, UcFirst(""), "")
	assert.Equal(t, UcFirst("test"), "Test")
	assert.Equal(t, UcFirst("ést"), "Ést")
}

func TestUniqueString(t *testing.T) {
	assert.Equal(t, []string{"foo", "bar"}, UniqueStrings([]string{"foo", "bar", "foo", "bar"}))
}

func TestToString(t *testing.T) {
	var f32nil *float32
	var mapnil map[string]string
	var strnil *string
	var slicenil []float64

	i8 := byte(7)
	ui8 := byte(7)
	i16 := int16(-15)
	ui16 := uint16(15)
	i32 := int32(-42)
	ui32 := uint32(42)
	i64 := int64(-42)
	ui64 := uint64(42)
	f32 := float32(3.14)
	f64 := float32(3.14)
	tme := time.UnixMilli(0)
	str := "hello world"

	tests := map[string]struct {
		input    interface{}
		expected string
	}{
		"int8":          {input: int8(-7), expected: "-7"},
		"pint8":         {input: &i8, expected: "7"},
		"ui8":           {input: uint8(7), expected: "7"},
		"pui8":          {input: &ui8, expected: "7"},
		"int16":         {input: int16(-15), expected: "-15"},
		"pint16":        {input: &i16, expected: "-15"},
		"uint16":        {input: uint16(15), expected: "15"},
		"puint16":       {input: &ui16, expected: "15"},
		"int32":         {input: int32(-42), expected: "-42"},
		"pint32":        {input: &i32, expected: "-42"},
		"uint32":        {input: uint32(42), expected: "42"},
		"puint32":       {input: &ui32, expected: "42"},
		"int64":         {input: int64(-42), expected: "-42"},
		"pint64":        {input: &i64, expected: "-42"},
		"uint64":        {input: uint64(42), expected: "42"},
		"puint64":       {input: &ui64, expected: "42"},
		"float32":       {input: float32(3.14), expected: "3.14"},
		"pfloat32":      {input: &f32, expected: "3.14"},
		"float32 round": {input: float32(3), expected: "3"},
		"float64":       {input: float64(3.14), expected: "3.14"},
		"pfloat64":      {input: &f64, expected: "3.14"},
		"float64 round": {input: float64(3), expected: "3"},
		"float64 round large no scientific notiation": {
			input: float64(107915343054010.0), expected: "107915343054010",
		},
		"positive infinity": {input: math.Inf(1), expected: "+Inf"},
		"negative infinity": {input: math.Inf(-1), expected: "-Inf"},
		"not a number":      {input: math.NaN(), expected: "NaN"},
		"nil":               {input: nil, expected: "-"},
		"float32 nil":       {input: f32nil, expected: "-"},
		"time":              {input: time.UnixMilli(0), expected: "Thu Jan 1 01:00:00 1970 CET"},
		"ptime":             {input: &tme, expected: "Thu Jan 1 01:00:00 1970 CET"},
		"map nil":           {input: mapnil, expected: "-"},
		"str nil":           {input: strnil, expected: "-"},
		"slice nil":         {input: slicenil, expected: "-"},
		"string":            {input: "hello world", expected: "hello world"},
		"pstring":           {input: &str, expected: "hello world"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			output := ToString(tc.input)
			assert.Equal(t, tc.expected, output)
		})
	}
}
