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

// SpanResponse The result of looking up a span. As an optimization to reduce duplication, resource attributes are inlined at the root-level of the schema.
type SpanResponse struct {
	Span Span `json:"span"`
	// Set of key/value pairs providing extra contextual information. Keys are unique.
	ResourceAttributes map[string]string `json:"resourceAttributes"`
}

// NewSpanResponse instantiates a new SpanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanResponse(span Span, resourceAttributes map[string]string) *SpanResponse {
	this := SpanResponse{}
	this.Span = span
	this.ResourceAttributes = resourceAttributes
	return &this
}

// NewSpanResponseWithDefaults instantiates a new SpanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanResponseWithDefaults() *SpanResponse {
	this := SpanResponse{}
	return &this
}

// GetSpan returns the Span field value
func (o *SpanResponse) GetSpan() Span {
	if o == nil {
		var ret Span
		return ret
	}

	return o.Span
}

// GetSpanOk returns a tuple with the Span field value
// and a boolean to check if the value has been set.
func (o *SpanResponse) GetSpanOk() (*Span, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Span, true
}

// SetSpan sets field value
func (o *SpanResponse) SetSpan(v Span) {
	o.Span = v
}

// GetResourceAttributes returns the ResourceAttributes field value
func (o *SpanResponse) GetResourceAttributes() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.ResourceAttributes
}

// GetResourceAttributesOk returns a tuple with the ResourceAttributes field value
// and a boolean to check if the value has been set.
func (o *SpanResponse) GetResourceAttributesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceAttributes, true
}

// SetResourceAttributes sets field value
func (o *SpanResponse) SetResourceAttributes(v map[string]string) {
	o.ResourceAttributes = v
}

func (o SpanResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["span"] = o.Span
	}
	if true {
		toSerialize["resourceAttributes"] = o.ResourceAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableSpanResponse struct {
	value *SpanResponse
	isSet bool
}

func (v NullableSpanResponse) Get() *SpanResponse {
	return v.value
}

func (v *NullableSpanResponse) Set(val *SpanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanResponse(val *SpanResponse) *NullableSpanResponse {
	return &NullableSpanResponse{value: val, isSet: true}
}

func (v NullableSpanResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
