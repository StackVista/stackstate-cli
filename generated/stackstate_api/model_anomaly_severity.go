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

// AnomalySeverity the model 'AnomalySeverity'
type AnomalySeverity string

// List of AnomalySeverity
const (
	ANOMALYSEVERITY_LOW    AnomalySeverity = "LOW"
	ANOMALYSEVERITY_MEDIUM AnomalySeverity = "MEDIUM"
	ANOMALYSEVERITY_HIGH   AnomalySeverity = "HIGH"
)

// All allowed values of AnomalySeverity enum
var AllowedAnomalySeverityEnumValues = []AnomalySeverity{
	"LOW",
	"MEDIUM",
	"HIGH",
}

func (v *AnomalySeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AnomalySeverity(value)
	for _, existing := range AllowedAnomalySeverityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AnomalySeverity", value)
}

// NewAnomalySeverityFromValue returns a pointer to a valid AnomalySeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAnomalySeverityFromValue(v string) (*AnomalySeverity, error) {
	ev := AnomalySeverity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AnomalySeverity: valid values are %v", v, AllowedAnomalySeverityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AnomalySeverity) IsValid() bool {
	for _, existing := range AllowedAnomalySeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnomalySeverity value
func (v AnomalySeverity) Ptr() *AnomalySeverity {
	return &v
}

type NullableAnomalySeverity struct {
	value *AnomalySeverity
	isSet bool
}

func (v NullableAnomalySeverity) Get() *AnomalySeverity {
	return v.value
}

func (v *NullableAnomalySeverity) Set(val *AnomalySeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableAnomalySeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableAnomalySeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnomalySeverity(val *AnomalySeverity) *NullableAnomalySeverity {
	return &NullableAnomalySeverity{value: val, isSet: true}
}

func (v NullableAnomalySeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnomalySeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
