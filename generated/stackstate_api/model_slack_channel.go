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

// SlackChannel struct for SlackChannel
type SlackChannel struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

// NewSlackChannel instantiates a new SlackChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackChannel(name string, id string) *SlackChannel {
	this := SlackChannel{}
	this.Name = name
	this.Id = id
	return &this
}

// NewSlackChannelWithDefaults instantiates a new SlackChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackChannelWithDefaults() *SlackChannel {
	this := SlackChannel{}
	return &this
}

// GetName returns the Name field value
func (o *SlackChannel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SlackChannel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SlackChannel) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *SlackChannel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SlackChannel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SlackChannel) SetId(v string) {
	o.Id = v
}

func (o SlackChannel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSlackChannel struct {
	value *SlackChannel
	isSet bool
}

func (v NullableSlackChannel) Get() *SlackChannel {
	return v.value
}

func (v *NullableSlackChannel) Set(val *SlackChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackChannel(val *SlackChannel) *NullableSlackChannel {
	return &NullableSlackChannel{value: val, isSet: true}
}

func (v NullableSlackChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}