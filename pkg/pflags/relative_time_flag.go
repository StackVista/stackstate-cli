package pflags

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/pflag"
)

const (
	RelativeTimeFlagType = "relative-time"
)

type Clock interface {
	Now() time.Time
}

type FixedTimeClock struct {
	now time.Time
}

// NewTestClock is used for setting a specific point in time for testing purposes.
func NewTestClock(now int64) Clock {
	return &FixedTimeClock{
		now: time.UnixMilli(now),
	}
}

/**
This clock is useful when parsing multiple time flags. In that case you do not want
a clock that still ticks, because the two or more flags could both be in relative time
and we should not want a situation where they both flags are relative to a moving target.
*/
func NewFixedTimeClock() Clock {
	return &FixedTimeClock{
		now: time.Now(),
	}
}

func (clock *FixedTimeClock) Now() time.Time {
	return clock.now
}

// RelativeTimeValue adapts time.Time for use as a flag. It can be specified as a relative time or absolute time
type RelativeTimeValue struct {
	clock Clock
	*time.Time
}

var _ pflag.DefaultZeroProvider = (*RelativeTimeValue)(nil)

func newRelativeTimeValue(val string, p *time.Time, clock Clock) *RelativeTimeValue {
	v, err := ParseRelativeTime(clock, val)
	if err != nil {
		panic(err)
	}

	*p = v
	tv := &RelativeTimeValue{
		clock: clock,
		Time:  p,
	}
	return tv
}

func (d *RelativeTimeValue) IsDefaultZero(defValue string) bool {
	return defValue == "" || durationZeroRegex.MatchString(defValue)
}

func ParseRelativeTime(clock Clock, s string) (time.Time, error) {
	if len(s) > 0 && s[0] == '-' {
		d, err := ParseDuration(s[1:])
		if err != nil {
			return time.Time{}, newInvalidRelativeTimeError(s, err)
		}

		result := clock.Now().Add(-d)
		return result, nil
	}

	unixTime, err := strconv.ParseInt(s, 0, 64) //nolint:gomnd
	if err != nil {
		result, err := time.Parse(time.RFC3339, s)
		if err != nil {
			return time.Time{}, newTimeParseError(s)
		}

		return result, nil
	}

	return time.UnixMilli(unixTime), nil
}

func newTimeParseError(strTime string) error {
	return fmt.Errorf(
		"invalid time format \"%s\"."+
			" Expected an ISO8601 date/time string (e.g. 2016-09-13T13:12:04Z), milliseconds since the unix epoch (e.g. 1652210038107) or relative time (e.g. -168h)",
		strTime,
	)
}

func newInvalidRelativeTimeError(strTime string, durationErr error) error {
	return fmt.Errorf(
		"invalid relative time format \"%s\" - %s", strTime, durationErr,
	)
}

// Set time.Time value from string based on accepted formats.
func (d *RelativeTimeValue) Set(s string) error {
	v, err := ParseRelativeTime(d.clock, s)
	*d.Time = v
	return err
}

// Type name for time.Time flags.
func (d *RelativeTimeValue) Type() string {
	return RelativeTimeFlagType
}

func (d *RelativeTimeValue) String() string { return d.Time.Format(time.RFC3339Nano) }

// GetRelativeTime return the time value of a flag with the given name
func GetRelativeTime(f *pflag.FlagSet, name string) (time.Time, error) {
	val, err := GetFlagType(f, name, RelativeTimeFlagType, func(f *pflag.Flag, sval string) (interface{}, error) {
		clock := f.Value.(*RelativeTimeValue).clock
		return ParseRelativeTime(clock, sval)
	})

	if err != nil {
		return time.Time{}, err
	}

	return val.(time.Time), nil
}

// RelativeTimeVar defines a time.Time flag with specified name, default value, and usage string.
// The argument p points to a time.Time variable in which to store the value of the flag.
func RelativeTimeVar(f *pflag.FlagSet, p *time.Time, name string, value string, clock Clock, usage string) {
	f.Var(newRelativeTimeValue(value, p, clock), name, usage)
}

// RelativeTimeVarP is like TimeVar, but accepts a shorthand letter that can be used after a single dash.
func RelativeTimeVarP(f *pflag.FlagSet, p *time.Time, name, shorthand string, value string, clock Clock, usage string) {
	f.VarP(newRelativeTimeValue(value, p, clock), name, shorthand, usage)
}

// RelativeTime defines a time.Time flag with specified name, default value, and usage string.
// The return value is the address of a time.Time variable that stores the value of the flag.
func RelativeTime(f *pflag.FlagSet, name string, value string, clock Clock, usage string) *time.Time {
	p := new(time.Time)
	RelativeTimeVarP(f, p, name, "", value, clock, usage)
	return p
}

// RelativeTimeP is like Time, but accepts a shorthand letter that can be used after a single dash.
func RelativeTimeP(f *pflag.FlagSet, name, shorthand string, value string, clock Clock, usage string) *time.Time {
	p := new(time.Time)
	RelativeTimeVarP(f, p, name, shorthand, value, clock, usage)
	return p
}
