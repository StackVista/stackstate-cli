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

// CausingEventsResult struct for CausingEventsResult
type CausingEventsResult struct {
	Type string `json:"_type"`
	Items []TopologyEvent `json:"items"`
}

// NewCausingEventsResult instantiates a new CausingEventsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCausingEventsResult(type_ string, items []TopologyEvent) *CausingEventsResult {
	this := CausingEventsResult{}
	this.Type = type_
	this.Items = items
	return &this
}

// NewCausingEventsResultWithDefaults instantiates a new CausingEventsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCausingEventsResultWithDefaults() *CausingEventsResult {
	this := CausingEventsResult{}
	return &this
}

// GetType returns the Type field value
func (o *CausingEventsResult) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CausingEventsResult) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CausingEventsResult) SetType(v string) {
	o.Type = v
}

// GetItems returns the Items field value
func (o *CausingEventsResult) GetItems() []TopologyEvent {
	if o == nil {
		var ret []TopologyEvent
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CausingEventsResult) GetItemsOk() ([]TopologyEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CausingEventsResult) SetItems(v []TopologyEvent) {
	o.Items = v
}

func (o CausingEventsResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableCausingEventsResult struct {
	value *CausingEventsResult
	isSet bool
}

func (v NullableCausingEventsResult) Get() *CausingEventsResult {
	return v.value
}

func (v *NullableCausingEventsResult) Set(val *CausingEventsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCausingEventsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCausingEventsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCausingEventsResult(val *CausingEventsResult) *NullableCausingEventsResult {
	return &NullableCausingEventsResult{value: val, isSet: true}
}

func (v NullableCausingEventsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCausingEventsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


