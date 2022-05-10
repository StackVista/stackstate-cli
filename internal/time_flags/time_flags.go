package time_flags

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

type Clock interface {
	Now() time.Time
}

type FixedTimeClock struct {
	now time.Time
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

func GetTimeFlag(cmd *cobra.Command, flagName string, clock Clock) (*time.Time, common.CLIError) {
	strTime, err := cmd.Flags().GetString(flagName)
	if err != nil {
		return nil, common.NewCLIArgParseError(err)
	}
	time, parseErr := ParseTime(clock, strTime)
	if parseErr != nil {
		return nil, common.NewCLIArgParseError(parseErr)
	}
	return time, nil
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

func ParseTime(clock Clock, strTime string) (*time.Time, error) {
	var result time.Time

	if len(strTime) > 0 && strTime[0] == '-' {
		duration, err := ParseDuration(strTime[1:])
		if err != nil {
			return nil, newInvalidRelativeTimeError(strTime, err)
		}
		result := clock.Now().Add(-duration)
		return &result, nil
	}

	unixTime, err := strconv.ParseInt(strTime, 0, 64) //nolint:gomnd
	if err != nil {
		result, err = time.Parse(time.RFC3339, strTime)
		if err != nil {
			return nil, newTimeParseError(strTime)
		}
	} else {
		result = time.UnixMilli(unixTime)
	}
	return &result, nil
}
