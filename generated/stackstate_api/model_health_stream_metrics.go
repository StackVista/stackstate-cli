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

// HealthStreamMetrics struct for HealthStreamMetrics
type HealthStreamMetrics struct {
	BucketSizeSeconds int32               `json:"bucketSizeSeconds"`
	LatencySeconds    []MetricBucketValue `json:"latencySeconds,omitempty"`
	MessagePerSecond  []MetricBucketValue `json:"messagePerSecond,omitempty"`
	CreatesPerSecond  []MetricBucketValue `json:"createsPerSecond,omitempty"`
	UpdatesPerSecond  []MetricBucketValue `json:"updatesPerSecond,omitempty"`
	DeletesPerSecond  []MetricBucketValue `json:"deletesPerSecond,omitempty"`
}

// NewHealthStreamMetrics instantiates a new HealthStreamMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthStreamMetrics(bucketSizeSeconds int32) *HealthStreamMetrics {
	this := HealthStreamMetrics{}
	this.BucketSizeSeconds = bucketSizeSeconds
	return &this
}

// NewHealthStreamMetricsWithDefaults instantiates a new HealthStreamMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthStreamMetricsWithDefaults() *HealthStreamMetrics {
	this := HealthStreamMetrics{}
	return &this
}

// GetBucketSizeSeconds returns the BucketSizeSeconds field value
func (o *HealthStreamMetrics) GetBucketSizeSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BucketSizeSeconds
}

// GetBucketSizeSecondsOk returns a tuple with the BucketSizeSeconds field value
// and a boolean to check if the value has been set.
func (o *HealthStreamMetrics) GetBucketSizeSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketSizeSeconds, true
}

// SetBucketSizeSeconds sets field value
func (o *HealthStreamMetrics) SetBucketSizeSeconds(v int32) {
	o.BucketSizeSeconds = v
}

// GetLatencySeconds returns the LatencySeconds field value if set, zero value otherwise.
func (o *HealthStreamMetrics) GetLatencySeconds() []MetricBucketValue {
	if o == nil || o.LatencySeconds == nil {
		var ret []MetricBucketValue
		return ret
	}
	return o.LatencySeconds
}

// GetLatencySecondsOk returns a tuple with the LatencySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthStreamMetrics) GetLatencySecondsOk() ([]MetricBucketValue, bool) {
	if o == nil || o.LatencySeconds == nil {
		return nil, false
	}
	return o.LatencySeconds, true
}

// HasLatencySeconds returns a boolean if a field has been set.
func (o *HealthStreamMetrics) HasLatencySeconds() bool {
	if o != nil && o.LatencySeconds != nil {
		return true
	}

	return false
}

// SetLatencySeconds gets a reference to the given []MetricBucketValue and assigns it to the LatencySeconds field.
func (o *HealthStreamMetrics) SetLatencySeconds(v []MetricBucketValue) {
	o.LatencySeconds = v
}

// GetMessagePerSecond returns the MessagePerSecond field value if set, zero value otherwise.
func (o *HealthStreamMetrics) GetMessagePerSecond() []MetricBucketValue {
	if o == nil || o.MessagePerSecond == nil {
		var ret []MetricBucketValue
		return ret
	}
	return o.MessagePerSecond
}

// GetMessagePerSecondOk returns a tuple with the MessagePerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthStreamMetrics) GetMessagePerSecondOk() ([]MetricBucketValue, bool) {
	if o == nil || o.MessagePerSecond == nil {
		return nil, false
	}
	return o.MessagePerSecond, true
}

// HasMessagePerSecond returns a boolean if a field has been set.
func (o *HealthStreamMetrics) HasMessagePerSecond() bool {
	if o != nil && o.MessagePerSecond != nil {
		return true
	}

	return false
}

// SetMessagePerSecond gets a reference to the given []MetricBucketValue and assigns it to the MessagePerSecond field.
func (o *HealthStreamMetrics) SetMessagePerSecond(v []MetricBucketValue) {
	o.MessagePerSecond = v
}

// GetCreatesPerSecond returns the CreatesPerSecond field value if set, zero value otherwise.
func (o *HealthStreamMetrics) GetCreatesPerSecond() []MetricBucketValue {
	if o == nil || o.CreatesPerSecond == nil {
		var ret []MetricBucketValue
		return ret
	}
	return o.CreatesPerSecond
}

