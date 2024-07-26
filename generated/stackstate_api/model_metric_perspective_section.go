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

// MetricPerspectiveSection struct for MetricPerspectiveSection
type MetricPerspectiveSection struct {
	Name   string  `json:"name"`
	Weight float32 `json:"weight"`
	Type   string  `json:"_type"`
	Tab    string  `json:"tab"`
}

// NewMetricPerspectiveSection instantiates a new MetricPerspectiveSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricPerspectiveSection(name string, weight float32, type_ string, tab string) *MetricPerspectiveSection {
	this := MetricPerspectiveSection{}
	this.Name = name
	this.Weight = weight
	this.Type = type_
	this.Tab = tab
	return &this
}

// NewMetricPerspectiveSectionWithDefaults instantiates a new MetricPerspectiveSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricPerspectiveSectionWithDefaults() *MetricPerspectiveSection {
	this := MetricPerspectiveSection{}
	return &this
}

// GetName returns the Name field value
func (o *MetricPerspectiveSection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetricPerspectiveSection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetricPerspectiveSection) SetName(v string) {
	o.Name = v
}

// GetWeight returns the Weight field value
func (o *MetricPerspectiveSection) GetWeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *MetricPerspectiveSection) GetWeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *MetricPerspectiveSection) SetWeight(v float32) {
	o.Weight = v
}

// GetType returns the Type field value
func (o *MetricPerspectiveSection) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MetricPerspectiveSection) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MetricPerspectiveSection) SetType(v string) {
	o.Type = v
}

// GetTab returns the Tab field value
func (o *MetricPerspectiveSection) GetTab() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tab
}

// GetTabOk returns a tuple with the Tab field value
// and a boolean to check if the value has been set.
func (o *MetricPerspectiveSection) GetTabOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tab, true
}

// SetTab sets field value
func (o *MetricPerspectiveSection) SetTab(v string) {
	o.Tab = v
}

func (o MetricPerspectiveSection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["weight"] = o.Weight
	}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["tab"] = o.Tab
	}
	return json.Marshal(toSerialize)
}

type NullableMetricPerspectiveSection struct {
	value *MetricPerspectiveSection
	isSet bool
}

func (v NullableMetricPerspectiveSection) Get() *MetricPerspectiveSection {
	return v.value
}

func (v *NullableMetricPerspectiveSection) Set(val *MetricPerspectiveSection) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricPerspectiveSection) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricPerspectiveSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricPerspectiveSection(val *MetricPerspectiveSection) *NullableMetricPerspectiveSection {
	return &NullableMetricPerspectiveSection{value: val, isSet: true}
}

func (v NullableMetricPerspectiveSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricPerspectiveSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}