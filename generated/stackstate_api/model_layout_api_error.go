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

// LayoutApiError struct for LayoutApiError
type LayoutApiError struct {
	StatusCode string `json:"statusCode"`
	Message    string `json:"message"`
}

// NewLayoutApiError instantiates a new LayoutApiError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLayoutApiError(statusCode string, message string) *LayoutApiError {
	this := LayoutApiError{}
	this.StatusCode = statusCode
	this.Message = message
	return &this
}

// NewLayoutApiErrorWithDefaults instantiates a new LayoutApiError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLayoutApiErrorWithDefaults() *LayoutApiError {
	this := LayoutApiError{}
	return &this
}

// GetStatusCode returns the StatusCode field value
func (o *LayoutApiError) GetStatusCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *LayoutApiError) GetStatusCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *LayoutApiError) SetStatusCode(v string) {
	o.StatusCode = v
}

// GetMessage returns the Message field value
func (o *LayoutApiError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *LayoutApiError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *LayoutApiError) SetMessage(v string) {
	o.Message = v
}

func (o LayoutApiError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["statusCode"] = o.StatusCode
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableLayoutApiError struct {
	value *LayoutApiError
	isSet bool
}

func (v NullableLayoutApiError) Get() *LayoutApiError {
	return v.value
}

func (v *NullableLayoutApiError) Set(val *LayoutApiError) {
	v.value = val
	v.isSet = true
}

func (v NullableLayoutApiError) IsSet() bool {
	return v.isSet
}

func (v *NullableLayoutApiError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLayoutApiError(val *LayoutApiError) *NullableLayoutApiError {
	return &NullableLayoutApiError{value: val, isSet: true}
}

func (v NullableLayoutApiError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLayoutApiError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
