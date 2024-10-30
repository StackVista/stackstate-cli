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

// TraceIdentifier Identifier for a \"service trace\", the branch of an OTEL trace that corresponds to a single service entry span
type TraceIdentifier struct {
	// The unique identifier for the trace, all spans of the same trace share the same trace_id
	TraceId string `json:"traceId"`
	// The unique identifier of the span within the trace
	SpanId string `json:"spanId"`
}

// NewTraceIdentifier instantiates a new TraceIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceIdentifier(traceId string, spanId string) *TraceIdentifier {
	this := TraceIdentifier{}
	this.TraceId = traceId
	this.SpanId = spanId
	return &this
}

// NewTraceIdentifierWithDefaults instantiates a new TraceIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceIdentifierWithDefaults() *TraceIdentifier {
	this := TraceIdentifier{}
	return &this
}

// GetTraceId returns the TraceId field value
func (o *TraceIdentifier) GetTraceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value
// and a boolean to check if the value has been set.
func (o *TraceIdentifier) GetTraceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceId, true
}

// SetTraceId sets field value
func (o *TraceIdentifier) SetTraceId(v string) {
	o.TraceId = v
}

// GetSpanId returns the SpanId field value
func (o *TraceIdentifier) GetSpanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value
// and a boolean to check if the value has been set.
func (o *TraceIdentifier) GetSpanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanId, true
}

// SetSpanId sets field value
func (o *TraceIdentifier) SetSpanId(v string) {
	o.SpanId = v
}

func (o TraceIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["traceId"] = o.TraceId
	}
	if true {
		toSerialize["spanId"] = o.SpanId
	}
	return json.Marshal(toSerialize)
}

type NullableTraceIdentifier struct {
	value *TraceIdentifier
	isSet bool
}

func (v NullableTraceIdentifier) Get() *TraceIdentifier {
	return v.value
}

func (v *NullableTraceIdentifier) Set(val *TraceIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceIdentifier(val *TraceIdentifier) *NullableTraceIdentifier {
	return &NullableTraceIdentifier{value: val, isSet: true}
}

func (v NullableTraceIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}