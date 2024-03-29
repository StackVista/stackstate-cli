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

// ExecuteScriptResponse struct for ExecuteScriptResponse
type ExecuteScriptResponse struct {
	// Contains a `value` and `_type` property. This is described this way, because `value` can hold any JSON type and most open api generators (but not all) have problems with such types or are inconsistent with the OpenAPI spec. The `_type` property describes the original type of the result value during script execution. The `value` property holds the resulting value serialized to JSON. The `value` property can be of any JSON type, i.e. null, undefined, boolean, number, string, array or object.
	Result map[string]interface{} `json:"result"`
}

// NewExecuteScriptResponse instantiates a new ExecuteScriptResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteScriptResponse(result map[string]interface{}) *ExecuteScriptResponse {
	this := ExecuteScriptResponse{}
	this.Result = result
	return &this
}

// NewExecuteScriptResponseWithDefaults instantiates a new ExecuteScriptResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteScriptResponseWithDefaults() *ExecuteScriptResponse {
	this := ExecuteScriptResponse{}
	return &this
}

// GetResult returns the Result field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *ExecuteScriptResponse) GetResult() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExecuteScriptResponse) GetResultOk() (map[string]interface{}, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *ExecuteScriptResponse) SetResult(v map[string]interface{}) {
	o.Result = v
}

func (o ExecuteScriptResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableExecuteScriptResponse struct {
	value *ExecuteScriptResponse
	isSet bool
}

func (v NullableExecuteScriptResponse) Get() *ExecuteScriptResponse {
	return v.value
}

func (v *NullableExecuteScriptResponse) Set(val *ExecuteScriptResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteScriptResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteScriptResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteScriptResponse(val *ExecuteScriptResponse) *NullableExecuteScriptResponse {
	return &NullableExecuteScriptResponse{value: val, isSet: true}
}

func (v NullableExecuteScriptResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteScriptResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
