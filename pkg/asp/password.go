package asp

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type PasswordValue struct {
	cmd      *cobra.Command
	flag     string
	question string
	field    *string
}

func NewPasswordValue(cmd *cobra.Command, flag string, question string, field *string) InteractiveValue {
	return &PasswordValue{
		cmd:      cmd,
		flag:     flag,
		question: question,
		field:    field,
	}
}

func (p *PasswordValue) Set() error {
	f := p.cmd.Flag(p.flag)
	if f == nil {
		return fmt.Errorf("unknown flag %s", p.flag)
	}

	if f.Changed {
		return nil
	}

	prompt := promptui.Prompt{
		Label:    p.question,
		Mask:     '*',
		Validate: Required,
	}

	pwd, err := prompt.Run()
	if err != nil {
		return err
	}

	fmt.Print()

	repeat := promptui.Prompt{
		Label: "Please repeat the password for confirmation: ",
		Mask:  '*',
		Validate: func(s string) error {
			if err := Required(s); err != nil {
				return err
			}

			if pwd != s {
				return fmt.Errorf("passwords mismatch")
			}
			return nil
		},
	}

	_, err = repeat.Run()
	if err != nil {
		return err
	}

	*p.field = pwd

	return nil
}
