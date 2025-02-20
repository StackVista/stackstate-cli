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

// TopologyStreamListItem struct for TopologyStreamListItem
type TopologyStreamListItem struct {
	SyncIdentifier    NullableString `json:"syncIdentifier,omitempty"`
	NodeId            int64          `json:"nodeId"`
	Name              string         `json:"name"`
	CreatedRelations  int64          `json:"createdRelations"`
	DeletedRelations  int64          `json:"deletedRelations"`
	CreatedComponents int64          `json:"createdComponents"`
	DeletedComponents int64          `json:"deletedComponents"`
	Errors            int64          `json:"errors"`
}

// NewTopologyStreamListItem instantiates a new TopologyStreamListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopologyStreamListItem(nodeId int64, name string, createdRelations int64, deletedRelations int64, createdComponents int64, deletedComponents int64, errors int64) *TopologyStreamListItem {
	this := TopologyStreamListItem{}
	this.NodeId = nodeId
	this.Name = name
	this.CreatedRelations = createdRelations
	this.DeletedRelations = deletedRelations
	this.CreatedComponents = createdComponents
	this.DeletedComponents = deletedComponents
	this.Errors = errors
	return &this
}

// NewTopologyStreamListItemWithDefaults instantiates a new TopologyStreamListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologyStreamListItemWithDefaults() *TopologyStreamListItem {
	this := TopologyStreamListItem{}
	return &this
}

// GetSyncIdentifier returns the SyncIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopologyStreamListItem) GetSyncIdentifier() string {
	if o == nil || o.SyncIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.SyncIdentifier.Get()
}

// GetSyncIdentifierOk returns a tuple with the SyncIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopologyStreamListItem) GetSyncIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyncIdentifier.Get(), o.SyncIdentifier.IsSet()
}

// HasSyncIdentifier returns a boolean if a field has been set.
func (o *TopologyStreamListItem) HasSyncIdentifier() bool {
	if o != nil && o.SyncIdentifier.IsSet() {
		return true
	}

	return false
}

// SetSyncIdentifier gets a reference to the given NullableString and assigns it to the SyncIdentifier field.
func (o *TopologyStreamListItem) SetSyncIdentifier(v string) {
	o.SyncIdentifier.Set(&v)
}

// SetSyncIdentifierNil sets the value for SyncIdentifier to be an explicit nil
func (o *TopologyStreamListItem) SetSyncIdentifierNil() {
	o.SyncIdentifier.Set(nil)
}

// UnsetSyncIdentifier ensures that no value is present for SyncIdentifier, not even an explicit nil
func (o *TopologyStreamListItem) UnsetSyncIdentifier() {
	o.SyncIdentifier.Unset()
}

// GetNodeId returns the NodeId field value
func (o *TopologyStreamListItem) GetNodeId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *TopologyStreamListItem) GetNodeIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *TopologyStreamListItem) SetNodeId(v int64) {
	o.NodeId = v
}

// GetName returns the Name field value
func (o *TopologyStreamListItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TopologyStreamListItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TopologyStreamListItem) SetName(v string) {
	o.Name = v
}

// GetCreatedRelations returns the CreatedRelations field value
func (o *TopologyStreamListItem) GetCreatedRelations() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedRelations
}

// GetCreatedRelationsOk returns a tuple with the CreatedRelations field value
// and a boolean to check if the value has been set.
func (o *TopologyStreamListItem) GetCreatedRelationsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedRelations, true
}

// SetCreatedRelations sets field value
func (o *TopologyStreamListItem) SetCreatedRelations(v int64) {
	o.CreatedRelations = v
}

// GetDeletedRelations returns the DeletedRelations field value
func (o *TopologyStreamListItem) GetDeletedRelations() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeletedRelations
}

// GetDeletedRelationsOk returns a tuple with the DeletedRelations field value
// and a boolean to check if the value has been set.
func (o *TopologyStreamListItem) GetDeletedRelationsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedRelations, true
}

// SetDeletedRelations sets field value
func (o *TopologyStreamListItem) SetDeletedRelations(v int64) {
	o.DeletedRelations = v
}

// GetCreatedComponents returns the CreatedComponents field value
func (o *TopologyStreamListItem) GetCreatedComponents() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedComponents
}

// GetCreatedComponentsOk returns a tuple with the CreatedComponents field value
// and a boolean to check if the value has been set.
func (o *TopologyStreamListItem) GetCreatedComponentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedComponents, true
}

// SetCreatedComponents sets field value
func (o *TopologyStreamListItem) SetCreatedComponents(v int64) {
	o.CreatedComponents = v
}

// GetDeletedComponents returns the DeletedComponents field value
func (o *TopologyStreamListItem) GetDeletedComponents() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeletedComponents
}

// GetDeletedComponentsOk returns a tuple with the DeletedComponents field value
// and a boolean to check if the value has been set.
func (o *TopologyStreamListItem) GetDeletedComponentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedComponents, true
}

// SetDeletedComponents sets field value
func (o *TopologyStreamListItem) SetDeletedComponents(v int64) {
	o.DeletedComponents = v
}

// GetErrors returns the Errors field value
func (o *TopologyStreamListItem) GetErrors() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *TopologyStreamListItem) GetErrorsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *TopologyStreamListItem) SetErrors(v int64) {
	o.Errors = v
}

func (o TopologyStreamListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SyncIdentifier.IsSet() {
		toSerialize["syncIdentifier"] = o.SyncIdentifier.Get()
	}
	if true {
		toSerialize["nodeId"] = o.NodeId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["createdRelations"] = o.CreatedRelations
	}
	if true {
		toSerialize["deletedRelations"] = o.DeletedRelations
	}
	if true {
		toSerialize["createdComponents"] = o.CreatedComponents
	}
	if true {
		toSerialize["deletedComponents"] = o.DeletedComponents
	}
	if true {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableTopologyStreamListItem struct {
	value *TopologyStreamListItem
	isSet bool
}

func (v NullableTopologyStreamListItem) Get() *TopologyStreamListItem {
	return v.value
}

func (v *NullableTopologyStreamListItem) Set(val *TopologyStreamListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTopologyStreamListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologyStreamListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologyStreamListItem(val *TopologyStreamListItem) *NullableTopologyStreamListItem {
	return &NullableTopologyStreamListItem{value: val, isSet: true}
}

func (v NullableTopologyStreamListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologyStreamListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
