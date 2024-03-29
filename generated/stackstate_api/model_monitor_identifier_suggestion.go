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

// MonitorIdentifierSuggestion struct for MonitorIdentifierSuggestion
type MonitorIdentifierSuggestion struct {
	IdentifierTemplate string  `json:"identifierTemplate"`
	Score              float32 `json:"score"`
	TotalMatches       int64   `json:"totalMatches"`
}

// NewMonitorIdentifierSuggestion instantiates a new MonitorIdentifierSuggestion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorIdentifierSuggestion(identifierTemplate string, score float32, totalMatches int64) *MonitorIdentifierSuggestion {
	this := MonitorIdentifierSuggestion{}
	this.IdentifierTemplate = identifierTemplate
	this.Score = score
	this.TotalMatches = totalMatches
	return &this
}

// NewMonitorIdentifierSuggestionWithDefaults instantiates a new MonitorIdentifierSuggestion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorIdentifierSuggestionWithDefaults() *MonitorIdentifierSuggestion {
	this := MonitorIdentifierSuggestion{}
	return &this
}

// GetIdentifierTemplate returns the IdentifierTemplate field value
func (o *MonitorIdentifierSuggestion) GetIdentifierTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentifierTemplate
}

// GetIdentifierTemplateOk returns a tuple with the IdentifierTemplate field value
// and a boolean to check if the value has been set.
func (o *MonitorIdentifierSuggestion) GetIdentifierTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentifierTemplate, true
}

// SetIdentifierTemplate sets field value
func (o *MonitorIdentifierSuggestion) SetIdentifierTemplate(v string) {
	o.IdentifierTemplate = v
}

// GetScore returns the Score field value
func (o *MonitorIdentifierSuggestion) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *MonitorIdentifierSuggestion) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *MonitorIdentifierSuggestion) SetScore(v float32) {
	o.Score = v
}

// GetTotalMatches returns the TotalMatches field value
func (o *MonitorIdentifierSuggestion) GetTotalMatches() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalMatches
}

// GetTotalMatchesOk returns a tuple with the TotalMatches field value
// and a boolean to check if the value has been set.
func (o *MonitorIdentifierSuggestion) GetTotalMatchesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMatches, true
}

// SetTotalMatches sets field value
func (o *MonitorIdentifierSuggestion) SetTotalMatches(v int64) {
	o.TotalMatches = v
}

func (o MonitorIdentifierSuggestion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identifierTemplate"] = o.IdentifierTemplate
	}
	if true {
		toSerialize["score"] = o.Score
	}
	if true {
		toSerialize["totalMatches"] = o.TotalMatches
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorIdentifierSuggestion struct {
	value *MonitorIdentifierSuggestion
	isSet bool
}

func (v NullableMonitorIdentifierSuggestion) Get() *MonitorIdentifierSuggestion {
	return v.value
}

func (v *NullableMonitorIdentifierSuggestion) Set(val *MonitorIdentifierSuggestion) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorIdentifierSuggestion) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorIdentifierSuggestion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorIdentifierSuggestion(val *MonitorIdentifierSuggestion) *NullableMonitorIdentifierSuggestion {
	return &NullableMonitorIdentifierSuggestion{value: val, isSet: true}
}

func (v NullableMonitorIdentifierSuggestion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorIdentifierSuggestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
