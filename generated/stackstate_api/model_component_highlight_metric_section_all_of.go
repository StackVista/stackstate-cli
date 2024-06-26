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

// ComponentHighlightMetricSectionAllOf struct for ComponentHighlightMetricSectionAllOf
type ComponentHighlightMetricSectionAllOf struct {
	Type string `json:"_type"`
}

// NewComponentHighlightMetricSectionAllOf instantiates a new ComponentHighlightMetricSectionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentHighlightMetricSectionAllOf(type_ string) *ComponentHighlightMetricSectionAllOf {
	this := ComponentHighlightMetricSectionAllOf{}
	this.Type = type_
	return &this
}

// NewComponentHighlightMetricSectionAllOfWithDefaults instantiates a new ComponentHighlightMetricSectionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentHighlightMetricSectionAllOfWithDefaults() *ComponentHighlightMetricSectionAllOf {
	this := ComponentHighlightMetricSectionAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *ComponentHighlightMetricSectionAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ComponentHighlightMetricSectionAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ComponentHighlightMetricSectionAllOf) SetType(v string) {
	o.Type = v
}

func (o ComponentHighlightMetricSectionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableComponentHighlightMetricSectionAllOf struct {
	value *ComponentHighlightMetricSectionAllOf
	isSet bool
}

func (v NullableComponentHighlightMetricSectionAllOf) Get() *ComponentHighlightMetricSectionAllOf {
	return v.value
}

func (v *NullableComponentHighlightMetricSectionAllOf) Set(val *ComponentHighlightMetricSectionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentHighlightMetricSectionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentHighlightMetricSectionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentHighlightMetricSectionAllOf(val *ComponentHighlightMetricSectionAllOf) *NullableComponentHighlightMetricSectionAllOf {
	return &NullableComponentHighlightMetricSectionAllOf{value: val, isSet: true}
}

func (v NullableComponentHighlightMetricSectionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentHighlightMetricSectionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
