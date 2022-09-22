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

type IdArgs struct {
	ID         int64
	Identifier string
}
