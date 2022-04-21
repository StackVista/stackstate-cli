package util

import "fmt"

func SafeIntPointerToString(number *int64) string {
	if number == nil {
		return "-"
	}
	return fmt.Sprintf("%d", *number)
}
