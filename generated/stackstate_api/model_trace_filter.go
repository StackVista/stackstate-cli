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

// TraceFilter Filter traces that have a span that matches the primary filter and a span that matches the secondary filter. These could be the same, or two different spans of the same trace.
type TraceFilter struct {
	PrimarySpanFilter   SpanFilter  `json:"primarySpanFilter"`
	SecondarySpanFilter *SpanFilter `json:"secondarySpanFilter,omitempty"`
}

// NewTraceFilter instantiates a new TraceFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceFilter(primarySpanFilter SpanFilter) *TraceFilter {
	this := TraceFilter{}
	this.PrimarySpanFilter = primarySpanFilter
	return &this
}

// NewTraceFilterWithDefaults instantiates a new TraceFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceFilterWithDefaults() *TraceFilter {
	this := TraceFilter{}
	return &this
}

// GetPrimarySpanFilter returns the PrimarySpanFilter field value
func (o *TraceFilter) GetPrimarySpanFilter() SpanFilter {
	if o == nil {
		var ret SpanFilter
		return ret
	}

	return o.PrimarySpanFilter
}

// GetPrimarySpanFilterOk returns a tuple with the PrimarySpanFilter field value
// and a boolean to check if the value has been set.
func (o *TraceFilter) GetPrimarySpanFilterOk() (*SpanFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimarySpanFilter, true
}

// SetPrimarySpanFilter sets field value
func (o *TraceFilter) SetPrimarySpanFilter(v SpanFilter) {
	o.PrimarySpanFilter = v
}

// GetSecondarySpanFilter returns the SecondarySpanFilter field value if set, zero value otherwise.
func (o *TraceFilter) GetSecondarySpanFilter() SpanFilter {
	if o == nil || o.SecondarySpanFilter == nil {
		var ret SpanFilter
		return ret
	}
	return *o.SecondarySpanFilter
}

// GetSecondarySpanFilterOk returns a tuple with the SecondarySpanFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceFilter) GetSecondarySpanFilterOk() (*SpanFilter, bool) {
	if o == nil || o.SecondarySpanFilter == nil {
		return nil, false
	}
	return o.SecondarySpanFilter, true
}

// HasSecondarySpanFilter returns a boolean if a field has been set.
func (o *TraceFilter) HasSecondarySpanFilter() bool {
	if o != nil && o.SecondarySpanFilter != nil {
		return true
	}

	return false
}

// SetSecondarySpanFilter gets a reference to the given SpanFilter and assigns it to the SecondarySpanFilter field.
func (o *TraceFilter) SetSecondarySpanFilter(v SpanFilter) {
	o.SecondarySpanFilter = &v
}

func (o TraceFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["primarySpanFilter"] = o.PrimarySpanFilter
	}
	if o.SecondarySpanFilter != nil {
		toSerialize["secondarySpanFilter"] = o.SecondarySpanFilter
	}
	return json.Marshal(toSerialize)
}

type NullableTraceFilter struct {
	value *TraceFilter
	isSet bool
}

func (v NullableTraceFilter) Get() *TraceFilter {
	return v.value
}

func (v *NullableTraceFilter) Set(val *TraceFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceFilter(val *TraceFilter) *NullableTraceFilter {
	return &NullableTraceFilter{value: val, isSet: true}
}

func (v NullableTraceFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
