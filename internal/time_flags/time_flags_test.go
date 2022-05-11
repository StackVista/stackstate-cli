package time_flags

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestClock struct {
	now int64
}

func newTestClock() Clock {
	return &TestClock{
		now: 1652108645000,
	}
}

func (clock *TestClock) Now() time.Time {
	return time.UnixMilli(clock.now)
}

func TestTimeParsing2(t *testing.T) {
	tests := map[string]struct {
		input   string
		outTime int64
		outErr  error
	}{
		"iso date":              {input: "2022-05-09T15:04:05Z", outTime: int64(1652108645000), outErr: nil},
		"relative time":         {input: "-2d", outTime: int64(1652108645000 - 2*86400000), outErr: nil},
		"unix time":             {input: "1652108645000", outTime: int64(1652108645000), outErr: nil},
		"garbage":               {input: "garbage", outTime: -1, outErr: newTimeParseError("garbage")},
		"empty ":                {input: "", outTime: -1, outErr: newTimeParseError("")},
		"invalid relative time": {input: "-2j", outTime: -1, outErr: newInvalidRelativeTimeError("-2j", newUnknownDurationUnitError("j"))},
	}

	clock := newTestClock()
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			parsedTime, err := ParseTime(clock, tc.input)
			parsedTimeMillis := int64(-1)
			if parsedTime != nil {
				parsedTimeMillis = parsedTime.UnixMilli()
			}
			assert.Equal(t, tc.outTime, parsedTimeMillis)
			assert.Equal(t, tc.outErr, err)
		})
	}
}
