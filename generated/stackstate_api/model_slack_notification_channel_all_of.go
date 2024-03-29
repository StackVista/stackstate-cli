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

// SlackNotificationChannelAllOf struct for SlackNotificationChannelAllOf
type SlackNotificationChannelAllOf struct {
	Type           string  `json:"_type"`
	SlackWorkspace string  `json:"slackWorkspace"`
	SlackChannel   *string `json:"slackChannel,omitempty"`
	SlackChannelId *string `json:"slackChannelId,omitempty"`
}

// NewSlackNotificationChannelAllOf instantiates a new SlackNotificationChannelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackNotificationChannelAllOf(type_ string, slackWorkspace string) *SlackNotificationChannelAllOf {
	this := SlackNotificationChannelAllOf{}
	this.Type = type_
	this.SlackWorkspace = slackWorkspace
	return &this
}

// NewSlackNotificationChannelAllOfWithDefaults instantiates a new SlackNotificationChannelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackNotificationChannelAllOfWithDefaults() *SlackNotificationChannelAllOf {
	this := SlackNotificationChannelAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *SlackNotificationChannelAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SlackNotificationChannelAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SlackNotificationChannelAllOf) SetType(v string) {
	o.Type = v
}

// GetSlackWorkspace returns the SlackWorkspace field value
func (o *SlackNotificationChannelAllOf) GetSlackWorkspace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SlackWorkspace
}

// GetSlackWorkspaceOk returns a tuple with the SlackWorkspace field value
// and a boolean to check if the value has been set.
func (o *SlackNotificationChannelAllOf) GetSlackWorkspaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlackWorkspace, true
}

// SetSlackWorkspace sets field value
func (o *SlackNotificationChannelAllOf) SetSlackWorkspace(v string) {
	o.SlackWorkspace = v
}

// GetSlackChannel returns the SlackChannel field value if set, zero value otherwise.
func (o *SlackNotificationChannelAllOf) GetSlackChannel() string {
	if o == nil || o.SlackChannel == nil {
		var ret string
		return ret
	}
	return *o.SlackChannel
}

// GetSlackChannelOk returns a tuple with the SlackChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackNotificationChannelAllOf) GetSlackChannelOk() (*string, bool) {
	if o == nil || o.SlackChannel == nil {
		return nil, false
	}
	return o.SlackChannel, true
}

// HasSlackChannel returns a boolean if a field has been set.
func (o *SlackNotificationChannelAllOf) HasSlackChannel() bool {
	if o != nil && o.SlackChannel != nil {
		return true
	}

	return false
}

// SetSlackChannel gets a reference to the given string and assigns it to the SlackChannel field.
func (o *SlackNotificationChannelAllOf) SetSlackChannel(v string) {
	o.SlackChannel = &v
}

// GetSlackChannelId returns the SlackChannelId field value if set, zero value otherwise.
func (o *SlackNotificationChannelAllOf) GetSlackChannelId() string {
	if o == nil || o.SlackChannelId == nil {
		var ret string
		return ret
	}
	return *o.SlackChannelId
}

// GetSlackChannelIdOk returns a tuple with the SlackChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackNotificationChannelAllOf) GetSlackChannelIdOk() (*string, bool) {
	if o == nil || o.SlackChannelId == nil {
		return nil, false
	}
	return o.SlackChannelId, true
}

// HasSlackChannelId returns a boolean if a field has been set.
func (o *SlackNotificationChannelAllOf) HasSlackChannelId() bool {
	if o != nil && o.SlackChannelId != nil {
		return true
	}

	return false
}

// SetSlackChannelId gets a reference to the given string and assigns it to the SlackChannelId field.
func (o *SlackNotificationChannelAllOf) SetSlackChannelId(v string) {
	o.SlackChannelId = &v
}

func (o SlackNotificationChannelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["slackWorkspace"] = o.SlackWorkspace
	}
	if o.SlackChannel != nil {
		toSerialize["slackChannel"] = o.SlackChannel
	}
	if o.SlackChannelId != nil {
		toSerialize["slackChannelId"] = o.SlackChannelId
	}
	return json.Marshal(toSerialize)
}

type NullableSlackNotificationChannelAllOf struct {
	value *SlackNotificationChannelAllOf
	isSet bool
}

func (v NullableSlackNotificationChannelAllOf) Get() *SlackNotificationChannelAllOf {
	return v.value
}

func (v *NullableSlackNotificationChannelAllOf) Set(val *SlackNotificationChannelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackNotificationChannelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackNotificationChannelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackNotificationChannelAllOf(val *SlackNotificationChannelAllOf) *NullableSlackNotificationChannelAllOf {
	return &NullableSlackNotificationChannelAllOf{value: val, isSet: true}
}

func (v NullableSlackNotificationChannelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackNotificationChannelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
