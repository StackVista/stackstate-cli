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

// MonitorRunResult struct for MonitorRunResult
type MonitorRunResult struct {
	Result map[string]interface{} `json:"result"`
}

// NewMonitorRunResult instantiates a new MonitorRunResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorRunResult(result map[string]interface{}) *MonitorRunResult {
	this := MonitorRunResult{}
	this.Result = result
	return &this
}

// NewMonitorRunResultWithDefaults instantiates a new MonitorRunResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorRunResultWithDefaults() *MonitorRunResult {
	this := MonitorRunResult{}
	return &this
}

// GetResult returns the Result field value
func (o *MonitorRunResult) GetResult() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *MonitorRunResult) GetResultOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *MonitorRunResult) SetResult(v map[string]interface{}) {
	o.Result = v
}

func (o MonitorRunResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorRunResult struct {
	value *MonitorRunResult
	isSet bool
}

func (v NullableMonitorRunResult) Get() *MonitorRunResult {
	return v.value
}

func (v *NullableMonitorRunResult) Set(val *MonitorRunResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorRunResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorRunResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorRunResult(val *MonitorRunResult) *NullableMonitorRunResult {
	return &NullableMonitorRunResult{value: val, isSet: true}
}

func (v NullableMonitorRunResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorRunResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


