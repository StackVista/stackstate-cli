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

// TopologySyncStatus the model 'TopologySyncStatus'
type TopologySyncStatus string

// List of TopologySyncStatus
const (
	TOPOLOGYSYNCSTATUS_RUNNING       TopologySyncStatus = "Running"
	TOPOLOGYSYNCSTATUS_RESETTING     TopologySyncStatus = "Resetting"
	TOPOLOGYSYNCSTATUS_DELETING      TopologySyncStatus = "Deleting"
	TOPOLOGYSYNCSTATUS_DELETE_FAILED TopologySyncStatus = "DeleteFailed"
)

// All allowed values of TopologySyncStatus enum
var AllowedTopologySyncStatusEnumValues = []TopologySyncStatus{
	"Running",
	"Resetting",
	"Deleting",
	"DeleteFailed",
}

func (v *TopologySyncStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TopologySyncStatus(value)
	for _, existing := range AllowedTopologySyncStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TopologySyncStatus", value)
}

// NewTopologySyncStatusFromValue returns a pointer to a valid TopologySyncStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTopologySyncStatusFromValue(v string) (*TopologySyncStatus, error) {
	ev := TopologySyncStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TopologySyncStatus: valid values are %v", v, AllowedTopologySyncStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TopologySyncStatus) IsValid() bool {
	for _, existing := range AllowedTopologySyncStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologySyncStatus value
func (v TopologySyncStatus) Ptr() *TopologySyncStatus {
	return &v
}

type NullableTopologySyncStatus struct {
	value *TopologySyncStatus
	isSet bool
}

func (v NullableTopologySyncStatus) Get() *TopologySyncStatus {
	return v.value
}

func (v *NullableTopologySyncStatus) Set(val *TopologySyncStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTopologySyncStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologySyncStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologySyncStatus(val *TopologySyncStatus) *NullableTopologySyncStatus {
	return &NullableTopologySyncStatus{value: val, isSet: true}
}

func (v NullableTopologySyncStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologySyncStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
