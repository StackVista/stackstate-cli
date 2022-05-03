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

// MultipleMatchesCheckState struct for MultipleMatchesCheckState
type MultipleMatchesCheckState struct {
	CheckStateId string `json:"checkStateId"`
	TopologyElementIdentifier string `json:"topologyElementIdentifier"`
	MatchCount int32 `json:"matchCount"`
}

// NewMultipleMatchesCheckState instantiates a new MultipleMatchesCheckState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleMatchesCheckState(checkStateId string, topologyElementIdentifier string, matchCount int32) *MultipleMatchesCheckState {
	this := MultipleMatchesCheckState{}
	this.CheckStateId = checkStateId
	this.TopologyElementIdentifier = topologyElementIdentifier
	this.MatchCount = matchCount
	return &this
}

// NewMultipleMatchesCheckStateWithDefaults instantiates a new MultipleMatchesCheckState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleMatchesCheckStateWithDefaults() *MultipleMatchesCheckState {
	this := MultipleMatchesCheckState{}
	return &this
}

// GetCheckStateId returns the CheckStateId field value
func (o *MultipleMatchesCheckState) GetCheckStateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckStateId
}

// GetCheckStateIdOk returns a tuple with the CheckStateId field value
// and a boolean to check if the value has been set.
func (o *MultipleMatchesCheckState) GetCheckStateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CheckStateId, true
}

// SetCheckStateId sets field value
func (o *MultipleMatchesCheckState) SetCheckStateId(v string) {
	o.CheckStateId = v
}

// GetTopologyElementIdentifier returns the TopologyElementIdentifier field value
func (o *MultipleMatchesCheckState) GetTopologyElementIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopologyElementIdentifier
}

// GetTopologyElementIdentifierOk returns a tuple with the TopologyElementIdentifier field value
// and a boolean to check if the value has been set.
func (o *MultipleMatchesCheckState) GetTopologyElementIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TopologyElementIdentifier, true
}

// SetTopologyElementIdentifier sets field value
func (o *MultipleMatchesCheckState) SetTopologyElementIdentifier(v string) {
	o.TopologyElementIdentifier = v
}

// GetMatchCount returns the MatchCount field value
func (o *MultipleMatchesCheckState) GetMatchCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MatchCount
}

// GetMatchCountOk returns a tuple with the MatchCount field value
// and a boolean to check if the value has been set.
func (o *MultipleMatchesCheckState) GetMatchCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MatchCount, true
}

// SetMatchCount sets field value
func (o *MultipleMatchesCheckState) SetMatchCount(v int32) {
	o.MatchCount = v
}

func (o MultipleMatchesCheckState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["checkStateId"] = o.CheckStateId
	}
	if true {
		toSerialize["topologyElementIdentifier"] = o.TopologyElementIdentifier
	}
	if true {
		toSerialize["matchCount"] = o.MatchCount
	}
	return json.Marshal(toSerialize)
}

type NullableMultipleMatchesCheckState struct {
	value *MultipleMatchesCheckState
	isSet bool
}

func (v NullableMultipleMatchesCheckState) Get() *MultipleMatchesCheckState {
	return v.value
}

func (v *NullableMultipleMatchesCheckState) Set(val *MultipleMatchesCheckState) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleMatchesCheckState) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleMatchesCheckState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleMatchesCheckState(val *MultipleMatchesCheckState) *NullableMultipleMatchesCheckState {
	return &NullableMultipleMatchesCheckState{value: val, isSet: true}
}

func (v NullableMultipleMatchesCheckState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleMatchesCheckState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


