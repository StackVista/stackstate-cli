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

// SimpleTrainingPeriodicity struct for SimpleTrainingPeriodicity
type SimpleTrainingPeriodicity struct {
	Type                  string `json:"_type"`
	Id                    *int64 `json:"id,omitempty"`
	FundamentalPeriod     int64  `json:"fundamentalPeriod"`
	LastUpdateTimestamp   *int64 `json:"lastUpdateTimestamp,omitempty"`
	TrainingWindowPeriods int64  `json:"trainingWindowPeriods"`
}

// NewSimpleTrainingPeriodicity instantiates a new SimpleTrainingPeriodicity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleTrainingPeriodicity(type_ string, fundamentalPeriod int64, trainingWindowPeriods int64) *SimpleTrainingPeriodicity {
	this := SimpleTrainingPeriodicity{}
	this.Type = type_
	this.FundamentalPeriod = fundamentalPeriod
	this.TrainingWindowPeriods = trainingWindowPeriods
	return &this
}

// NewSimpleTrainingPeriodicityWithDefaults instantiates a new SimpleTrainingPeriodicity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleTrainingPeriodicityWithDefaults() *SimpleTrainingPeriodicity {
	this := SimpleTrainingPeriodicity{}
	return &this
}

// GetType returns the Type field value
func (o *SimpleTrainingPeriodicity) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SimpleTrainingPeriodicity) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SimpleTrainingPeriodicity) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SimpleTrainingPeriodicity) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleTrainingPeriodicity) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SimpleTrainingPeriodicity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SimpleTrainingPeriodicity) SetId(v int64) {
	o.Id = &v
}

// GetFundamentalPeriod returns the FundamentalPeriod field value
func (o *SimpleTrainingPeriodicity) GetFundamentalPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FundamentalPeriod
}

// GetFundamentalPeriodOk returns a tuple with the FundamentalPeriod field value
// and a boolean to check if the value has been set.
func (o *SimpleTrainingPeriodicity) GetFundamentalPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundamentalPeriod, true
}

// SetFundamentalPeriod sets field value
func (o *SimpleTrainingPeriodicity) SetFundamentalPeriod(v int64) {
	o.FundamentalPeriod = v
}

// GetLastUpdateTimestamp returns the LastUpdateTimestamp field value if set, zero value otherwise.
func (o *SimpleTrainingPeriodicity) GetLastUpdateTimestamp() int64 {
	if o == nil || o.LastUpdateTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdateTimestamp
}

// GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleTrainingPeriodicity) GetLastUpdateTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdateTimestamp == nil {
		return nil, false
	}
	return o.LastUpdateTimestamp, true
}

// HasLastUpdateTimestamp returns a boolean if a field has been set.
func (o *SimpleTrainingPeriodicity) HasLastUpdateTimestamp() bool {
	if o != nil && o.LastUpdateTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdateTimestamp gets a reference to the given int64 and assigns it to the LastUpdateTimestamp field.
func (o *SimpleTrainingPeriodicity) SetLastUpdateTimestamp(v int64) {
	o.LastUpdateTimestamp = &v
}

// GetTrainingWindowPeriods returns the TrainingWindowPeriods field value
func (o *SimpleTrainingPeriodicity) GetTrainingWindowPeriods() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TrainingWindowPeriods
}

// GetTrainingWindowPeriodsOk returns a tuple with the TrainingWindowPeriods field value
// and a boolean to check if the value has been set.
func (o *SimpleTrainingPeriodicity) GetTrainingWindowPeriodsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrainingWindowPeriods, true
}

// SetTrainingWindowPeriods sets field value
func (o *SimpleTrainingPeriodicity) SetTrainingWindowPeriods(v int64) {
	o.TrainingWindowPeriods = v
}

func (o SimpleTrainingPeriodicity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["fundamentalPeriod"] = o.FundamentalPeriod
	}
	if o.LastUpdateTimestamp != nil {
		toSerialize["lastUpdateTimestamp"] = o.LastUpdateTimestamp
	}
	if true {
		toSerialize["trainingWindowPeriods"] = o.TrainingWindowPeriods
	}
	return json.Marshal(toSerialize)
}

type NullableSimpleTrainingPeriodicity struct {
	value *SimpleTrainingPeriodicity
	isSet bool
}

func (v NullableSimpleTrainingPeriodicity) Get() *SimpleTrainingPeriodicity {
	return v.value
}

func (v *NullableSimpleTrainingPeriodicity) Set(val *SimpleTrainingPeriodicity) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleTrainingPeriodicity) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleTrainingPeriodicity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleTrainingPeriodicity(val *SimpleTrainingPeriodicity) *NullableSimpleTrainingPeriodicity {
	return &NullableSimpleTrainingPeriodicity{value: val, isSet: true}
}

func (v NullableSimpleTrainingPeriodicity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleTrainingPeriodicity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
