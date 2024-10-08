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

// TraceQuery struct for TraceQuery
type TraceQuery struct {
	SpanFilter SpanFilter `json:"spanFilter"`
	// Filter traces by 1 or more attributes
	TraceAttributes map[string][]string `json:"traceAttributes"`
	SortBy          []SpanSortOption    `json:"sortBy,omitempty"`
}

// NewTraceQuery instantiates a new TraceQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceQuery(spanFilter SpanFilter, traceAttributes map[string][]string) *TraceQuery {
	this := TraceQuery{}
	this.SpanFilter = spanFilter
	this.TraceAttributes = traceAttributes
	return &this
}

// NewTraceQueryWithDefaults instantiates a new TraceQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceQueryWithDefaults() *TraceQuery {
	this := TraceQuery{}
	return &this
}

// GetSpanFilter returns the SpanFilter field value
func (o *TraceQuery) GetSpanFilter() SpanFilter {
	if o == nil {
		var ret SpanFilter
		return ret
	}

	return o.SpanFilter
}

// GetSpanFilterOk returns a tuple with the SpanFilter field value
// and a boolean to check if the value has been set.
func (o *TraceQuery) GetSpanFilterOk() (*SpanFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanFilter, true
}

// SetSpanFilter sets field value
func (o *TraceQuery) SetSpanFilter(v SpanFilter) {
	o.SpanFilter = v
}

// GetTraceAttributes returns the TraceAttributes field value
func (o *TraceQuery) GetTraceAttributes() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.TraceAttributes
}

// GetTraceAttributesOk returns a tuple with the TraceAttributes field value
// and a boolean to check if the value has been set.
func (o *TraceQuery) GetTraceAttributesOk() (*map[string][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceAttributes, true
}

// SetTraceAttributes sets field value
func (o *TraceQuery) SetTraceAttributes(v map[string][]string) {
	o.TraceAttributes = v
}

// GetSortBy returns the SortBy field value if set, zero value otherwise.
func (o *TraceQuery) GetSortBy() []SpanSortOption {
	if o == nil || o.SortBy == nil {
		var ret []SpanSortOption
		return ret
	}
	return o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceQuery) GetSortByOk() ([]SpanSortOption, bool) {
	if o == nil || o.SortBy == nil {
		return nil, false
	}
	return o.SortBy, true
}

// HasSortBy returns a boolean if a field has been set.
func (o *TraceQuery) HasSortBy() bool {
	if o != nil && o.SortBy != nil {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given []SpanSortOption and assigns it to the SortBy field.
func (o *TraceQuery) SetSortBy(v []SpanSortOption) {
	o.SortBy = v
}

func (o TraceQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["spanFilter"] = o.SpanFilter
	}
	if true {
		toSerialize["traceAttributes"] = o.TraceAttributes
	}
	if o.SortBy != nil {
		toSerialize["sortBy"] = o.SortBy
	}
	return json.Marshal(toSerialize)
}

type NullableTraceQuery struct {
	value *TraceQuery
	isSet bool
}

func (v NullableTraceQuery) Get() *TraceQuery {
	return v.value
}

func (v *NullableTraceQuery) Set(val *TraceQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceQuery(val *TraceQuery) *NullableTraceQuery {
	return &NullableTraceQuery{value: val, isSet: true}
}

func (v NullableTraceQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
