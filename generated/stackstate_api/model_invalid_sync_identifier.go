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

// InvalidSyncIdentifier struct for InvalidSyncIdentifier
type InvalidSyncIdentifier struct {
	Message string `json:"message"`
}

// NewInvalidSyncIdentifier instantiates a new InvalidSyncIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidSyncIdentifier(message string) *InvalidSyncIdentifier {
	this := InvalidSyncIdentifier{}
	this.Message = message
	return &this
}

// NewInvalidSyncIdentifierWithDefaults instantiates a new InvalidSyncIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidSyncIdentifierWithDefaults() *InvalidSyncIdentifier {
	this := InvalidSyncIdentifier{}
	return &this
}

// GetMessage returns the Message field value
func (o *InvalidSyncIdentifier) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InvalidSyncIdentifier) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InvalidSyncIdentifier) SetMessage(v string) {
	o.Message = v
}

func (o InvalidSyncIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableInvalidSyncIdentifier struct {
	value *InvalidSyncIdentifier
	isSet bool
}

func (v NullableInvalidSyncIdentifier) Get() *InvalidSyncIdentifier {
	return v.value
}

func (v *NullableInvalidSyncIdentifier) Set(val *InvalidSyncIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidSyncIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidSyncIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidSyncIdentifier(val *InvalidSyncIdentifier) *NullableInvalidSyncIdentifier {
	return &NullableInvalidSyncIdentifier{value: val, isSet: true}
}

func (v NullableInvalidSyncIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidSyncIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
