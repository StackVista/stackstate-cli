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

// TraceApiInvalidDurationFilter struct for TraceApiInvalidDurationFilter
type TraceApiInvalidDurationFilter struct {
	Type    string `json:"_type"`
	Message string `json:"message"`
	From    int64  `json:"from"`
	To      int64  `json:"to"`
}

// NewTraceApiInvalidDurationFilter instantiates a new TraceApiInvalidDurationFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceApiInvalidDurationFilter(type_ string, message string, from int64, to int64) *TraceApiInvalidDurationFilter {
	this := TraceApiInvalidDurationFilter{}
	this.Type = type_
	this.Message = message
	this.From = from
	this.To = to
	return &this
}

// NewTraceApiInvalidDurationFilterWithDefaults instantiates a new TraceApiInvalidDurationFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceApiInvalidDurationFilterWithDefaults() *TraceApiInvalidDurationFilter {
	this := TraceApiInvalidDurationFilter{}
	return &this
}

// GetType returns the Type field value
func (o *TraceApiInvalidDurationFilter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TraceApiInvalidDurationFilter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TraceApiInvalidDurationFilter) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *TraceApiInvalidDurationFilter) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *TraceApiInvalidDurationFilter) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *TraceApiInvalidDurationFilter) SetMessage(v string) {
	o.Message = v
}

// GetFrom returns the From field value
func (o *TraceApiInvalidDurationFilter) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *TraceApiInvalidDurationFilter) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *TraceApiInvalidDurationFilter) SetFrom(v int64) {
	o.From = v
}

// GetTo returns the To field value
func (o *TraceApiInvalidDurationFilter) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *TraceApiInvalidDurationFilter) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *TraceApiInvalidDurationFilter) SetTo(v int64) {
	o.To = v
}

func (o TraceApiInvalidDurationFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableTraceApiInvalidDurationFilter struct {
	value *TraceApiInvalidDurationFilter
	isSet bool
}

func (v NullableTraceApiInvalidDurationFilter) Get() *TraceApiInvalidDurationFilter {
	return v.value
}

func (v *NullableTraceApiInvalidDurationFilter) Set(val *TraceApiInvalidDurationFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceApiInvalidDurationFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceApiInvalidDurationFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceApiInvalidDurationFilter(val *TraceApiInvalidDurationFilter) *NullableTraceApiInvalidDurationFilter {
	return &NullableTraceApiInvalidDurationFilter{value: val, isSet: true}
}

func (v NullableTraceApiInvalidDurationFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceApiInvalidDurationFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
