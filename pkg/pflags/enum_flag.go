package pflags

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

const (
	EnumFlagType = "enum"
)

type enumValue struct {
	value         *string
	allowedValues []string
}

func newEnumValue(val string, p *string, allowed []string) *enumValue {
	*p = val

	return &enumValue{
		value:         p,
		allowedValues: allowed,
	}
}

func IsAllowed(val string, allowed []string) bool {
	for _, v := range allowed {
		if v == val {
			return true
		}
	}

	return false
}

func (e *enumValue) Set(val string) error {
	if !IsAllowed(val, e.allowedValues) {
		return newInvalidEnumStringError(val, e.allowedValues)
	}

	*e.value = val

	return nil
}

// GetEnum return the string value of a flag with the given name
func GetEnum(f *pflag.FlagSet, name string) (string, error) {
	val, err := GetFlagType(f, name, EnumFlagType, func(f *pflag.Flag, sval string) (interface{}, error) {
		allowed := f.Value.(*enumValue).allowedValues
		if !IsAllowed(sval, allowed) {
			return nil, newInvalidEnumStringError(sval, allowed)
		}
		return sval, nil
	})

	if err != nil {
		return "", err
	}

	return val.(string), nil
}

func newInvalidEnumStringError(val string, allowedValues []string) error {
	return fmt.Errorf(
		"invalid enum value \"%s\", allowed values are [%s]", val, strings.Join(allowedValues, ", "))
}

func (e *enumValue) Type() string {
	return EnumFlagType
}

func (e *enumValue) String() string {
	return *e.value
}

// EnumVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func EnumVar(f *pflag.FlagSet, p *string, name string, value string, allowed []string, usage string) {
	f.Var(newEnumValue(value, p, allowed), name, usage)
}

// EnumVarP is like EnumVar, but accepts a shorthand letter that can be used after a single dash.
func EnumVarP(f *pflag.FlagSet, p *string, name, shorthand string, value string, allowed []string, usage string) {
	f.VarP(newEnumValue(value, p, allowed), name, shorthand, usage)
}

// Enum defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func EnumP(f *pflag.FlagSet, name, shorthand string, value string, allowed []string, usage string) *string {
	p := new(string)
	EnumVarP(f, p, name, shorthand, value, allowed, usage)
	return p
}

// Enum defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func Enum(f *pflag.FlagSet, name string, value string, allowed []string, usage string) *string {
	p := new(string)
	EnumVar(f, p, name, value, allowed, usage)
	return p
}
