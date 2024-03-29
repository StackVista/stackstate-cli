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

// BaseMonitorError struct for BaseMonitorError
type BaseMonitorError struct {
	MonitorId     *string `json:"monitorId,omitempty"`
	MonitorIdType *string `json:"monitorIdType,omitempty"`
}

// NewBaseMonitorError instantiates a new BaseMonitorError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseMonitorError() *BaseMonitorError {
	this := BaseMonitorError{}
	return &this
}

// NewBaseMonitorErrorWithDefaults instantiates a new BaseMonitorError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseMonitorErrorWithDefaults() *BaseMonitorError {
	this := BaseMonitorError{}
	return &this
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise.
func (o *BaseMonitorError) GetMonitorId() string {
	if o == nil || o.MonitorId == nil {
		var ret string
		return ret
	}
	return *o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseMonitorError) GetMonitorIdOk() (*string, bool) {
	if o == nil || o.MonitorId == nil {
		return nil, false
	}
	return o.MonitorId, true
}

// HasMonitorId returns a boolean if a field has been set.
func (o *BaseMonitorError) HasMonitorId() bool {
	if o != nil && o.MonitorId != nil {
		return true
	}

	return false
}

// SetMonitorId gets a reference to the given string and assigns it to the MonitorId field.
func (o *BaseMonitorError) SetMonitorId(v string) {
	o.MonitorId = &v
}

// GetMonitorIdType returns the MonitorIdType field value if set, zero value otherwise.
func (o *BaseMonitorError) GetMonitorIdType() string {
	if o == nil || o.MonitorIdType == nil {
		var ret string
		return ret
	}
	return *o.MonitorIdType
}

// GetMonitorIdTypeOk returns a tuple with the MonitorIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseMonitorError) GetMonitorIdTypeOk() (*string, bool) {
	if o == nil || o.MonitorIdType == nil {
		return nil, false
	}
	return o.MonitorIdType, true
}

// HasMonitorIdType returns a boolean if a field has been set.
func (o *BaseMonitorError) HasMonitorIdType() bool {
	if o != nil && o.MonitorIdType != nil {
		return true
	}

	return false
}

// SetMonitorIdType gets a reference to the given string and assigns it to the MonitorIdType field.
func (o *BaseMonitorError) SetMonitorIdType(v string) {
	o.MonitorIdType = &v
}

func (o BaseMonitorError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MonitorId != nil {
		toSerialize["monitorId"] = o.MonitorId
	}
	if o.MonitorIdType != nil {
		toSerialize["monitorIdType"] = o.MonitorIdType
	}
	return json.Marshal(toSerialize)
}

type NullableBaseMonitorError struct {
	value *BaseMonitorError
	isSet bool
}

func (v NullableBaseMonitorError) Get() *BaseMonitorError {
	return v.value
}

func (v *NullableBaseMonitorError) Set(val *BaseMonitorError) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseMonitorError) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseMonitorError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseMonitorError(val *BaseMonitorError) *NullableBaseMonitorError {
	return &NullableBaseMonitorError{value: val, isSet: true}
}

func (v NullableBaseMonitorError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseMonitorError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
