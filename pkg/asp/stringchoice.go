package asp

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/pkg/util"
)

type ChoicesProvider func() ([]string, error)

type StringChoice struct {
	cmd             *cobra.Command
	flag            string
	question        string
	field           *string
	choices         []string
	choicesProvider ChoicesProvider
	defaultVal      *string
}

func NewStringChoice(cmd *cobra.Command, flag string, question string, field *string, choices []string) InteractiveValue {
	return &StringChoice{
		cmd:      cmd,
		flag:     flag,
		question: question,
		field:    field,
		choices:  choices,
	}
}

func NewStringChoiceP(cmd *cobra.Command, flag string, question string, field *string, choicesProvider ChoicesProvider) InteractiveValue {
	return &StringChoice{
		cmd:             cmd,
		flag:            flag,
		question:        question,
		field:           field,
		choicesProvider: choicesProvider,
	}
}

func NewStringChoiceDef(cmd *cobra.Command, flag string, question string, field *string, choices []string, defaulVal string) InteractiveValue {
	return &StringChoice{
		cmd:        cmd,
		flag:       flag,
		question:   question,
		field:      field,
		choices:    choices,
		defaultVal: &defaulVal,
	}
}

func NewStringChoiceDefP(cmd *cobra.Command, flag string, question string, field *string, choicesProvider ChoicesProvider, defaulVal string) InteractiveValue {
	return &StringChoice{
		cmd:             cmd,
		flag:            flag,
		question:        question,
		field:           field,
		choicesProvider: choicesProvider,
		defaultVal:      &defaulVal,
	}
}

func (c *StringChoice) Set() error {
	f := c.cmd.Flag(c.flag)
	if f == nil {
		return fmt.Errorf("unknown flag %s", c.flag)
	}

	if c.choicesProvider != nil {
		choices, err := c.choicesProvider()
		if err != nil {
			return err
		}

		c.choices = choices
	}

	if f.Changed && util.ContainsString(c.choices, *c.field) {
		return nil
	}

	idx := 0
	if c.defaultVal != nil {
		idx = util.StringSliceIndex(c.choices, *c.defaultVal)
	}

	p := promptui.Select{
		Label:     c.question,
		Items:     c.choices,
		CursorPos: idx,
	}

	_, v, err := p.Run()
	if err != nil {
		return err
	}

	*c.field = v

	return nil
}
