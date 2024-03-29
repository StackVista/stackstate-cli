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

// IngestionApiKey struct for IngestionApiKey
type IngestionApiKey struct {
	Id                  int64   `json:"id"`
	LastUpdateTimestamp int64   `json:"lastUpdateTimestamp"`
	Name                string  `json:"name"`
	Description         *string `json:"description,omitempty"`
	Expiration          *int64  `json:"expiration,omitempty"`
}

// NewIngestionApiKey instantiates a new IngestionApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestionApiKey(id int64, lastUpdateTimestamp int64, name string) *IngestionApiKey {
	this := IngestionApiKey{}
	this.Id = id
	this.LastUpdateTimestamp = lastUpdateTimestamp
	this.Name = name
	return &this
}

// NewIngestionApiKeyWithDefaults instantiates a new IngestionApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestionApiKeyWithDefaults() *IngestionApiKey {
	this := IngestionApiKey{}
	return &this
}

// GetId returns the Id field value
func (o *IngestionApiKey) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IngestionApiKey) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IngestionApiKey) SetId(v int64) {
	o.Id = v
}

// GetLastUpdateTimestamp returns the LastUpdateTimestamp field value
func (o *IngestionApiKey) GetLastUpdateTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LastUpdateTimestamp
}

// GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field value
// and a boolean to check if the value has been set.
func (o *IngestionApiKey) GetLastUpdateTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdateTimestamp, true
}

// SetLastUpdateTimestamp sets field value
func (o *IngestionApiKey) SetLastUpdateTimestamp(v int64) {
	o.LastUpdateTimestamp = v
}

// GetName returns the Name field value
func (o *IngestionApiKey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IngestionApiKey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IngestionApiKey) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IngestionApiKey) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionApiKey) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IngestionApiKey) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IngestionApiKey) SetDescription(v string) {
	o.Description = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *IngestionApiKey) GetExpiration() int64 {
	if o == nil || o.Expiration == nil {
		var ret int64
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionApiKey) GetExpirationOk() (*int64, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *IngestionApiKey) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given int64 and assigns it to the Expiration field.
func (o *IngestionApiKey) SetExpiration(v int64) {
	o.Expiration = &v
}

func (o IngestionApiKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["lastUpdateTimestamp"] = o.LastUpdateTimestamp
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	return json.Marshal(toSerialize)
}

type NullableIngestionApiKey struct {
	value *IngestionApiKey
	isSet bool
}

func (v NullableIngestionApiKey) Get() *IngestionApiKey {
	return v.value
}

func (v *NullableIngestionApiKey) Set(val *IngestionApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestionApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestionApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestionApiKey(val *IngestionApiKey) *NullableIngestionApiKey {
	return &NullableIngestionApiKey{value: val, isSet: true}
}

func (v NullableIngestionApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestionApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
