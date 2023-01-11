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

// RunStateValue the model 'RunStateValue'
type RunStateValue string

// List of RunStateValue
const (
	RUNSTATEVALUE_UNKNOWN   RunStateValue = "UNKNOWN"
	RUNSTATEVALUE_STARTING  RunStateValue = "STARTING"
	RUNSTATEVALUE_RUNNING   RunStateValue = "RUNNING"
	RUNSTATEVALUE_STOPPING  RunStateValue = "STOPPING"
	RUNSTATEVALUE_STOPPED   RunStateValue = "STOPPED"
	RUNSTATEVALUE_DEPLOYED  RunStateValue = "DEPLOYED"
	RUNSTATEVALUE_DEPLOYING RunStateValue = "DEPLOYING"
)

// All allowed values of RunStateValue enum
var AllowedRunStateValueEnumValues = []RunStateValue{
	"UNKNOWN",
	"STARTING",
	"RUNNING",
	"STOPPING",
	"STOPPED",
	"DEPLOYED",
	"DEPLOYING",
}

func (v *RunStateValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RunStateValue(value)
	for _, existing := range AllowedRunStateValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RunStateValue", value)
}

// NewRunStateValueFromValue returns a pointer to a valid RunStateValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRunStateValueFromValue(v string) (*RunStateValue, error) {
	ev := RunStateValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RunStateValue: valid values are %v", v, AllowedRunStateValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RunStateValue) IsValid() bool {
	for _, existing := range AllowedRunStateValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RunStateValue value
func (v RunStateValue) Ptr() *RunStateValue {
	return &v
}

type NullableRunStateValue struct {
	value *RunStateValue
	isSet bool
}

func (v NullableRunStateValue) Get() *RunStateValue {
	return v.value
}

func (v *NullableRunStateValue) Set(val *RunStateValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStateValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStateValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStateValue(val *RunStateValue) *NullableRunStateValue {
	return &NullableRunStateValue{value: val, isSet: true}
}

func (v NullableRunStateValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStateValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
