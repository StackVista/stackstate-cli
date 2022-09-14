package monitor

import (
	"fmt"
)

func IdOrIdentifier(id int64, identifier string) string {
	if id == 0 {
		return identifier
	}
	return fmt.Sprintf("%d", id)
}