// GetCreatesPerSecondOk returns a tuple with the CreatesPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthStreamMetrics) GetCreatesPerSecondOk() ([]MetricBucketValue, bool) {
	if o == nil || o.CreatesPerSecond == nil {
		return nil, false
	}
	return o.CreatesPerSecond, true
}

// HasCreatesPerSecond returns a boolean if a field has been set.
func (o *HealthStreamMetrics) HasCreatesPerSecond() bool {
	if o != nil && o.CreatesPerSecond != nil {
		return true
	}

	return false
}

// SetCreatesPerSecond gets a reference to the given []MetricBucketValue and assigns it to the CreatesPerSecond field.
func (o *HealthStreamMetrics) SetCreatesPerSecond(v []MetricBucketValue) {
	o.CreatesPerSecond = v
}

// GetUpdatesPerSecond returns the UpdatesPerSecond field value if set, zero value otherwise.
func (o *HealthStreamMetrics) GetUpdatesPerSecond() []MetricBucketValue {
	if o == nil || o.UpdatesPerSecond == nil {
		var ret []MetricBucketValue
		return ret
	}
	return o.UpdatesPerSecond
}

// GetUpdatesPerSecondOk returns a tuple with the UpdatesPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthStreamMetrics) GetUpdatesPerSecondOk() ([]MetricBucketValue, bool) {
	if o == nil || o.UpdatesPerSecond == nil {
		return nil, false
	}
	return o.UpdatesPerSecond, true
}

// HasUpdatesPerSecond returns a boolean if a field has been set.
func (o *HealthStreamMetrics) HasUpdatesPerSecond() bool {
	if o != nil && o.UpdatesPerSecond != nil {
		return true
	}

	return false
}

// SetUpdatesPerSecond gets a reference to the given []MetricBucketValue and assigns it to the UpdatesPerSecond field.
func (o *HealthStreamMetrics) SetUpdatesPerSecond(v []MetricBucketValue) {
	o.UpdatesPerSecond = v
}

// GetDeletesPerSecond returns the DeletesPerSecond field value if set, zero value otherwise.
func (o *HealthStreamMetrics) GetDeletesPerSecond() []MetricBucketValue {
	if o == nil || o.DeletesPerSecond == nil {
		var ret []MetricBucketValue
		return ret
	}
	return o.DeletesPerSecond
}

// GetDeletesPerSecondOk returns a tuple with the DeletesPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthStreamMetrics) GetDeletesPerSecondOk() ([]MetricBucketValue, bool) {
	if o == nil || o.DeletesPerSecond == nil {
		return nil, false
	}
	return o.DeletesPerSecond, true
}

// HasDeletesPerSecond returns a boolean if a field has been set.
func (o *HealthStreamMetrics) HasDeletesPerSecond() bool {
	if o != nil && o.DeletesPerSecond != nil {
		return true
	}

	return false
}

// SetDeletesPerSecond gets a reference to the given []MetricBucketValue and assigns it to the DeletesPerSecond field.
func (o *HealthStreamMetrics) SetDeletesPerSecond(v []MetricBucketValue) {
	o.DeletesPerSecond = v
}

func (o HealthStreamMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bucketSizeSeconds"] = o.BucketSizeSeconds
	}
	if o.LatencySeconds != nil {
		toSerialize["latencySeconds"] = o.LatencySeconds
	}
	if o.MessagePerSecond != nil {
		toSerialize["messagePerSecond"] = o.MessagePerSecond
	}
	if o.CreatesPerSecond != nil {
		toSerialize["createsPerSecond"] = o.CreatesPerSecond
	}
	if o.UpdatesPerSecond != nil {
		toSerialize["updatesPerSecond"] = o.UpdatesPerSecond
	}
	if o.DeletesPerSecond != nil {
		toSerialize["deletesPerSecond"] = o.DeletesPerSecond
	}
	return json.Marshal(toSerialize)
}

type NullableHealthStreamMetrics struct {
	value *HealthStreamMetrics
	isSet bool
}

func (v NullableHealthStreamMetrics) Get() *HealthStreamMetrics {
	return v.value
}

func (v *NullableHealthStreamMetrics) Set(val *HealthStreamMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthStreamMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthStreamMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthStreamMetrics(val *HealthStreamMetrics) *NullableHealthStreamMetrics {
	return &NullableHealthStreamMetrics{value: val, isSet: true}
}

func (v NullableHealthStreamMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthStreamMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
