package pflags

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

type timeValue struct {
	value  *time.Time
	format string
}

func (v *timeValue) Set(s string) error {
	s = strings.TrimSpace(s)
	parsed, err := time.Parse(v.format, s)
	if err != nil {
		return fmt.Errorf("invalid time format `%s` must be one of: `%s`", s, v.format)
	}
	*v.value = parsed
	return nil
}

func (v *timeValue) Type() string { return "time" }

func (v *timeValue) String() string { return v.value.Format(time.RFC3339Nano) }

func TimeVar(f *pflag.FlagSet, p *time.Time, name string, value time.Time, format, usage string) {
	*p = value
	f.Var(&timeValue{value: p, format: format}, name, usage)
	if value.IsZero() {
		f.Lookup(name).DefValue = ""
	}
}
