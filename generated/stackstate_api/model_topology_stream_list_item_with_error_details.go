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

// TopologyStreamListItemWithErrorDetails struct for TopologyStreamListItemWithErrorDetails
type TopologyStreamListItemWithErrorDetails struct {
	Item TopologyStreamListItem `json:"item"`
	ErrorDetails []TopologyStreamError `json:"errorDetails"`
}

// NewTopologyStreamListItemWithErrorDetails instantiates a new TopologyStreamListItemWithErrorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopologyStreamListItemWithErrorDetails(item TopologyStreamListItem, errorDetails []TopologyStreamError) *TopologyStreamListItemWithErrorDetails {
	this := TopologyStreamListItemWithErrorDetails{}
	this.Item = item
	this.ErrorDetails = errorDetails
	return &this
}

// NewTopologyStreamListItemWithErrorDetailsWithDefaults instantiates a new TopologyStreamListItemWithErrorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologyStreamListItemWithErrorDetailsWithDefaults() *TopologyStreamListItemWithErrorDetails {
	this := TopologyStreamListItemWithErrorDetails{}
	return &this
}

// GetItem returns the Item field value
func (o *TopologyStreamListItemWithErrorDetails) GetItem() TopologyStreamListItem {
	if o == nil {
		var ret TopologyStreamListItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *TopologyStreamListItemWithErrorDetails) GetItemOk() (*TopologyStreamListItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *TopologyStreamListItemWithErrorDetails) SetItem(v TopologyStreamListItem) {
	o.Item = v
}

// GetErrorDetails returns the ErrorDetails field value
func (o *TopologyStreamListItemWithErrorDetails) GetErrorDetails() []TopologyStreamError {
	if o == nil {
		var ret []TopologyStreamError
		return ret
	}

	return o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value
// and a boolean to check if the value has been set.
func (o *TopologyStreamListItemWithErrorDetails) GetErrorDetailsOk() (*[]TopologyStreamError, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorDetails, true
}

// SetErrorDetails sets field value
func (o *TopologyStreamListItemWithErrorDetails) SetErrorDetails(v []TopologyStreamError) {
	o.ErrorDetails = v
}

func (o TopologyStreamListItemWithErrorDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	if true {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	return json.Marshal(toSerialize)
}

type NullableTopologyStreamListItemWithErrorDetails struct {
	value *TopologyStreamListItemWithErrorDetails
	isSet bool
}

func (v NullableTopologyStreamListItemWithErrorDetails) Get() *TopologyStreamListItemWithErrorDetails {
	return v.value
}

func (v *NullableTopologyStreamListItemWithErrorDetails) Set(val *TopologyStreamListItemWithErrorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTopologyStreamListItemWithErrorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologyStreamListItemWithErrorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologyStreamListItemWithErrorDetails(val *TopologyStreamListItemWithErrorDetails) *NullableTopologyStreamListItemWithErrorDetails {
	return &NullableTopologyStreamListItemWithErrorDetails{value: val, isSet: true}
}

func (v NullableTopologyStreamListItemWithErrorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologyStreamListItemWithErrorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


