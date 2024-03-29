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

// InstantNanoPrecision A custom representation for a date/time that needs better than milliseconds precision. Simply using nanoseconds since epoch results in integers that are too big to be represented correctly in Javascript (which is limited to 2^53-1). Instead this uses the standard representation of milliseconds since epoch with a nanosecond offset. Calculate nanoseconds since epoch like this `nanosSinceEpoch = timestamp * 1000000 + offsetNanos`.
type InstantNanoPrecision struct {
	// Date/time representation in milliseconds since epoch (1970-01-01 00:00:00)
	Timestamp int32 `json:"timestamp"`
	// Offset in nanoseconds (relative to the timestamp). Especially useful when comparing start and/or end times of spans, for example when rendering a trace chart.
	OffsetNanos int32 `json:"offsetNanos"`
}

// NewInstantNanoPrecision instantiates a new InstantNanoPrecision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstantNanoPrecision(timestamp int32, offsetNanos int32) *InstantNanoPrecision {
	this := InstantNanoPrecision{}
	this.Timestamp = timestamp
	this.OffsetNanos = offsetNanos
	return &this
}

// NewInstantNanoPrecisionWithDefaults instantiates a new InstantNanoPrecision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstantNanoPrecisionWithDefaults() *InstantNanoPrecision {
	this := InstantNanoPrecision{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *InstantNanoPrecision) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *InstantNanoPrecision) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *InstantNanoPrecision) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetOffsetNanos returns the OffsetNanos field value
func (o *InstantNanoPrecision) GetOffsetNanos() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OffsetNanos
}

// GetOffsetNanosOk returns a tuple with the OffsetNanos field value
// and a boolean to check if the value has been set.
func (o *InstantNanoPrecision) GetOffsetNanosOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OffsetNanos, true
}

// SetOffsetNanos sets field value
func (o *InstantNanoPrecision) SetOffsetNanos(v int32) {
	o.OffsetNanos = v
}

func (o InstantNanoPrecision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["offsetNanos"] = o.OffsetNanos
	}
	return json.Marshal(toSerialize)
}

type NullableInstantNanoPrecision struct {
	value *InstantNanoPrecision
	isSet bool
}

func (v NullableInstantNanoPrecision) Get() *InstantNanoPrecision {
	return v.value
}

func (v *NullableInstantNanoPrecision) Set(val *InstantNanoPrecision) {
	v.value = val
	v.isSet = true
}

func (v NullableInstantNanoPrecision) IsSet() bool {
	return v.isSet
}

func (v *NullableInstantNanoPrecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstantNanoPrecision(val *InstantNanoPrecision) *NullableInstantNanoPrecision {
	return &NullableInstantNanoPrecision{value: val, isSet: true}
}

func (v NullableInstantNanoPrecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstantNanoPrecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
