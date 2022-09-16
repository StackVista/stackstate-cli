/*
StackState API

StackState's API specification

API version: 0.0.1
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackstate_api

import (
	"encoding/json"
	"fmt"
)

// SubscriptionState - struct for SubscriptionState
type SubscriptionState struct {
	ExpiredSubscription *ExpiredSubscription
	LicensedSubscription *LicensedSubscription
	UnlicensedSubscription *UnlicensedSubscription
}

// ExpiredSubscriptionAsSubscriptionState is a convenience function that returns ExpiredSubscription wrapped in SubscriptionState
func ExpiredSubscriptionAsSubscriptionState(v *ExpiredSubscription) SubscriptionState {
	return SubscriptionState{
		ExpiredSubscription: v,
	}
}

// LicensedSubscriptionAsSubscriptionState is a convenience function that returns LicensedSubscription wrapped in SubscriptionState
func LicensedSubscriptionAsSubscriptionState(v *LicensedSubscription) SubscriptionState {
	return SubscriptionState{
		LicensedSubscription: v,
	}
}

// UnlicensedSubscriptionAsSubscriptionState is a convenience function that returns UnlicensedSubscription wrapped in SubscriptionState
func UnlicensedSubscriptionAsSubscriptionState(v *UnlicensedSubscription) SubscriptionState {
	return SubscriptionState{
		UnlicensedSubscription: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubscriptionState) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'ExpiredSubscription'
	if jsonDict["_type"] == "ExpiredSubscription" {
		// try to unmarshal JSON data into ExpiredSubscription
		err = json.Unmarshal(data, &dst.ExpiredSubscription)
		if err == nil {
			return nil // data stored in dst.ExpiredSubscription, return on the first match
		} else {
			dst.ExpiredSubscription = nil
			return fmt.Errorf("Failed to unmarshal SubscriptionState as ExpiredSubscription: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LicensedSubscription'
	if jsonDict["_type"] == "LicensedSubscription" {
		// try to unmarshal JSON data into LicensedSubscription
		err = json.Unmarshal(data, &dst.LicensedSubscription)
		if err == nil {
			return nil // data stored in dst.LicensedSubscription, return on the first match
		} else {
			dst.LicensedSubscription = nil
			return fmt.Errorf("Failed to unmarshal SubscriptionState as LicensedSubscription: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UnlicensedSubscription'
	if jsonDict["_type"] == "UnlicensedSubscription" {
		// try to unmarshal JSON data into UnlicensedSubscription
		err = json.Unmarshal(data, &dst.UnlicensedSubscription)
		if err == nil {
			return nil // data stored in dst.UnlicensedSubscription, return on the first match
		} else {
			dst.UnlicensedSubscription = nil
			return fmt.Errorf("Failed to unmarshal SubscriptionState as UnlicensedSubscription: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubscriptionState) MarshalJSON() ([]byte, error) {
	if src.ExpiredSubscription != nil {
		return json.Marshal(&src.ExpiredSubscription)
	}

	if src.LicensedSubscription != nil {
		return json.Marshal(&src.LicensedSubscription)
	}

	if src.UnlicensedSubscription != nil {
		return json.Marshal(&src.UnlicensedSubscription)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubscriptionState) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ExpiredSubscription != nil {
		return obj.ExpiredSubscription
	}

	if obj.LicensedSubscription != nil {
		return obj.LicensedSubscription
	}

	if obj.UnlicensedSubscription != nil {
		return obj.UnlicensedSubscription
	}

	// all schemas are nil
	return nil
}

type NullableSubscriptionState struct {
	value *SubscriptionState
	isSet bool
}

func (v NullableSubscriptionState) Get() *SubscriptionState {
	return v.value
}

func (v *NullableSubscriptionState) Set(val *SubscriptionState) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionState) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionState(val *SubscriptionState) *NullableSubscriptionState {
	return &NullableSubscriptionState{value: val, isSet: true}
}

func (v NullableSubscriptionState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


