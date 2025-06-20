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

// DashboardReadSchema - struct for DashboardReadSchema
type DashboardReadSchema struct {
	DashboardReadFullSchema     *DashboardReadFullSchema
	DashboardReadMetadataSchema *DashboardReadMetadataSchema
}

// DashboardReadFullSchemaAsDashboardReadSchema is a convenience function that returns DashboardReadFullSchema wrapped in DashboardReadSchema
func DashboardReadFullSchemaAsDashboardReadSchema(v *DashboardReadFullSchema) DashboardReadSchema {
	return DashboardReadSchema{
		DashboardReadFullSchema: v,
	}
}

// DashboardReadMetadataSchemaAsDashboardReadSchema is a convenience function that returns DashboardReadMetadataSchema wrapped in DashboardReadSchema
func DashboardReadMetadataSchemaAsDashboardReadSchema(v *DashboardReadMetadataSchema) DashboardReadSchema {
	return DashboardReadSchema{
		DashboardReadMetadataSchema: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DashboardReadSchema) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'DashboardReadFullSchema'
	if jsonDict["_type"] == "DashboardReadFullSchema" {
		// try to unmarshal JSON data into DashboardReadFullSchema
		err = json.Unmarshal(data, &dst.DashboardReadFullSchema)
		if err == nil {
			return nil // data stored in dst.DashboardReadFullSchema, return on the first match
		} else {
			dst.DashboardReadFullSchema = nil
			return fmt.Errorf("Failed to unmarshal DashboardReadSchema as DashboardReadFullSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DashboardReadMetadataSchema'
	if jsonDict["_type"] == "DashboardReadMetadataSchema" {
		// try to unmarshal JSON data into DashboardReadMetadataSchema
		err = json.Unmarshal(data, &dst.DashboardReadMetadataSchema)
		if err == nil {
			return nil // data stored in dst.DashboardReadMetadataSchema, return on the first match
		} else {
			dst.DashboardReadMetadataSchema = nil
			return fmt.Errorf("Failed to unmarshal DashboardReadSchema as DashboardReadMetadataSchema: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DashboardReadSchema) MarshalJSON() ([]byte, error) {
	if src.DashboardReadFullSchema != nil {
		return json.Marshal(&src.DashboardReadFullSchema)
	}

	if src.DashboardReadMetadataSchema != nil {
		return json.Marshal(&src.DashboardReadMetadataSchema)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DashboardReadSchema) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DashboardReadFullSchema != nil {
		return obj.DashboardReadFullSchema
	}

	if obj.DashboardReadMetadataSchema != nil {
		return obj.DashboardReadMetadataSchema
	}

	// all schemas are nil
	return nil
}

type NullableDashboardReadSchema struct {
	value *DashboardReadSchema
	isSet bool
}

func (v NullableDashboardReadSchema) Get() *DashboardReadSchema {
	return v.value
}

func (v *NullableDashboardReadSchema) Set(val *DashboardReadSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardReadSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardReadSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardReadSchema(val *DashboardReadSchema) *NullableDashboardReadSchema {
	return &NullableDashboardReadSchema{value: val, isSet: true}
}

func (v NullableDashboardReadSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardReadSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
