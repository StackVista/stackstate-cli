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

// EventFilters struct for EventFilters
type EventFilters struct {
	Types      []string `json:"types,omitempty"`
	Tags       []string `json:"tags,omitempty"`
	Categories []string `json:"categories,omitempty"`
	Sources    []string `json:"sources,omitempty"`
}

// NewEventFilters instantiates a new EventFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventFilters() *EventFilters {
	this := EventFilters{}
	return &this
}

// NewEventFiltersWithDefaults instantiates a new EventFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventFiltersWithDefaults() *EventFilters {
	this := EventFilters{}
	return &this
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *EventFilters) GetTypes() []string {
	if o == nil || o.Types == nil {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilters) GetTypesOk() ([]string, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *EventFilters) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *EventFilters) SetTypes(v []string) {
	o.Types = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EventFilters) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilters) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EventFilters) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *EventFilters) SetTags(v []string) {
	o.Tags = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *EventFilters) GetCategories() []string {
	if o == nil || o.Categories == nil {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilters) GetCategoriesOk() ([]string, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *EventFilters) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *EventFilters) SetCategories(v []string) {
	o.Categories = v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *EventFilters) GetSources() []string {
	if o == nil || o.Sources == nil {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilters) GetSourcesOk() ([]string, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *EventFilters) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []string and assigns it to the Sources field.
func (o *EventFilters) SetSources(v []string) {
	o.Sources = v
}

func (o EventFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	return json.Marshal(toSerialize)
}

type NullableEventFilters struct {
	value *EventFilters
	isSet bool
}

func (v NullableEventFilters) Get() *EventFilters {
	return v.value
}

func (v *NullableEventFilters) Set(val *EventFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableEventFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableEventFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventFilters(val *EventFilters) *NullableEventFilters {
	return &NullableEventFilters{value: val, isSet: true}
}

func (v NullableEventFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
