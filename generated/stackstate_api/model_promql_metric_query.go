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

// PromqlMetricQuery struct for PromqlMetricQuery
type PromqlMetricQuery struct {
	Query                       string  `json:"query"`
	Alias                       *string `json:"alias,omitempty"`
	Description                 *string `json:"description,omitempty"`
	StartTime                   *int32  `json:"startTime,omitempty"`
	EndTime                     *int32  `json:"endTime,omitempty"`
	Step                        *string `json:"step,omitempty"`
	Unit                        *string `json:"unit,omitempty"`
	ComponentIdentifierTemplate *string `json:"componentIdentifierTemplate,omitempty"`
}

// NewPromqlMetricQuery instantiates a new PromqlMetricQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromqlMetricQuery(query string) *PromqlMetricQuery {
	this := PromqlMetricQuery{}
	this.Query = query
	return &this
}

// NewPromqlMetricQueryWithDefaults instantiates a new PromqlMetricQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromqlMetricQueryWithDefaults() *PromqlMetricQuery {
	this := PromqlMetricQuery{}
	return &this
}

// GetQuery returns the Query field value
func (o *PromqlMetricQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *PromqlMetricQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *PromqlMetricQuery) SetQuery(v string) {
	o.Query = v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *PromqlMetricQuery) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromqlMetricQuery) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *PromqlMetricQuery) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *PromqlMetricQuery) SetAlias(v string) {
	o.Alias = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PromqlMetricQuery) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromqlMetricQuery) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PromqlMetricQuery) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PromqlMetricQuery) SetDescription(v string) {
	o.Description = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *PromqlMetricQuery) GetStartTime() int32 {
	if o == nil || o.StartTime == nil {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromqlMetricQuery) GetStartTimeOk() (*int32, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *PromqlMetricQuery) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *PromqlMetricQuery) SetStartTime(v int32) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *PromqlMetricQuery) GetEndTime() int32 {
	if o == nil || o.EndTime == nil {
		var ret int32
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromqlMetricQuery) GetEndTimeOk() (*int32, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *PromqlMetricQuery) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int32 and assigns it to the EndTime field.
func (o *PromqlMetricQuery) SetEndTime(v int32) {
	o.EndTime = &v
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *PromqlMetricQuery) GetStep() string {
	if o == nil || o.Step == nil {
		var ret string
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromqlMetricQuery) GetStepOk() (*string, bool) {
	if o == nil || o.Step == nil {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *PromqlMetricQuery) HasStep() bool {
	if o != nil && o.Step != nil {
		return true
	}

	return false
}

// SetStep gets a reference to the given string and assigns it to the Step field.
func (o *PromqlMetricQuery) SetStep(v string) {
	o.Step = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *PromqlMetricQuery) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromqlMetricQuery) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *PromqlMetricQuery) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *PromqlMetricQuery) SetUnit(v string) {
	o.Unit = &v
}

// GetComponentIdentifierTemplate returns the ComponentIdentifierTemplate field value if set, zero value otherwise.
func (o *PromqlMetricQuery) GetComponentIdentifierTemplate() string {
	if o == nil || o.ComponentIdentifierTemplate == nil {
		var ret string
		return ret
	}
	return *o.ComponentIdentifierTemplate
}

// GetComponentIdentifierTemplateOk returns a tuple with the ComponentIdentifierTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromqlMetricQuery) GetComponentIdentifierTemplateOk() (*string, bool) {
	if o == nil || o.ComponentIdentifierTemplate == nil {
		return nil, false
	}
	return o.ComponentIdentifierTemplate, true
}

// HasComponentIdentifierTemplate returns a boolean if a field has been set.
func (o *PromqlMetricQuery) HasComponentIdentifierTemplate() bool {
	if o != nil && o.ComponentIdentifierTemplate != nil {
		return true
	}

	return false
}

// SetComponentIdentifierTemplate gets a reference to the given string and assigns it to the ComponentIdentifierTemplate field.
func (o *PromqlMetricQuery) SetComponentIdentifierTemplate(v string) {
	o.ComponentIdentifierTemplate = &v
}

func (o PromqlMetricQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["query"] = o.Query
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["endTime"] = o.EndTime
	}
	if o.Step != nil {
		toSerialize["step"] = o.Step
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.ComponentIdentifierTemplate != nil {
		toSerialize["componentIdentifierTemplate"] = o.ComponentIdentifierTemplate
	}
	return json.Marshal(toSerialize)
}

type NullablePromqlMetricQuery struct {
	value *PromqlMetricQuery
	isSet bool
}

func (v NullablePromqlMetricQuery) Get() *PromqlMetricQuery {
	return v.value
}

func (v *NullablePromqlMetricQuery) Set(val *PromqlMetricQuery) {
	v.value = val
	v.isSet = true
}

func (v NullablePromqlMetricQuery) IsSet() bool {
	return v.isSet
}

func (v *NullablePromqlMetricQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromqlMetricQuery(val *PromqlMetricQuery) *NullablePromqlMetricQuery {
	return &NullablePromqlMetricQuery{value: val, isSet: true}
}

func (v NullablePromqlMetricQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromqlMetricQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}