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

// PromMatrix struct for PromMatrix
type PromMatrix struct {
	ResultType string `json:"resultType"`
	Result []PromDataResult `json:"result"`
}

// NewPromMatrix instantiates a new PromMatrix object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromMatrix(resultType string, result []PromDataResult) *PromMatrix {
	this := PromMatrix{}
	this.ResultType = resultType
	this.Result = result
	return &this
}

// NewPromMatrixWithDefaults instantiates a new PromMatrix object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromMatrixWithDefaults() *PromMatrix {
	this := PromMatrix{}
	return &this
}

// GetResultType returns the ResultType field value
func (o *PromMatrix) GetResultType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultType
}

// GetResultTypeOk returns a tuple with the ResultType field value
// and a boolean to check if the value has been set.
func (o *PromMatrix) GetResultTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultType, true
}

// SetResultType sets field value
func (o *PromMatrix) SetResultType(v string) {
	o.ResultType = v
}

// GetResult returns the Result field value
func (o *PromMatrix) GetResult() []PromDataResult {
	if o == nil {
		var ret []PromDataResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *PromMatrix) GetResultOk() ([]PromDataResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *PromMatrix) SetResult(v []PromDataResult) {
	o.Result = v
}

func (o PromMatrix) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultType"] = o.ResultType
	}
	if true {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullablePromMatrix struct {
	value *PromMatrix
	isSet bool
}

func (v NullablePromMatrix) Get() *PromMatrix {
	return v.value
}

func (v *NullablePromMatrix) Set(val *PromMatrix) {
	v.value = val
	v.isSet = true
}

func (v NullablePromMatrix) IsSet() bool {
	return v.isSet
}

func (v *NullablePromMatrix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromMatrix(val *PromMatrix) *NullablePromMatrix {
	return &NullablePromMatrix{value: val, isSet: true}
}

func (v NullablePromMatrix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromMatrix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

