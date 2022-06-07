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

// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	Type *string `json:"_type,omitempty"`
	StreamSnaphots []LatestTelemetryStreamMetrics `json:"streamSnaphots,omitempty"`
}

// NewInlineResponse200 instantiates a new InlineResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200WithDefaults() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200) SetType(v string) {
	o.Type = &v
}

// GetStreamSnaphots returns the StreamSnaphots field value if set, zero value otherwise.
func (o *InlineResponse200) GetStreamSnaphots() []LatestTelemetryStreamMetrics {
	if o == nil || o.StreamSnaphots == nil {
		var ret []LatestTelemetryStreamMetrics
		return ret
	}
	return o.StreamSnaphots
}

// GetStreamSnaphotsOk returns a tuple with the StreamSnaphots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetStreamSnaphotsOk() ([]LatestTelemetryStreamMetrics, bool) {
	if o == nil || o.StreamSnaphots == nil {
		return nil, false
	}
	return o.StreamSnaphots, true
}

// HasStreamSnaphots returns a boolean if a field has been set.
func (o *InlineResponse200) HasStreamSnaphots() bool {
	if o != nil && o.StreamSnaphots != nil {
		return true
	}

	return false
}

// SetStreamSnaphots gets a reference to the given []LatestTelemetryStreamMetrics and assigns it to the StreamSnaphots field.
func (o *InlineResponse200) SetStreamSnaphots(v []LatestTelemetryStreamMetrics) {
	o.StreamSnaphots = v
}

func (o InlineResponse200) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["_type"] = o.Type
	}
	if o.StreamSnaphots != nil {
		toSerialize["streamSnaphots"] = o.StreamSnaphots
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200 struct {
	value *InlineResponse200
	isSet bool
}

func (v NullableInlineResponse200) Get() *InlineResponse200 {
	return v.value
}

func (v *NullableInlineResponse200) Set(val *InlineResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200(val *InlineResponse200) *NullableInlineResponse200 {
	return &NullableInlineResponse200{value: val, isSet: true}
}

func (v NullableInlineResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


