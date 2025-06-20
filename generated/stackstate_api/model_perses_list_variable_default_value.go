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

// PersesListVariableDefaultValue - struct for PersesListVariableDefaultValue
type PersesListVariableDefaultValue struct {
	PersesListVariableDefaultSingleValue *PersesListVariableDefaultSingleValue
	PersesListVariableDefaultSliceValues *PersesListVariableDefaultSliceValues
}

// PersesListVariableDefaultSingleValueAsPersesListVariableDefaultValue is a convenience function that returns PersesListVariableDefaultSingleValue wrapped in PersesListVariableDefaultValue
func PersesListVariableDefaultSingleValueAsPersesListVariableDefaultValue(v *PersesListVariableDefaultSingleValue) PersesListVariableDefaultValue {
	return PersesListVariableDefaultValue{
		PersesListVariableDefaultSingleValue: v,
	}
}

// PersesListVariableDefaultSliceValuesAsPersesListVariableDefaultValue is a convenience function that returns PersesListVariableDefaultSliceValues wrapped in PersesListVariableDefaultValue
func PersesListVariableDefaultSliceValuesAsPersesListVariableDefaultValue(v *PersesListVariableDefaultSliceValues) PersesListVariableDefaultValue {
	return PersesListVariableDefaultValue{
		PersesListVariableDefaultSliceValues: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PersesListVariableDefaultValue) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'PersesListVariableDefaultSingleValue'
	if jsonDict["kind"] == "PersesListVariableDefaultSingleValue" {
		// try to unmarshal JSON data into PersesListVariableDefaultSingleValue
		err = json.Unmarshal(data, &dst.PersesListVariableDefaultSingleValue)
		if err == nil {
			return nil // data stored in dst.PersesListVariableDefaultSingleValue, return on the first match
		} else {
			dst.PersesListVariableDefaultSingleValue = nil
			return fmt.Errorf("Failed to unmarshal PersesListVariableDefaultValue as PersesListVariableDefaultSingleValue: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PersesListVariableDefaultSliceValues'
	if jsonDict["kind"] == "PersesListVariableDefaultSliceValues" {
		// try to unmarshal JSON data into PersesListVariableDefaultSliceValues
		err = json.Unmarshal(data, &dst.PersesListVariableDefaultSliceValues)
		if err == nil {
			return nil // data stored in dst.PersesListVariableDefaultSliceValues, return on the first match
		} else {
			dst.PersesListVariableDefaultSliceValues = nil
			return fmt.Errorf("Failed to unmarshal PersesListVariableDefaultValue as PersesListVariableDefaultSliceValues: %s", err.Error())
		}
	}

	// check if the discriminator value is 'singleValue'
	if jsonDict["kind"] == "singleValue" {
		// try to unmarshal JSON data into PersesListVariableDefaultSingleValue
		err = json.Unmarshal(data, &dst.PersesListVariableDefaultSingleValue)
		if err == nil {
			return nil // data stored in dst.PersesListVariableDefaultSingleValue, return on the first match
		} else {
			dst.PersesListVariableDefaultSingleValue = nil
			return fmt.Errorf("Failed to unmarshal PersesListVariableDefaultValue as PersesListVariableDefaultSingleValue: %s", err.Error())
		}
	}

	// check if the discriminator value is 'sliceValues'
	if jsonDict["kind"] == "sliceValues" {
		// try to unmarshal JSON data into PersesListVariableDefaultSliceValues
		err = json.Unmarshal(data, &dst.PersesListVariableDefaultSliceValues)
		if err == nil {
			return nil // data stored in dst.PersesListVariableDefaultSliceValues, return on the first match
		} else {
			dst.PersesListVariableDefaultSliceValues = nil
			return fmt.Errorf("Failed to unmarshal PersesListVariableDefaultValue as PersesListVariableDefaultSliceValues: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PersesListVariableDefaultValue) MarshalJSON() ([]byte, error) {
	if src.PersesListVariableDefaultSingleValue != nil {
		return json.Marshal(&src.PersesListVariableDefaultSingleValue)
	}

	if src.PersesListVariableDefaultSliceValues != nil {
		return json.Marshal(&src.PersesListVariableDefaultSliceValues)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PersesListVariableDefaultValue) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PersesListVariableDefaultSingleValue != nil {
		return obj.PersesListVariableDefaultSingleValue
	}

	if obj.PersesListVariableDefaultSliceValues != nil {
		return obj.PersesListVariableDefaultSliceValues
	}

	// all schemas are nil
	return nil
}

type NullablePersesListVariableDefaultValue struct {
	value *PersesListVariableDefaultValue
	isSet bool
}

func (v NullablePersesListVariableDefaultValue) Get() *PersesListVariableDefaultValue {
	return v.value
}

func (v *NullablePersesListVariableDefaultValue) Set(val *PersesListVariableDefaultValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePersesListVariableDefaultValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePersesListVariableDefaultValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersesListVariableDefaultValue(val *PersesListVariableDefaultValue) *NullablePersesListVariableDefaultValue {
	return &NullablePersesListVariableDefaultValue{value: val, isSet: true}
}

func (v NullablePersesListVariableDefaultValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersesListVariableDefaultValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
