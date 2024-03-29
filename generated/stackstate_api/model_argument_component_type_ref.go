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

// ArgumentComponentTypeRef struct for ArgumentComponentTypeRef
type ArgumentComponentTypeRef struct {
	Type                string `json:"_type"`
	Id                  *int64 `json:"id,omitempty"`
	LastUpdateTimestamp *int64 `json:"lastUpdateTimestamp,omitempty"`
	Parameter           int64  `json:"parameter"`
	ComponentType       int64  `json:"componentType"`
}

// NewArgumentComponentTypeRef instantiates a new ArgumentComponentTypeRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArgumentComponentTypeRef(type_ string, parameter int64, componentType int64) *ArgumentComponentTypeRef {
	this := ArgumentComponentTypeRef{}
	this.Type = type_
	this.Parameter = parameter
	this.ComponentType = componentType
	return &this
}

// NewArgumentComponentTypeRefWithDefaults instantiates a new ArgumentComponentTypeRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArgumentComponentTypeRefWithDefaults() *ArgumentComponentTypeRef {
	this := ArgumentComponentTypeRef{}
	return &this
}

// GetType returns the Type field value
func (o *ArgumentComponentTypeRef) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArgumentComponentTypeRef) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArgumentComponentTypeRef) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArgumentComponentTypeRef) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgumentComponentTypeRef) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArgumentComponentTypeRef) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ArgumentComponentTypeRef) SetId(v int64) {
	o.Id = &v
}

// GetLastUpdateTimestamp returns the LastUpdateTimestamp field value if set, zero value otherwise.
func (o *ArgumentComponentTypeRef) GetLastUpdateTimestamp() int64 {
	if o == nil || o.LastUpdateTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdateTimestamp
}

// GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgumentComponentTypeRef) GetLastUpdateTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdateTimestamp == nil {
		return nil, false
	}
	return o.LastUpdateTimestamp, true
}

// HasLastUpdateTimestamp returns a boolean if a field has been set.
func (o *ArgumentComponentTypeRef) HasLastUpdateTimestamp() bool {
	if o != nil && o.LastUpdateTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdateTimestamp gets a reference to the given int64 and assigns it to the LastUpdateTimestamp field.
func (o *ArgumentComponentTypeRef) SetLastUpdateTimestamp(v int64) {
	o.LastUpdateTimestamp = &v
}

// GetParameter returns the Parameter field value
func (o *ArgumentComponentTypeRef) GetParameter() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value
// and a boolean to check if the value has been set.
func (o *ArgumentComponentTypeRef) GetParameterOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameter, true
}

// SetParameter sets field value
func (o *ArgumentComponentTypeRef) SetParameter(v int64) {
	o.Parameter = v
}

// GetComponentType returns the ComponentType field value
func (o *ArgumentComponentTypeRef) GetComponentType() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value
// and a boolean to check if the value has been set.
func (o *ArgumentComponentTypeRef) GetComponentTypeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentType, true
}

// SetComponentType sets field value
func (o *ArgumentComponentTypeRef) SetComponentType(v int64) {
	o.ComponentType = v
}

func (o ArgumentComponentTypeRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdateTimestamp != nil {
		toSerialize["lastUpdateTimestamp"] = o.LastUpdateTimestamp
	}
	if true {
		toSerialize["parameter"] = o.Parameter
	}
	if true {
		toSerialize["componentType"] = o.ComponentType
	}
	return json.Marshal(toSerialize)
}

type NullableArgumentComponentTypeRef struct {
	value *ArgumentComponentTypeRef
	isSet bool
}

func (v NullableArgumentComponentTypeRef) Get() *ArgumentComponentTypeRef {
	return v.value
}

func (v *NullableArgumentComponentTypeRef) Set(val *ArgumentComponentTypeRef) {
	v.value = val
	v.isSet = true
}

func (v NullableArgumentComponentTypeRef) IsSet() bool {
	return v.isSet
}

func (v *NullableArgumentComponentTypeRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArgumentComponentTypeRef(val *ArgumentComponentTypeRef) *NullableArgumentComponentTypeRef {
	return &NullableArgumentComponentTypeRef{value: val, isSet: true}
}

func (v NullableArgumentComponentTypeRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArgumentComponentTypeRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
