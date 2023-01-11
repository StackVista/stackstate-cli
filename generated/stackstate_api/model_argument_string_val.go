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

// ArgumentStringVal struct for ArgumentStringVal
type ArgumentStringVal struct {
	Type                string `json:"_type"`
	Id                  *int64 `json:"id,omitempty"`
	LastUpdateTimestamp *int64 `json:"lastUpdateTimestamp,omitempty"`
	Parameter           int64  `json:"parameter"`
	Value               string `json:"value"`
}

// NewArgumentStringVal instantiates a new ArgumentStringVal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArgumentStringVal(type_ string, parameter int64, value string) *ArgumentStringVal {
	this := ArgumentStringVal{}
	this.Type = type_
	this.Parameter = parameter
	this.Value = value
	return &this
}

// NewArgumentStringValWithDefaults instantiates a new ArgumentStringVal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArgumentStringValWithDefaults() *ArgumentStringVal {
	this := ArgumentStringVal{}
	return &this
}

// GetType returns the Type field value
func (o *ArgumentStringVal) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArgumentStringVal) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArgumentStringVal) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArgumentStringVal) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgumentStringVal) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArgumentStringVal) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ArgumentStringVal) SetId(v int64) {
	o.Id = &v
}

// GetLastUpdateTimestamp returns the LastUpdateTimestamp field value if set, zero value otherwise.
func (o *ArgumentStringVal) GetLastUpdateTimestamp() int64 {
	if o == nil || o.LastUpdateTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdateTimestamp
}

// GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArgumentStringVal) GetLastUpdateTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdateTimestamp == nil {
		return nil, false
	}
	return o.LastUpdateTimestamp, true
}

// HasLastUpdateTimestamp returns a boolean if a field has been set.
func (o *ArgumentStringVal) HasLastUpdateTimestamp() bool {
	if o != nil && o.LastUpdateTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdateTimestamp gets a reference to the given int64 and assigns it to the LastUpdateTimestamp field.
func (o *ArgumentStringVal) SetLastUpdateTimestamp(v int64) {
	o.LastUpdateTimestamp = &v
}

// GetParameter returns the Parameter field value
func (o *ArgumentStringVal) GetParameter() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value
// and a boolean to check if the value has been set.
func (o *ArgumentStringVal) GetParameterOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameter, true
}

// SetParameter sets field value
func (o *ArgumentStringVal) SetParameter(v int64) {
	o.Parameter = v
}

// GetValue returns the Value field value
func (o *ArgumentStringVal) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ArgumentStringVal) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ArgumentStringVal) SetValue(v string) {
	o.Value = v
}

func (o ArgumentStringVal) MarshalJSON() ([]byte, error) {
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
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableArgumentStringVal struct {
	value *ArgumentStringVal
	isSet bool
}

func (v NullableArgumentStringVal) Get() *ArgumentStringVal {
	return v.value
}

func (v *NullableArgumentStringVal) Set(val *ArgumentStringVal) {
	v.value = val
	v.isSet = true
}

func (v NullableArgumentStringVal) IsSet() bool {
	return v.isSet
}

func (v *NullableArgumentStringVal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArgumentStringVal(val *ArgumentStringVal) *NullableArgumentStringVal {
	return &NullableArgumentStringVal{value: val, isSet: true}
}

func (v NullableArgumentStringVal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArgumentStringVal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
