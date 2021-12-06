package util

import (
	"strings"
	"unicode"
)

// Find string in list of strings
// From https://stackoverflow.com/a/15323988/1860591
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Find string in list of strings, case insensitive
func StringInSliceIgnoreCase(a string, list []string) bool {
	for _, b := range list {
		if strings.ToLower(b) == strings.ToLower(a) {
			return true
		}
	}
	return false
}

// Capitalises the first character of a string
// from https://github.com/tommysolsen/capitalise/blob/master/first.go
func UcFirst(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}

// Add a new line to the string, if it exists
func WithNewLine(s string) string {
	if s != "" {
		return s + "\n"
	} else {
		return s
	}
}
