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

// NotificationChannel - struct for NotificationChannel
type NotificationChannel struct {
	OpsgenieNotificationChannel *OpsgenieNotificationChannel
	SlackNotificationChannel    *SlackNotificationChannel
	WebhookNotificationChannel  *WebhookNotificationChannel
}

// OpsgenieNotificationChannelAsNotificationChannel is a convenience function that returns OpsgenieNotificationChannel wrapped in NotificationChannel
func OpsgenieNotificationChannelAsNotificationChannel(v *OpsgenieNotificationChannel) NotificationChannel {
	return NotificationChannel{
		OpsgenieNotificationChannel: v,
	}
}

// SlackNotificationChannelAsNotificationChannel is a convenience function that returns SlackNotificationChannel wrapped in NotificationChannel
func SlackNotificationChannelAsNotificationChannel(v *SlackNotificationChannel) NotificationChannel {
	return NotificationChannel{
		SlackNotificationChannel: v,
	}
}

// WebhookNotificationChannelAsNotificationChannel is a convenience function that returns WebhookNotificationChannel wrapped in NotificationChannel
func WebhookNotificationChannelAsNotificationChannel(v *WebhookNotificationChannel) NotificationChannel {
	return NotificationChannel{
		WebhookNotificationChannel: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NotificationChannel) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'OpsgenieNotificationChannel'
	if jsonDict["_type"] == "OpsgenieNotificationChannel" {
		// try to unmarshal JSON data into OpsgenieNotificationChannel
		err = json.Unmarshal(data, &dst.OpsgenieNotificationChannel)
		if err == nil {
			return nil // data stored in dst.OpsgenieNotificationChannel, return on the first match
		} else {
			dst.OpsgenieNotificationChannel = nil
			return fmt.Errorf("Failed to unmarshal NotificationChannel as OpsgenieNotificationChannel: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SlackNotificationChannel'
	if jsonDict["_type"] == "SlackNotificationChannel" {
		// try to unmarshal JSON data into SlackNotificationChannel
		err = json.Unmarshal(data, &dst.SlackNotificationChannel)
		if err == nil {
			return nil // data stored in dst.SlackNotificationChannel, return on the first match
		} else {
			dst.SlackNotificationChannel = nil
			return fmt.Errorf("Failed to unmarshal NotificationChannel as SlackNotificationChannel: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WebhookNotificationChannel'
	if jsonDict["_type"] == "WebhookNotificationChannel" {
		// try to unmarshal JSON data into WebhookNotificationChannel
		err = json.Unmarshal(data, &dst.WebhookNotificationChannel)
		if err == nil {
			return nil // data stored in dst.WebhookNotificationChannel, return on the first match
		} else {
			dst.WebhookNotificationChannel = nil
			return fmt.Errorf("Failed to unmarshal NotificationChannel as WebhookNotificationChannel: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NotificationChannel) MarshalJSON() ([]byte, error) {
	if src.OpsgenieNotificationChannel != nil {
		return json.Marshal(&src.OpsgenieNotificationChannel)
	}

	if src.SlackNotificationChannel != nil {
		return json.Marshal(&src.SlackNotificationChannel)
	}

	if src.WebhookNotificationChannel != nil {
		return json.Marshal(&src.WebhookNotificationChannel)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NotificationChannel) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OpsgenieNotificationChannel != nil {
		return obj.OpsgenieNotificationChannel
	}

	if obj.SlackNotificationChannel != nil {
		return obj.SlackNotificationChannel
	}

	if obj.WebhookNotificationChannel != nil {
		return obj.WebhookNotificationChannel
	}

	// all schemas are nil
	return nil
}

type NullableNotificationChannel struct {
	value *NotificationChannel
	isSet bool
}

func (v NullableNotificationChannel) Get() *NotificationChannel {
	return v.value
}

func (v *NullableNotificationChannel) Set(val *NotificationChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationChannel(val *NotificationChannel) *NullableNotificationChannel {
	return &NullableNotificationChannel{value: val, isSet: true}
}

func (v NullableNotificationChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}