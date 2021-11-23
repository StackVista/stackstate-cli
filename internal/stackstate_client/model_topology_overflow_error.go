/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_client

import (
	"encoding/json"
)

// TopologyOverflowError struct for TopologyOverflowError
type TopologyOverflowError struct {
	MaxSize int32 `json:"maxSize"`
}

// NewTopologyOverflowError instantiates a new TopologyOverflowError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopologyOverflowError(maxSize int32) *TopologyOverflowError {
	this := TopologyOverflowError{}
	this.MaxSize = maxSize
	return &this
}

// NewTopologyOverflowErrorWithDefaults instantiates a new TopologyOverflowError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologyOverflowErrorWithDefaults() *TopologyOverflowError {
	this := TopologyOverflowError{}
	return &this
}

// GetMaxSize returns the MaxSize field value
func (o *TopologyOverflowError) GetMaxSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value
// and a boolean to check if the value has been set.
func (o *TopologyOverflowError) GetMaxSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxSize, true
}

// SetMaxSize sets field value
func (o *TopologyOverflowError) SetMaxSize(v int32) {
	o.MaxSize = v
}

func (o TopologyOverflowError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["maxSize"] = o.MaxSize
	}
	return json.Marshal(toSerialize)
}

type NullableTopologyOverflowError struct {
	value *TopologyOverflowError
	isSet bool
}

func (v NullableTopologyOverflowError) Get() *TopologyOverflowError {
	return v.value
}

func (v *NullableTopologyOverflowError) Set(val *TopologyOverflowError) {
	v.value = val
	v.isSet = true
}

func (v NullableTopologyOverflowError) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologyOverflowError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologyOverflowError(val *TopologyOverflowError) *NullableTopologyOverflowError {
	return &NullableTopologyOverflowError{value: val, isSet: true}
}

func (v NullableTopologyOverflowError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologyOverflowError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

