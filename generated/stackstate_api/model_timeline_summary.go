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

// TimelineSummary struct for TimelineSummary
type TimelineSummary struct {
	Buckets       []TimelineSummaryEventBucket  `json:"buckets"`
	HealthHistory []TimelineSummaryHealthChange `json:"healthHistory"`
	FromTime      int64                         `json:"fromTime"`
}

// NewTimelineSummary instantiates a new TimelineSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineSummary(buckets []TimelineSummaryEventBucket, healthHistory []TimelineSummaryHealthChange, fromTime int64) *TimelineSummary {
	this := TimelineSummary{}
	this.Buckets = buckets
	this.HealthHistory = healthHistory
	this.FromTime = fromTime
	return &this
}

// NewTimelineSummaryWithDefaults instantiates a new TimelineSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineSummaryWithDefaults() *TimelineSummary {
	this := TimelineSummary{}
	return &this
}

// GetBuckets returns the Buckets field value
func (o *TimelineSummary) GetBuckets() []TimelineSummaryEventBucket {
	if o == nil {
		var ret []TimelineSummaryEventBucket
		return ret
	}

	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value
// and a boolean to check if the value has been set.
func (o *TimelineSummary) GetBucketsOk() ([]TimelineSummaryEventBucket, bool) {
	if o == nil {
		return nil, false
	}
	return o.Buckets, true
}

// SetBuckets sets field value
func (o *TimelineSummary) SetBuckets(v []TimelineSummaryEventBucket) {
	o.Buckets = v
}

// GetHealthHistory returns the HealthHistory field value
func (o *TimelineSummary) GetHealthHistory() []TimelineSummaryHealthChange {
	if o == nil {
		var ret []TimelineSummaryHealthChange
		return ret
	}

	return o.HealthHistory
}

// GetHealthHistoryOk returns a tuple with the HealthHistory field value
// and a boolean to check if the value has been set.
func (o *TimelineSummary) GetHealthHistoryOk() ([]TimelineSummaryHealthChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.HealthHistory, true
}

// SetHealthHistory sets field value
func (o *TimelineSummary) SetHealthHistory(v []TimelineSummaryHealthChange) {
	o.HealthHistory = v
}

// GetFromTime returns the FromTime field value
func (o *TimelineSummary) GetFromTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FromTime
}

// GetFromTimeOk returns a tuple with the FromTime field value
// and a boolean to check if the value has been set.
func (o *TimelineSummary) GetFromTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromTime, true
}

// SetFromTime sets field value
func (o *TimelineSummary) SetFromTime(v int64) {
	o.FromTime = v
}

func (o TimelineSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["buckets"] = o.Buckets
	}
	if true {
		toSerialize["healthHistory"] = o.HealthHistory
	}
	if true {
		toSerialize["fromTime"] = o.FromTime
	}
	return json.Marshal(toSerialize)
}

type NullableTimelineSummary struct {
	value *TimelineSummary
	isSet bool
}

func (v NullableTimelineSummary) Get() *TimelineSummary {
	return v.value
}

func (v *NullableTimelineSummary) Set(val *TimelineSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineSummary(val *TimelineSummary) *NullableTimelineSummary {
	return &NullableTimelineSummary{value: val, isSet: true}
}

func (v NullableTimelineSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
