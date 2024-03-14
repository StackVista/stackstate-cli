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

// MonitorFunctionTest struct for MonitorFunctionTest
type MonitorFunctionTest struct {
	Arguments []Argument `json:"arguments"`
}

// NewMonitorFunctionTest instantiates a new MonitorFunctionTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorFunctionTest(arguments []Argument) *MonitorFunctionTest {
	this := MonitorFunctionTest{}
	this.Arguments = arguments
	return &this
}

// NewMonitorFunctionTestWithDefaults instantiates a new MonitorFunctionTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorFunctionTestWithDefaults() *MonitorFunctionTest {
	this := MonitorFunctionTest{}
	return &this
}

// GetArguments returns the Arguments field value
func (o *MonitorFunctionTest) GetArguments() []Argument {
	if o == nil {
		var ret []Argument
		return ret
	}

	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *MonitorFunctionTest) GetArgumentsOk() ([]Argument, bool) {
	if o == nil {
		return nil, false
	}
	return o.Arguments, true
}

// SetArguments sets field value
func (o *MonitorFunctionTest) SetArguments(v []Argument) {
	o.Arguments = v
}

func (o MonitorFunctionTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["arguments"] = o.Arguments
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorFunctionTest struct {
	value *MonitorFunctionTest
	isSet bool
}

func (v NullableMonitorFunctionTest) Get() *MonitorFunctionTest {
	return v.value
}

func (v *NullableMonitorFunctionTest) Set(val *MonitorFunctionTest) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorFunctionTest) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorFunctionTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorFunctionTest(val *MonitorFunctionTest) *NullableMonitorFunctionTest {
	return &NullableMonitorFunctionTest{value: val, isSet: true}
}

func (v NullableMonitorFunctionTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorFunctionTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}