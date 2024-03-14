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

// GenerateIngestionApiKeyRequest struct for GenerateIngestionApiKeyRequest
type GenerateIngestionApiKeyRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Expiration  *int64  `json:"expiration,omitempty"`
}

// NewGenerateIngestionApiKeyRequest instantiates a new GenerateIngestionApiKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateIngestionApiKeyRequest(name string) *GenerateIngestionApiKeyRequest {
	this := GenerateIngestionApiKeyRequest{}
	this.Name = name
	return &this
}

// NewGenerateIngestionApiKeyRequestWithDefaults instantiates a new GenerateIngestionApiKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateIngestionApiKeyRequestWithDefaults() *GenerateIngestionApiKeyRequest {
	this := GenerateIngestionApiKeyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *GenerateIngestionApiKeyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GenerateIngestionApiKeyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GenerateIngestionApiKeyRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GenerateIngestionApiKeyRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateIngestionApiKeyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GenerateIngestionApiKeyRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GenerateIngestionApiKeyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *GenerateIngestionApiKeyRequest) GetExpiration() int64 {
	if o == nil || o.Expiration == nil {
		var ret int64
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateIngestionApiKeyRequest) GetExpirationOk() (*int64, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *GenerateIngestionApiKeyRequest) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given int64 and assigns it to the Expiration field.
func (o *GenerateIngestionApiKeyRequest) SetExpiration(v int64) {
	o.Expiration = &v
}

func (o GenerateIngestionApiKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableGenerateIngestionApiKeyRequest struct {
	value *GenerateIngestionApiKeyRequest
	isSet bool
}

func (v NullableGenerateIngestionApiKeyRequest) Get() *GenerateIngestionApiKeyRequest {
	return v.value
}

func (v *NullableGenerateIngestionApiKeyRequest) Set(val *GenerateIngestionApiKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateIngestionApiKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateIngestionApiKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateIngestionApiKeyRequest(val *GenerateIngestionApiKeyRequest) *NullableGenerateIngestionApiKeyRequest {
	return &NullableGenerateIngestionApiKeyRequest{value: val, isSet: true}
}

func (v NullableGenerateIngestionApiKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateIngestionApiKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
