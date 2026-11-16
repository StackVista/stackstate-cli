package pflags

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

var durationTests = map[string]struct {
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

func TestDurationParses(t *testing.T) {
	for name, tc := range durationTests {
		t.Run(name, func(t *testing.T) {
			d, err := ParseDuration(tc.input)
			assert.Equal(t, tc.outDuration, d)
			assert.Equal(t, tc.outErr, err)
		})
	}
}

func setupDuration(t *time.Duration, defValue time.Duration) *pflag.FlagSet {
	f := pflag.NewFlagSet("test", pflag.ContinueOnError)
	DurationVar(f, t, "duration", defValue, "Duration")
	return f
}

func TestDuration(t *testing.T) {
	devnull, _ := os.Open(os.DevNull)
	os.Stderr = devnull
	for name, tc := range durationTests {
		t.Run(name, func(t *testing.T) {
			var durationVar time.Duration
			f := setupDuration(&durationVar, 0)
			arg := fmt.Sprintf("--duration=%s", tc.input)
			err := f.Parse([]string{arg})

			if tc.outErr == nil {
				assert.NoError(t, err)
				durationResult, err := GetDuration(f, "duration")
				assert.NoError(t, err)
				assert.Equal(t, tc.outDuration, durationResult)
			} else {
				// Cobra wraps our error message with a message containing the flag name and value. So check whether our error message is contained instead of equal.
				assert.Contains(t, err.Error(), tc.outErr.Error())
			}
		})
	}
}
