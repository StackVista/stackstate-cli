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

// ExecuteScriptSyntaxErrorsErrorsInner struct for ExecuteScriptSyntaxErrorsErrorsInner
type ExecuteScriptSyntaxErrorsErrorsInner struct {
	Message  string         `json:"message"`
	Location ScriptLocation `json:"location"`
}

// NewExecuteScriptSyntaxErrorsErrorsInner instantiates a new ExecuteScriptSyntaxErrorsErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteScriptSyntaxErrorsErrorsInner(message string, location ScriptLocation) *ExecuteScriptSyntaxErrorsErrorsInner {
	this := ExecuteScriptSyntaxErrorsErrorsInner{}
	this.Message = message
	this.Location = location
	return &this
}

// NewExecuteScriptSyntaxErrorsErrorsInnerWithDefaults instantiates a new ExecuteScriptSyntaxErrorsErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteScriptSyntaxErrorsErrorsInnerWithDefaults() *ExecuteScriptSyntaxErrorsErrorsInner {
	this := ExecuteScriptSyntaxErrorsErrorsInner{}
	return &this
}

// GetMessage returns the Message field value
func (o *ExecuteScriptSyntaxErrorsErrorsInner) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ExecuteScriptSyntaxErrorsErrorsInner) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ExecuteScriptSyntaxErrorsErrorsInner) SetMessage(v string) {
	o.Message = v
}

// GetLocation returns the Location field value
func (o *ExecuteScriptSyntaxErrorsErrorsInner) GetLocation() ScriptLocation {
	if o == nil {
		var ret ScriptLocation
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *ExecuteScriptSyntaxErrorsErrorsInner) GetLocationOk() (*ScriptLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *ExecuteScriptSyntaxErrorsErrorsInner) SetLocation(v ScriptLocation) {
	o.Location = v
}

func (o ExecuteScriptSyntaxErrorsErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableExecuteScriptSyntaxErrorsErrorsInner struct {
	value *ExecuteScriptSyntaxErrorsErrorsInner
	isSet bool
}

func (v NullableExecuteScriptSyntaxErrorsErrorsInner) Get() *ExecuteScriptSyntaxErrorsErrorsInner {
	return v.value
}

func (v *NullableExecuteScriptSyntaxErrorsErrorsInner) Set(val *ExecuteScriptSyntaxErrorsErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteScriptSyntaxErrorsErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteScriptSyntaxErrorsErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteScriptSyntaxErrorsErrorsInner(val *ExecuteScriptSyntaxErrorsErrorsInner) *NullableExecuteScriptSyntaxErrorsErrorsInner {
	return &NullableExecuteScriptSyntaxErrorsErrorsInner{value: val, isSet: true}
}

func (v NullableExecuteScriptSyntaxErrorsErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteScriptSyntaxErrorsErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
