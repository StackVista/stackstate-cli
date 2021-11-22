package asp

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type OptionalStringValue struct {
	cmd      *cobra.Command
	flag     string
	question string
	field    *string
}

func NewOptionalStringValue(cmd *cobra.Command, flag string, question string, field *string) InteractiveValue {
	return &OptionalStringValue{
		cmd:      cmd,
		flag:     flag,
		question: question,
		field:    field,
	}
}

func (s *OptionalStringValue) Set() error {
	f := s.cmd.Flag(s.flag)
	if f == nil {
		return fmt.Errorf("unknown flag %s", s.flag)
	}

	if !f.Changed {
		p := promptui.Prompt{
			Label: s.question,
		}

		result, err := p.Run()
		if err != nil {
			return err
		}

		*s.field = strings.TrimSpace(result)
	}

	return nil
}
