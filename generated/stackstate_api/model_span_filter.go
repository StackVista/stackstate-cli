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

// SpanFilter struct for SpanFilter
type SpanFilter struct {
	// Filter spans by service name
	ServiceName []string `json:"serviceName,omitempty"`
	// Filter spans by name
	SpanName []string `json:"spanName,omitempty"`
	// Filter spans by 1 or more attributes
	Attributes *map[string][]string `json:"attributes,omitempty"`
	// Filter span by kind
	SpanKind []SpanKind `json:"spanKind,omitempty"`
	// Filter span by parent type
	SpanParentType []SpanParentType `json:"spanParentType,omitempty"`
	// Filter spans by duration >= value, in nanoseconds
	DurationFromNanos *int64 `json:"durationFromNanos,omitempty"`
	// Filter spans by duration < value, in nanoseconds
	DurationToNanos *int64 `json:"durationToNanos,omitempty"`
	// Filter spans by the StatusCode
	StatusCode []StatusCode `json:"statusCode,omitempty"`
	// Filter spans by trace id, use only this filter to get a complete trace
	TraceId []string `json:"traceId,omitempty"`
	// Filter spans by span id, use only this filter to get a single span
	SpanId []string `json:"spanId,omitempty"`
	// Filter spans by scope
	ScopeName []string `json:"scopeName,omitempty"`
	// Filter spans by scope version
	ScopeVersion []string `json:"scopeVersion,omitempty"`
}

// NewSpanFilter instantiates a new SpanFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanFilter() *SpanFilter {
	this := SpanFilter{}
	return &this
}

// NewSpanFilterWithDefaults instantiates a new SpanFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanFilterWithDefaults() *SpanFilter {
	this := SpanFilter{}
	return &this
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *SpanFilter) GetServiceName() []string {
	if o == nil || o.ServiceName == nil {
		var ret []string
		return ret
	}
	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanFilter) GetServiceNameOk() ([]string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *SpanFilter) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given []string and assigns it to the ServiceName field.
func (o *SpanFilter) SetServiceName(v []string) {
	o.ServiceName = v
}

// GetSpanName returns the SpanName field value if set, zero value otherwise.
func (o *SpanFilter) GetSpanName() []string {
	if o == nil || o.SpanName == nil {
		var ret []string
		return ret
	}
	return o.SpanName
}

// GetSpanNameOk returns a tuple with the SpanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanFilter) GetSpanNameOk() ([]string, bool) {
	if o == nil || o.SpanName == nil {
		return nil, false
	}
	return o.SpanName, true
}

// HasSpanName returns a boolean if a field has been set.
func (o *SpanFilter) HasSpanName() bool {
	if o != nil && o.SpanName != nil {
		return true
	}

	return false
}

// SetSpanName gets a reference to the given []string and assigns it to the SpanName field.
func (o *SpanFilter) SetSpanName(v []string) {
	o.SpanName = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SpanFilter) GetAttributes() map[string][]string {
	if o == nil || o.Attributes == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanFilter) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SpanFilter) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string][]string and assigns it to the Attributes field.
func (o *SpanFilter) SetAttributes(v map[string][]string) {
	o.Attributes = &v
}

// GetSpanKind returns the SpanKind field value if set, zero value otherwise.
func (o *SpanFilter) GetSpanKind() []SpanKind {
	if o == nil || o.SpanKind == nil {
		var ret []SpanKind
		return ret
	}
	return o.SpanKind
}

// GetSpanKindOk returns a tuple with the SpanKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanFilter) GetSpanKindOk() ([]SpanKind, bool) {
	if o == nil || o.SpanKind == nil {
		return nil, false
	}
	return o.SpanKind, true
}

// HasSpanKind returns a boolean if a field has been set.
func (o *SpanFilter) HasSpanKind() bool {
	if o != nil && o.SpanKind != nil {
		return true
	}

	return false
}

