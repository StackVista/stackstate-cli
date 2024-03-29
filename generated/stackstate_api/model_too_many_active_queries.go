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

// TooManyActiveQueries struct for TooManyActiveQueries
type TooManyActiveQueries struct {
	Type    string `json:"_type"`
	Message string `json:"message"`
}

// NewTooManyActiveQueries instantiates a new TooManyActiveQueries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTooManyActiveQueries(type_ string, message string) *TooManyActiveQueries {
	this := TooManyActiveQueries{}
	this.Type = type_
	this.Message = message
	return &this
}

// NewTooManyActiveQueriesWithDefaults instantiates a new TooManyActiveQueries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTooManyActiveQueriesWithDefaults() *TooManyActiveQueries {
	this := TooManyActiveQueries{}
	return &this
}

// GetType returns the Type field value
func (o *TooManyActiveQueries) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TooManyActiveQueries) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TooManyActiveQueries) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *TooManyActiveQueries) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *TooManyActiveQueries) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *TooManyActiveQueries) SetMessage(v string) {
	o.Message = v
}

func (o TooManyActiveQueries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableTooManyActiveQueries struct {
	value *TooManyActiveQueries
	isSet bool
}

func (v NullableTooManyActiveQueries) Get() *TooManyActiveQueries {
	return v.value
}

func (v *NullableTooManyActiveQueries) Set(val *TooManyActiveQueries) {
	v.value = val
	v.isSet = true
}

func (v NullableTooManyActiveQueries) IsSet() bool {
	return v.isSet
}

func (v *NullableTooManyActiveQueries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTooManyActiveQueries(val *TooManyActiveQueries) *NullableTooManyActiveQueries {
	return &NullableTooManyActiveQueries{value: val, isSet: true}
}

func (v NullableTooManyActiveQueries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTooManyActiveQueries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
