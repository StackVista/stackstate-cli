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
	"fmt"
)

// IngestionApiKeyCreateError - struct for IngestionApiKeyCreateError
type IngestionApiKeyCreateError struct {
	IngestionApiKeyInvalidExpiryError *IngestionApiKeyInvalidExpiryError
}

// IngestionApiKeyInvalidExpiryErrorAsIngestionApiKeyCreateError is a convenience function that returns IngestionApiKeyInvalidExpiryError wrapped in IngestionApiKeyCreateError
func IngestionApiKeyInvalidExpiryErrorAsIngestionApiKeyCreateError(v *IngestionApiKeyInvalidExpiryError) IngestionApiKeyCreateError {
	return IngestionApiKeyCreateError{
		IngestionApiKeyInvalidExpiryError: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IngestionApiKeyCreateError) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'IngestionApiKeyInvalidExpiryError'
	if jsonDict["_type"] == "IngestionApiKeyInvalidExpiryError" {
		// try to unmarshal JSON data into IngestionApiKeyInvalidExpiryError
		err = json.Unmarshal(data, &dst.IngestionApiKeyInvalidExpiryError)
		if err == nil {
			return nil // data stored in dst.IngestionApiKeyInvalidExpiryError, return on the first match
		} else {
			dst.IngestionApiKeyInvalidExpiryError = nil
			return fmt.Errorf("Failed to unmarshal IngestionApiKeyCreateError as IngestionApiKeyInvalidExpiryError: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IngestionApiKeyCreateError) MarshalJSON() ([]byte, error) {
	if src.IngestionApiKeyInvalidExpiryError != nil {
		return json.Marshal(&src.IngestionApiKeyInvalidExpiryError)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IngestionApiKeyCreateError) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IngestionApiKeyInvalidExpiryError != nil {
		return obj.IngestionApiKeyInvalidExpiryError
	}

	// all schemas are nil
	return nil
}

type NullableIngestionApiKeyCreateError struct {
	value *IngestionApiKeyCreateError
	isSet bool
}

func (v NullableIngestionApiKeyCreateError) Get() *IngestionApiKeyCreateError {
	return v.value
}

func (v *NullableIngestionApiKeyCreateError) Set(val *IngestionApiKeyCreateError) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestionApiKeyCreateError) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestionApiKeyCreateError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestionApiKeyCreateError(val *IngestionApiKeyCreateError) *NullableIngestionApiKeyCreateError {
	return &NullableIngestionApiKeyCreateError{value: val, isSet: true}
}

func (v NullableIngestionApiKeyCreateError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestionApiKeyCreateError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}