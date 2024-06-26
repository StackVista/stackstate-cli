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

// BaseLayoutHint struct for BaseLayoutHint
type BaseLayoutHint struct {
	Name   string  `json:"name"`
	Weight float32 `json:"weight"`
}

// NewBaseLayoutHint instantiates a new BaseLayoutHint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseLayoutHint(name string, weight float32) *BaseLayoutHint {
	this := BaseLayoutHint{}
	this.Name = name
	this.Weight = weight
	return &this
}

// NewBaseLayoutHintWithDefaults instantiates a new BaseLayoutHint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseLayoutHintWithDefaults() *BaseLayoutHint {
	this := BaseLayoutHint{}
	return &this
}

// GetName returns the Name field value
func (o *BaseLayoutHint) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BaseLayoutHint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BaseLayoutHint) SetName(v string) {
	o.Name = v
}

// GetWeight returns the Weight field value
func (o *BaseLayoutHint) GetWeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *BaseLayoutHint) GetWeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *BaseLayoutHint) SetWeight(v float32) {
	o.Weight = v
}

func (o BaseLayoutHint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableBaseLayoutHint struct {
	value *BaseLayoutHint
	isSet bool
}

func (v NullableBaseLayoutHint) Get() *BaseLayoutHint {
	return v.value
}

func (v *NullableBaseLayoutHint) Set(val *BaseLayoutHint) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseLayoutHint) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseLayoutHint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseLayoutHint(val *BaseLayoutHint) *NullableBaseLayoutHint {
	return &NullableBaseLayoutHint{value: val, isSet: true}
}

func (v NullableBaseLayoutHint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseLayoutHint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
