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

// DashboardScope Scope of the dashboard. 'publicDashboard' for accessible by everyone, 'privateDashboard' for owned and accessible by the current user.
type DashboardScope string

// List of DashboardScope
const (
	DASHBOARDSCOPE_PUBLIC_DASHBOARD  DashboardScope = "publicDashboard"
	DASHBOARDSCOPE_PRIVATE_DASHBOARD DashboardScope = "privateDashboard"
)

// All allowed values of DashboardScope enum
var AllowedDashboardScopeEnumValues = []DashboardScope{
	"publicDashboard",
	"privateDashboard",
}

func (v *DashboardScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DashboardScope(value)
	for _, existing := range AllowedDashboardScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DashboardScope", value)
}

// NewDashboardScopeFromValue returns a pointer to a valid DashboardScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDashboardScopeFromValue(v string) (*DashboardScope, error) {
	ev := DashboardScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DashboardScope: valid values are %v", v, AllowedDashboardScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DashboardScope) IsValid() bool {
	for _, existing := range AllowedDashboardScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardScope value
func (v DashboardScope) Ptr() *DashboardScope {
	return &v
}

type NullableDashboardScope struct {
	value *DashboardScope
	isSet bool
}

func (v NullableDashboardScope) Get() *DashboardScope {
	return v.value
}

func (v *NullableDashboardScope) Set(val *DashboardScope) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardScope) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardScope(val *DashboardScope) *NullableDashboardScope {
	return &NullableDashboardScope{value: val, isSet: true}
}

func (v NullableDashboardScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
