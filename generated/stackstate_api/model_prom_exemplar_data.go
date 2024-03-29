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

// PromExemplarData struct for PromExemplarData
type PromExemplarData struct {
	SeriesLabels map[string]string `json:"seriesLabels"`
	Exemplars    []PromExemplar    `json:"exemplars"`
}

// NewPromExemplarData instantiates a new PromExemplarData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromExemplarData(seriesLabels map[string]string, exemplars []PromExemplar) *PromExemplarData {
	this := PromExemplarData{}
	this.SeriesLabels = seriesLabels
	this.Exemplars = exemplars
	return &this
}

// NewPromExemplarDataWithDefaults instantiates a new PromExemplarData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromExemplarDataWithDefaults() *PromExemplarData {
	this := PromExemplarData{}
	return &this
}

// GetSeriesLabels returns the SeriesLabels field value
func (o *PromExemplarData) GetSeriesLabels() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.SeriesLabels
}

// GetSeriesLabelsOk returns a tuple with the SeriesLabels field value
// and a boolean to check if the value has been set.
func (o *PromExemplarData) GetSeriesLabelsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesLabels, true
}

// SetSeriesLabels sets field value
func (o *PromExemplarData) SetSeriesLabels(v map[string]string) {
	o.SeriesLabels = v
}

// GetExemplars returns the Exemplars field value
func (o *PromExemplarData) GetExemplars() []PromExemplar {
	if o == nil {
		var ret []PromExemplar
		return ret
	}

	return o.Exemplars
}

// GetExemplarsOk returns a tuple with the Exemplars field value
// and a boolean to check if the value has been set.
func (o *PromExemplarData) GetExemplarsOk() ([]PromExemplar, bool) {
	if o == nil {
		return nil, false
	}
	return o.Exemplars, true
}

// SetExemplars sets field value
func (o *PromExemplarData) SetExemplars(v []PromExemplar) {
	o.Exemplars = v
}

func (o PromExemplarData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["seriesLabels"] = o.SeriesLabels
	}
	if true {
		toSerialize["exemplars"] = o.Exemplars
	}
	return json.Marshal(toSerialize)
}

type NullablePromExemplarData struct {
	value *PromExemplarData
	isSet bool
}

func (v NullablePromExemplarData) Get() *PromExemplarData {
	return v.value
}

func (v *NullablePromExemplarData) Set(val *PromExemplarData) {
	v.value = val
	v.isSet = true
}

func (v NullablePromExemplarData) IsSet() bool {
	return v.isSet
}

func (v *NullablePromExemplarData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromExemplarData(val *PromExemplarData) *NullablePromExemplarData {
	return &NullablePromExemplarData{value: val, isSet: true}
}

func (v NullablePromExemplarData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromExemplarData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
