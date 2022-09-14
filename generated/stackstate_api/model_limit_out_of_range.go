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

// LimitOutOfRange struct for LimitOutOfRange
type LimitOutOfRange struct {
	Type string `json:"_type"`
	Message string `json:"message"`
	Limit int32 `json:"limit"`
	LowerBound int32 `json:"lowerBound"`
	UpperBound int32 `json:"upperBound"`
}

// NewLimitOutOfRange instantiates a new LimitOutOfRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimitOutOfRange(type_ string, message string, limit int32, lowerBound int32, upperBound int32) *LimitOutOfRange {
	this := LimitOutOfRange{}
	this.Type = type_
	this.Message = message
	this.Limit = limit
	this.LowerBound = lowerBound
	this.UpperBound = upperBound
	return &this
}

// NewLimitOutOfRangeWithDefaults instantiates a new LimitOutOfRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitOutOfRangeWithDefaults() *LimitOutOfRange {
	this := LimitOutOfRange{}
	return &this
}

// GetType returns the Type field value
func (o *LimitOutOfRange) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LimitOutOfRange) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LimitOutOfRange) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *LimitOutOfRange) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *LimitOutOfRange) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *LimitOutOfRange) SetMessage(v string) {
	o.Message = v
}

// GetLimit returns the Limit field value
func (o *LimitOutOfRange) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *LimitOutOfRange) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *LimitOutOfRange) SetLimit(v int32) {
	o.Limit = v
}

// GetLowerBound returns the LowerBound field value
func (o *LimitOutOfRange) GetLowerBound() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LowerBound
}

// GetLowerBoundOk returns a tuple with the LowerBound field value
// and a boolean to check if the value has been set.
func (o *LimitOutOfRange) GetLowerBoundOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LowerBound, true
}

// SetLowerBound sets field value
func (o *LimitOutOfRange) SetLowerBound(v int32) {
	o.LowerBound = v
}

// GetUpperBound returns the UpperBound field value
func (o *LimitOutOfRange) GetUpperBound() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpperBound
}

// GetUpperBoundOk returns a tuple with the UpperBound field value
// and a boolean to check if the value has been set.
func (o *LimitOutOfRange) GetUpperBoundOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpperBound, true
}

// SetUpperBound sets field value
func (o *LimitOutOfRange) SetUpperBound(v int32) {
	o.UpperBound = v
}

func (o LimitOutOfRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["lowerBound"] = o.LowerBound
	}
	if true {
		toSerialize["upperBound"] = o.UpperBound
	}
	return json.Marshal(toSerialize)
}

type NullableLimitOutOfRange struct {
	value *LimitOutOfRange
	isSet bool
}

func (v NullableLimitOutOfRange) Get() *LimitOutOfRange {
	return v.value
}

func (v *NullableLimitOutOfRange) Set(val *LimitOutOfRange) {
	v.value = val
	v.isSet = true
}

func (v NullableLimitOutOfRange) IsSet() bool {
	return v.isSet
}

func (v *NullableLimitOutOfRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimitOutOfRange(val *LimitOutOfRange) *NullableLimitOutOfRange {
	return &NullableLimitOutOfRange{value: val, isSet: true}
}

func (v NullableLimitOutOfRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimitOutOfRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

