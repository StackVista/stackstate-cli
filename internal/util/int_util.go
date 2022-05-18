package util

import "fmt"

func SafeIntPointerToString(number *int64) string {
	if number == nil {
		return "-"
	}
	return fmt.Sprintf("%d", *number)
}

func Int32NilP(nr int32) *int32 {
	if nr == 0 {
		return nil
	}
	return &nr
}
