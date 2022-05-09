/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
	"fmt"
)

// DownsamplingMethod the model 'DownsamplingMethod'
type DownsamplingMethod string

// List of DownsamplingMethod
const (
	MEAN DownsamplingMethod = "MEAN"
	PERCENTILE_25 DownsamplingMethod = "PERCENTILE_25"
	PERCENTILE_50 DownsamplingMethod = "PERCENTILE_50"
	PERCENTILE_75 DownsamplingMethod = "PERCENTILE_75"
	PERCENTILE_90 DownsamplingMethod = "PERCENTILE_90"
	PERCENTILE_95 DownsamplingMethod = "PERCENTILE_95"
	PERCENTILE_98 DownsamplingMethod = "PERCENTILE_98"
	PERCENTILE_99 DownsamplingMethod = "PERCENTILE_99"
	MAX DownsamplingMethod = "MAX"
	MIN DownsamplingMethod = "MIN"
	SUM DownsamplingMethod = "SUM"
	EVENT_COUNT DownsamplingMethod = "EVENT_COUNT"
	SUM_NO_ZEROS DownsamplingMethod = "SUM_NO_ZEROS"
	EVENT_COUNT_NO_ZEROS DownsamplingMethod = "EVENT_COUNT_NO_ZEROS"
)

// All allowed values of DownsamplingMethod enum
var AllowedDownsamplingMethodEnumValues = []DownsamplingMethod{
	"MEAN",
	"PERCENTILE_25",
	"PERCENTILE_50",
	"PERCENTILE_75",
	"PERCENTILE_90",
	"PERCENTILE_95",
	"PERCENTILE_98",
	"PERCENTILE_99",
	"MAX",
	"MIN",
	"SUM",
	"EVENT_COUNT",
	"SUM_NO_ZEROS",
	"EVENT_COUNT_NO_ZEROS",
}

func (v *DownsamplingMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DownsamplingMethod(value)
	for _, existing := range AllowedDownsamplingMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DownsamplingMethod", value)
}

// NewDownsamplingMethodFromValue returns a pointer to a valid DownsamplingMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDownsamplingMethodFromValue(v string) (*DownsamplingMethod, error) {
	ev := DownsamplingMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DownsamplingMethod: valid values are %v", v, AllowedDownsamplingMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DownsamplingMethod) IsValid() bool {
	for _, existing := range AllowedDownsamplingMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DownsamplingMethod value
func (v DownsamplingMethod) Ptr() *DownsamplingMethod {
	return &v
}

type NullableDownsamplingMethod struct {
	value *DownsamplingMethod
	isSet bool
}

func (v NullableDownsamplingMethod) Get() *DownsamplingMethod {
	return v.value
}

func (v *NullableDownsamplingMethod) Set(val *DownsamplingMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableDownsamplingMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableDownsamplingMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownsamplingMethod(val *DownsamplingMethod) *NullableDownsamplingMethod {
	return &NullableDownsamplingMethod{value: val, isSet: true}
}

func (v NullableDownsamplingMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownsamplingMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
