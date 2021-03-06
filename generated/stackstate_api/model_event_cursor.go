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

// EventCursor struct for EventCursor
type EventCursor struct {
	LastEventTimestampMs int64 `json:"lastEventTimestampMs"`
	LastEventId string `json:"lastEventId"`
}

// NewEventCursor instantiates a new EventCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventCursor(lastEventTimestampMs int64, lastEventId string) *EventCursor {
	this := EventCursor{}
	this.LastEventTimestampMs = lastEventTimestampMs
	this.LastEventId = lastEventId
	return &this
}

// NewEventCursorWithDefaults instantiates a new EventCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventCursorWithDefaults() *EventCursor {
	this := EventCursor{}
	return &this
}

// GetLastEventTimestampMs returns the LastEventTimestampMs field value
func (o *EventCursor) GetLastEventTimestampMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LastEventTimestampMs
}

// GetLastEventTimestampMsOk returns a tuple with the LastEventTimestampMs field value
// and a boolean to check if the value has been set.
func (o *EventCursor) GetLastEventTimestampMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEventTimestampMs, true
}

// SetLastEventTimestampMs sets field value
func (o *EventCursor) SetLastEventTimestampMs(v int64) {
	o.LastEventTimestampMs = v
}

// GetLastEventId returns the LastEventId field value
func (o *EventCursor) GetLastEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEventId
}

// GetLastEventIdOk returns a tuple with the LastEventId field value
// and a boolean to check if the value has been set.
func (o *EventCursor) GetLastEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEventId, true
}

// SetLastEventId sets field value
func (o *EventCursor) SetLastEventId(v string) {
	o.LastEventId = v
}

func (o EventCursor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["lastEventTimestampMs"] = o.LastEventTimestampMs
	}
	if true {
		toSerialize["lastEventId"] = o.LastEventId
	}
	return json.Marshal(toSerialize)
}

type NullableEventCursor struct {
	value *EventCursor
	isSet bool
}

func (v NullableEventCursor) Get() *EventCursor {
	return v.value
}

func (v *NullableEventCursor) Set(val *EventCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCursor(val *EventCursor) *NullableEventCursor {
	return &NullableEventCursor{value: val, isSet: true}
}

func (v NullableEventCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


