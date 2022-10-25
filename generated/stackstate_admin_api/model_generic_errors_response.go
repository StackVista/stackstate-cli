/*
StackState Admin API

StackState's Admin API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
)

// GenericErrorsResponse struct for GenericErrorsResponse
type GenericErrorsResponse struct {
	Type            *string           `json:"_type,omitempty"`
	TrackingKey     string            `json:"trackingKey"`
	ServerTimestamp int64             `json:"serverTimestamp"`
	Errors          []GenericApiError `json:"errors"`
}

// NewGenericErrorsResponse instantiates a new GenericErrorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericErrorsResponse(trackingKey string, serverTimestamp int64, errors []GenericApiError) *GenericErrorsResponse {
	this := GenericErrorsResponse{}
	this.TrackingKey = trackingKey
	this.ServerTimestamp = serverTimestamp
	this.Errors = errors
	return &this
}

// NewGenericErrorsResponseWithDefaults instantiates a new GenericErrorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericErrorsResponseWithDefaults() *GenericErrorsResponse {
	this := GenericErrorsResponse{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GenericErrorsResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorsResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GenericErrorsResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GenericErrorsResponse) SetType(v string) {
	o.Type = &v
}

// GetTrackingKey returns the TrackingKey field value
func (o *GenericErrorsResponse) GetTrackingKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrackingKey
}

// GetTrackingKeyOk returns a tuple with the TrackingKey field value
// and a boolean to check if the value has been set.
func (o *GenericErrorsResponse) GetTrackingKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackingKey, true
}

// SetTrackingKey sets field value
func (o *GenericErrorsResponse) SetTrackingKey(v string) {
	o.TrackingKey = v
}

// GetServerTimestamp returns the ServerTimestamp field value
func (o *GenericErrorsResponse) GetServerTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServerTimestamp
}

// GetServerTimestampOk returns a tuple with the ServerTimestamp field value
// and a boolean to check if the value has been set.
func (o *GenericErrorsResponse) GetServerTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerTimestamp, true
}

// SetServerTimestamp sets field value
func (o *GenericErrorsResponse) SetServerTimestamp(v int64) {
	o.ServerTimestamp = v
}

// GetErrors returns the Errors field value
func (o *GenericErrorsResponse) GetErrors() []GenericApiError {
	if o == nil {
		var ret []GenericApiError
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *GenericErrorsResponse) GetErrorsOk() ([]GenericApiError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *GenericErrorsResponse) SetErrors(v []GenericApiError) {
	o.Errors = v
}

func (o GenericErrorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["trackingKey"] = o.TrackingKey
	}
	if true {
		toSerialize["serverTimestamp"] = o.ServerTimestamp
	}
	if true {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableGenericErrorsResponse struct {
	value *GenericErrorsResponse
	isSet bool
}

func (v NullableGenericErrorsResponse) Get() *GenericErrorsResponse {
	return v.value
}

func (v *NullableGenericErrorsResponse) Set(val *GenericErrorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericErrorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericErrorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericErrorsResponse(val *GenericErrorsResponse) *NullableGenericErrorsResponse {
	return &NullableGenericErrorsResponse{value: val, isSet: true}
}

func (v NullableGenericErrorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericErrorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
