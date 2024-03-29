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

// MonitorCheckStatusNotFoundError struct for MonitorCheckStatusNotFoundError
type MonitorCheckStatusNotFoundError struct {
	Id   int64  `json:"id"`
	Type string `json:"_type"`
}

// NewMonitorCheckStatusNotFoundError instantiates a new MonitorCheckStatusNotFoundError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorCheckStatusNotFoundError(id int64, type_ string) *MonitorCheckStatusNotFoundError {
	this := MonitorCheckStatusNotFoundError{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewMonitorCheckStatusNotFoundErrorWithDefaults instantiates a new MonitorCheckStatusNotFoundError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorCheckStatusNotFoundErrorWithDefaults() *MonitorCheckStatusNotFoundError {
	this := MonitorCheckStatusNotFoundError{}
	return &this
}

// GetId returns the Id field value
func (o *MonitorCheckStatusNotFoundError) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MonitorCheckStatusNotFoundError) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MonitorCheckStatusNotFoundError) SetId(v int64) {
	o.Id = v
}

// GetType returns the Type field value
func (o *MonitorCheckStatusNotFoundError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MonitorCheckStatusNotFoundError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MonitorCheckStatusNotFoundError) SetType(v string) {
	o.Type = v
}

func (o MonitorCheckStatusNotFoundError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["_type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorCheckStatusNotFoundError struct {
	value *MonitorCheckStatusNotFoundError
	isSet bool
}

func (v NullableMonitorCheckStatusNotFoundError) Get() *MonitorCheckStatusNotFoundError {
	return v.value
}

func (v *NullableMonitorCheckStatusNotFoundError) Set(val *MonitorCheckStatusNotFoundError) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorCheckStatusNotFoundError) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorCheckStatusNotFoundError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorCheckStatusNotFoundError(val *MonitorCheckStatusNotFoundError) *NullableMonitorCheckStatusNotFoundError {
	return &NullableMonitorCheckStatusNotFoundError{value: val, isSet: true}
}

func (v NullableMonitorCheckStatusNotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorCheckStatusNotFoundError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
