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

// EventsHistogramRequest struct for EventsHistogramRequest
type EventsHistogramRequest struct {
	TopologyTimestamp          *int32          `json:"topologyTimestamp,omitempty"`
	StartTimestampMs           int32           `json:"startTimestampMs"`
	EndTimestampMs             int32           `json:"endTimestampMs"`
	TopologyQuery              string          `json:"topologyQuery"`
	IncludeConnectedComponents *bool           `json:"includeConnectedComponents,omitempty"`
	EventTypes                 []string        `json:"eventTypes,omitempty"`
	EventCategories            []EventCategory `json:"eventCategories,omitempty"`
	HistogramBucketCount       int32           `json:"histogramBucketCount"`
}

// NewEventsHistogramRequest instantiates a new EventsHistogramRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsHistogramRequest(startTimestampMs int32, endTimestampMs int32, topologyQuery string, histogramBucketCount int32) *EventsHistogramRequest {
	this := EventsHistogramRequest{}
	this.StartTimestampMs = startTimestampMs
	this.EndTimestampMs = endTimestampMs
	this.TopologyQuery = topologyQuery
	this.HistogramBucketCount = histogramBucketCount
	return &this
}

// NewEventsHistogramRequestWithDefaults instantiates a new EventsHistogramRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsHistogramRequestWithDefaults() *EventsHistogramRequest {
	this := EventsHistogramRequest{}
	return &this
}

// GetTopologyTimestamp returns the TopologyTimestamp field value if set, zero value otherwise.
func (o *EventsHistogramRequest) GetTopologyTimestamp() int32 {
	if o == nil || o.TopologyTimestamp == nil {
		var ret int32
		return ret
	}
	return *o.TopologyTimestamp
}

// GetTopologyTimestampOk returns a tuple with the TopologyTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsHistogramRequest) GetTopologyTimestampOk() (*int32, bool) {
	if o == nil || o.TopologyTimestamp == nil {
		return nil, false
	}
	return o.TopologyTimestamp, true
}

// HasTopologyTimestamp returns a boolean if a field has been set.
func (o *EventsHistogramRequest) HasTopologyTimestamp() bool {
	if o != nil && o.TopologyTimestamp != nil {
		return true
	}

	return false
}

// SetTopologyTimestamp gets a reference to the given int32 and assigns it to the TopologyTimestamp field.
func (o *EventsHistogramRequest) SetTopologyTimestamp(v int32) {
	o.TopologyTimestamp = &v
}

// GetStartTimestampMs returns the StartTimestampMs field value
func (o *EventsHistogramRequest) GetStartTimestampMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartTimestampMs
}

// GetStartTimestampMsOk returns a tuple with the StartTimestampMs field value
// and a boolean to check if the value has been set.
func (o *EventsHistogramRequest) GetStartTimestampMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTimestampMs, true
}

// SetStartTimestampMs sets field value
func (o *EventsHistogramRequest) SetStartTimestampMs(v int32) {
	o.StartTimestampMs = v
}

// GetEndTimestampMs returns the EndTimestampMs field value
func (o *EventsHistogramRequest) GetEndTimestampMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EndTimestampMs
}

// GetEndTimestampMsOk returns a tuple with the EndTimestampMs field value
// and a boolean to check if the value has been set.
func (o *EventsHistogramRequest) GetEndTimestampMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTimestampMs, true
}

// SetEndTimestampMs sets field value
func (o *EventsHistogramRequest) SetEndTimestampMs(v int32) {
	o.EndTimestampMs = v
}

// GetTopologyQuery returns the TopologyQuery field value
func (o *EventsHistogramRequest) GetTopologyQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopologyQuery
}

// GetTopologyQueryOk returns a tuple with the TopologyQuery field value
// and a boolean to check if the value has been set.
func (o *EventsHistogramRequest) GetTopologyQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopologyQuery, true
}

// SetTopologyQuery sets field value
func (o *EventsHistogramRequest) SetTopologyQuery(v string) {
	o.TopologyQuery = v
}

// GetIncludeConnectedComponents returns the IncludeConnectedComponents field value if set, zero value otherwise.
func (o *EventsHistogramRequest) GetIncludeConnectedComponents() bool {
	if o == nil || o.IncludeConnectedComponents == nil {
		var ret bool
		return ret
	}
	return *o.IncludeConnectedComponents
}

