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

// PropagatedHealthStateValue the model 'PropagatedHealthStateValue'
type PropagatedHealthStateValue string

// List of PropagatedHealthStateValue
const (
	PROPAGATEDHEALTHSTATEVALUE_UNKNOWN PropagatedHealthStateValue = "UNKNOWN"
	PROPAGATEDHEALTHSTATEVALUE_PROPAGATION_ERROR PropagatedHealthStateValue = "PROPAGATION_ERROR"
	PROPAGATEDHEALTHSTATEVALUE_DEVIATING PropagatedHealthStateValue = "DEVIATING"
	PROPAGATEDHEALTHSTATEVALUE_FLAPPING PropagatedHealthStateValue = "FLAPPING"
	PROPAGATEDHEALTHSTATEVALUE_CRITICAL PropagatedHealthStateValue = "CRITICAL"
)

// All allowed values of PropagatedHealthStateValue enum
var AllowedPropagatedHealthStateValueEnumValues = []PropagatedHealthStateValue{
	"UNKNOWN",
	"PROPAGATION_ERROR",
	"DEVIATING",
	"FLAPPING",
	"CRITICAL",
}

func (v *PropagatedHealthStateValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PropagatedHealthStateValue(value)
	for _, existing := range AllowedPropagatedHealthStateValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PropagatedHealthStateValue", value)
}

// NewPropagatedHealthStateValueFromValue returns a pointer to a valid PropagatedHealthStateValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPropagatedHealthStateValueFromValue(v string) (*PropagatedHealthStateValue, error) {
	ev := PropagatedHealthStateValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PropagatedHealthStateValue: valid values are %v", v, AllowedPropagatedHealthStateValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PropagatedHealthStateValue) IsValid() bool {
	for _, existing := range AllowedPropagatedHealthStateValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PropagatedHealthStateValue value
func (v PropagatedHealthStateValue) Ptr() *PropagatedHealthStateValue {
	return &v
}

type NullablePropagatedHealthStateValue struct {
	value *PropagatedHealthStateValue
	isSet bool
}

func (v NullablePropagatedHealthStateValue) Get() *PropagatedHealthStateValue {
	return v.value
}

func (v *NullablePropagatedHealthStateValue) Set(val *PropagatedHealthStateValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePropagatedHealthStateValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePropagatedHealthStateValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropagatedHealthStateValue(val *PropagatedHealthStateValue) *NullablePropagatedHealthStateValue {
	return &NullablePropagatedHealthStateValue{value: val, isSet: true}
}

func (v NullablePropagatedHealthStateValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropagatedHealthStateValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
