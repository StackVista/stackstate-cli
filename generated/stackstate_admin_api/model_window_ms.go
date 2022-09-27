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

// WindowMs struct for WindowMs
type WindowMs struct {
	WindowMs *int64 `json:"windowMs,omitempty"`
}

// NewWindowMs instantiates a new WindowMs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWindowMs() *WindowMs {
	this := WindowMs{}
	return &this
}

// NewWindowMsWithDefaults instantiates a new WindowMs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWindowMsWithDefaults() *WindowMs {
	this := WindowMs{}
	return &this
}

// GetWindowMs returns the WindowMs field value if set, zero value otherwise.
func (o *WindowMs) GetWindowMs() int64 {
	if o == nil || o.WindowMs == nil {
		var ret int64
		return ret
	}
	return *o.WindowMs
}

// GetWindowMsOk returns a tuple with the WindowMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowMs) GetWindowMsOk() (*int64, bool) {
	if o == nil || o.WindowMs == nil {
		return nil, false
	}
	return o.WindowMs, true
}

// HasWindowMs returns a boolean if a field has been set.
func (o *WindowMs) HasWindowMs() bool {
	if o != nil && o.WindowMs != nil {
		return true
	}

	return false
}

// SetWindowMs gets a reference to the given int64 and assigns it to the WindowMs field.
func (o *WindowMs) SetWindowMs(v int64) {
	o.WindowMs = &v
}

func (o WindowMs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WindowMs != nil {
		toSerialize["windowMs"] = o.WindowMs
	}
	return json.Marshal(toSerialize)
}

type NullableWindowMs struct {
	value *WindowMs
	isSet bool
}

func (v NullableWindowMs) Get() *WindowMs {
	return v.value
}

func (v *NullableWindowMs) Set(val *WindowMs) {
	v.value = val
	v.isSet = true
}

func (v NullableWindowMs) IsSet() bool {
	return v.isSet
}

func (v *NullableWindowMs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindowMs(val *WindowMs) *NullableWindowMs {
	return &NullableWindowMs{value: val, isSet: true}
}

func (v NullableWindowMs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWindowMs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

