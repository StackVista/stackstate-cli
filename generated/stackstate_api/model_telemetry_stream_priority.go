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

// TelemetryStreamPriority the model 'TelemetryStreamPriority'
type TelemetryStreamPriority string

// List of TelemetryStreamPriority
const (
	TELEMETRYSTREAMPRIORITY_NONE TelemetryStreamPriority = "None"
	TELEMETRYSTREAMPRIORITY_HIGH TelemetryStreamPriority = "High"
	TELEMETRYSTREAMPRIORITY_MEDIUM TelemetryStreamPriority = "Medium"
	TELEMETRYSTREAMPRIORITY_LOW TelemetryStreamPriority = "Low"
)

// All allowed values of TelemetryStreamPriority enum
var AllowedTelemetryStreamPriorityEnumValues = []TelemetryStreamPriority{
	"None",
	"High",
	"Medium",
	"Low",
}

func (v *TelemetryStreamPriority) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TelemetryStreamPriority(value)
	for _, existing := range AllowedTelemetryStreamPriorityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TelemetryStreamPriority", value)
}

// NewTelemetryStreamPriorityFromValue returns a pointer to a valid TelemetryStreamPriority
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTelemetryStreamPriorityFromValue(v string) (*TelemetryStreamPriority, error) {
	ev := TelemetryStreamPriority(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TelemetryStreamPriority: valid values are %v", v, AllowedTelemetryStreamPriorityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TelemetryStreamPriority) IsValid() bool {
	for _, existing := range AllowedTelemetryStreamPriorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TelemetryStreamPriority value
func (v TelemetryStreamPriority) Ptr() *TelemetryStreamPriority {
	return &v
}

type NullableTelemetryStreamPriority struct {
	value *TelemetryStreamPriority
	isSet bool
}

func (v NullableTelemetryStreamPriority) Get() *TelemetryStreamPriority {
	return v.value
}

func (v *NullableTelemetryStreamPriority) Set(val *TelemetryStreamPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryStreamPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryStreamPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryStreamPriority(val *TelemetryStreamPriority) *NullableTelemetryStreamPriority {
	return &NullableTelemetryStreamPriority{value: val, isSet: true}
}

func (v NullableTelemetryStreamPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryStreamPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

