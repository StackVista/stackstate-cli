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

// Node struct for Node
type Node struct {
	TypeName            string  `json:"typeName"`
	Id                  int64   `json:"id"`
	LastUpdateTimestamp int64   `json:"lastUpdateTimestamp"`
	Identifier          *string `json:"identifier,omitempty"`
	Name                *string `json:"name,omitempty"`
	Description         *string `json:"description,omitempty"`
	OwnedBy             *string `json:"ownedBy,omitempty"`
	Manual              *bool   `json:"manual,omitempty"`
	IsSettingsNode      *bool   `json:"isSettingsNode,omitempty"`
}

// NewNode instantiates a new Node object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNode(typeName string, id int64, lastUpdateTimestamp int64) *Node {
	this := Node{}
	this.TypeName = typeName
	this.Id = id
	this.LastUpdateTimestamp = lastUpdateTimestamp
	return &this
}

// NewNodeWithDefaults instantiates a new Node object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeWithDefaults() *Node {
	this := Node{}
	return &this
}

// GetTypeName returns the TypeName field value
func (o *Node) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *Node) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *Node) SetTypeName(v string) {
	o.TypeName = v
}

// GetId returns the Id field value
func (o *Node) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Node) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Node) SetId(v int64) {
	o.Id = v
}

// GetLastUpdateTimestamp returns the LastUpdateTimestamp field value
func (o *Node) GetLastUpdateTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LastUpdateTimestamp
}

// GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field value
// and a boolean to check if the value has been set.
func (o *Node) GetLastUpdateTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdateTimestamp, true
}

// SetLastUpdateTimestamp sets field value
func (o *Node) SetLastUpdateTimestamp(v int64) {
	o.LastUpdateTimestamp = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *Node) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Node) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *Node) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Node) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Node) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Node) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Node) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Node) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Node) SetDescription(v string) {
	o.Description = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *Node) GetOwnedBy() string {
	if o == nil || o.OwnedBy == nil {
		var ret string
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetOwnedByOk() (*string, bool) {
	if o == nil || o.OwnedBy == nil {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *Node) HasOwnedBy() bool {
	if o != nil && o.OwnedBy != nil {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given string and assigns it to the OwnedBy field.
func (o *Node) SetOwnedBy(v string) {
	o.OwnedBy = &v
}

// GetManual returns the Manual field value if set, zero value otherwise.
func (o *Node) GetManual() bool {
	if o == nil || o.Manual == nil {
		var ret bool
		return ret
	}
	return *o.Manual
}

// GetManualOk returns a tuple with the Manual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetManualOk() (*bool, bool) {
	if o == nil || o.Manual == nil {
		return nil, false
	}
	return o.Manual, true
}

// HasManual returns a boolean if a field has been set.
func (o *Node) HasManual() bool {
	if o != nil && o.Manual != nil {
		return true
	}

	return false
}

// SetManual gets a reference to the given bool and assigns it to the Manual field.
func (o *Node) SetManual(v bool) {
	o.Manual = &v
}

// GetIsSettingsNode returns the IsSettingsNode field value if set, zero value otherwise.
func (o *Node) GetIsSettingsNode() bool {
	if o == nil || o.IsSettingsNode == nil {
		var ret bool
		return ret
	}
	return *o.IsSettingsNode
}

// GetIsSettingsNodeOk returns a tuple with the IsSettingsNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetIsSettingsNodeOk() (*bool, bool) {
	if o == nil || o.IsSettingsNode == nil {
		return nil, false
	}
	return o.IsSettingsNode, true
}

// HasIsSettingsNode returns a boolean if a field has been set.
func (o *Node) HasIsSettingsNode() bool {
	if o != nil && o.IsSettingsNode != nil {
		return true
	}

	return false
}

// SetIsSettingsNode gets a reference to the given bool and assigns it to the IsSettingsNode field.
func (o *Node) SetIsSettingsNode(v bool) {
	o.IsSettingsNode = &v
}

func (o Node) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["typeName"] = o.TypeName
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["lastUpdateTimestamp"] = o.LastUpdateTimestamp
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.OwnedBy != nil {
		toSerialize["ownedBy"] = o.OwnedBy
	}
	if o.Manual != nil {
		toSerialize["manual"] = o.Manual
	}
	if o.IsSettingsNode != nil {
		toSerialize["isSettingsNode"] = o.IsSettingsNode
	}
	return json.Marshal(toSerialize)
}

type NullableNode struct {
	value *Node
	isSet bool
}

func (v NullableNode) Get() *Node {
	return v.value
}

func (v *NullableNode) Set(val *Node) {
	v.value = val
	v.isSet = true
}

func (v NullableNode) IsSet() bool {
	return v.isSet
}

func (v *NullableNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNode(val *Node) *NullableNode {
	return &NullableNode{value: val, isSet: true}
}

func (v NullableNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
