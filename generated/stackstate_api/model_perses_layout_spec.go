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

// PersesLayoutSpec struct for PersesLayoutSpec
type PersesLayoutSpec struct {
	Display *PersesGridLayoutDisplay `json:"display,omitempty"`
	Items   []PersesGridItem         `json:"items"`
}

// NewPersesLayoutSpec instantiates a new PersesLayoutSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersesLayoutSpec(items []PersesGridItem) *PersesLayoutSpec {
	this := PersesLayoutSpec{}
	this.Items = items
	return &this
}

// NewPersesLayoutSpecWithDefaults instantiates a new PersesLayoutSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersesLayoutSpecWithDefaults() *PersesLayoutSpec {
	this := PersesLayoutSpec{}
	return &this
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PersesLayoutSpec) GetDisplay() PersesGridLayoutDisplay {
	if o == nil || o.Display == nil {
		var ret PersesGridLayoutDisplay
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersesLayoutSpec) GetDisplayOk() (*PersesGridLayoutDisplay, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PersesLayoutSpec) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given PersesGridLayoutDisplay and assigns it to the Display field.
func (o *PersesLayoutSpec) SetDisplay(v PersesGridLayoutDisplay) {
	o.Display = &v
}

// GetItems returns the Items field value
func (o *PersesLayoutSpec) GetItems() []PersesGridItem {
	if o == nil {
		var ret []PersesGridItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PersesLayoutSpec) GetItemsOk() ([]PersesGridItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *PersesLayoutSpec) SetItems(v []PersesGridItem) {
	o.Items = v
}

func (o PersesLayoutSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullablePersesLayoutSpec struct {
	value *PersesLayoutSpec
	isSet bool
}

func (v NullablePersesLayoutSpec) Get() *PersesLayoutSpec {
	return v.value
}

func (v *NullablePersesLayoutSpec) Set(val *PersesLayoutSpec) {
	v.value = val
	v.isSet = true
}

func (v NullablePersesLayoutSpec) IsSet() bool {
	return v.isSet
}

func (v *NullablePersesLayoutSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersesLayoutSpec(val *PersesLayoutSpec) *NullablePersesLayoutSpec {
	return &NullablePersesLayoutSpec{value: val, isSet: true}
}

func (v NullablePersesLayoutSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersesLayoutSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
