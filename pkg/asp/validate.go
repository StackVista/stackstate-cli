package asp

import (
	"fmt"
	"strings"
)

func Required(s string) error {
	if strings.TrimSpace(s) == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}
