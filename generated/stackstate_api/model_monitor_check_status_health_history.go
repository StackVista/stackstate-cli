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

// MonitorCheckStatusHealthHistory struct for MonitorCheckStatusHealthHistory
type MonitorCheckStatusHealthHistory struct {
	CheckStatusId int64 `json:"checkStatusId"`
	StartTime     int32 `json:"startTime"`
	EndTime       int32 `json:"endTime"`
	// List of health state changes ordered from most recent to oldest.
	HealthStateChanges []MonitorCheckStatusHealthChange `json:"healthStateChanges"`
}

// NewMonitorCheckStatusHealthHistory instantiates a new MonitorCheckStatusHealthHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorCheckStatusHealthHistory(checkStatusId int64, startTime int32, endTime int32, healthStateChanges []MonitorCheckStatusHealthChange) *MonitorCheckStatusHealthHistory {
	this := MonitorCheckStatusHealthHistory{}
	this.CheckStatusId = checkStatusId
	this.StartTime = startTime
	this.EndTime = endTime
	this.HealthStateChanges = healthStateChanges
	return &this
}

// NewMonitorCheckStatusHealthHistoryWithDefaults instantiates a new MonitorCheckStatusHealthHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorCheckStatusHealthHistoryWithDefaults() *MonitorCheckStatusHealthHistory {
	this := MonitorCheckStatusHealthHistory{}
	return &this
}

// GetCheckStatusId returns the CheckStatusId field value
func (o *MonitorCheckStatusHealthHistory) GetCheckStatusId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CheckStatusId
}

// GetCheckStatusIdOk returns a tuple with the CheckStatusId field value
// and a boolean to check if the value has been set.
func (o *MonitorCheckStatusHealthHistory) GetCheckStatusIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckStatusId, true
}

// SetCheckStatusId sets field value
func (o *MonitorCheckStatusHealthHistory) SetCheckStatusId(v int64) {
	o.CheckStatusId = v
}

// GetStartTime returns the StartTime field value
func (o *MonitorCheckStatusHealthHistory) GetStartTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *MonitorCheckStatusHealthHistory) GetStartTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *MonitorCheckStatusHealthHistory) SetStartTime(v int32) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *MonitorCheckStatusHealthHistory) GetEndTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *MonitorCheckStatusHealthHistory) GetEndTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *MonitorCheckStatusHealthHistory) SetEndTime(v int32) {
	o.EndTime = v
}

// GetHealthStateChanges returns the HealthStateChanges field value
func (o *MonitorCheckStatusHealthHistory) GetHealthStateChanges() []MonitorCheckStatusHealthChange {
	if o == nil {
		var ret []MonitorCheckStatusHealthChange
		return ret
	}

	return o.HealthStateChanges
}

// GetHealthStateChangesOk returns a tuple with the HealthStateChanges field value
// and a boolean to check if the value has been set.
func (o *MonitorCheckStatusHealthHistory) GetHealthStateChangesOk() ([]MonitorCheckStatusHealthChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.HealthStateChanges, true
}

// SetHealthStateChanges sets field value
func (o *MonitorCheckStatusHealthHistory) SetHealthStateChanges(v []MonitorCheckStatusHealthChange) {
	o.HealthStateChanges = v
}

func (o MonitorCheckStatusHealthHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["checkStatusId"] = o.CheckStatusId
	}
	if true {
		toSerialize["startTime"] = o.StartTime
	}
	if true {
		toSerialize["endTime"] = o.EndTime
	}
	if true {
		toSerialize["healthStateChanges"] = o.HealthStateChanges
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorCheckStatusHealthHistory struct {
	value *MonitorCheckStatusHealthHistory
	isSet bool
}

func (v NullableMonitorCheckStatusHealthHistory) Get() *MonitorCheckStatusHealthHistory {
	return v.value
}

func (v *NullableMonitorCheckStatusHealthHistory) Set(val *MonitorCheckStatusHealthHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorCheckStatusHealthHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorCheckStatusHealthHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorCheckStatusHealthHistory(val *MonitorCheckStatusHealthHistory) *NullableMonitorCheckStatusHealthHistory {
	return &NullableMonitorCheckStatusHealthHistory{value: val, isSet: true}
}

func (v NullableMonitorCheckStatusHealthHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorCheckStatusHealthHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}