package pflags

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

const (
	TimeZero = int64(-62135596800000)
)

var timeTests = map[string]struct {
	input   string
	outTime int64
	outErr  error
}{
	"iso date":              {input: "2022-05-09T15:04:05Z", outTime: int64(1652108645000), outErr: nil},
	"relative time":         {input: "-2d", outTime: int64(1652108645000 - 2*86400000), outErr: nil},
	"unix time":             {input: "1652108645000", outTime: int64(1652108645000), outErr: nil},
	"garbage":               {input: "garbage", outTime: TimeZero, outErr: newTimeParseError("garbage")},
	"empty ":                {input: "", outTime: TimeZero, outErr: newTimeParseError("")},
	"invalid relative time": {input: "-2j", outTime: TimeZero, outErr: newInvalidRelativeTimeError("-2j", newUnknownDurationUnitError("j"))},
}

func TestRelativeTimeParsing(t *testing.T) {
	clock := NewTestClock(1652108645000)
	for name, tc := range timeTests {
		t.Run(name, func(t *testing.T) {
			parsedTime, err := ParseRelativeTime(clock, tc.input)
			parsedTimeMillis := parsedTime.UnixMilli()
			assert.Equal(t, tc.outTime, parsedTimeMillis)
			assert.Equal(t, tc.outErr, err)
		})
	}
}

func setupRelativeTime(t *time.Time, defValue string) *pflag.FlagSet {
	f := pflag.NewFlagSet("test", pflag.ContinueOnError)
	RelativeTimeVar(f, t, "time", defValue, NewTestClock(1652108645000), "Time")
	return f
}

func TestRelativeTime(t *testing.T) {
	devnull, _ := os.Open(os.DevNull)
	os.Stderr = devnull
	for name, tc := range timeTests {
		t.Run(name, func(t *testing.T) {
			var timeVar time.Time
			f := setupRelativeTime(&timeVar, "0")
			arg := fmt.Sprintf("--time=%s", tc.input)
			err := f.Parse([]string{arg})

			if tc.outErr == nil {
				assert.NoError(t, err)
				timeResult, err := GetRelativeTime(f, "time")
				assert.NoError(t, err)
				assert.Equal(t, tc.outTime, timeResult.UnixMilli())
			} else {
				// Cobra wraps our error message with a message containing the flag name and value. So check whether our error message is contained instead of equal.
				assert.Contains(t, err.Error(), tc.outErr.Error())
			}
		})
	}
}
