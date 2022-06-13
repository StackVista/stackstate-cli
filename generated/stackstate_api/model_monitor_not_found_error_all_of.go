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

// MonitorNotFoundErrorAllOf struct for MonitorNotFoundErrorAllOf
type MonitorNotFoundErrorAllOf struct {
	Type *string `json:"_type,omitempty"`
}

// NewMonitorNotFoundErrorAllOf instantiates a new MonitorNotFoundErrorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorNotFoundErrorAllOf() *MonitorNotFoundErrorAllOf {
	this := MonitorNotFoundErrorAllOf{}
	return &this
}

// NewMonitorNotFoundErrorAllOfWithDefaults instantiates a new MonitorNotFoundErrorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorNotFoundErrorAllOfWithDefaults() *MonitorNotFoundErrorAllOf {
	this := MonitorNotFoundErrorAllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MonitorNotFoundErrorAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorNotFoundErrorAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MonitorNotFoundErrorAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MonitorNotFoundErrorAllOf) SetType(v string) {
	o.Type = &v
}

func (o MonitorNotFoundErrorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["_type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorNotFoundErrorAllOf struct {
	value *MonitorNotFoundErrorAllOf
	isSet bool
}

func (v NullableMonitorNotFoundErrorAllOf) Get() *MonitorNotFoundErrorAllOf {
	return v.value
}

func (v *NullableMonitorNotFoundErrorAllOf) Set(val *MonitorNotFoundErrorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorNotFoundErrorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorNotFoundErrorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorNotFoundErrorAllOf(val *MonitorNotFoundErrorAllOf) *NullableMonitorNotFoundErrorAllOf {
	return &NullableMonitorNotFoundErrorAllOf{value: val, isSet: true}
}

func (v NullableMonitorNotFoundErrorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorNotFoundErrorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

