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

// EmailNotificationStatus struct for EmailNotificationStatus
type EmailNotificationStatus struct {
	ValidConfiguration bool `json:"validConfiguration"`
}

// NewEmailNotificationStatus instantiates a new EmailNotificationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailNotificationStatus(validConfiguration bool) *EmailNotificationStatus {
	this := EmailNotificationStatus{}
	this.ValidConfiguration = validConfiguration
	return &this
}

// NewEmailNotificationStatusWithDefaults instantiates a new EmailNotificationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailNotificationStatusWithDefaults() *EmailNotificationStatus {
	this := EmailNotificationStatus{}
	return &this
}

// GetValidConfiguration returns the ValidConfiguration field value
func (o *EmailNotificationStatus) GetValidConfiguration() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ValidConfiguration
}

// GetValidConfigurationOk returns a tuple with the ValidConfiguration field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationStatus) GetValidConfigurationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidConfiguration, true
}

// SetValidConfiguration sets field value
func (o *EmailNotificationStatus) SetValidConfiguration(v bool) {
	o.ValidConfiguration = v
}

func (o EmailNotificationStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["validConfiguration"] = o.ValidConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableEmailNotificationStatus struct {
	value *EmailNotificationStatus
	isSet bool
}

func (v NullableEmailNotificationStatus) Get() *EmailNotificationStatus {
	return v.value
}

func (v *NullableEmailNotificationStatus) Set(val *EmailNotificationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailNotificationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailNotificationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailNotificationStatus(val *EmailNotificationStatus) *NullableEmailNotificationStatus {
	return &NullableEmailNotificationStatus{value: val, isSet: true}
}

func (v NullableEmailNotificationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailNotificationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
