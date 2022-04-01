package util

import (
	"fmt"
	"time"
)

func TimeToString(val interface{}) string {
	t, ok := val.(time.Time)
	if !ok {
		return fmt.Sprintf("invalid time %v", val)
	}
	timeZone, _ := t.Zone()
	str := fmt.Sprintf("%s %s", t.Format("Mon Jan _2 15:04:05 2006"), timeZone)
	return str
}
