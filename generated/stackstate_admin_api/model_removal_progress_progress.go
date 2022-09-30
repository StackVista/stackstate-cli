/*
StackState Admin API

StackState's Admin API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
)

// RemovalProgressProgress struct for RemovalProgressProgress
type RemovalProgressProgress struct {
	Type string `json:"_type"`
}

// NewRemovalProgressProgress instantiates a new RemovalProgressProgress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemovalProgressProgress(type_ string) *RemovalProgressProgress {
	this := RemovalProgressProgress{}
	this.Type = type_
	return &this
}

// NewRemovalProgressProgressWithDefaults instantiates a new RemovalProgressProgress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemovalProgressProgressWithDefaults() *RemovalProgressProgress {
	this := RemovalProgressProgress{}
	return &this
}

// GetType returns the Type field value
func (o *RemovalProgressProgress) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RemovalProgressProgress) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RemovalProgressProgress) SetType(v string) {
	o.Type = v
}

func (o RemovalProgressProgress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRemovalProgressProgress struct {
	value *RemovalProgressProgress
	isSet bool
}

func (v NullableRemovalProgressProgress) Get() *RemovalProgressProgress {
	return v.value
}

func (v *NullableRemovalProgressProgress) Set(val *RemovalProgressProgress) {
	v.value = val
	v.isSet = true
}

func (v NullableRemovalProgressProgress) IsSet() bool {
	return v.isSet
}

func (v *NullableRemovalProgressProgress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemovalProgressProgress(val *RemovalProgressProgress) *NullableRemovalProgressProgress {
	return &NullableRemovalProgressProgress{value: val, isSet: true}
}

func (v NullableRemovalProgressProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemovalProgressProgress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

