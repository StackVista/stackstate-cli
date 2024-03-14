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

// Suggestions struct for Suggestions
type Suggestions struct {
	Suggestions []string `json:"suggestions"`
}

// NewSuggestions instantiates a new Suggestions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestions(suggestions []string) *Suggestions {
	this := Suggestions{}
	this.Suggestions = suggestions
	return &this
}

// NewSuggestionsWithDefaults instantiates a new Suggestions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestionsWithDefaults() *Suggestions {
	this := Suggestions{}
	return &this
}

// GetSuggestions returns the Suggestions field value
func (o *Suggestions) GetSuggestions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Suggestions
}

// GetSuggestionsOk returns a tuple with the Suggestions field value
// and a boolean to check if the value has been set.
func (o *Suggestions) GetSuggestionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suggestions, true
}

// SetSuggestions sets field value
func (o *Suggestions) SetSuggestions(v []string) {
	o.Suggestions = v
}

func (o Suggestions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["suggestions"] = o.Suggestions
	}
	return json.Marshal(toSerialize)
}

type NullableSuggestions struct {
	value *Suggestions
	isSet bool
}

func (v NullableSuggestions) Get() *Suggestions {
	return v.value
}

func (v *NullableSuggestions) Set(val *Suggestions) {
	v.value = val
	v.isSet = true
}

func (v NullableSuggestions) IsSet() bool {
	return v.isSet
}

func (v *NullableSuggestions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuggestions(val *Suggestions) *NullableSuggestions {
	return &NullableSuggestions{value: val, isSet: true}
}

func (v NullableSuggestions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuggestions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}