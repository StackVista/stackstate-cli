package util

import "strings"

// From https://stackoverflow.com/a/15323988/1860591
// Oef, Go is such a shitty language
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func StringInSliceIgnoreCase(a string, list []string) bool {
	for _, b := range list {
		if strings.ToLower(b) == strings.ToLower(a) {
			return true
		}
	}
	return false
}
