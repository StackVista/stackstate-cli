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

// MonitorCheckStatusRelatedFailures struct for MonitorCheckStatusRelatedFailures
type MonitorCheckStatusRelatedFailures struct {
	CheckStatusId int64                                          `json:"checkStatusId"`
	TopologyTime  int32                                          `json:"topologyTime"`
	CheckStatuses []MonitorCheckStatusRelatedFailuresCheckStatus `json:"checkStatuses"`
}

// NewMonitorCheckStatusRelatedFailures instantiates a new MonitorCheckStatusRelatedFailures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorCheckStatusRelatedFailures(checkStatusId int64, topologyTime int32, checkStatuses []MonitorCheckStatusRelatedFailuresCheckStatus) *MonitorCheckStatusRelatedFailures {
	this := MonitorCheckStatusRelatedFailures{}
	this.CheckStatusId = checkStatusId
	this.TopologyTime = topologyTime
	this.CheckStatuses = checkStatuses
	return &this
}

// NewMonitorCheckStatusRelatedFailuresWithDefaults instantiates a new MonitorCheckStatusRelatedFailures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorCheckStatusRelatedFailuresWithDefaults() *MonitorCheckStatusRelatedFailures {
	this := MonitorCheckStatusRelatedFailures{}
	return &this
}

// GetCheckStatusId returns the CheckStatusId field value
func (o *MonitorCheckStatusRelatedFailures) GetCheckStatusId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CheckStatusId
}

// GetCheckStatusIdOk returns a tuple with the CheckStatusId field value
// and a boolean to check if the value has been set.
func (o *MonitorCheckStatusRelatedFailures) GetCheckStatusIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckStatusId, true
}

// SetCheckStatusId sets field value
func (o *MonitorCheckStatusRelatedFailures) SetCheckStatusId(v int64) {
	o.CheckStatusId = v
}

// GetTopologyTime returns the TopologyTime field value
func (o *MonitorCheckStatusRelatedFailures) GetTopologyTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TopologyTime
}

// GetTopologyTimeOk returns a tuple with the TopologyTime field value
// and a boolean to check if the value has been set.
func (o *MonitorCheckStatusRelatedFailures) GetTopologyTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopologyTime, true
}

// SetTopologyTime sets field value
func (o *MonitorCheckStatusRelatedFailures) SetTopologyTime(v int32) {
	o.TopologyTime = v
}

// GetCheckStatuses returns the CheckStatuses field value
func (o *MonitorCheckStatusRelatedFailures) GetCheckStatuses() []MonitorCheckStatusRelatedFailuresCheckStatus {
	if o == nil {
		var ret []MonitorCheckStatusRelatedFailuresCheckStatus
		return ret
	}

	return o.CheckStatuses
}

// GetCheckStatusesOk returns a tuple with the CheckStatuses field value
// and a boolean to check if the value has been set.
func (o *MonitorCheckStatusRelatedFailures) GetCheckStatusesOk() ([]MonitorCheckStatusRelatedFailuresCheckStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckStatuses, true
}

// SetCheckStatuses sets field value
func (o *MonitorCheckStatusRelatedFailures) SetCheckStatuses(v []MonitorCheckStatusRelatedFailuresCheckStatus) {
	o.CheckStatuses = v
}

func (o MonitorCheckStatusRelatedFailures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["checkStatusId"] = o.CheckStatusId
	}
	if true {
		toSerialize["topologyTime"] = o.TopologyTime
	}
	if true {
		toSerialize["checkStatuses"] = o.CheckStatuses
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorCheckStatusRelatedFailures struct {
	value *MonitorCheckStatusRelatedFailures
	isSet bool
}

func (v NullableMonitorCheckStatusRelatedFailures) Get() *MonitorCheckStatusRelatedFailures {
	return v.value
}

func (v *NullableMonitorCheckStatusRelatedFailures) Set(val *MonitorCheckStatusRelatedFailures) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorCheckStatusRelatedFailures) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorCheckStatusRelatedFailures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorCheckStatusRelatedFailures(val *MonitorCheckStatusRelatedFailures) *NullableMonitorCheckStatusRelatedFailures {
	return &NullableMonitorCheckStatusRelatedFailures{value: val, isSet: true}
}

func (v NullableMonitorCheckStatusRelatedFailures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorCheckStatusRelatedFailures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}