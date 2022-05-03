/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
)

// EventNotFoundError struct for EventNotFoundError
type EventNotFoundError struct {
	Type string `json:"_type"`
	EventId string `json:"eventId"`
}

// NewEventNotFoundError instantiates a new EventNotFoundError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventNotFoundError(type_ string, eventId string) *EventNotFoundError {
	this := EventNotFoundError{}
	this.Type = type_
	this.EventId = eventId
	return &this
}

// NewEventNotFoundErrorWithDefaults instantiates a new EventNotFoundError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventNotFoundErrorWithDefaults() *EventNotFoundError {
	this := EventNotFoundError{}
	return &this
}

// GetType returns the Type field value
func (o *EventNotFoundError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventNotFoundError) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventNotFoundError) SetType(v string) {
	o.Type = v
}

// GetEventId returns the EventId field value
func (o *EventNotFoundError) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *EventNotFoundError) GetEventIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *EventNotFoundError) SetEventId(v string) {
	o.EventId = v
}

func (o EventNotFoundError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["eventId"] = o.EventId
	}
	return json.Marshal(toSerialize)
}

type NullableEventNotFoundError struct {
	value *EventNotFoundError
	isSet bool
}

func (v NullableEventNotFoundError) Get() *EventNotFoundError {
	return v.value
}

func (v *NullableEventNotFoundError) Set(val *EventNotFoundError) {
	v.value = val
	v.isSet = true
}

func (v NullableEventNotFoundError) IsSet() bool {
	return v.isSet
}

func (v *NullableEventNotFoundError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventNotFoundError(val *EventNotFoundError) *NullableEventNotFoundError {
	return &NullableEventNotFoundError{value: val, isSet: true}
}

func (v NullableEventNotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventNotFoundError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


