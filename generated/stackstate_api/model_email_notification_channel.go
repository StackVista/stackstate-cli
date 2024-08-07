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

// EmailNotificationChannel struct for EmailNotificationChannel
type EmailNotificationChannel struct {
	Id                          int64                     `json:"id"`
	NotificationConfigurationId *int64                    `json:"notificationConfigurationId,omitempty"`
	Status                      NotificationChannelStatus `json:"status"`
	To                          []string                  `json:"to"`
	Cc                          []string                  `json:"cc"`
	SubjectPrefix               *string                   `json:"subjectPrefix,omitempty"`
	Type                        string                    `json:"_type"`
}

// NewEmailNotificationChannel instantiates a new EmailNotificationChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailNotificationChannel(id int64, status NotificationChannelStatus, to []string, cc []string, type_ string) *EmailNotificationChannel {
	this := EmailNotificationChannel{}
	this.Id = id
	this.Status = status
	this.To = to
	this.Cc = cc
	this.Type = type_
	return &this
}

// NewEmailNotificationChannelWithDefaults instantiates a new EmailNotificationChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailNotificationChannelWithDefaults() *EmailNotificationChannel {
	this := EmailNotificationChannel{}
	return &this
}

// GetId returns the Id field value
func (o *EmailNotificationChannel) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationChannel) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EmailNotificationChannel) SetId(v int64) {
	o.Id = v
}

// GetNotificationConfigurationId returns the NotificationConfigurationId field value if set, zero value otherwise.
func (o *EmailNotificationChannel) GetNotificationConfigurationId() int64 {
	if o == nil || o.NotificationConfigurationId == nil {
		var ret int64
		return ret
	}
	return *o.NotificationConfigurationId
}

// GetNotificationConfigurationIdOk returns a tuple with the NotificationConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationChannel) GetNotificationConfigurationIdOk() (*int64, bool) {
	if o == nil || o.NotificationConfigurationId == nil {
		return nil, false
	}
	return o.NotificationConfigurationId, true
}

// HasNotificationConfigurationId returns a boolean if a field has been set.
func (o *EmailNotificationChannel) HasNotificationConfigurationId() bool {
	if o != nil && o.NotificationConfigurationId != nil {
		return true
	}

	return false
}

// SetNotificationConfigurationId gets a reference to the given int64 and assigns it to the NotificationConfigurationId field.
func (o *EmailNotificationChannel) SetNotificationConfigurationId(v int64) {
	o.NotificationConfigurationId = &v
}

// GetStatus returns the Status field value
func (o *EmailNotificationChannel) GetStatus() NotificationChannelStatus {
	if o == nil {
		var ret NotificationChannelStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationChannel) GetStatusOk() (*NotificationChannelStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EmailNotificationChannel) SetStatus(v NotificationChannelStatus) {
	o.Status = v
}

// GetTo returns the To field value
func (o *EmailNotificationChannel) GetTo() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationChannel) GetToOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.To, true
}

// SetTo sets field value
func (o *EmailNotificationChannel) SetTo(v []string) {
	o.To = v
}

// GetCc returns the Cc field value
func (o *EmailNotificationChannel) GetCc() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Cc
}

// GetCcOk returns a tuple with the Cc field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationChannel) GetCcOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cc, true
}

// SetCc sets field value
func (o *EmailNotificationChannel) SetCc(v []string) {
	o.Cc = v
}

// GetSubjectPrefix returns the SubjectPrefix field value if set, zero value otherwise.
func (o *EmailNotificationChannel) GetSubjectPrefix() string {
	if o == nil || o.SubjectPrefix == nil {
		var ret string
		return ret
	}
	return *o.SubjectPrefix
}

// GetSubjectPrefixOk returns a tuple with the SubjectPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationChannel) GetSubjectPrefixOk() (*string, bool) {
	if o == nil || o.SubjectPrefix == nil {
		return nil, false
	}
	return o.SubjectPrefix, true
}

// HasSubjectPrefix returns a boolean if a field has been set.
func (o *EmailNotificationChannel) HasSubjectPrefix() bool {
	if o != nil && o.SubjectPrefix != nil {
		return true
	}

	return false
}

// SetSubjectPrefix gets a reference to the given string and assigns it to the SubjectPrefix field.
func (o *EmailNotificationChannel) SetSubjectPrefix(v string) {
	o.SubjectPrefix = &v
}

// GetType returns the Type field value
func (o *EmailNotificationChannel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationChannel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EmailNotificationChannel) SetType(v string) {
	o.Type = v
}

func (o EmailNotificationChannel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.NotificationConfigurationId != nil {
		toSerialize["notificationConfigurationId"] = o.NotificationConfigurationId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["to"] = o.To
	}
	if true {
		toSerialize["cc"] = o.Cc
	}
	if o.SubjectPrefix != nil {
		toSerialize["subjectPrefix"] = o.SubjectPrefix
	}
	if true {
		toSerialize["_type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableEmailNotificationChannel struct {
	value *EmailNotificationChannel
	isSet bool
}

func (v NullableEmailNotificationChannel) Get() *EmailNotificationChannel {
	return v.value
}

func (v *NullableEmailNotificationChannel) Set(val *EmailNotificationChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailNotificationChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailNotificationChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailNotificationChannel(val *EmailNotificationChannel) *NullableEmailNotificationChannel {
	return &NullableEmailNotificationChannel{value: val, isSet: true}
}

func (v NullableEmailNotificationChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailNotificationChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
