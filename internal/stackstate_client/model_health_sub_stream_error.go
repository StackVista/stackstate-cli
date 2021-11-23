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

// HealthSubStreamError - struct for HealthSubStreamError
type HealthSubStreamError struct {
	HealthStreamNotFound *HealthStreamNotFound
	HealthSubStreamNotFound *HealthSubStreamNotFound
}

// HealthStreamNotFoundAsHealthSubStreamError is a convenience function that returns HealthStreamNotFound wrapped in HealthSubStreamError
func HealthStreamNotFoundAsHealthSubStreamError(v *HealthStreamNotFound) HealthSubStreamError {
	return HealthSubStreamError{ HealthStreamNotFound: v}
}

// HealthSubStreamNotFoundAsHealthSubStreamError is a convenience function that returns HealthSubStreamNotFound wrapped in HealthSubStreamError
func HealthSubStreamNotFoundAsHealthSubStreamError(v *HealthSubStreamNotFound) HealthSubStreamError {
	return HealthSubStreamError{ HealthSubStreamNotFound: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *HealthSubStreamError) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HealthStreamNotFound
	err = json.Unmarshal(data, &dst.HealthStreamNotFound)
	if err == nil {
		jsonHealthStreamNotFound, _ := json.Marshal(dst.HealthStreamNotFound)
		if string(jsonHealthStreamNotFound) == "{}" { // empty struct
			dst.HealthStreamNotFound = nil
		} else {
			match++
		}
	} else {
		dst.HealthStreamNotFound = nil
	}

	// try to unmarshal data into HealthSubStreamNotFound
	err = json.Unmarshal(data, &dst.HealthSubStreamNotFound)
	if err == nil {
		jsonHealthSubStreamNotFound, _ := json.Marshal(dst.HealthSubStreamNotFound)
		if string(jsonHealthSubStreamNotFound) == "{}" { // empty struct
			dst.HealthSubStreamNotFound = nil
		} else {
			match++
		}
	} else {
		dst.HealthSubStreamNotFound = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.HealthStreamNotFound = nil
		dst.HealthSubStreamNotFound = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(HealthSubStreamError)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(HealthSubStreamError)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HealthSubStreamError) MarshalJSON() ([]byte, error) {
	if src.HealthStreamNotFound != nil {
		return json.Marshal(&src.HealthStreamNotFound)
	}

	if src.HealthSubStreamNotFound != nil {
		return json.Marshal(&src.HealthSubStreamNotFound)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HealthSubStreamError) GetActualInstance() (interface{}) {
	if obj.HealthStreamNotFound != nil {
		return obj.HealthStreamNotFound
	}

	if obj.HealthSubStreamNotFound != nil {
		return obj.HealthSubStreamNotFound
	}

	// all schemas are nil
	return nil
}

type NullableHealthSubStreamError struct {
	value *HealthSubStreamError
	isSet bool
}

func (v NullableHealthSubStreamError) Get() *HealthSubStreamError {
	return v.value
}

func (v *NullableHealthSubStreamError) Set(val *HealthSubStreamError) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthSubStreamError) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthSubStreamError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthSubStreamError(val *HealthSubStreamError) *NullableHealthSubStreamError {
	return &NullableHealthSubStreamError{value: val, isSet: true}
}

func (v NullableHealthSubStreamError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthSubStreamError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

