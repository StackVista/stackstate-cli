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

// HealthSubStreamStatus struct for HealthSubStreamStatus
type HealthSubStreamStatus struct {
	Errors          []HealthStreamError             `json:"errors,omitempty"`
	Metrics         HealthStreamMetrics             `json:"metrics"`
	SubStreamState  HealthSubStreamConsistencyState `json:"subStreamState"`
	CheckStateCount int32                           `json:"checkStateCount"`
}

// NewHealthSubStreamStatus instantiates a new HealthSubStreamStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthSubStreamStatus(metrics HealthStreamMetrics, subStreamState HealthSubStreamConsistencyState, checkStateCount int32) *HealthSubStreamStatus {
	this := HealthSubStreamStatus{}
	this.Metrics = metrics
	this.SubStreamState = subStreamState
	this.CheckStateCount = checkStateCount
	return &this
}

// NewHealthSubStreamStatusWithDefaults instantiates a new HealthSubStreamStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthSubStreamStatusWithDefaults() *HealthSubStreamStatus {
	this := HealthSubStreamStatus{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *HealthSubStreamStatus) GetErrors() []HealthStreamError {
	if o == nil || o.Errors == nil {
		var ret []HealthStreamError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthSubStreamStatus) GetErrorsOk() ([]HealthStreamError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *HealthSubStreamStatus) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []HealthStreamError and assigns it to the Errors field.
func (o *HealthSubStreamStatus) SetErrors(v []HealthStreamError) {
	o.Errors = v
}

// GetMetrics returns the Metrics field value
func (o *HealthSubStreamStatus) GetMetrics() HealthStreamMetrics {
	if o == nil {
		var ret HealthStreamMetrics
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *HealthSubStreamStatus) GetMetricsOk() (*HealthStreamMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *HealthSubStreamStatus) SetMetrics(v HealthStreamMetrics) {
	o.Metrics = v
}

// GetSubStreamState returns the SubStreamState field value
func (o *HealthSubStreamStatus) GetSubStreamState() HealthSubStreamConsistencyState {
	if o == nil {
		var ret HealthSubStreamConsistencyState
		return ret
	}

	return o.SubStreamState
}

// GetSubStreamStateOk returns a tuple with the SubStreamState field value
// and a boolean to check if the value has been set.
func (o *HealthSubStreamStatus) GetSubStreamStateOk() (*HealthSubStreamConsistencyState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubStreamState, true
}

// SetSubStreamState sets field value
func (o *HealthSubStreamStatus) SetSubStreamState(v HealthSubStreamConsistencyState) {
	o.SubStreamState = v
}

// GetCheckStateCount returns the CheckStateCount field value
func (o *HealthSubStreamStatus) GetCheckStateCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CheckStateCount
}

// GetCheckStateCountOk returns a tuple with the CheckStateCount field value
// and a boolean to check if the value has been set.
func (o *HealthSubStreamStatus) GetCheckStateCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckStateCount, true
}

// SetCheckStateCount sets field value
func (o *HealthSubStreamStatus) SetCheckStateCount(v int32) {
	o.CheckStateCount = v
}

func (o HealthSubStreamStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if true {
		toSerialize["metrics"] = o.Metrics
	}
	if true {
		toSerialize["subStreamState"] = o.SubStreamState
	}
	if true {
		toSerialize["checkStateCount"] = o.CheckStateCount
	}
	return json.Marshal(toSerialize)
}

type NullableHealthSubStreamStatus struct {
	value *HealthSubStreamStatus
	isSet bool
}

func (v NullableHealthSubStreamStatus) Get() *HealthSubStreamStatus {
	return v.value
}

func (v *NullableHealthSubStreamStatus) Set(val *HealthSubStreamStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthSubStreamStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthSubStreamStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthSubStreamStatus(val *HealthSubStreamStatus) *NullableHealthSubStreamStatus {
	return &NullableHealthSubStreamStatus{value: val, isSet: true}
}

func (v NullableHealthSubStreamStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthSubStreamStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
