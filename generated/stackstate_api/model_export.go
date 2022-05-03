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

// Export struct for Export
type Export struct {
	NodesWithIds *[]int64 `json:"nodesWithIds,omitempty"`
	AllNodesOfTypes *[]string `json:"allNodesOfTypes,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	AllowReferences *[]string `json:"allowReferences,omitempty"`
}

// NewExport instantiates a new Export object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExport() *Export {
	this := Export{}
	return &this
}

// NewExportWithDefaults instantiates a new Export object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportWithDefaults() *Export {
	this := Export{}
	return &this
}

// GetNodesWithIds returns the NodesWithIds field value if set, zero value otherwise.
func (o *Export) GetNodesWithIds() []int64 {
	if o == nil || o.NodesWithIds == nil {
		var ret []int64
		return ret
	}
	return *o.NodesWithIds
}

// GetNodesWithIdsOk returns a tuple with the NodesWithIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetNodesWithIdsOk() (*[]int64, bool) {
	if o == nil || o.NodesWithIds == nil {
		return nil, false
	}
	return o.NodesWithIds, true
}

// HasNodesWithIds returns a boolean if a field has been set.
func (o *Export) HasNodesWithIds() bool {
	if o != nil && o.NodesWithIds != nil {
		return true
	}

	return false
}

// SetNodesWithIds gets a reference to the given []int64 and assigns it to the NodesWithIds field.
func (o *Export) SetNodesWithIds(v []int64) {
	o.NodesWithIds = &v
}

// GetAllNodesOfTypes returns the AllNodesOfTypes field value if set, zero value otherwise.
func (o *Export) GetAllNodesOfTypes() []string {
	if o == nil || o.AllNodesOfTypes == nil {
		var ret []string
		return ret
	}
	return *o.AllNodesOfTypes
}

// GetAllNodesOfTypesOk returns a tuple with the AllNodesOfTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetAllNodesOfTypesOk() (*[]string, bool) {
	if o == nil || o.AllNodesOfTypes == nil {
		return nil, false
	}
	return o.AllNodesOfTypes, true
}

// HasAllNodesOfTypes returns a boolean if a field has been set.
func (o *Export) HasAllNodesOfTypes() bool {
	if o != nil && o.AllNodesOfTypes != nil {
		return true
	}

	return false
}

// SetAllNodesOfTypes gets a reference to the given []string and assigns it to the AllNodesOfTypes field.
func (o *Export) SetAllNodesOfTypes(v []string) {
	o.AllNodesOfTypes = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *Export) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Export) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *Export) SetNamespace(v string) {
	o.Namespace = &v
}

// GetAllowReferences returns the AllowReferences field value if set, zero value otherwise.
func (o *Export) GetAllowReferences() []string {
	if o == nil || o.AllowReferences == nil {
		var ret []string
		return ret
	}
	return *o.AllowReferences
}

// GetAllowReferencesOk returns a tuple with the AllowReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetAllowReferencesOk() (*[]string, bool) {
	if o == nil || o.AllowReferences == nil {
		return nil, false
	}
	return o.AllowReferences, true
}

// HasAllowReferences returns a boolean if a field has been set.
func (o *Export) HasAllowReferences() bool {
	if o != nil && o.AllowReferences != nil {
		return true
	}

	return false
}

// SetAllowReferences gets a reference to the given []string and assigns it to the AllowReferences field.
func (o *Export) SetAllowReferences(v []string) {
	o.AllowReferences = &v
}

func (o Export) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NodesWithIds != nil {
		toSerialize["nodesWithIds"] = o.NodesWithIds
	}
	if o.AllNodesOfTypes != nil {
		toSerialize["allNodesOfTypes"] = o.AllNodesOfTypes
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.AllowReferences != nil {
		toSerialize["allowReferences"] = o.AllowReferences
	}
	return json.Marshal(toSerialize)
}

type NullableExport struct {
	value *Export
	isSet bool
}

func (v NullableExport) Get() *Export {
	return v.value
}

func (v *NullableExport) Set(val *Export) {
	v.value = val
	v.isSet = true
}

func (v NullableExport) IsSet() bool {
	return v.isSet
}

func (v *NullableExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExport(val *Export) *NullableExport {
	return &NullableExport{value: val, isSet: true}
}

func (v NullableExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


