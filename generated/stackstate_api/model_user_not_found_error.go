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

// UserNotFoundError struct for UserNotFoundError
type UserNotFoundError struct {
	Type string `json:"_type"`
	Name string `json:"name"`
}

// NewUserNotFoundError instantiates a new UserNotFoundError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserNotFoundError(type_ string, name string) *UserNotFoundError {
	this := UserNotFoundError{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewUserNotFoundErrorWithDefaults instantiates a new UserNotFoundError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserNotFoundErrorWithDefaults() *UserNotFoundError {
	this := UserNotFoundError{}
	return &this
}

// GetType returns the Type field value
func (o *UserNotFoundError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserNotFoundError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserNotFoundError) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *UserNotFoundError) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserNotFoundError) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserNotFoundError) SetName(v string) {
	o.Name = v
}

func (o UserNotFoundError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableUserNotFoundError struct {
	value *UserNotFoundError
	isSet bool
}

func (v NullableUserNotFoundError) Get() *UserNotFoundError {
	return v.value
}

func (v *NullableUserNotFoundError) Set(val *UserNotFoundError) {
	v.value = val
	v.isSet = true
}

func (v NullableUserNotFoundError) IsSet() bool {
	return v.isSet
}

func (v *NullableUserNotFoundError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserNotFoundError(val *UserNotFoundError) *NullableUserNotFoundError {
	return &NullableUserNotFoundError{value: val, isSet: true}
}

func (v NullableUserNotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserNotFoundError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
