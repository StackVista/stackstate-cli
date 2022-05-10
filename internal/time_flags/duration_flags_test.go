package time_flags

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDurationParses(t *testing.T) {
	tests := map[string]struct {
		input       string
		outDuration time.Duration
		outErr      error
	}{
		"1 second":              {input: "1s", outDuration: time.Second, outErr: nil},
		"15 second":             {input: "15s", outDuration: 15 * time.Second, outErr: nil},
		"10 and a half minutes": {input: "10.5m", outDuration: 10*time.Minute + (30 * time.Second), outErr: nil},
		"2 hours":               {input: "2h", outDuration: 2 * time.Hour, outErr: nil},
		"2 day":                 {input: "2d", outDuration: 48 * time.Hour, outErr: nil},
		"1 week":                {input: "1w", outDuration: 24 * time.Hour * 7, outErr: nil},
		"empty string":          {input: "", outDuration: 0, outErr: newInvalidDurationStringError()},
		"garbage":               {input: "woot", outDuration: 0, outErr: newInvalidDurationStringError()},
		"missing unit single":   {input: "1", outDuration: 0, outErr: newInvalidDurationStringError()},
		"missing unit multi":    {input: "122", outDuration: 0, outErr: newInvalidDurationStringError()},
		"wrong unit":            {input: "1j", outDuration: 0, outErr: newUnknownDurationUnitError("j")},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			d, err := ParseDuration(tc.input)
			assert.Equal(t, tc.outDuration, d)
			assert.Equal(t, tc.outErr, err)
		})
	}
}
