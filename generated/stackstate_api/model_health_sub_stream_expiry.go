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

// HealthSubStreamExpiry struct for HealthSubStreamExpiry
type HealthSubStreamExpiry struct {
	Type string `json:"_type"`
	ExpiryIntervalMs int32 `json:"expiryIntervalMs"`
	RepeatIntervalMs int32 `json:"repeatIntervalMs"`
}

// NewHealthSubStreamExpiry instantiates a new HealthSubStreamExpiry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthSubStreamExpiry(type_ string, expiryIntervalMs int32, repeatIntervalMs int32) *HealthSubStreamExpiry {
	this := HealthSubStreamExpiry{}
	this.Type = type_
	this.ExpiryIntervalMs = expiryIntervalMs
	this.RepeatIntervalMs = repeatIntervalMs
	return &this
}

// NewHealthSubStreamExpiryWithDefaults instantiates a new HealthSubStreamExpiry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthSubStreamExpiryWithDefaults() *HealthSubStreamExpiry {
	this := HealthSubStreamExpiry{}
	return &this
}

// GetType returns the Type field value
func (o *HealthSubStreamExpiry) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HealthSubStreamExpiry) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HealthSubStreamExpiry) SetType(v string) {
	o.Type = v
}

// GetExpiryIntervalMs returns the ExpiryIntervalMs field value
func (o *HealthSubStreamExpiry) GetExpiryIntervalMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiryIntervalMs
}

// GetExpiryIntervalMsOk returns a tuple with the ExpiryIntervalMs field value
// and a boolean to check if the value has been set.
func (o *HealthSubStreamExpiry) GetExpiryIntervalMsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpiryIntervalMs, true
}

// SetExpiryIntervalMs sets field value
func (o *HealthSubStreamExpiry) SetExpiryIntervalMs(v int32) {
	o.ExpiryIntervalMs = v
}

// GetRepeatIntervalMs returns the RepeatIntervalMs field value
func (o *HealthSubStreamExpiry) GetRepeatIntervalMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RepeatIntervalMs
}

// GetRepeatIntervalMsOk returns a tuple with the RepeatIntervalMs field value
// and a boolean to check if the value has been set.
func (o *HealthSubStreamExpiry) GetRepeatIntervalMsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RepeatIntervalMs, true
}

// SetRepeatIntervalMs sets field value
func (o *HealthSubStreamExpiry) SetRepeatIntervalMs(v int32) {
	o.RepeatIntervalMs = v
}

func (o HealthSubStreamExpiry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["expiryIntervalMs"] = o.ExpiryIntervalMs
	}
	if true {
		toSerialize["repeatIntervalMs"] = o.RepeatIntervalMs
	}
	return json.Marshal(toSerialize)
}

type NullableHealthSubStreamExpiry struct {
	value *HealthSubStreamExpiry
	isSet bool
}

func (v NullableHealthSubStreamExpiry) Get() *HealthSubStreamExpiry {
	return v.value
}

func (v *NullableHealthSubStreamExpiry) Set(val *HealthSubStreamExpiry) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthSubStreamExpiry) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthSubStreamExpiry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthSubStreamExpiry(val *HealthSubStreamExpiry) *NullableHealthSubStreamExpiry {
	return &NullableHealthSubStreamExpiry{value: val, isSet: true}
}

func (v NullableHealthSubStreamExpiry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthSubStreamExpiry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


