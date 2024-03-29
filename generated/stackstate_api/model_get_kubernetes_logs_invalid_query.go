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
)

// GetKubernetesLogsInvalidQuery struct for GetKubernetesLogsInvalidQuery
type GetKubernetesLogsInvalidQuery struct {
	Type    string `json:"_type"`
	Message string `json:"message"`
	Query   string `json:"query"`
}

// NewGetKubernetesLogsInvalidQuery instantiates a new GetKubernetesLogsInvalidQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetKubernetesLogsInvalidQuery(type_ string, message string, query string) *GetKubernetesLogsInvalidQuery {
	this := GetKubernetesLogsInvalidQuery{}
	this.Type = type_
	this.Message = message
	this.Query = query
	return &this
}

// NewGetKubernetesLogsInvalidQueryWithDefaults instantiates a new GetKubernetesLogsInvalidQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetKubernetesLogsInvalidQueryWithDefaults() *GetKubernetesLogsInvalidQuery {
	this := GetKubernetesLogsInvalidQuery{}
	return &this
}

// GetType returns the Type field value
func (o *GetKubernetesLogsInvalidQuery) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetKubernetesLogsInvalidQuery) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetKubernetesLogsInvalidQuery) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *GetKubernetesLogsInvalidQuery) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GetKubernetesLogsInvalidQuery) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *GetKubernetesLogsInvalidQuery) SetMessage(v string) {
	o.Message = v
}

// GetQuery returns the Query field value
func (o *GetKubernetesLogsInvalidQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GetKubernetesLogsInvalidQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *GetKubernetesLogsInvalidQuery) SetQuery(v string) {
	o.Query = v
}

func (o GetKubernetesLogsInvalidQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableGetKubernetesLogsInvalidQuery struct {
	value *GetKubernetesLogsInvalidQuery
	isSet bool
}

func (v NullableGetKubernetesLogsInvalidQuery) Get() *GetKubernetesLogsInvalidQuery {
	return v.value
}

func (v *NullableGetKubernetesLogsInvalidQuery) Set(val *GetKubernetesLogsInvalidQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableGetKubernetesLogsInvalidQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableGetKubernetesLogsInvalidQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetKubernetesLogsInvalidQuery(val *GetKubernetesLogsInvalidQuery) *NullableGetKubernetesLogsInvalidQuery {
	return &NullableGetKubernetesLogsInvalidQuery{value: val, isSet: true}
}

func (v NullableGetKubernetesLogsInvalidQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetKubernetesLogsInvalidQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
