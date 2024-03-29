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

// PromQLMetric struct for PromQLMetric
type PromQLMetric struct {
	Type          string `json:"_type"`
	Id            *int64 `json:"id,omitempty"`
	Query         string `json:"query"`
	Unit          string `json:"unit"`
	AliasTemplate string `json:"aliasTemplate"`
}

// NewPromQLMetric instantiates a new PromQLMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromQLMetric(type_ string, query string, unit string, aliasTemplate string) *PromQLMetric {
	this := PromQLMetric{}
	this.Type = type_
	this.Query = query
	this.Unit = unit
	this.AliasTemplate = aliasTemplate
	return &this
}

// NewPromQLMetricWithDefaults instantiates a new PromQLMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromQLMetricWithDefaults() *PromQLMetric {
	this := PromQLMetric{}
	return &this
}

// GetType returns the Type field value
func (o *PromQLMetric) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PromQLMetric) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PromQLMetric) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PromQLMetric) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromQLMetric) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PromQLMetric) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PromQLMetric) SetId(v int64) {
	o.Id = &v
}

// GetQuery returns the Query field value
func (o *PromQLMetric) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *PromQLMetric) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *PromQLMetric) SetQuery(v string) {
	o.Query = v
}

// GetUnit returns the Unit field value
func (o *PromQLMetric) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *PromQLMetric) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *PromQLMetric) SetUnit(v string) {
	o.Unit = v
}

// GetAliasTemplate returns the AliasTemplate field value
func (o *PromQLMetric) GetAliasTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AliasTemplate
}

// GetAliasTemplateOk returns a tuple with the AliasTemplate field value
// and a boolean to check if the value has been set.
func (o *PromQLMetric) GetAliasTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AliasTemplate, true
}

// SetAliasTemplate sets field value
func (o *PromQLMetric) SetAliasTemplate(v string) {
	o.AliasTemplate = v
}

func (o PromQLMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	if true {
		toSerialize["aliasTemplate"] = o.AliasTemplate
	}
	return json.Marshal(toSerialize)
}

type NullablePromQLMetric struct {
	value *PromQLMetric
	isSet bool
}

func (v NullablePromQLMetric) Get() *PromQLMetric {
	return v.value
}

func (v *NullablePromQLMetric) Set(val *PromQLMetric) {
	v.value = val
	v.isSet = true
}

func (v NullablePromQLMetric) IsSet() bool {
	return v.isSet
}

func (v *NullablePromQLMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromQLMetric(val *PromQLMetric) *NullablePromQLMetric {
	return &NullablePromQLMetric{value: val, isSet: true}
}

func (v NullablePromQLMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromQLMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
