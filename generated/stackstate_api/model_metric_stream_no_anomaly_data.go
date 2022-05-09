/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
)

// MetricStreamNoAnomalyData struct for MetricStreamNoAnomalyData
type MetricStreamNoAnomalyData struct {
	Type string `json:"_type"`
	CheckedInterval TimeRange `json:"checkedInterval"`
	Explanation string `json:"explanation"`
	ModelInfo map[string]interface{} `json:"modelInfo"`
	Query *AnnotationMetricQuery `json:"query,omitempty"`
}

// NewMetricStreamNoAnomalyData instantiates a new MetricStreamNoAnomalyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricStreamNoAnomalyData(type_ string, checkedInterval TimeRange, explanation string, modelInfo map[string]interface{}) *MetricStreamNoAnomalyData {
	this := MetricStreamNoAnomalyData{}
	this.Type = type_
	this.CheckedInterval = checkedInterval
	this.Explanation = explanation
	this.ModelInfo = modelInfo
	return &this
}

// NewMetricStreamNoAnomalyDataWithDefaults instantiates a new MetricStreamNoAnomalyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricStreamNoAnomalyDataWithDefaults() *MetricStreamNoAnomalyData {
	this := MetricStreamNoAnomalyData{}
	return &this
}

// GetType returns the Type field value
func (o *MetricStreamNoAnomalyData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MetricStreamNoAnomalyData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MetricStreamNoAnomalyData) SetType(v string) {
	o.Type = v
}

// GetCheckedInterval returns the CheckedInterval field value
func (o *MetricStreamNoAnomalyData) GetCheckedInterval() TimeRange {
	if o == nil {
		var ret TimeRange
		return ret
	}

	return o.CheckedInterval
}

// GetCheckedIntervalOk returns a tuple with the CheckedInterval field value
// and a boolean to check if the value has been set.
func (o *MetricStreamNoAnomalyData) GetCheckedIntervalOk() (*TimeRange, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CheckedInterval, true
}

// SetCheckedInterval sets field value
func (o *MetricStreamNoAnomalyData) SetCheckedInterval(v TimeRange) {
	o.CheckedInterval = v
}

// GetExplanation returns the Explanation field value
func (o *MetricStreamNoAnomalyData) GetExplanation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Explanation
}

// GetExplanationOk returns a tuple with the Explanation field value
// and a boolean to check if the value has been set.
func (o *MetricStreamNoAnomalyData) GetExplanationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Explanation, true
}

// SetExplanation sets field value
func (o *MetricStreamNoAnomalyData) SetExplanation(v string) {
	o.Explanation = v
}

// GetModelInfo returns the ModelInfo field value
func (o *MetricStreamNoAnomalyData) GetModelInfo() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ModelInfo
}

// GetModelInfoOk returns a tuple with the ModelInfo field value
// and a boolean to check if the value has been set.
func (o *MetricStreamNoAnomalyData) GetModelInfoOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModelInfo, true
}

// SetModelInfo sets field value
func (o *MetricStreamNoAnomalyData) SetModelInfo(v map[string]interface{}) {
	o.ModelInfo = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *MetricStreamNoAnomalyData) GetQuery() AnnotationMetricQuery {
	if o == nil || o.Query == nil {
		var ret AnnotationMetricQuery
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricStreamNoAnomalyData) GetQueryOk() (*AnnotationMetricQuery, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *MetricStreamNoAnomalyData) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given AnnotationMetricQuery and assigns it to the Query field.
func (o *MetricStreamNoAnomalyData) SetQuery(v AnnotationMetricQuery) {
	o.Query = &v
}

func (o MetricStreamNoAnomalyData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["checkedInterval"] = o.CheckedInterval
	}
	if true {
		toSerialize["explanation"] = o.Explanation
	}
	if true {
		toSerialize["modelInfo"] = o.ModelInfo
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableMetricStreamNoAnomalyData struct {
	value *MetricStreamNoAnomalyData
	isSet bool
}

func (v NullableMetricStreamNoAnomalyData) Get() *MetricStreamNoAnomalyData {
	return v.value
}

func (v *NullableMetricStreamNoAnomalyData) Set(val *MetricStreamNoAnomalyData) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricStreamNoAnomalyData) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricStreamNoAnomalyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricStreamNoAnomalyData(val *MetricStreamNoAnomalyData) *NullableMetricStreamNoAnomalyData {
	return &NullableMetricStreamNoAnomalyData{value: val, isSet: true}
}

func (v NullableMetricStreamNoAnomalyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricStreamNoAnomalyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


