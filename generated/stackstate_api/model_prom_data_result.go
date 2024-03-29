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

// PromDataResult struct for PromDataResult
type PromDataResult struct {
	Metric map[string]string   `json:"metric"`
	Values [][]PromSampleInner `json:"values"`
}

// NewPromDataResult instantiates a new PromDataResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromDataResult(metric map[string]string, values [][]PromSampleInner) *PromDataResult {
	this := PromDataResult{}
	this.Metric = metric
	this.Values = values
	return &this
}

// NewPromDataResultWithDefaults instantiates a new PromDataResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromDataResultWithDefaults() *PromDataResult {
	this := PromDataResult{}
	return &this
}

// GetMetric returns the Metric field value
func (o *PromDataResult) GetMetric() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *PromDataResult) GetMetricOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *PromDataResult) SetMetric(v map[string]string) {
	o.Metric = v
}

// GetValues returns the Values field value
func (o *PromDataResult) GetValues() [][]PromSampleInner {
	if o == nil {
		var ret [][]PromSampleInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *PromDataResult) GetValuesOk() ([][]PromSampleInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *PromDataResult) SetValues(v [][]PromSampleInner) {
	o.Values = v
}

func (o PromDataResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metric"] = o.Metric
	}
	if true {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullablePromDataResult struct {
	value *PromDataResult
	isSet bool
}

func (v NullablePromDataResult) Get() *PromDataResult {
	return v.value
}

func (v *NullablePromDataResult) Set(val *PromDataResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePromDataResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePromDataResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromDataResult(val *PromDataResult) *NullablePromDataResult {
	return &NullablePromDataResult{value: val, isSet: true}
}

func (v NullablePromDataResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromDataResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
