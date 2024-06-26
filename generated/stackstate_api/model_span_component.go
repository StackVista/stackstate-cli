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

// SpanComponent struct for SpanComponent
type SpanComponent struct {
	Id          int64            `json:"id"`
	Identifier  string           `json:"identifier"`
	Name        string           `json:"name"`
	Type        string           `json:"type"`
	HealthState HealthStateValue `json:"healthState"`
}

// NewSpanComponent instantiates a new SpanComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanComponent(id int64, identifier string, name string, type_ string, healthState HealthStateValue) *SpanComponent {
	this := SpanComponent{}
	this.Id = id
	this.Identifier = identifier
	this.Name = name
	this.Type = type_
	this.HealthState = healthState
	return &this
}

// NewSpanComponentWithDefaults instantiates a new SpanComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanComponentWithDefaults() *SpanComponent {
	this := SpanComponent{}
	return &this
}

// GetId returns the Id field value
func (o *SpanComponent) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SpanComponent) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SpanComponent) SetId(v int64) {
	o.Id = v
}

// GetIdentifier returns the Identifier field value
func (o *SpanComponent) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *SpanComponent) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *SpanComponent) SetIdentifier(v string) {
	o.Identifier = v
}

// GetName returns the Name field value
func (o *SpanComponent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SpanComponent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SpanComponent) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *SpanComponent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpanComponent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpanComponent) SetType(v string) {
	o.Type = v
}

// GetHealthState returns the HealthState field value
func (o *SpanComponent) GetHealthState() HealthStateValue {
	if o == nil {
		var ret HealthStateValue
		return ret
	}

	return o.HealthState
}

// GetHealthStateOk returns a tuple with the HealthState field value
// and a boolean to check if the value has been set.
func (o *SpanComponent) GetHealthStateOk() (*HealthStateValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthState, true
}

// SetHealthState sets field value
func (o *SpanComponent) SetHealthState(v HealthStateValue) {
	o.HealthState = v
}

func (o SpanComponent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["healthState"] = o.HealthState
	}
	return json.Marshal(toSerialize)
}

type NullableSpanComponent struct {
	value *SpanComponent
	isSet bool
}

func (v NullableSpanComponent) Get() *SpanComponent {
	return v.value
}

func (v *NullableSpanComponent) Set(val *SpanComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanComponent(val *SpanComponent) *NullableSpanComponent {
	return &NullableSpanComponent{value: val, isSet: true}
}

func (v NullableSpanComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
