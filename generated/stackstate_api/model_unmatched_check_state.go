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

// UnmatchedCheckState struct for UnmatchedCheckState
type UnmatchedCheckState struct {
	CheckStateId string `json:"checkStateId"`
	TopologyElementIdentifier string `json:"topologyElementIdentifier"`
}

// NewUnmatchedCheckState instantiates a new UnmatchedCheckState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnmatchedCheckState(checkStateId string, topologyElementIdentifier string) *UnmatchedCheckState {
	this := UnmatchedCheckState{}
	this.CheckStateId = checkStateId
	this.TopologyElementIdentifier = topologyElementIdentifier
	return &this
}

// NewUnmatchedCheckStateWithDefaults instantiates a new UnmatchedCheckState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnmatchedCheckStateWithDefaults() *UnmatchedCheckState {
	this := UnmatchedCheckState{}
	return &this
}

// GetCheckStateId returns the CheckStateId field value
func (o *UnmatchedCheckState) GetCheckStateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckStateId
}

// GetCheckStateIdOk returns a tuple with the CheckStateId field value
// and a boolean to check if the value has been set.
func (o *UnmatchedCheckState) GetCheckStateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CheckStateId, true
}

// SetCheckStateId sets field value
func (o *UnmatchedCheckState) SetCheckStateId(v string) {
	o.CheckStateId = v
}

// GetTopologyElementIdentifier returns the TopologyElementIdentifier field value
func (o *UnmatchedCheckState) GetTopologyElementIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopologyElementIdentifier
}

// GetTopologyElementIdentifierOk returns a tuple with the TopologyElementIdentifier field value
// and a boolean to check if the value has been set.
func (o *UnmatchedCheckState) GetTopologyElementIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TopologyElementIdentifier, true
}

// SetTopologyElementIdentifier sets field value
func (o *UnmatchedCheckState) SetTopologyElementIdentifier(v string) {
	o.TopologyElementIdentifier = v
}

func (o UnmatchedCheckState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["checkStateId"] = o.CheckStateId
	}
	if true {
		toSerialize["topologyElementIdentifier"] = o.TopologyElementIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableUnmatchedCheckState struct {
	value *UnmatchedCheckState
	isSet bool
}

func (v NullableUnmatchedCheckState) Get() *UnmatchedCheckState {
	return v.value
}

func (v *NullableUnmatchedCheckState) Set(val *UnmatchedCheckState) {
	v.value = val
	v.isSet = true
}

func (v NullableUnmatchedCheckState) IsSet() bool {
	return v.isSet
}

func (v *NullableUnmatchedCheckState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnmatchedCheckState(val *UnmatchedCheckState) *NullableUnmatchedCheckState {
	return &NullableUnmatchedCheckState{value: val, isSet: true}
}

func (v NullableUnmatchedCheckState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnmatchedCheckState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


