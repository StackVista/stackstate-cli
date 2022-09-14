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

// NodeTypesNodeTypesInner struct for NodeTypesNodeTypesInner
type NodeTypesNodeTypesInner struct {
	TypeName string `json:"typeName"`
	Description string `json:"description"`
}

// NewNodeTypesNodeTypesInner instantiates a new NodeTypesNodeTypesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeTypesNodeTypesInner(typeName string, description string) *NodeTypesNodeTypesInner {
	this := NodeTypesNodeTypesInner{}
	this.TypeName = typeName
	this.Description = description
	return &this
}

// NewNodeTypesNodeTypesInnerWithDefaults instantiates a new NodeTypesNodeTypesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeTypesNodeTypesInnerWithDefaults() *NodeTypesNodeTypesInner {
	this := NodeTypesNodeTypesInner{}
	return &this
}

// GetTypeName returns the TypeName field value
func (o *NodeTypesNodeTypesInner) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *NodeTypesNodeTypesInner) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *NodeTypesNodeTypesInner) SetTypeName(v string) {
	o.TypeName = v
}

// GetDescription returns the Description field value
func (o *NodeTypesNodeTypesInner) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *NodeTypesNodeTypesInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *NodeTypesNodeTypesInner) SetDescription(v string) {
	o.Description = v
}

func (o NodeTypesNodeTypesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["typeName"] = o.TypeName
	}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableNodeTypesNodeTypesInner struct {
	value *NodeTypesNodeTypesInner
	isSet bool
}

func (v NullableNodeTypesNodeTypesInner) Get() *NodeTypesNodeTypesInner {
	return v.value
}

func (v *NullableNodeTypesNodeTypesInner) Set(val *NodeTypesNodeTypesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeTypesNodeTypesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeTypesNodeTypesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeTypesNodeTypesInner(val *NodeTypesNodeTypesInner) *NullableNodeTypesNodeTypesInner {
	return &NullableNodeTypesNodeTypesInner{value: val, isSet: true}
}

func (v NullableNodeTypesNodeTypesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeTypesNodeTypesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

