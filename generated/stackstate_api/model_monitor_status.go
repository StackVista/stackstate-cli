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

// MonitorStatus struct for MonitorStatus
type MonitorStatus struct {
	Monitor             Monitor              `json:"monitor"`
	Function            MonitorFunction      `json:"function"`
	Errors              []MonitorError       `json:"errors,omitempty"`
	Metrics             MonitorMetrics       `json:"metrics"`
	TopologyMatchResult *TopologyMatchResult `json:"topologyMatchResult,omitempty"`
}

// NewMonitorStatus instantiates a new MonitorStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorStatus(monitor Monitor, function MonitorFunction, metrics MonitorMetrics) *MonitorStatus {
	this := MonitorStatus{}
	this.Monitor = monitor
	this.Function = function
	this.Metrics = metrics
	return &this
}

// NewMonitorStatusWithDefaults instantiates a new MonitorStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorStatusWithDefaults() *MonitorStatus {
	this := MonitorStatus{}
	return &this
}

// GetMonitor returns the Monitor field value
func (o *MonitorStatus) GetMonitor() Monitor {
	if o == nil {
		var ret Monitor
		return ret
	}

	return o.Monitor
}

// GetMonitorOk returns a tuple with the Monitor field value
// and a boolean to check if the value has been set.
func (o *MonitorStatus) GetMonitorOk() (*Monitor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Monitor, true
}

// SetMonitor sets field value
func (o *MonitorStatus) SetMonitor(v Monitor) {
	o.Monitor = v
}

// GetFunction returns the Function field value
func (o *MonitorStatus) GetFunction() MonitorFunction {
	if o == nil {
		var ret MonitorFunction
		return ret
	}

	return o.Function
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
func (o *MonitorStatus) GetFunctionOk() (*MonitorFunction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Function, true
}

// SetFunction sets field value
func (o *MonitorStatus) SetFunction(v MonitorFunction) {
	o.Function = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *MonitorStatus) GetErrors() []MonitorError {
	if o == nil || o.Errors == nil {
		var ret []MonitorError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStatus) GetErrorsOk() ([]MonitorError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *MonitorStatus) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []MonitorError and assigns it to the Errors field.
func (o *MonitorStatus) SetErrors(v []MonitorError) {
	o.Errors = v
}

// GetMetrics returns the Metrics field value
func (o *MonitorStatus) GetMetrics() MonitorMetrics {
	if o == nil {
		var ret MonitorMetrics
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *MonitorStatus) GetMetricsOk() (*MonitorMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *MonitorStatus) SetMetrics(v MonitorMetrics) {
	o.Metrics = v
}

// GetTopologyMatchResult returns the TopologyMatchResult field value if set, zero value otherwise.
func (o *MonitorStatus) GetTopologyMatchResult() TopologyMatchResult {
	if o == nil || o.TopologyMatchResult == nil {
		var ret TopologyMatchResult
		return ret
	}
	return *o.TopologyMatchResult
}

// GetTopologyMatchResultOk returns a tuple with the TopologyMatchResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStatus) GetTopologyMatchResultOk() (*TopologyMatchResult, bool) {
	if o == nil || o.TopologyMatchResult == nil {
		return nil, false
	}
	return o.TopologyMatchResult, true
}

// HasTopologyMatchResult returns a boolean if a field has been set.
func (o *MonitorStatus) HasTopologyMatchResult() bool {
	if o != nil && o.TopologyMatchResult != nil {
		return true
	}

	return false
}

// SetTopologyMatchResult gets a reference to the given TopologyMatchResult and assigns it to the TopologyMatchResult field.
func (o *MonitorStatus) SetTopologyMatchResult(v TopologyMatchResult) {
	o.TopologyMatchResult = &v
}

func (o MonitorStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["monitor"] = o.Monitor
	}
	if true {
		toSerialize["function"] = o.Function
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if true {
		toSerialize["metrics"] = o.Metrics
	}
	if o.TopologyMatchResult != nil {
		toSerialize["topologyMatchResult"] = o.TopologyMatchResult
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorStatus struct {
	value *MonitorStatus
	isSet bool
}

func (v NullableMonitorStatus) Get() *MonitorStatus {
	return v.value
}

func (v *NullableMonitorStatus) Set(val *MonitorStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorStatus(val *MonitorStatus) *NullableMonitorStatus {
	return &NullableMonitorStatus{value: val, isSet: true}
}

func (v NullableMonitorStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
