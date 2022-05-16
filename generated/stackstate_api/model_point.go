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

// Point struct for Point
type Point struct {
	Ts float32 `json:"ts"`
	V float64 `json:"v"`
}

// NewPoint instantiates a new Point object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoint(ts float32, v float64) *Point {
	this := Point{}
	this.Ts = ts
	this.V = v
	return &this
}

// NewPointWithDefaults instantiates a new Point object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointWithDefaults() *Point {
	this := Point{}
	return &this
}

// GetTs returns the Ts field value
func (o *Point) GetTs() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Ts
}

// GetTsOk returns a tuple with the Ts field value
// and a boolean to check if the value has been set.
func (o *Point) GetTsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ts, true
}

// SetTs sets field value
func (o *Point) SetTs(v float32) {
	o.Ts = v
}

// GetV returns the V field value
func (o *Point) GetV() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.V
}

// GetVOk returns a tuple with the V field value
// and a boolean to check if the value has been set.
func (o *Point) GetVOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.V, true
}

// SetV sets field value
func (o *Point) SetV(v float64) {
	o.V = v
}

func (o Point) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ts"] = o.Ts
	}
	if true {
		toSerialize["v"] = o.V
	}
	return json.Marshal(toSerialize)
}

type NullablePoint struct {
	value *Point
	isSet bool
}

func (v NullablePoint) Get() *Point {
	return v.value
}

func (v *NullablePoint) Set(val *Point) {
	v.value = val
	v.isSet = true
}

func (v NullablePoint) IsSet() bool {
	return v.isSet
}

func (v *NullablePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoint(val *Point) *NullablePoint {
	return &NullablePoint{value: val, isSet: true}
}

func (v NullablePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


