/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_client

import (
	"encoding/json"
	"fmt"
)

// GenericServerError - struct for GenericServerError
type GenericServerError struct {
	InternalServerError *InternalServerError
}

// InternalServerErrorAsGenericServerError is a convenience function that returns InternalServerError wrapped in GenericServerError
func InternalServerErrorAsGenericServerError(v *InternalServerError) GenericServerError {
	return GenericServerError{ InternalServerError: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GenericServerError) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InternalServerError
	err = json.Unmarshal(data, &dst.InternalServerError)
	if err == nil {
		jsonInternalServerError, _ := json.Marshal(dst.InternalServerError)
		if string(jsonInternalServerError) == "{}" { // empty struct
			dst.InternalServerError = nil
		} else {
			match++
		}
	} else {
		dst.InternalServerError = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InternalServerError = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GenericServerError)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GenericServerError)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GenericServerError) MarshalJSON() ([]byte, error) {
	if src.InternalServerError != nil {
		return json.Marshal(&src.InternalServerError)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GenericServerError) GetActualInstance() (interface{}) {
	if obj.InternalServerError != nil {
		return obj.InternalServerError
	}

	// all schemas are nil
	return nil
}

type NullableGenericServerError struct {
	value *GenericServerError
	isSet bool
}

func (v NullableGenericServerError) Get() *GenericServerError {
	return v.value
}

func (v *NullableGenericServerError) Set(val *GenericServerError) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericServerError) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericServerError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericServerError(val *GenericServerError) *NullableGenericServerError {
	return &NullableGenericServerError{value: val, isSet: true}
}

func (v NullableGenericServerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericServerError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

