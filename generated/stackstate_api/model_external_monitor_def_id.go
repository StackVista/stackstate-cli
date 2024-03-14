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

// ExternalMonitorDefId struct for ExternalMonitorDefId
type ExternalMonitorDefId struct {
	Type string `json:"_type"`
	Id   int64  `json:"id"`
}

// NewExternalMonitorDefId instantiates a new ExternalMonitorDefId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalMonitorDefId(type_ string, id int64) *ExternalMonitorDefId {
	this := ExternalMonitorDefId{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewExternalMonitorDefIdWithDefaults instantiates a new ExternalMonitorDefId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalMonitorDefIdWithDefaults() *ExternalMonitorDefId {
	this := ExternalMonitorDefId{}
	return &this
}

// GetType returns the Type field value
func (o *ExternalMonitorDefId) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalMonitorDefId) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalMonitorDefId) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ExternalMonitorDefId) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalMonitorDefId) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalMonitorDefId) SetId(v int64) {
	o.Id = v
}

func (o ExternalMonitorDefId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableExternalMonitorDefId struct {
	value *ExternalMonitorDefId
	isSet bool
}

func (v NullableExternalMonitorDefId) Get() *ExternalMonitorDefId {
	return v.value
}

func (v *NullableExternalMonitorDefId) Set(val *ExternalMonitorDefId) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalMonitorDefId) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalMonitorDefId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalMonitorDefId(val *ExternalMonitorDefId) *NullableExternalMonitorDefId {
	return &NullableExternalMonitorDefId{value: val, isSet: true}
}

func (v NullableExternalMonitorDefId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalMonitorDefId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}