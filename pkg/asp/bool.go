package asp

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type BoolValue struct {
	cmd      *cobra.Command
	flag     string
	question string
	field    *bool
}

func NewBoolValue(cmd *cobra.Command, flag string, question string, field *bool) InteractiveValue {
	return &BoolValue{
		cmd:      cmd,
		flag:     flag,
		question: question,
		field:    field,
	}
}

func (b *BoolValue) Set() error {
	f := b.cmd.Flag(b.flag)
	if f == nil {
		return fmt.Errorf("unknown flag %s", b.flag)
	}

	if f.Changed {
		return nil
	}

	p := promptui.Prompt{
		Label: b.question + " [y/N]",
		Validate: func(s string) error {
			l := strings.ToLower(s)
			if l != "" && l != "y" && l != "n" {
				return fmt.Errorf("answer can be either 'y' or 'n'")
			}
			return nil
		},
	}

	v, err := p.Run()
	if err != nil {
		return err
	}

	trimmed := strings.ToLower(strings.TrimSpace(v))
	*b.field = trimmed == "y"

	return nil
}
