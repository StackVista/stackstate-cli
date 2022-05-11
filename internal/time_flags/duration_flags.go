package time_flags

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

const (
	Day  = 24 * time.Hour
	Week = 7 * Day
)

var (
	durationRegex = regexp.MustCompile(`([\d|.]+)([a-z])`)
)

func GetDurationFlag(cmd *cobra.Command, flagName string) (time.Duration, common.CLIError) {
	strDuration, err := cmd.Flags().GetString(flagName)
	if err != nil {
		return 0, common.NewCLIArgParseError(err)
	}
	duration, parseErr := ParseDuration(strDuration)
	if parseErr != nil {
		return 0, common.NewCLIArgParseError(parseErr)
	}
	return duration, nil
}

func ParseDuration(strDuration string) (time.Duration, error) {
	match := durationRegex.FindStringSubmatch(strDuration)
	if match == nil {
		return 0, newInvalidDurationStringError()
	}

	unit := match[2]
	var unitLength time.Duration
	switch unit {
	case "s":
		unitLength = time.Second
	case "m":
		unitLength = time.Minute
	case "h":
		unitLength = time.Hour
	case "d":
		unitLength = Day
	case "w":
		unitLength = Week
	default:
		return 0, newUnknownDurationUnitError(unit)
	}

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
