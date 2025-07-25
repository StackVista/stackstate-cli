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

// PersesDashboardSpec struct for PersesDashboardSpec
type PersesDashboardSpec struct {
	// Datasources is an optional list of datasource definition.
	Datasources *map[string]PersesDatasourceSpec `json:"datasources,omitempty"`
	Display     *PersesDashboardDisplaySpec      `json:"display,omitempty"`
	// A Duration represents the elapsed time between two instants. It must be provided as a string like \"1h\", \"30m\", \"15s\".
	Duration *string                 `json:"duration,omitempty"`
	Layouts  []PersesLayout          `json:"layouts,omitempty"`
	Panels   *map[string]PersesPanel `json:"panels,omitempty"`
	// A Duration represents the elapsed time between two instants. It must be provided as a string like \"1h\", \"30m\", \"15s\".
	RefreshInterval *string               `json:"refreshInterval,omitempty"`
	Variables       []PersesVariableTypes `json:"variables,omitempty"`
}

// NewPersesDashboardSpec instantiates a new PersesDashboardSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersesDashboardSpec() *PersesDashboardSpec {
	this := PersesDashboardSpec{}
	return &this
}

// NewPersesDashboardSpecWithDefaults instantiates a new PersesDashboardSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersesDashboardSpecWithDefaults() *PersesDashboardSpec {
	this := PersesDashboardSpec{}
	return &this
}

// GetDatasources returns the Datasources field value if set, zero value otherwise.
func (o *PersesDashboardSpec) GetDatasources() map[string]PersesDatasourceSpec {
	if o == nil || o.Datasources == nil {
		var ret map[string]PersesDatasourceSpec
		return ret
	}
	return *o.Datasources
}

// GetDatasourcesOk returns a tuple with the Datasources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersesDashboardSpec) GetDatasourcesOk() (*map[string]PersesDatasourceSpec, bool) {
	if o == nil || o.Datasources == nil {
		return nil, false
	}
	return o.Datasources, true
}

// HasDatasources returns a boolean if a field has been set.
func (o *PersesDashboardSpec) HasDatasources() bool {
	if o != nil && o.Datasources != nil {
		return true
	}

	return false
}

// SetDatasources gets a reference to the given map[string]PersesDatasourceSpec and assigns it to the Datasources field.
func (o *PersesDashboardSpec) SetDatasources(v map[string]PersesDatasourceSpec) {
	o.Datasources = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PersesDashboardSpec) GetDisplay() PersesDashboardDisplaySpec {
	if o == nil || o.Display == nil {
		var ret PersesDashboardDisplaySpec
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersesDashboardSpec) GetDisplayOk() (*PersesDashboardDisplaySpec, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PersesDashboardSpec) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given PersesDashboardDisplaySpec and assigns it to the Display field.
func (o *PersesDashboardSpec) SetDisplay(v PersesDashboardDisplaySpec) {
	o.Display = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *PersesDashboardSpec) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersesDashboardSpec) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *PersesDashboardSpec) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *PersesDashboardSpec) SetDuration(v string) {
	o.Duration = &v
}

// GetLayouts returns the Layouts field value if set, zero value otherwise.
func (o *PersesDashboardSpec) GetLayouts() []PersesLayout {
	if o == nil || o.Layouts == nil {
		var ret []PersesLayout
		return ret
	}
	return o.Layouts
}

// GetLayoutsOk returns a tuple with the Layouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersesDashboardSpec) GetLayoutsOk() ([]PersesLayout, bool) {
	if o == nil || o.Layouts == nil {
		return nil, false
	}
	return o.Layouts, true
}

// HasLayouts returns a boolean if a field has been set.
func (o *PersesDashboardSpec) HasLayouts() bool {
	if o != nil && o.Layouts != nil {
		return true
	}

	return false
}

// SetLayouts gets a reference to the given []PersesLayout and assigns it to the Layouts field.
func (o *PersesDashboardSpec) SetLayouts(v []PersesLayout) {
	o.Layouts = v
}

// GetPanels returns the Panels field value if set, zero value otherwise.
func (o *PersesDashboardSpec) GetPanels() map[string]PersesPanel {
	if o == nil || o.Panels == nil {
		var ret map[string]PersesPanel
		return ret
	}
	return *o.Panels
}

// GetPanelsOk returns a tuple with the Panels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersesDashboardSpec) GetPanelsOk() (*map[string]PersesPanel, bool) {
	if o == nil || o.Panels == nil {
		return nil, false
	}
	return o.Panels, true
}

// HasPanels returns a boolean if a field has been set.
func (o *PersesDashboardSpec) HasPanels() bool {
	if o != nil && o.Panels != nil {
		return true
	}

	return false
}

// SetPanels gets a reference to the given map[string]PersesPanel and assigns it to the Panels field.
func (o *PersesDashboardSpec) SetPanels(v map[string]PersesPanel) {
	o.Panels = &v
}

// GetRefreshInterval returns the RefreshInterval field value if set, zero value otherwise.
func (o *PersesDashboardSpec) GetRefreshInterval() string {
	if o == nil || o.RefreshInterval == nil {
		var ret string
		return ret
	}
	return *o.RefreshInterval
}

// GetRefreshIntervalOk returns a tuple with the RefreshInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersesDashboardSpec) GetRefreshIntervalOk() (*string, bool) {
	if o == nil || o.RefreshInterval == nil {
		return nil, false
	}
	return o.RefreshInterval, true
}

// HasRefreshInterval returns a boolean if a field has been set.
func (o *PersesDashboardSpec) HasRefreshInterval() bool {
	if o != nil && o.RefreshInterval != nil {
		return true
	}

	return false
}

// SetRefreshInterval gets a reference to the given string and assigns it to the RefreshInterval field.
func (o *PersesDashboardSpec) SetRefreshInterval(v string) {
	o.RefreshInterval = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *PersesDashboardSpec) GetVariables() []PersesVariableTypes {
	if o == nil || o.Variables == nil {
		var ret []PersesVariableTypes
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersesDashboardSpec) GetVariablesOk() ([]PersesVariableTypes, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *PersesDashboardSpec) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []PersesVariableTypes and assigns it to the Variables field.
func (o *PersesDashboardSpec) SetVariables(v []PersesVariableTypes) {
	o.Variables = v
}

func (o PersesDashboardSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Datasources != nil {
		toSerialize["datasources"] = o.Datasources
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.Layouts != nil {
		toSerialize["layouts"] = o.Layouts
	}
	if o.Panels != nil {
		toSerialize["panels"] = o.Panels
	}
	if o.RefreshInterval != nil {
		toSerialize["refreshInterval"] = o.RefreshInterval
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	return json.Marshal(toSerialize)
}

type NullablePersesDashboardSpec struct {
	value *PersesDashboardSpec
	isSet bool
}

func (v NullablePersesDashboardSpec) Get() *PersesDashboardSpec {
	return v.value
}

func (v *NullablePersesDashboardSpec) Set(val *PersesDashboardSpec) {
	v.value = val
	v.isSet = true
}

func (v NullablePersesDashboardSpec) IsSet() bool {
	return v.isSet
}

func (v *NullablePersesDashboardSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersesDashboardSpec(val *PersesDashboardSpec) *NullablePersesDashboardSpec {
	return &NullablePersesDashboardSpec{value: val, isSet: true}
}

func (v NullablePersesDashboardSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersesDashboardSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
