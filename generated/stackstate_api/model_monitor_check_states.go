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
)

// MonitorCheckStates struct for MonitorCheckStates
type MonitorCheckStates struct {
	States []ViewCheckState `json:"states"`
}

// NewMonitorCheckStates instantiates a new MonitorCheckStates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorCheckStates(states []ViewCheckState) *MonitorCheckStates {
	this := MonitorCheckStates{}
	this.States = states
	return &this
}

// NewMonitorCheckStatesWithDefaults instantiates a new MonitorCheckStates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorCheckStatesWithDefaults() *MonitorCheckStates {
	this := MonitorCheckStates{}
	return &this
}

// GetStates returns the States field value
func (o *MonitorCheckStates) GetStates() []ViewCheckState {
	if o == nil {
		var ret []ViewCheckState
		return ret
	}

	return o.States
}

// GetStatesOk returns a tuple with the States field value
// and a boolean to check if the value has been set.
func (o *MonitorCheckStates) GetStatesOk() ([]ViewCheckState, bool) {
	if o == nil {
		return nil, false
	}
	return o.States, true
}

// SetStates sets field value
func (o *MonitorCheckStates) SetStates(v []ViewCheckState) {
	o.States = v
}

func (o MonitorCheckStates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["states"] = o.States
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorCheckStates struct {
	value *MonitorCheckStates
	isSet bool
}

func (v NullableMonitorCheckStates) Get() *MonitorCheckStates {
	return v.value
}

func (v *NullableMonitorCheckStates) Set(val *MonitorCheckStates) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorCheckStates) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorCheckStates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorCheckStates(val *MonitorCheckStates) *NullableMonitorCheckStates {
	return &NullableMonitorCheckStates{value: val, isSet: true}
}

func (v NullableMonitorCheckStates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorCheckStates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
