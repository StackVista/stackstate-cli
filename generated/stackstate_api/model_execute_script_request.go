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

// ExecuteScriptRequest struct for ExecuteScriptRequest
type ExecuteScriptRequest struct {
	TimeoutMs *int32 `json:"timeoutMs,omitempty"`
	Script string `json:"script"`
	ArgumentsScript *string `json:"argumentsScript,omitempty"`
}

// NewExecuteScriptRequest instantiates a new ExecuteScriptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteScriptRequest(script string) *ExecuteScriptRequest {
	this := ExecuteScriptRequest{}
	this.Script = script
	return &this
}

// NewExecuteScriptRequestWithDefaults instantiates a new ExecuteScriptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteScriptRequestWithDefaults() *ExecuteScriptRequest {
	this := ExecuteScriptRequest{}
	return &this
}

// GetTimeoutMs returns the TimeoutMs field value if set, zero value otherwise.
func (o *ExecuteScriptRequest) GetTimeoutMs() int32 {
	if o == nil || o.TimeoutMs == nil {
		var ret int32
		return ret
	}
	return *o.TimeoutMs
}

// GetTimeoutMsOk returns a tuple with the TimeoutMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteScriptRequest) GetTimeoutMsOk() (*int32, bool) {
	if o == nil || o.TimeoutMs == nil {
		return nil, false
	}
	return o.TimeoutMs, true
}

// HasTimeoutMs returns a boolean if a field has been set.
func (o *ExecuteScriptRequest) HasTimeoutMs() bool {
	if o != nil && o.TimeoutMs != nil {
		return true
	}

	return false
}

// SetTimeoutMs gets a reference to the given int32 and assigns it to the TimeoutMs field.
func (o *ExecuteScriptRequest) SetTimeoutMs(v int32) {
	o.TimeoutMs = &v
}

// GetScript returns the Script field value
func (o *ExecuteScriptRequest) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *ExecuteScriptRequest) GetScriptOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *ExecuteScriptRequest) SetScript(v string) {
	o.Script = v
}

// GetArgumentsScript returns the ArgumentsScript field value if set, zero value otherwise.
func (o *ExecuteScriptRequest) GetArgumentsScript() string {
	if o == nil || o.ArgumentsScript == nil {
		var ret string
		return ret
	}
	return *o.ArgumentsScript
}

// GetArgumentsScriptOk returns a tuple with the ArgumentsScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteScriptRequest) GetArgumentsScriptOk() (*string, bool) {
	if o == nil || o.ArgumentsScript == nil {
		return nil, false
	}
	return o.ArgumentsScript, true
}

// HasArgumentsScript returns a boolean if a field has been set.
func (o *ExecuteScriptRequest) HasArgumentsScript() bool {
	if o != nil && o.ArgumentsScript != nil {
		return true
	}

	return false
}

// SetArgumentsScript gets a reference to the given string and assigns it to the ArgumentsScript field.
func (o *ExecuteScriptRequest) SetArgumentsScript(v string) {
	o.ArgumentsScript = &v
}

func (o ExecuteScriptRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TimeoutMs != nil {
		toSerialize["timeoutMs"] = o.TimeoutMs
	}
	if true {
		toSerialize["script"] = o.Script
	}
	if o.ArgumentsScript != nil {
		toSerialize["argumentsScript"] = o.ArgumentsScript
	}
	return json.Marshal(toSerialize)
}

type NullableExecuteScriptRequest struct {
	value *ExecuteScriptRequest
	isSet bool
}

func (v NullableExecuteScriptRequest) Get() *ExecuteScriptRequest {
	return v.value
}

func (v *NullableExecuteScriptRequest) Set(val *ExecuteScriptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteScriptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteScriptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteScriptRequest(val *ExecuteScriptRequest) *NullableExecuteScriptRequest {
	return &NullableExecuteScriptRequest{value: val, isSet: true}
}

func (v NullableExecuteScriptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteScriptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

