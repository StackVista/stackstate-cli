package pflags

import (
	"testing"
	"time"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeVar(t *testing.T) {
	var expiration time.Time
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	TimeVar(flags, &expiration, "expiration", time.Time{}, time.DateOnly, "Expiration date")
	require.NoError(t, flags.Set("expiration", " 2024-02-29 "))
	want := time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, want, expiration)
	assert.Equal(t, "time", flags.Lookup("expiration").Value.Type())
	assert.Equal(t, "2024-02-29T00:00:00Z", flags.Lookup("expiration").Value.String())
	for _, input := range []string{"2023-02-29", "invalid", "2024-02-29T00:00:00Z"} {
		assert.ErrorContains(t, flags.Set("expiration", input), "invalid time format")
		assert.Equal(t, want, expiration)
	}
}

func TestCustomFlagZeroDefaults(t *testing.T) {
	var duration time.Duration
	var relative, expiration time.Time
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	DurationVarP(flags, &duration, "duration", "d", 0, "Duration")
	RelativeTimeVar(flags, &relative, "relative", "-0m", NewTestClock(0), "Relative time")
	TimeVar(flags, &expiration, "expiration", time.Time{}, time.DateOnly, "Expiration date")
	assert.NotContains(t, flags.FlagUsages(), "(default")
	require.NoError(t, flags.Parse([]string{"-d", "1d", "--relative=-1m", "--expiration=2024-02-29"}))
	assert.NotContains(t, flags.FlagUsages(), "(default")
}

func TestCustomFlagNonZeroDefaults(t *testing.T) {
	var duration time.Duration
	var relative, expiration time.Time
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	DurationVar(flags, &duration, "duration", time.Hour, "Duration")
	RelativeTimeVarP(flags, &relative, "relative", "r", "-1m", NewTestClock(0), "Relative time")
	TimeVar(flags, &expiration, "expiration", time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC), time.DateOnly, "Expiration date")
	assert.Contains(t, flags.FlagUsages(), "(default 1h0m0s)")
	assert.Contains(t, flags.FlagUsages(), "1969-12-31T23:59:00Z")
	assert.Contains(t, flags.FlagUsages(), "2024-02-29T00:00:00Z")
}
