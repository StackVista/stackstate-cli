/*
StackState API

This API documentation page describes the StackState server API. The StackState UI and CLI use the StackState server API to configure and query StackState.  You can use the API for similar purposes.  Each request sent to the StackState server API must be authenticated using one of the available authentication methods.   *Note that the StackState receiver API, used to send topology, telemetry and traces to StackState, is not described on this page and requires a different authentication method.*  For more information on StackState, refer to the [StackState documentation](https://docs.stackstate.com).

API version: 5.2.0
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
	"fmt"
)

// HealthStateValue the model 'HealthStateValue'
type HealthStateValue string

// List of HealthStateValue
const (
	HEALTHSTATEVALUE_UNINITIALIZED HealthStateValue = "UNINITIALIZED"
	HEALTHSTATEVALUE_UNKNOWN       HealthStateValue = "UNKNOWN"
	HEALTHSTATEVALUE_CLEAR         HealthStateValue = "CLEAR"
	HEALTHSTATEVALUE_DISABLED      HealthStateValue = "DISABLED"
	HEALTHSTATEVALUE_DEVIATING     HealthStateValue = "DEVIATING"
	HEALTHSTATEVALUE_FLAPPING      HealthStateValue = "FLAPPING"
	HEALTHSTATEVALUE_CRITICAL      HealthStateValue = "CRITICAL"
)

// All allowed values of HealthStateValue enum
var AllowedHealthStateValueEnumValues = []HealthStateValue{
	"UNINITIALIZED",
	"UNKNOWN",
	"CLEAR",
	"DISABLED",
	"DEVIATING",
	"FLAPPING",
	"CRITICAL",
}

func (v *HealthStateValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HealthStateValue(value)
	for _, existing := range AllowedHealthStateValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HealthStateValue", value)
}

// NewHealthStateValueFromValue returns a pointer to a valid HealthStateValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHealthStateValueFromValue(v string) (*HealthStateValue, error) {
	ev := HealthStateValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HealthStateValue: valid values are %v", v, AllowedHealthStateValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HealthStateValue) IsValid() bool {
	for _, existing := range AllowedHealthStateValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HealthStateValue value
func (v HealthStateValue) Ptr() *HealthStateValue {
	return &v
}

type NullableHealthStateValue struct {
	value *HealthStateValue
	isSet bool
}

func (v NullableHealthStateValue) Get() *HealthStateValue {
	return v.value
}

func (v *NullableHealthStateValue) Set(val *HealthStateValue) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthStateValue) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthStateValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthStateValue(val *HealthStateValue) *NullableHealthStateValue {
	return &NullableHealthStateValue{value: val, isSet: true}
}

func (v NullableHealthStateValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthStateValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
