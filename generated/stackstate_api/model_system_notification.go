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

// SystemNotification struct for SystemNotification
type SystemNotification struct {
	NotificationId          string                     `json:"notificationId"`
	Title                   string                     `json:"title"`
	Severity                SystemNotificationSeverity `json:"severity"`
	NotificationTimeEpochMs int64                      `json:"notificationTimeEpochMs"`
	Content                 string                     `json:"content"`
	Toast                   bool                       `json:"toast"`
}

// NewSystemNotification instantiates a new SystemNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemNotification(notificationId string, title string, severity SystemNotificationSeverity, notificationTimeEpochMs int64, content string, toast bool) *SystemNotification {
	this := SystemNotification{}
	this.NotificationId = notificationId
	this.Title = title
	this.Severity = severity
	this.NotificationTimeEpochMs = notificationTimeEpochMs
	this.Content = content
	this.Toast = toast
	return &this
}

// NewSystemNotificationWithDefaults instantiates a new SystemNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemNotificationWithDefaults() *SystemNotification {
	this := SystemNotification{}
	return &this
}

// GetNotificationId returns the NotificationId field value
func (o *SystemNotification) GetNotificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value
// and a boolean to check if the value has been set.
func (o *SystemNotification) GetNotificationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationId, true
}

// SetNotificationId sets field value
func (o *SystemNotification) SetNotificationId(v string) {
	o.NotificationId = v
}

// GetTitle returns the Title field value
func (o *SystemNotification) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SystemNotification) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *SystemNotification) SetTitle(v string) {
	o.Title = v
}

// GetSeverity returns the Severity field value
func (o *SystemNotification) GetSeverity() SystemNotificationSeverity {
	if o == nil {
		var ret SystemNotificationSeverity
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *SystemNotification) GetSeverityOk() (*SystemNotificationSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *SystemNotification) SetSeverity(v SystemNotificationSeverity) {
	o.Severity = v
}

// GetNotificationTimeEpochMs returns the NotificationTimeEpochMs field value
func (o *SystemNotification) GetNotificationTimeEpochMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NotificationTimeEpochMs
}

// GetNotificationTimeEpochMsOk returns a tuple with the NotificationTimeEpochMs field value
// and a boolean to check if the value has been set.
func (o *SystemNotification) GetNotificationTimeEpochMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationTimeEpochMs, true
}

// SetNotificationTimeEpochMs sets field value
func (o *SystemNotification) SetNotificationTimeEpochMs(v int64) {
	o.NotificationTimeEpochMs = v
}

// GetContent returns the Content field value
func (o *SystemNotification) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *SystemNotification) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *SystemNotification) SetContent(v string) {
	o.Content = v
}

// GetToast returns the Toast field value
func (o *SystemNotification) GetToast() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Toast
}

// GetToastOk returns a tuple with the Toast field value
// and a boolean to check if the value has been set.
func (o *SystemNotification) GetToastOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Toast, true
}

// SetToast sets field value
func (o *SystemNotification) SetToast(v bool) {
	o.Toast = v
}

func (o SystemNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["notificationId"] = o.NotificationId
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["severity"] = o.Severity
	}
	if true {
		toSerialize["notificationTimeEpochMs"] = o.NotificationTimeEpochMs
	}
	if true {
		toSerialize["content"] = o.Content
	}
	if true {
		toSerialize["toast"] = o.Toast
	}
	return json.Marshal(toSerialize)
}

type NullableSystemNotification struct {
	value *SystemNotification
	isSet bool
}

func (v NullableSystemNotification) Get() *SystemNotification {
	return v.value
}

func (v *NullableSystemNotification) Set(val *SystemNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemNotification(val *SystemNotification) *NullableSystemNotification {
	return &NullableSystemNotification{value: val, isSet: true}
}

func (v NullableSystemNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
