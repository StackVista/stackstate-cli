/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// HealthSubStreamConsistencyState - struct for HealthSubStreamConsistencyState
type HealthSubStreamConsistencyState struct {
	HealthSubStreamExpiry *HealthSubStreamExpiry
	HealthSubStreamSnapshot *HealthSubStreamSnapshot
	HealthSubStreamTransactionalIncrements *HealthSubStreamTransactionalIncrements
}

// HealthSubStreamExpiryAsHealthSubStreamConsistencyState is a convenience function that returns HealthSubStreamExpiry wrapped in HealthSubStreamConsistencyState
func HealthSubStreamExpiryAsHealthSubStreamConsistencyState(v *HealthSubStreamExpiry) HealthSubStreamConsistencyState {
	return HealthSubStreamConsistencyState{ HealthSubStreamExpiry: v}
}

// HealthSubStreamSnapshotAsHealthSubStreamConsistencyState is a convenience function that returns HealthSubStreamSnapshot wrapped in HealthSubStreamConsistencyState
func HealthSubStreamSnapshotAsHealthSubStreamConsistencyState(v *HealthSubStreamSnapshot) HealthSubStreamConsistencyState {
	return HealthSubStreamConsistencyState{ HealthSubStreamSnapshot: v}
}

// HealthSubStreamTransactionalIncrementsAsHealthSubStreamConsistencyState is a convenience function that returns HealthSubStreamTransactionalIncrements wrapped in HealthSubStreamConsistencyState
func HealthSubStreamTransactionalIncrementsAsHealthSubStreamConsistencyState(v *HealthSubStreamTransactionalIncrements) HealthSubStreamConsistencyState {
	return HealthSubStreamConsistencyState{ HealthSubStreamTransactionalIncrements: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *HealthSubStreamConsistencyState) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HealthSubStreamExpiry
	err = json.Unmarshal(data, &dst.HealthSubStreamExpiry)
	if err == nil {
		jsonHealthSubStreamExpiry, _ := json.Marshal(dst.HealthSubStreamExpiry)
		if string(jsonHealthSubStreamExpiry) == "{}" { // empty struct
			dst.HealthSubStreamExpiry = nil
		} else {
			match++
		}
	} else {
		dst.HealthSubStreamExpiry = nil
	}

	// try to unmarshal data into HealthSubStreamSnapshot
	err = json.Unmarshal(data, &dst.HealthSubStreamSnapshot)
	if err == nil {
		jsonHealthSubStreamSnapshot, _ := json.Marshal(dst.HealthSubStreamSnapshot)
		if string(jsonHealthSubStreamSnapshot) == "{}" { // empty struct
			dst.HealthSubStreamSnapshot = nil
		} else {
			match++
		}
	} else {
		dst.HealthSubStreamSnapshot = nil
	}

	// try to unmarshal data into HealthSubStreamTransactionalIncrements
	err = json.Unmarshal(data, &dst.HealthSubStreamTransactionalIncrements)
	if err == nil {
		jsonHealthSubStreamTransactionalIncrements, _ := json.Marshal(dst.HealthSubStreamTransactionalIncrements)
		if string(jsonHealthSubStreamTransactionalIncrements) == "{}" { // empty struct
			dst.HealthSubStreamTransactionalIncrements = nil
		} else {
			match++
		}
	} else {
		dst.HealthSubStreamTransactionalIncrements = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.HealthSubStreamExpiry = nil
		dst.HealthSubStreamSnapshot = nil
		dst.HealthSubStreamTransactionalIncrements = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(HealthSubStreamConsistencyState)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(HealthSubStreamConsistencyState)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HealthSubStreamConsistencyState) MarshalJSON() ([]byte, error) {
	if src.HealthSubStreamExpiry != nil {
		return json.Marshal(&src.HealthSubStreamExpiry)
	}

	if src.HealthSubStreamSnapshot != nil {
		return json.Marshal(&src.HealthSubStreamSnapshot)
	}

	if src.HealthSubStreamTransactionalIncrements != nil {
		return json.Marshal(&src.HealthSubStreamTransactionalIncrements)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HealthSubStreamConsistencyState) GetActualInstance() (interface{}) {
	if obj.HealthSubStreamExpiry != nil {
		return obj.HealthSubStreamExpiry
	}

	if obj.HealthSubStreamSnapshot != nil {
		return obj.HealthSubStreamSnapshot
	}

	if obj.HealthSubStreamTransactionalIncrements != nil {
		return obj.HealthSubStreamTransactionalIncrements
	}

	// all schemas are nil
	return nil
}

type NullableHealthSubStreamConsistencyState struct {
	value *HealthSubStreamConsistencyState
	isSet bool
}

func (v NullableHealthSubStreamConsistencyState) Get() *HealthSubStreamConsistencyState {
	return v.value
}

func (v *NullableHealthSubStreamConsistencyState) Set(val *HealthSubStreamConsistencyState) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthSubStreamConsistencyState) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthSubStreamConsistencyState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthSubStreamConsistencyState(val *HealthSubStreamConsistencyState) *NullableHealthSubStreamConsistencyState {
	return &NullableHealthSubStreamConsistencyState{value: val, isSet: true}
}

func (v NullableHealthSubStreamConsistencyState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthSubStreamConsistencyState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


