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

// RequestValidationError struct for RequestValidationError
type RequestValidationError struct {
	Type string `json:"_type"`
	Message string `json:"message"`
}

// NewRequestValidationError instantiates a new RequestValidationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestValidationError(type_ string, message string) *RequestValidationError {
	this := RequestValidationError{}
	this.Type = type_
	this.Message = message
	return &this
}

// NewRequestValidationErrorWithDefaults instantiates a new RequestValidationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestValidationErrorWithDefaults() *RequestValidationError {
	this := RequestValidationError{}
	return &this
}

// GetType returns the Type field value
func (o *RequestValidationError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestValidationError) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestValidationError) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *RequestValidationError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RequestValidationError) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RequestValidationError) SetMessage(v string) {
	o.Message = v
}

func (o RequestValidationError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableRequestValidationError struct {
	value *RequestValidationError
	isSet bool
}

func (v NullableRequestValidationError) Get() *RequestValidationError {
	return v.value
}

func (v *NullableRequestValidationError) Set(val *RequestValidationError) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestValidationError) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestValidationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestValidationError(val *RequestValidationError) *NullableRequestValidationError {
	return &NullableRequestValidationError{value: val, isSet: true}
}

func (v NullableRequestValidationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestValidationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