// GetIncludeConnectedComponentsOk returns a tuple with the IncludeConnectedComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsHistogramRequest) GetIncludeConnectedComponentsOk() (*bool, bool) {
	if o == nil || o.IncludeConnectedComponents == nil {
		return nil, false
	}
	return o.IncludeConnectedComponents, true
}

// HasIncludeConnectedComponents returns a boolean if a field has been set.
func (o *EventsHistogramRequest) HasIncludeConnectedComponents() bool {
	if o != nil && o.IncludeConnectedComponents != nil {
		return true
	}

	return false
}

// SetIncludeConnectedComponents gets a reference to the given bool and assigns it to the IncludeConnectedComponents field.
func (o *EventsHistogramRequest) SetIncludeConnectedComponents(v bool) {
	o.IncludeConnectedComponents = &v
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *EventsHistogramRequest) GetEventTypes() []string {
	if o == nil || o.EventTypes == nil {
		var ret []string
		return ret
	}
	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsHistogramRequest) GetEventTypesOk() ([]string, bool) {
	if o == nil || o.EventTypes == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *EventsHistogramRequest) HasEventTypes() bool {
	if o != nil && o.EventTypes != nil {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []string and assigns it to the EventTypes field.
func (o *EventsHistogramRequest) SetEventTypes(v []string) {
	o.EventTypes = v
}

// GetEventCategories returns the EventCategories field value if set, zero value otherwise.
func (o *EventsHistogramRequest) GetEventCategories() []EventCategory {
	if o == nil || o.EventCategories == nil {
		var ret []EventCategory
		return ret
	}
	return o.EventCategories
}

// GetEventCategoriesOk returns a tuple with the EventCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsHistogramRequest) GetEventCategoriesOk() ([]EventCategory, bool) {
	if o == nil || o.EventCategories == nil {
		return nil, false
	}
	return o.EventCategories, true
}

// HasEventCategories returns a boolean if a field has been set.
func (o *EventsHistogramRequest) HasEventCategories() bool {
	if o != nil && o.EventCategories != nil {
		return true
	}

	return false
}

// SetEventCategories gets a reference to the given []EventCategory and assigns it to the EventCategories field.
func (o *EventsHistogramRequest) SetEventCategories(v []EventCategory) {
	o.EventCategories = v
}

// GetHistogramBucketCount returns the HistogramBucketCount field value
func (o *EventsHistogramRequest) GetHistogramBucketCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HistogramBucketCount
}

// GetHistogramBucketCountOk returns a tuple with the HistogramBucketCount field value
// and a boolean to check if the value has been set.
func (o *EventsHistogramRequest) GetHistogramBucketCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HistogramBucketCount, true
}

// SetHistogramBucketCount sets field value
func (o *EventsHistogramRequest) SetHistogramBucketCount(v int32) {
	o.HistogramBucketCount = v
}

func (o EventsHistogramRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TopologyTimestamp != nil {
		toSerialize["topologyTimestamp"] = o.TopologyTimestamp
	}
	if true {
		toSerialize["startTimestampMs"] = o.StartTimestampMs
	}
	if true {
		toSerialize["endTimestampMs"] = o.EndTimestampMs
	}
	if true {
		toSerialize["topologyQuery"] = o.TopologyQuery
	}
	if o.IncludeConnectedComponents != nil {
		toSerialize["includeConnectedComponents"] = o.IncludeConnectedComponents
	}
	if o.EventTypes != nil {
		toSerialize["eventTypes"] = o.EventTypes
	}
	if o.EventCategories != nil {
		toSerialize["eventCategories"] = o.EventCategories
	}
	if true {
		toSerialize["histogramBucketCount"] = o.HistogramBucketCount
	}
	return json.Marshal(toSerialize)
}

type NullableEventsHistogramRequest struct {
	value *EventsHistogramRequest
	isSet bool
}

func (v NullableEventsHistogramRequest) Get() *EventsHistogramRequest {
	return v.value
}

func (v *NullableEventsHistogramRequest) Set(val *EventsHistogramRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsHistogramRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsHistogramRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsHistogramRequest(val *EventsHistogramRequest) *NullableEventsHistogramRequest {
	return &NullableEventsHistogramRequest{value: val, isSet: true}
}

func (v NullableEventsHistogramRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsHistogramRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
