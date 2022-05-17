package pflags

import (
	"fmt"

	"github.com/spf13/pflag"
)

func GetFlagType(f *pflag.FlagSet, name string, ftype string, convFunc func(sval string) (interface{}, error)) (interface{}, error) {
	flag := f.Lookup(name)
	if flag == nil {
		err := fmt.Errorf("flag accessed but not defined: %s", name)
		return nil, err
	}

	if flag.Value.Type() != ftype {
		err := fmt.Errorf("trying to get %s value of flag of type %s", ftype, flag.Value.Type())
		return nil, err
	}

	sval := flag.Value.String()
	result, err := convFunc(sval)
	if err != nil {
		return nil, err
	}
	return result, nil
}
