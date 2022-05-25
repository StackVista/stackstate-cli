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
)

// MonitorNotFoundError struct for MonitorNotFoundError
type MonitorNotFoundError struct {
	Type string `json:"_type"`
	MonitorId int64 `json:"monitorId"`
}

// NewMonitorNotFoundError instantiates a new MonitorNotFoundError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorNotFoundError(type_ string, monitorId int64) *MonitorNotFoundError {
	this := MonitorNotFoundError{}
	this.Type = type_
	this.MonitorId = monitorId
	return &this
}

// NewMonitorNotFoundErrorWithDefaults instantiates a new MonitorNotFoundError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorNotFoundErrorWithDefaults() *MonitorNotFoundError {
	this := MonitorNotFoundError{}
	return &this
}

// GetType returns the Type field value
func (o *MonitorNotFoundError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MonitorNotFoundError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MonitorNotFoundError) SetType(v string) {
	o.Type = v
}

// GetMonitorId returns the MonitorId field value
func (o *MonitorNotFoundError) GetMonitorId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value
// and a boolean to check if the value has been set.
func (o *MonitorNotFoundError) GetMonitorIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorId, true
}

// SetMonitorId sets field value
func (o *MonitorNotFoundError) SetMonitorId(v int64) {
	o.MonitorId = v
}

func (o MonitorNotFoundError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["monitorId"] = o.MonitorId
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorNotFoundError struct {
	value *MonitorNotFoundError
	isSet bool
}

func (v NullableMonitorNotFoundError) Get() *MonitorNotFoundError {
	return v.value
}

func (v *NullableMonitorNotFoundError) Set(val *MonitorNotFoundError) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorNotFoundError) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorNotFoundError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorNotFoundError(val *MonitorNotFoundError) *NullableMonitorNotFoundError {
	return &NullableMonitorNotFoundError{value: val, isSet: true}
}

func (v NullableMonitorNotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorNotFoundError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