// SetSpanKind gets a reference to the given []SpanKind and assigns it to the SpanKind field.
func (o *SpanFilter) SetSpanKind(v []SpanKind) {
	o.SpanKind = v
}

// GetSpanParentType returns the SpanParentType field value if set, zero value otherwise.
func (o *SpanFilter) GetSpanParentType() []SpanParentType {
	if o == nil || o.SpanParentType == nil {
		var ret []SpanParentType
		return ret
	}
	return o.SpanParentType
}

// GetSpanParentTypeOk returns a tuple with the SpanParentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanFilter) GetSpanParentTypeOk() ([]SpanParentType, bool) {
	if o == nil || o.SpanParentType == nil {
		return nil, false
	}
	return o.SpanParentType, true
}

// HasSpanParentType returns a boolean if a field has been set.
func (o *SpanFilter) HasSpanParentType() bool {
	if o != nil && o.SpanParentType != nil {
		return true
	}

	return false
}

// SetSpanParentType gets a reference to the given []SpanParentType and assigns it to the SpanParentType field.
func (o *SpanFilter) SetSpanParentType(v []SpanParentType) {
	o.SpanParentType = v
}

// GetDurationFromNanos returns the DurationFromNanos field value if set, zero value otherwise.
func (o *SpanFilter) GetDurationFromNanos() int64 {
	if o == nil || o.DurationFromNanos == nil {
		var ret int64
		return ret
	}
	return *o.DurationFromNanos
}

// GetDurationFromNanosOk returns a tuple with the DurationFromNanos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanFilter) GetDurationFromNanosOk() (*int64, bool) {
	if o == nil || o.DurationFromNanos == nil {
		return nil, false
	}
	return o.DurationFromNanos, true
}

// HasDurationFromNanos returns a boolean if a field has been set.
func (o *SpanFilter) HasDurationFromNanos() bool {
	if o != nil && o.DurationFromNanos != nil {
		return true
	}

	return false
}

// SetDurationFromNanos gets a reference to the given int64 and assigns it to the DurationFromNanos field.
func (o *SpanFilter) SetDurationFromNanos(v int64) {
	o.DurationFromNanos = &v
}

// GetDurationToNanos returns the DurationToNanos field value if set, zero value otherwise.
func (o *SpanFilter) GetDurationToNanos() int64 {
	if o == nil || o.DurationToNanos == nil {
		var ret int64
		return ret
	}
	return *o.DurationToNanos
}

// GetDurationToNanosOk returns a tuple with the DurationToNanos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanFilter) GetDurationToNanosOk() (*int64, bool) {
	if o == nil || o.DurationToNanos == nil {
		return nil, false
	}
	return o.DurationToNanos, true
}

// HasDurationToNanos returns a boolean if a field has been set.
func (o *SpanFilter) HasDurationToNanos() bool {
	if o != nil && o.DurationToNanos != nil {
		return true
	}

	return false
}

// SetDurationToNanos gets a reference to the given int64 and assigns it to the DurationToNanos field.
func (o *SpanFilter) SetDurationToNanos(v int64) {
	o.DurationToNanos = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *SpanFilter) GetStatusCode() []StatusCode {
	if o == nil || o.StatusCode == nil {
		var ret []StatusCode
		return ret
	}
	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanFilter) GetStatusCodeOk() ([]StatusCode, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *SpanFilter) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given []StatusCode and assigns it to the StatusCode field.
