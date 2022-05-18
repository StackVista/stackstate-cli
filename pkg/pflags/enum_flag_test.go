package pflags

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

var enumTests = map[string]struct {
	input   string
	allowed []string
	output  string
	outErr  error
}{
	"correct input":                      {input: "a", allowed: []string{"a", "b", "c"}, output: "a", outErr: nil},
	"disallowed input":                   {input: "not-in-list", allowed: []string{"a", "b", "c"}, output: "", outErr: newInvalidEnumStringError("not-in-list", []string{"a", "b", "c"})},
	"empty disallowed if not in allowed": {input: "", allowed: []string{"a", "b", "c"}, output: "", outErr: newInvalidEnumStringError("", []string{"a", "b", "c"})},
	"empty allowed if in allowed":        {input: "", allowed: []string{"", "b", "c"}, output: "", outErr: nil},
}

func setupEnum(t *string, defValue string, allowed []string) *pflag.FlagSet {
	f := pflag.NewFlagSet("test", pflag.ContinueOnError)
	EnumVar(f, t, "enum", defValue, allowed, "Enum")
	return f
}

func TestEnum(t *testing.T) {
	devnull, _ := os.Open(os.DevNull)
	os.Stderr = devnull
	for name, tc := range enumTests {
		t.Run(name, func(t *testing.T) {
			var enumVar string
			f := setupEnum(&enumVar, "", tc.allowed)
			arg := fmt.Sprintf("--enum=%s", tc.input)
			err := f.Parse([]string{arg})

			if tc.outErr == nil {
				assert.NoError(t, err)
				enumResult, err := GetEnum(f, "enum")
				assert.NoError(t, err)
				assert.Equal(t, tc.output, enumResult)
			} else {
				// Cobra wraps our error message with a message containing the flag name and value. So check whether our error message is contained instead of equal.
				assert.Contains(t, err.Error(), tc.outErr.Error())
			}
		})
	}
}
