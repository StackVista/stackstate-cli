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

// GetKubernetesLogsBadRequest - struct for GetKubernetesLogsBadRequest
type GetKubernetesLogsBadRequest struct {
	GetKubernetesLogsInvalidPagination *GetKubernetesLogsInvalidPagination
	GetKubernetesLogsInvalidQuery      *GetKubernetesLogsInvalidQuery
	GetKubernetesLogsInvalidTimeRange  *GetKubernetesLogsInvalidTimeRange
}

// GetKubernetesLogsInvalidPaginationAsGetKubernetesLogsBadRequest is a convenience function that returns GetKubernetesLogsInvalidPagination wrapped in GetKubernetesLogsBadRequest
func GetKubernetesLogsInvalidPaginationAsGetKubernetesLogsBadRequest(v *GetKubernetesLogsInvalidPagination) GetKubernetesLogsBadRequest {
	return GetKubernetesLogsBadRequest{
		GetKubernetesLogsInvalidPagination: v,
	}
}

// GetKubernetesLogsInvalidQueryAsGetKubernetesLogsBadRequest is a convenience function that returns GetKubernetesLogsInvalidQuery wrapped in GetKubernetesLogsBadRequest
func GetKubernetesLogsInvalidQueryAsGetKubernetesLogsBadRequest(v *GetKubernetesLogsInvalidQuery) GetKubernetesLogsBadRequest {
	return GetKubernetesLogsBadRequest{
		GetKubernetesLogsInvalidQuery: v,
	}
}

// GetKubernetesLogsInvalidTimeRangeAsGetKubernetesLogsBadRequest is a convenience function that returns GetKubernetesLogsInvalidTimeRange wrapped in GetKubernetesLogsBadRequest
func GetKubernetesLogsInvalidTimeRangeAsGetKubernetesLogsBadRequest(v *GetKubernetesLogsInvalidTimeRange) GetKubernetesLogsBadRequest {
	return GetKubernetesLogsBadRequest{
		GetKubernetesLogsInvalidTimeRange: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetKubernetesLogsBadRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'GetKubernetesLogsInvalidPagination'
	if jsonDict["_type"] == "GetKubernetesLogsInvalidPagination" {
		// try to unmarshal JSON data into GetKubernetesLogsInvalidPagination
		err = json.Unmarshal(data, &dst.GetKubernetesLogsInvalidPagination)
		if err == nil {
			return nil // data stored in dst.GetKubernetesLogsInvalidPagination, return on the first match
		} else {
			dst.GetKubernetesLogsInvalidPagination = nil
			return fmt.Errorf("Failed to unmarshal GetKubernetesLogsBadRequest as GetKubernetesLogsInvalidPagination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GetKubernetesLogsInvalidQuery'
	if jsonDict["_type"] == "GetKubernetesLogsInvalidQuery" {
		// try to unmarshal JSON data into GetKubernetesLogsInvalidQuery
		err = json.Unmarshal(data, &dst.GetKubernetesLogsInvalidQuery)
		if err == nil {
			return nil // data stored in dst.GetKubernetesLogsInvalidQuery, return on the first match
		} else {
			dst.GetKubernetesLogsInvalidQuery = nil
			return fmt.Errorf("Failed to unmarshal GetKubernetesLogsBadRequest as GetKubernetesLogsInvalidQuery: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GetKubernetesLogsInvalidTimeRange'
	if jsonDict["_type"] == "GetKubernetesLogsInvalidTimeRange" {
		// try to unmarshal JSON data into GetKubernetesLogsInvalidTimeRange
		err = json.Unmarshal(data, &dst.GetKubernetesLogsInvalidTimeRange)
		if err == nil {
			return nil // data stored in dst.GetKubernetesLogsInvalidTimeRange, return on the first match
		} else {
			dst.GetKubernetesLogsInvalidTimeRange = nil
			return fmt.Errorf("Failed to unmarshal GetKubernetesLogsBadRequest as GetKubernetesLogsInvalidTimeRange: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetKubernetesLogsBadRequest) MarshalJSON() ([]byte, error) {
	if src.GetKubernetesLogsInvalidPagination != nil {
		return json.Marshal(&src.GetKubernetesLogsInvalidPagination)
	}

	if src.GetKubernetesLogsInvalidQuery != nil {
		return json.Marshal(&src.GetKubernetesLogsInvalidQuery)
	}

	if src.GetKubernetesLogsInvalidTimeRange != nil {
		return json.Marshal(&src.GetKubernetesLogsInvalidTimeRange)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetKubernetesLogsBadRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetKubernetesLogsInvalidPagination != nil {
		return obj.GetKubernetesLogsInvalidPagination
	}

	if obj.GetKubernetesLogsInvalidQuery != nil {
		return obj.GetKubernetesLogsInvalidQuery
	}

	if obj.GetKubernetesLogsInvalidTimeRange != nil {
		return obj.GetKubernetesLogsInvalidTimeRange
	}

	// all schemas are nil
	return nil
}

type NullableGetKubernetesLogsBadRequest struct {
	value *GetKubernetesLogsBadRequest
	isSet bool
}

func (v NullableGetKubernetesLogsBadRequest) Get() *GetKubernetesLogsBadRequest {
	return v.value
}

func (v *NullableGetKubernetesLogsBadRequest) Set(val *GetKubernetesLogsBadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetKubernetesLogsBadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetKubernetesLogsBadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetKubernetesLogsBadRequest(val *GetKubernetesLogsBadRequest) *NullableGetKubernetesLogsBadRequest {
	return &NullableGetKubernetesLogsBadRequest{value: val, isSet: true}
}

func (v NullableGetKubernetesLogsBadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetKubernetesLogsBadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
