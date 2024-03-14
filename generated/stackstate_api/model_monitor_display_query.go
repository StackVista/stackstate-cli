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

// MonitorDisplayQuery struct for MonitorDisplayQuery
type MonitorDisplayQuery struct {
	Query                       string  `json:"query"`
	Alias                       *string `json:"alias,omitempty"`
	ComponentIdentifierTemplate *string `json:"componentIdentifierTemplate,omitempty"`
}

// NewMonitorDisplayQuery instantiates a new MonitorDisplayQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorDisplayQuery(query string) *MonitorDisplayQuery {
	this := MonitorDisplayQuery{}
	this.Query = query
	return &this
}

// NewMonitorDisplayQueryWithDefaults instantiates a new MonitorDisplayQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorDisplayQueryWithDefaults() *MonitorDisplayQuery {
	this := MonitorDisplayQuery{}
	return &this
}

// GetQuery returns the Query field value
func (o *MonitorDisplayQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *MonitorDisplayQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *MonitorDisplayQuery) SetQuery(v string) {
	o.Query = v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *MonitorDisplayQuery) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorDisplayQuery) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *MonitorDisplayQuery) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *MonitorDisplayQuery) SetAlias(v string) {
	o.Alias = &v
}

// GetComponentIdentifierTemplate returns the ComponentIdentifierTemplate field value if set, zero value otherwise.
func (o *MonitorDisplayQuery) GetComponentIdentifierTemplate() string {
	if o == nil || o.ComponentIdentifierTemplate == nil {
		var ret string
		return ret
	}
	return *o.ComponentIdentifierTemplate
}

// GetComponentIdentifierTemplateOk returns a tuple with the ComponentIdentifierTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorDisplayQuery) GetComponentIdentifierTemplateOk() (*string, bool) {
	if o == nil || o.ComponentIdentifierTemplate == nil {
		return nil, false
	}
	return o.ComponentIdentifierTemplate, true
}

// HasComponentIdentifierTemplate returns a boolean if a field has been set.
func (o *MonitorDisplayQuery) HasComponentIdentifierTemplate() bool {
	if o != nil && o.ComponentIdentifierTemplate != nil {
		return true
	}

	return false
}

// SetComponentIdentifierTemplate gets a reference to the given string and assigns it to the ComponentIdentifierTemplate field.
func (o *MonitorDisplayQuery) SetComponentIdentifierTemplate(v string) {
	o.ComponentIdentifierTemplate = &v
}

func (o MonitorDisplayQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["query"] = o.Query
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.ComponentIdentifierTemplate != nil {
		toSerialize["componentIdentifierTemplate"] = o.ComponentIdentifierTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorDisplayQuery struct {
	value *MonitorDisplayQuery
	isSet bool
}

func (v NullableMonitorDisplayQuery) Get() *MonitorDisplayQuery {
	return v.value
}

func (v *NullableMonitorDisplayQuery) Set(val *MonitorDisplayQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorDisplayQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorDisplayQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorDisplayQuery(val *MonitorDisplayQuery) *NullableMonitorDisplayQuery {
	return &NullableMonitorDisplayQuery{value: val, isSet: true}
}

func (v NullableMonitorDisplayQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorDisplayQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
