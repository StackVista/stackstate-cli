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

// PromScalar struct for PromScalar
type PromScalar struct {
	ResultType string `json:"resultType"`
	// This is always a tuple represented as an array with in first position the unix timestamp as  a float with precision 3 in seconds) and in second position the sample value as a string. 
	Result []PromSampleInner `json:"result"`
}

// NewPromScalar instantiates a new PromScalar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromScalar(resultType string, result []PromSampleInner) *PromScalar {
	this := PromScalar{}
	this.ResultType = resultType
	this.Result = result
	return &this
}

// NewPromScalarWithDefaults instantiates a new PromScalar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromScalarWithDefaults() *PromScalar {
	this := PromScalar{}
	return &this
}

// GetResultType returns the ResultType field value
func (o *PromScalar) GetResultType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultType
}

// GetResultTypeOk returns a tuple with the ResultType field value
// and a boolean to check if the value has been set.
func (o *PromScalar) GetResultTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultType, true
}

// SetResultType sets field value
func (o *PromScalar) SetResultType(v string) {
	o.ResultType = v
}

// GetResult returns the Result field value
func (o *PromScalar) GetResult() []PromSampleInner {
	if o == nil {
		var ret []PromSampleInner
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *PromScalar) GetResultOk() ([]PromSampleInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *PromScalar) SetResult(v []PromSampleInner) {
	o.Result = v
}

func (o PromScalar) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultType"] = o.ResultType
	}
	if true {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullablePromScalar struct {
	value *PromScalar
	isSet bool
}

func (v NullablePromScalar) Get() *PromScalar {
	return v.value
}

func (v *NullablePromScalar) Set(val *PromScalar) {
	v.value = val
	v.isSet = true
}

func (v NullablePromScalar) IsSet() bool {
	return v.isSet
}

func (v *NullablePromScalar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromScalar(val *PromScalar) *NullablePromScalar {
	return &NullablePromScalar{value: val, isSet: true}
}

func (v NullablePromScalar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromScalar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

