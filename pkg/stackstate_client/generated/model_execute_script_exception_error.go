/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ExecuteScriptExceptionError struct for ExecuteScriptExceptionError
type ExecuteScriptExceptionError struct {
	Type string `json:"_type"`
	Message string `json:"message"`
}

// NewExecuteScriptExceptionError instantiates a new ExecuteScriptExceptionError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteScriptExceptionError(type_ string, message string) *ExecuteScriptExceptionError {
	this := ExecuteScriptExceptionError{}
	this.Type = type_
	this.Message = message
	return &this
}

// NewExecuteScriptExceptionErrorWithDefaults instantiates a new ExecuteScriptExceptionError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteScriptExceptionErrorWithDefaults() *ExecuteScriptExceptionError {
	this := ExecuteScriptExceptionError{}
	return &this
}

// GetType returns the Type field value
func (o *ExecuteScriptExceptionError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExecuteScriptExceptionError) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExecuteScriptExceptionError) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *ExecuteScriptExceptionError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ExecuteScriptExceptionError) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ExecuteScriptExceptionError) SetMessage(v string) {
	o.Message = v
}

func (o ExecuteScriptExceptionError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableExecuteScriptExceptionError struct {
	value *ExecuteScriptExceptionError
	isSet bool
}

func (v NullableExecuteScriptExceptionError) Get() *ExecuteScriptExceptionError {
	return v.value
}

func (v *NullableExecuteScriptExceptionError) Set(val *ExecuteScriptExceptionError) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteScriptExceptionError) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteScriptExceptionError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteScriptExceptionError(val *ExecuteScriptExceptionError) *NullableExecuteScriptExceptionError {
	return &NullableExecuteScriptExceptionError{value: val, isSet: true}
}

func (v NullableExecuteScriptExceptionError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteScriptExceptionError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


