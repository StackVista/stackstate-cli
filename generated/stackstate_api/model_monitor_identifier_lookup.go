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

// MonitorIdentifierLookup struct for MonitorIdentifierLookup
type MonitorIdentifierLookup struct {
	MetricQuery   string                            `json:"metricQuery"`
	ComponentType int64                             `json:"componentType"`
	TopN          *int32                            `json:"topN,omitempty"`
	Overrides     *MonitorIdentifierLookupOverrides `json:"overrides,omitempty"`
}

// NewMonitorIdentifierLookup instantiates a new MonitorIdentifierLookup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorIdentifierLookup(metricQuery string, componentType int64) *MonitorIdentifierLookup {
	this := MonitorIdentifierLookup{}
	this.MetricQuery = metricQuery
	this.ComponentType = componentType
	return &this
}

// NewMonitorIdentifierLookupWithDefaults instantiates a new MonitorIdentifierLookup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorIdentifierLookupWithDefaults() *MonitorIdentifierLookup {
	this := MonitorIdentifierLookup{}
	return &this
}

// GetMetricQuery returns the MetricQuery field value
func (o *MonitorIdentifierLookup) GetMetricQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricQuery
}

// GetMetricQueryOk returns a tuple with the MetricQuery field value
// and a boolean to check if the value has been set.
func (o *MonitorIdentifierLookup) GetMetricQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricQuery, true
}

// SetMetricQuery sets field value
func (o *MonitorIdentifierLookup) SetMetricQuery(v string) {
	o.MetricQuery = v
}

// GetComponentType returns the ComponentType field value
func (o *MonitorIdentifierLookup) GetComponentType() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value
// and a boolean to check if the value has been set.
func (o *MonitorIdentifierLookup) GetComponentTypeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentType, true
}

// SetComponentType sets field value
func (o *MonitorIdentifierLookup) SetComponentType(v int64) {
	o.ComponentType = v
}

// GetTopN returns the TopN field value if set, zero value otherwise.
func (o *MonitorIdentifierLookup) GetTopN() int32 {
	if o == nil || o.TopN == nil {
		var ret int32
		return ret
	}
	return *o.TopN
}

// GetTopNOk returns a tuple with the TopN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorIdentifierLookup) GetTopNOk() (*int32, bool) {
	if o == nil || o.TopN == nil {
		return nil, false
	}
	return o.TopN, true
}

// HasTopN returns a boolean if a field has been set.
func (o *MonitorIdentifierLookup) HasTopN() bool {
	if o != nil && o.TopN != nil {
		return true
	}

	return false
}

// SetTopN gets a reference to the given int32 and assigns it to the TopN field.
func (o *MonitorIdentifierLookup) SetTopN(v int32) {
	o.TopN = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *MonitorIdentifierLookup) GetOverrides() MonitorIdentifierLookupOverrides {
	if o == nil || o.Overrides == nil {
		var ret MonitorIdentifierLookupOverrides
		return ret
	}
	return *o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorIdentifierLookup) GetOverridesOk() (*MonitorIdentifierLookupOverrides, bool) {
	if o == nil || o.Overrides == nil {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *MonitorIdentifierLookup) HasOverrides() bool {
	if o != nil && o.Overrides != nil {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given MonitorIdentifierLookupOverrides and assigns it to the Overrides field.
func (o *MonitorIdentifierLookup) SetOverrides(v MonitorIdentifierLookupOverrides) {
	o.Overrides = &v
}

func (o MonitorIdentifierLookup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metricQuery"] = o.MetricQuery
	}
	if true {
		toSerialize["componentType"] = o.ComponentType
	}
	if o.TopN != nil {
		toSerialize["topN"] = o.TopN
	}
	if o.Overrides != nil {
		toSerialize["overrides"] = o.Overrides
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorIdentifierLookup struct {
	value *MonitorIdentifierLookup
	isSet bool
}

func (v NullableMonitorIdentifierLookup) Get() *MonitorIdentifierLookup {
	return v.value
}

func (v *NullableMonitorIdentifierLookup) Set(val *MonitorIdentifierLookup) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorIdentifierLookup) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorIdentifierLookup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorIdentifierLookup(val *MonitorIdentifierLookup) *NullableMonitorIdentifierLookup {
	return &NullableMonitorIdentifierLookup{value: val, isSet: true}
}

func (v NullableMonitorIdentifierLookup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorIdentifierLookup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
