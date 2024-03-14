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

// ChannelReferenceId - struct for ChannelReferenceId
type ChannelReferenceId struct {
	OpsgenieChannelRefId *OpsgenieChannelRefId
	SlackChannelRefId    *SlackChannelRefId
	WebhookChannelRefId  *WebhookChannelRefId
}

// OpsgenieChannelRefIdAsChannelReferenceId is a convenience function that returns OpsgenieChannelRefId wrapped in ChannelReferenceId
func OpsgenieChannelRefIdAsChannelReferenceId(v *OpsgenieChannelRefId) ChannelReferenceId {
	return ChannelReferenceId{
		OpsgenieChannelRefId: v,
	}
}

// SlackChannelRefIdAsChannelReferenceId is a convenience function that returns SlackChannelRefId wrapped in ChannelReferenceId
func SlackChannelRefIdAsChannelReferenceId(v *SlackChannelRefId) ChannelReferenceId {
	return ChannelReferenceId{
		SlackChannelRefId: v,
	}
}

// WebhookChannelRefIdAsChannelReferenceId is a convenience function that returns WebhookChannelRefId wrapped in ChannelReferenceId
func WebhookChannelRefIdAsChannelReferenceId(v *WebhookChannelRefId) ChannelReferenceId {
	return ChannelReferenceId{
		WebhookChannelRefId: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ChannelReferenceId) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'OpsgenieChannelRefId'
	if jsonDict["_type"] == "OpsgenieChannelRefId" {
		// try to unmarshal JSON data into OpsgenieChannelRefId
		err = json.Unmarshal(data, &dst.OpsgenieChannelRefId)
		if err == nil {
			return nil // data stored in dst.OpsgenieChannelRefId, return on the first match
		} else {
			dst.OpsgenieChannelRefId = nil
			return fmt.Errorf("Failed to unmarshal ChannelReferenceId as OpsgenieChannelRefId: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SlackChannelRefId'
	if jsonDict["_type"] == "SlackChannelRefId" {
		// try to unmarshal JSON data into SlackChannelRefId
		err = json.Unmarshal(data, &dst.SlackChannelRefId)
		if err == nil {
			return nil // data stored in dst.SlackChannelRefId, return on the first match
		} else {
			dst.SlackChannelRefId = nil
			return fmt.Errorf("Failed to unmarshal ChannelReferenceId as SlackChannelRefId: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WebhookChannelRefId'
	if jsonDict["_type"] == "WebhookChannelRefId" {
		// try to unmarshal JSON data into WebhookChannelRefId
		err = json.Unmarshal(data, &dst.WebhookChannelRefId)
		if err == nil {
			return nil // data stored in dst.WebhookChannelRefId, return on the first match
		} else {
			dst.WebhookChannelRefId = nil
			return fmt.Errorf("Failed to unmarshal ChannelReferenceId as WebhookChannelRefId: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChannelReferenceId) MarshalJSON() ([]byte, error) {
	if src.OpsgenieChannelRefId != nil {
		return json.Marshal(&src.OpsgenieChannelRefId)
	}

	if src.SlackChannelRefId != nil {
		return json.Marshal(&src.SlackChannelRefId)
	}

	if src.WebhookChannelRefId != nil {
		return json.Marshal(&src.WebhookChannelRefId)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ChannelReferenceId) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OpsgenieChannelRefId != nil {
		return obj.OpsgenieChannelRefId
	}

	if obj.SlackChannelRefId != nil {
		return obj.SlackChannelRefId
	}

	if obj.WebhookChannelRefId != nil {
		return obj.WebhookChannelRefId
	}

	// all schemas are nil
	return nil
}

type NullableChannelReferenceId struct {
	value *ChannelReferenceId
	isSet bool
}

func (v NullableChannelReferenceId) Get() *ChannelReferenceId {
	return v.value
}

func (v *NullableChannelReferenceId) Set(val *ChannelReferenceId) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelReferenceId) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelReferenceId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelReferenceId(val *ChannelReferenceId) *NullableChannelReferenceId {
	return &NullableChannelReferenceId{value: val, isSet: true}
}

func (v NullableChannelReferenceId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelReferenceId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
