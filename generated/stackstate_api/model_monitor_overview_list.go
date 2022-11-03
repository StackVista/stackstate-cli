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

// MonitorOverviewList struct for MonitorOverviewList
type MonitorOverviewList struct {
	Monitors []MonitorOverview `json:"monitors"`
}

// NewMonitorOverviewList instantiates a new MonitorOverviewList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorOverviewList(monitors []MonitorOverview) *MonitorOverviewList {
	this := MonitorOverviewList{}
	this.Monitors = monitors
	return &this
}

// NewMonitorOverviewListWithDefaults instantiates a new MonitorOverviewList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorOverviewListWithDefaults() *MonitorOverviewList {
	this := MonitorOverviewList{}
	return &this
}

// GetMonitors returns the Monitors field value
func (o *MonitorOverviewList) GetMonitors() []MonitorOverview {
	if o == nil {
		var ret []MonitorOverview
		return ret
	}

	return o.Monitors
}

// GetMonitorsOk returns a tuple with the Monitors field value
// and a boolean to check if the value has been set.
func (o *MonitorOverviewList) GetMonitorsOk() ([]MonitorOverview, bool) {
	if o == nil {
		return nil, false
	}
	return o.Monitors, true
}

// SetMonitors sets field value
func (o *MonitorOverviewList) SetMonitors(v []MonitorOverview) {
	o.Monitors = v
}

func (o MonitorOverviewList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["monitors"] = o.Monitors
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorOverviewList struct {
	value *MonitorOverviewList
	isSet bool
}

func (v NullableMonitorOverviewList) Get() *MonitorOverviewList {
	return v.value
}

func (v *NullableMonitorOverviewList) Set(val *MonitorOverviewList) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorOverviewList) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorOverviewList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorOverviewList(val *MonitorOverviewList) *NullableMonitorOverviewList {
	return &NullableMonitorOverviewList{value: val, isSet: true}
}

func (v NullableMonitorOverviewList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorOverviewList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