func (o *SpanFilter) SetStatusCode(v []StatusCode) {
	o.StatusCode = v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *SpanFilter) GetTraceId() []string {
	if o == nil || o.TraceId == nil {
		var ret []string
		return ret
	}
	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanFilter) GetTraceIdOk() ([]string, bool) {
	if o == nil || o.TraceId == nil {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *SpanFilter) HasTraceId() bool {
	if o != nil && o.TraceId != nil {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given []string and assigns it to the TraceId field.
func (o *SpanFilter) SetTraceId(v []string) {
	o.TraceId = v
}

// GetSpanId returns the SpanId field value if set, zero value otherwise.
func (o *SpanFilter) GetSpanId() []string {
	if o == nil || o.SpanId == nil {
		var ret []string
		return ret
	}
	return o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanFilter) GetSpanIdOk() ([]string, bool) {
	if o == nil || o.SpanId == nil {
		return nil, false
	}
	return o.SpanId, true
}

// HasSpanId returns a boolean if a field has been set.
func (o *SpanFilter) HasSpanId() bool {
	if o != nil && o.SpanId != nil {
		return true
	}

	return false
}

// SetSpanId gets a reference to the given []string and assigns it to the SpanId field.
func (o *SpanFilter) SetSpanId(v []string) {
	o.SpanId = v
}

// GetScopeName returns the ScopeName field value if set, zero value otherwise.
func (o *SpanFilter) GetScopeName() []string {
	if o == nil || o.ScopeName == nil {
		var ret []string
		return ret
	}
	return o.ScopeName
}

// GetScopeNameOk returns a tuple with the ScopeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanFilter) GetScopeNameOk() ([]string, bool) {
	if o == nil || o.ScopeName == nil {
		return nil, false
	}
	return o.ScopeName, true
}

// HasScopeName returns a boolean if a field has been set.
func (o *SpanFilter) HasScopeName() bool {
	if o != nil && o.ScopeName != nil {
		return true
	}

	return false
}

// SetScopeName gets a reference to the given []string and assigns it to the ScopeName field.
func (o *SpanFilter) SetScopeName(v []string) {
	o.ScopeName = v
}

// GetScopeVersion returns the ScopeVersion field value if set, zero value otherwise.
func (o *SpanFilter) GetScopeVersion() []string {
	if o == nil || o.ScopeVersion == nil {
		var ret []string
		return ret
	}
	return o.ScopeVersion
}

// GetScopeVersionOk returns a tuple with the ScopeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanFilter) GetScopeVersionOk() ([]string, bool) {
	if o == nil || o.ScopeVersion == nil {
		return nil, false
	}
	return o.ScopeVersion, true
}

// HasScopeVersion returns a boolean if a field has been set.
func (o *SpanFilter) HasScopeVersion() bool {
	if o != nil && o.ScopeVersion != nil {
		return true
	}

	return false
}

// SetScopeVersion gets a reference to the given []string and assigns it to the ScopeVersion field.
func (o *SpanFilter) SetScopeVersion(v []string) {
	o.ScopeVersion = v
}

func (o SpanFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceName != nil {
		toSerialize["serviceName"] = o.ServiceName
	}
	if o.SpanName != nil {
		toSerialize["spanName"] = o.SpanName
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.SpanKind != nil {
		toSerialize["spanKind"] = o.SpanKind
	}
	if o.SpanParentType != nil {
		toSerialize["spanParentType"] = o.SpanParentType
	}
	if o.DurationFromNanos != nil {
		toSerialize["durationFromNanos"] = o.DurationFromNanos
	}
	if o.DurationToNanos != nil {
		toSerialize["durationToNanos"] = o.DurationToNanos
	}
	if o.StatusCode != nil {
		toSerialize["statusCode"] = o.StatusCode
	}
	if o.TraceId != nil {
		toSerialize["traceId"] = o.TraceId
	}
	if o.SpanId != nil {
		toSerialize["spanId"] = o.SpanId
	}
	if o.ScopeName != nil {
		toSerialize["scopeName"] = o.ScopeName
	}
	if o.ScopeVersion != nil {
		toSerialize["scopeVersion"] = o.ScopeVersion
	}
	return json.Marshal(toSerialize)
}

type NullableSpanFilter struct {
	value *SpanFilter
	isSet bool
}

func (v NullableSpanFilter) Get() *SpanFilter {
	return v.value
}

func (v *NullableSpanFilter) Set(val *SpanFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanFilter(val *SpanFilter) *NullableSpanFilter {
	return &NullableSpanFilter{value: val, isSet: true}
}

func (v NullableSpanFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
