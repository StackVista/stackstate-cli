package pflags

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/spf13/pflag"
)

const (
	Day  = 24 * time.Hour
	Week = 7 * Day
)

var (
	durationRegex     = regexp.MustCompile(`([\d|.]+)([a-z])`)
	durationZeroRegex = regexp.MustCompile(`^[-]?0[smhdw]?$`)
	durationMapping   = map[string]time.Duration{
		"s": time.Second,
		"m": time.Minute,
		"h": time.Hour,
		"d": Day,
		"w": Week,
	}
)

var _ pflag.DefaultZeroProvider = (*durationValue)(nil)

// -- time.Duration Value
type durationValue time.Duration

func newDurationValue(val time.Duration, p *time.Duration) *durationValue {
	*p = val
	return (*durationValue)(p)
}

func (d *durationValue) Set(s string) error {
	v, err := ParseDuration(s)
	*d = durationValue(v)
	return err
}

func (d *durationValue) Type() string {
	return "duration"
}

func (d *durationValue) String() string { return (*time.Duration)(d).String() }

// IsDefaultZero is true if the default is empty, or positive or negative zero
func (d *durationValue) IsDefaultZero(defValue string) bool {
	return defValue == "" || durationZeroRegex.MatchString(defValue)
}

func ParseDuration(strDuration string) (time.Duration, error) {
	match := durationRegex.FindStringSubmatch(strDuration)
	if match == nil {
		return 0, newInvalidDurationStringError()
	}

	unit := match[2]
	if _, ok := durationMapping[unit]; !ok {
		return 0, newUnknownDurationUnitError(unit)
	}
	unitLength := durationMapping[unit]

	value, parseFloatErr := strconv.ParseFloat(match[1], 64) //nolint:gomnd
	if parseFloatErr != nil {
		return 0, newInvalidDurationValueError(strDuration, parseFloatErr)
	}
	return time.Duration(float64(unitLength.Nanoseconds()) * value), nil
}

func newInvalidDurationStringError() error {
	return fmt.Errorf("invalid duration - expected a value and a unit (s, m, h, d or w), e.g. 2h or 3.5d")
}

func newUnknownDurationUnitError(unit string) error {
	return fmt.Errorf("unknown unit \"%s\" - expected s, m, h, d or w (second, minute, day, hour or week)", unit)
}

func newInvalidDurationValueError(strDuration string, parseFloatErr error) error {
	length := len(strDuration)
	return fmt.Errorf("invalid value %s - expected a floating point value (%s)", strDuration[0:length-1], parseFloatErr)
}

// Needed because ParseDuration returns a time.Duration and not an interface{}
func durationConv(f *pflag.Flag, sval string) (interface{}, error) {
	// We should be using time.ParseDuration here, as our string is now rendered as a complex duration, and no longer includes the 'd' or 'w' unit
	return time.ParseDuration(sval)
}

// GetDuration return the duration value of a flag with the given name
func GetDuration(f *pflag.FlagSet, name string) (time.Duration, error) {
	val, err := GetFlagType(f, name, "duration", durationConv)
	if err != nil {
		return 0, err
	}
	return val.(time.Duration), nil
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
func DurationVar(f *pflag.FlagSet, p *time.Duration, name string, value time.Duration, usage string) {
	f.VarP(newDurationValue(value, p), name, "", usage)
}

// DurationVarP is like DurationVar, but accepts a shorthand letter that can be used after a single dash.
func DurationVarP(f *pflag.FlagSet, p *time.Duration, name, shorthand string, value time.Duration, usage string) {
	f.VarP(newDurationValue(value, p), name, shorthand, usage)
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
func Duration(f *pflag.FlagSet, name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	DurationVarP(f, p, name, "", value, usage)
	return p
}

// DurationP is like Duration, but accepts a shorthand letter that can be used after a single dash.
func DurationP(f *pflag.FlagSet, name, shorthand string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	DurationVarP(f, p, name, shorthand, value, usage)
	return p
}
