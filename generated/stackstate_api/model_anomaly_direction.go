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

// AnomalyDirection the model 'AnomalyDirection'
type AnomalyDirection string

// List of AnomalyDirection
const (
	ANOMALYDIRECTION_RISE AnomalyDirection = "RISE"
	ANOMALYDIRECTION_DROP AnomalyDirection = "DROP"
	ANOMALYDIRECTION_RUNNING AnomalyDirection = "RUNNING"
	ANOMALYDIRECTION_ANY AnomalyDirection = "ANY"
)

// All allowed values of AnomalyDirection enum
var AllowedAnomalyDirectionEnumValues = []AnomalyDirection{
	"RISE",
	"DROP",
	"RUNNING",
	"ANY",
}

func (v *AnomalyDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AnomalyDirection(value)
	for _, existing := range AllowedAnomalyDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AnomalyDirection", value)
}

// NewAnomalyDirectionFromValue returns a pointer to a valid AnomalyDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAnomalyDirectionFromValue(v string) (*AnomalyDirection, error) {
	ev := AnomalyDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AnomalyDirection: valid values are %v", v, AllowedAnomalyDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AnomalyDirection) IsValid() bool {
	for _, existing := range AllowedAnomalyDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnomalyDirection value
func (v AnomalyDirection) Ptr() *AnomalyDirection {
	return &v
}

type NullableAnomalyDirection struct {
	value *AnomalyDirection
	isSet bool
}

func (v NullableAnomalyDirection) Get() *AnomalyDirection {
	return v.value
}

func (v *NullableAnomalyDirection) Set(val *AnomalyDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableAnomalyDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableAnomalyDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnomalyDirection(val *AnomalyDirection) *NullableAnomalyDirection {
	return &NullableAnomalyDirection{value: val, isSet: true}
}

func (v NullableAnomalyDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnomalyDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
