package util

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"
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

func ToString(x interface{}) string {
	switch v := x.(type) {
	case *string:
		return SafeStringPtrToString(v)
	case float64:
		i, err := safeConvertFloat64ToInt64(v)
		if err != nil {
			return fmt.Sprintf("%f", v)
		} else {
			return fmt.Sprintf("%d", i)
		}
	default:
		return fmt.Sprintf("%v", v)
	}
}

func safeConvertFloat64ToInt64(f float64) (int64, error) {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return 0, fmt.Errorf("NaN or Inf")
	}
	v := int64(f)
	if float64(v)-f != 0 {
		return 0, fmt.Errorf("lost precision during conversion")
	} else {
		return v, nil
	}
}

func SafeStringPtrToString(s *string) string {
	if s == nil {
		return ""
	} else {
		return *s
	}
}

func ToStringSlice(data [][]interface{}) [][]string {
	result := make([][]string, 0)
	for _, row := range data {
		columns := make([]string, 0)
		for _, v := range row {
			if reflect.ValueOf(v).Kind() == reflect.Int ||
				reflect.ValueOf(v).Kind() == reflect.Int16 ||
				reflect.ValueOf(v).Kind() == reflect.Int32 ||
				reflect.ValueOf(v).Kind() == reflect.Int64 {
				columns = append(columns, fmt.Sprintf("%d", reflect.ValueOf(v).Int()))
			} else if reflect.ValueOf(v).Kind() == reflect.Float32 || reflect.ValueOf(v).Kind() == reflect.Float64 {
				columns = append(columns, fmt.Sprintf("%f", reflect.ValueOf(v).Float()))
			} else if reflect.ValueOf(v).Kind() == reflect.String {
				columns = append(columns, reflect.ValueOf(v).String())
			} else if reflect.ValueOf(v).Type() == reflect.TypeOf(time.Time{}) {
				columns = append(columns, TimeToString(v))
			} else {
				columns = append(columns, fmt.Sprintf("%v", v))
			}
		}
		result = append(result, columns)
	}
	return result
}
