package util

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	Int64BitSize      = 64
	StringToInt64Base = 0
)

func DefaultIfEmpty(val, def string) string {
	if val == "" {
		return def
	}
	return val
}

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
		if strings.EqualFold(a, b) {
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

// Best effort conversion from any type to a end-user readable string. Not a developer readable string.
func ToString(x interface{}) string {
	r := reflect.ValueOf(x)
	k := r.Kind()

	// this is the only way to check whether `x == nill`, because nil is specific to the type in Golang :(
	isPointer := k == reflect.Chan || k == reflect.Func || k == reflect.Map || k == reflect.Ptr ||
		k == reflect.UnsafePointer || k == reflect.Slice || k == reflect.Interface
	if k == reflect.Invalid || (isPointer && r.IsNil()) {
		return "-"
	}

	// by this time we know that all null pointer are already handled
	// so we can simply dereference any pointer we have
	// if there is a more elegant way of dereferencing pointers you're welcome to do to
	switch p := x.(type) {
	case *int8:
		x = *p
	case *uint8:
		x = *p
	case *int16:
		x = *p
	case *uint16:
		x = *p
	case *int32:
		x = *p
	case *uint32:
		x = *p
	case *int64:
		x = *p
	case *uint64:
		x = *p
	case *float32:
		x = *p
	case *float64:
		x = *p
	case *time.Time:
		x = *p
	case *string:
		x = *p
	}

	switch v := x.(type) {
	case float64:
		if math.Floor(v)-v == 0.0 {
			// print large decimal numbers without scientific notation
			// this can be annoying because often ids in StackState are big numbers
			// and these are often deserialized as float64
			return fmt.Sprintf("%.0f", v)
		} else {
			// %g = necessary digits only
			return fmt.Sprintf("%g", v)
		}
	case time.Time:
		return v.Format("Mon Jan 2 15:04:05 2006 MST")
	default:
		return fmt.Sprintf("%v", v)
	}
}

func StringSliceToInterfaceSlice(a []string) []interface{} {
	var b = make([]interface{}, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}
	return b
}

func ToStringSlice(data [][]interface{}) [][]string {
	result := make([][]string, 0)
	for _, row := range data {
		columns := make([]string, 0)
		for _, v := range row {
			columns = append(columns, ToString(v))
		}
		result = append(result, columns)
	}
	return result
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, StringToInt64Base, Int64BitSize)
}

func UniqueStrings(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func StringNilP(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
