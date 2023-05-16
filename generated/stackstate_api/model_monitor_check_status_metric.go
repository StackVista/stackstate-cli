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

// MonitorCheckStatusMetric struct for MonitorCheckStatusMetric
type MonitorCheckStatusMetric struct {
	Type  string            `json:"_type"`
	Name  string            `json:"name"`
	Query PromqlMetricQuery `json:"query"`
}

// NewMonitorCheckStatusMetric instantiates a new MonitorCheckStatusMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorCheckStatusMetric(type_ string, name string, query PromqlMetricQuery) *MonitorCheckStatusMetric {
	this := MonitorCheckStatusMetric{}
	this.Type = type_
	this.Name = name
	this.Query = query
	return &this
}

// NewMonitorCheckStatusMetricWithDefaults instantiates a new MonitorCheckStatusMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorCheckStatusMetricWithDefaults() *MonitorCheckStatusMetric {
	this := MonitorCheckStatusMetric{}
	return &this
}

// GetType returns the Type field value
func (o *MonitorCheckStatusMetric) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MonitorCheckStatusMetric) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MonitorCheckStatusMetric) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *MonitorCheckStatusMetric) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MonitorCheckStatusMetric) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MonitorCheckStatusMetric) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value
func (o *MonitorCheckStatusMetric) GetQuery() PromqlMetricQuery {
	if o == nil {
		var ret PromqlMetricQuery
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *MonitorCheckStatusMetric) GetQueryOk() (*PromqlMetricQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *MonitorCheckStatusMetric) SetQuery(v PromqlMetricQuery) {
	o.Query = v
}

func (o MonitorCheckStatusMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorCheckStatusMetric struct {
	value *MonitorCheckStatusMetric
	isSet bool
}

func (v NullableMonitorCheckStatusMetric) Get() *MonitorCheckStatusMetric {
	return v.value
}

func (v *NullableMonitorCheckStatusMetric) Set(val *MonitorCheckStatusMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorCheckStatusMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorCheckStatusMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorCheckStatusMetric(val *MonitorCheckStatusMetric) *NullableMonitorCheckStatusMetric {
	return &NullableMonitorCheckStatusMetric{value: val, isSet: true}
}

func (v NullableMonitorCheckStatusMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorCheckStatusMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
