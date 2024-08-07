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

// EmailNotificationChannelAllOf struct for EmailNotificationChannelAllOf
type EmailNotificationChannelAllOf struct {
	Type string `json:"_type"`
}

// NewEmailNotificationChannelAllOf instantiates a new EmailNotificationChannelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailNotificationChannelAllOf(type_ string) *EmailNotificationChannelAllOf {
	this := EmailNotificationChannelAllOf{}
	this.Type = type_
	return &this
}

// NewEmailNotificationChannelAllOfWithDefaults instantiates a new EmailNotificationChannelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailNotificationChannelAllOfWithDefaults() *EmailNotificationChannelAllOf {
	this := EmailNotificationChannelAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *EmailNotificationChannelAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationChannelAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EmailNotificationChannelAllOf) SetType(v string) {
	o.Type = v
}

func (o EmailNotificationChannelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableEmailNotificationChannelAllOf struct {
	value *EmailNotificationChannelAllOf
	isSet bool
}

func (v NullableEmailNotificationChannelAllOf) Get() *EmailNotificationChannelAllOf {
	return v.value
}

func (v *NullableEmailNotificationChannelAllOf) Set(val *EmailNotificationChannelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailNotificationChannelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailNotificationChannelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailNotificationChannelAllOf(val *EmailNotificationChannelAllOf) *NullableEmailNotificationChannelAllOf {
	return &NullableEmailNotificationChannelAllOf{value: val, isSet: true}
}

func (v NullableEmailNotificationChannelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailNotificationChannelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
