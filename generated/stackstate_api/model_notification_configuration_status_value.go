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
	"fmt"
)

// NotificationConfigurationStatusValue the model 'NotificationConfigurationStatusValue'
type NotificationConfigurationStatusValue string

// List of NotificationConfigurationStatusValue
const (
	NOTIFICATIONCONFIGURATIONSTATUSVALUE_ENABLED  NotificationConfigurationStatusValue = "ENABLED"
	NOTIFICATIONCONFIGURATIONSTATUSVALUE_DISABLED NotificationConfigurationStatusValue = "DISABLED"
)

// All allowed values of NotificationConfigurationStatusValue enum
var AllowedNotificationConfigurationStatusValueEnumValues = []NotificationConfigurationStatusValue{
	"ENABLED",
	"DISABLED",
}

func (v *NotificationConfigurationStatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationConfigurationStatusValue(value)
	for _, existing := range AllowedNotificationConfigurationStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationConfigurationStatusValue", value)
}

// NewNotificationConfigurationStatusValueFromValue returns a pointer to a valid NotificationConfigurationStatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationConfigurationStatusValueFromValue(v string) (*NotificationConfigurationStatusValue, error) {
	ev := NotificationConfigurationStatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationConfigurationStatusValue: valid values are %v", v, AllowedNotificationConfigurationStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationConfigurationStatusValue) IsValid() bool {
	for _, existing := range AllowedNotificationConfigurationStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationConfigurationStatusValue value
func (v NotificationConfigurationStatusValue) Ptr() *NotificationConfigurationStatusValue {
	return &v
}

type NullableNotificationConfigurationStatusValue struct {
	value *NotificationConfigurationStatusValue
	isSet bool
}

func (v NullableNotificationConfigurationStatusValue) Get() *NotificationConfigurationStatusValue {
	return v.value
}

func (v *NullableNotificationConfigurationStatusValue) Set(val *NotificationConfigurationStatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationConfigurationStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationConfigurationStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationConfigurationStatusValue(val *NotificationConfigurationStatusValue) *NullableNotificationConfigurationStatusValue {
	return &NullableNotificationConfigurationStatusValue{value: val, isSet: true}
}

func (v NullableNotificationConfigurationStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationConfigurationStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}