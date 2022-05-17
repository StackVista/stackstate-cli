package servicetoken

import "time"

type Args struct {
	ID         string
	Expiration time.Time
	Name       string
	Roles      []string
}
