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

// ExecuteScriptSyntaxErrors struct for ExecuteScriptSyntaxErrors
type ExecuteScriptSyntaxErrors struct {
	Type string `json:"_type"`
	Message string `json:"message"`
	Errors []ExecuteScriptSyntaxErrorsErrors `json:"errors"`
}

// NewExecuteScriptSyntaxErrors instantiates a new ExecuteScriptSyntaxErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteScriptSyntaxErrors(type_ string, message string, errors []ExecuteScriptSyntaxErrorsErrors) *ExecuteScriptSyntaxErrors {
	this := ExecuteScriptSyntaxErrors{}
	this.Type = type_
	this.Message = message
	this.Errors = errors
	return &this
}

// NewExecuteScriptSyntaxErrorsWithDefaults instantiates a new ExecuteScriptSyntaxErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteScriptSyntaxErrorsWithDefaults() *ExecuteScriptSyntaxErrors {
	this := ExecuteScriptSyntaxErrors{}
	return &this
}

// GetType returns the Type field value
func (o *ExecuteScriptSyntaxErrors) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExecuteScriptSyntaxErrors) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExecuteScriptSyntaxErrors) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *ExecuteScriptSyntaxErrors) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ExecuteScriptSyntaxErrors) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ExecuteScriptSyntaxErrors) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value
func (o *ExecuteScriptSyntaxErrors) GetErrors() []ExecuteScriptSyntaxErrorsErrors {
	if o == nil {
		var ret []ExecuteScriptSyntaxErrorsErrors
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *ExecuteScriptSyntaxErrors) GetErrorsOk() (*[]ExecuteScriptSyntaxErrorsErrors, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *ExecuteScriptSyntaxErrors) SetErrors(v []ExecuteScriptSyntaxErrorsErrors) {
	o.Errors = v
}

func (o ExecuteScriptSyntaxErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableExecuteScriptSyntaxErrors struct {
	value *ExecuteScriptSyntaxErrors
	isSet bool
}

func (v NullableExecuteScriptSyntaxErrors) Get() *ExecuteScriptSyntaxErrors {
	return v.value
}

func (v *NullableExecuteScriptSyntaxErrors) Set(val *ExecuteScriptSyntaxErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteScriptSyntaxErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteScriptSyntaxErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteScriptSyntaxErrors(val *ExecuteScriptSyntaxErrors) *NullableExecuteScriptSyntaxErrors {
	return &NullableExecuteScriptSyntaxErrors{value: val, isSet: true}
}

func (v NullableExecuteScriptSyntaxErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteScriptSyntaxErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


