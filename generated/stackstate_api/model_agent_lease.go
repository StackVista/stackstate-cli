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

// AgentLease the model 'AgentLease'
type AgentLease string

// List of AgentLease
const (
	AGENTLEASE_ACTIVE  AgentLease = "Active"
	AGENTLEASE_LIMITED AgentLease = "Limited"
	AGENTLEASE_STALE   AgentLease = "Stale"
)

// All allowed values of AgentLease enum
var AllowedAgentLeaseEnumValues = []AgentLease{
	"Active",
	"Limited",
	"Stale",
}

func (v *AgentLease) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AgentLease(value)
	for _, existing := range AllowedAgentLeaseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AgentLease", value)
}

// NewAgentLeaseFromValue returns a pointer to a valid AgentLease
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentLeaseFromValue(v string) (*AgentLease, error) {
	ev := AgentLease(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentLease: valid values are %v", v, AllowedAgentLeaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentLease) IsValid() bool {
	for _, existing := range AllowedAgentLeaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentLease value
func (v AgentLease) Ptr() *AgentLease {
	return &v
}

type NullableAgentLease struct {
	value *AgentLease
	isSet bool
}

func (v NullableAgentLease) Get() *AgentLease {
	return v.value
}

func (v *NullableAgentLease) Set(val *AgentLease) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentLease) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentLease(val *AgentLease) *NullableAgentLease {
	return &NullableAgentLease{value: val, isSet: true}
}

func (v NullableAgentLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}