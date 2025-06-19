# PersesDashboardSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasources** | Pointer to [**map[string]PersesDatasourceSpec**](PersesDatasourceSpec.md) | Datasources is an optional list of datasource definition. | [optional] 
**Display** | Pointer to [**PersesDashboardDisplaySpec**](PersesDashboardDisplaySpec.md) |  | [optional] 
**Duration** | Pointer to **string** | A Duration represents the elapsed time between two instants. It must be provided as a string like \&quot;1h\&quot;, \&quot;30m\&quot;, \&quot;15s\&quot;.  | [optional] 
**Layouts** | Pointer to [**[]PersesLayout**](PersesLayout.md) |  | [optional] 
**Panels** | Pointer to [**map[string]PersesPanel**](PersesPanel.md) |  | [optional] 
**RefreshInterval** | Pointer to **string** | A Duration represents the elapsed time between two instants. It must be provided as a string like \&quot;1h\&quot;, \&quot;30m\&quot;, \&quot;15s\&quot;.  | [optional] 
**Variables** | Pointer to [**[]PersesVariableTypes**](PersesVariableTypes.md) |  | [optional] 

## Methods

### NewPersesDashboardSpec

`func NewPersesDashboardSpec() *PersesDashboardSpec`

NewPersesDashboardSpec instantiates a new PersesDashboardSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersesDashboardSpecWithDefaults

`func NewPersesDashboardSpecWithDefaults() *PersesDashboardSpec`

NewPersesDashboardSpecWithDefaults instantiates a new PersesDashboardSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasources

`func (o *PersesDashboardSpec) GetDatasources() map[string]PersesDatasourceSpec`

GetDatasources returns the Datasources field if non-nil, zero value otherwise.

### GetDatasourcesOk

`func (o *PersesDashboardSpec) GetDatasourcesOk() (*map[string]PersesDatasourceSpec, bool)`

GetDatasourcesOk returns a tuple with the Datasources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasources

`func (o *PersesDashboardSpec) SetDatasources(v map[string]PersesDatasourceSpec)`

SetDatasources sets Datasources field to given value.

### HasDatasources

`func (o *PersesDashboardSpec) HasDatasources() bool`

HasDatasources returns a boolean if a field has been set.

### GetDisplay

`func (o *PersesDashboardSpec) GetDisplay() PersesDashboardDisplaySpec`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PersesDashboardSpec) GetDisplayOk() (*PersesDashboardDisplaySpec, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PersesDashboardSpec) SetDisplay(v PersesDashboardDisplaySpec)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PersesDashboardSpec) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDuration

`func (o *PersesDashboardSpec) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PersesDashboardSpec) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PersesDashboardSpec) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PersesDashboardSpec) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetLayouts

`func (o *PersesDashboardSpec) GetLayouts() []PersesLayout`

GetLayouts returns the Layouts field if non-nil, zero value otherwise.

### GetLayoutsOk

`func (o *PersesDashboardSpec) GetLayoutsOk() (*[]PersesLayout, bool)`

GetLayoutsOk returns a tuple with the Layouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayouts

`func (o *PersesDashboardSpec) SetLayouts(v []PersesLayout)`

SetLayouts sets Layouts field to given value.

### HasLayouts

`func (o *PersesDashboardSpec) HasLayouts() bool`

HasLayouts returns a boolean if a field has been set.

### GetPanels

`func (o *PersesDashboardSpec) GetPanels() map[string]PersesPanel`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *PersesDashboardSpec) GetPanelsOk() (*map[string]PersesPanel, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *PersesDashboardSpec) SetPanels(v map[string]PersesPanel)`

SetPanels sets Panels field to given value.

### HasPanels

`func (o *PersesDashboardSpec) HasPanels() bool`

HasPanels returns a boolean if a field has been set.

### GetRefreshInterval

`func (o *PersesDashboardSpec) GetRefreshInterval() string`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *PersesDashboardSpec) GetRefreshIntervalOk() (*string, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *PersesDashboardSpec) SetRefreshInterval(v string)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *PersesDashboardSpec) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### GetVariables

`func (o *PersesDashboardSpec) GetVariables() []PersesVariableTypes`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PersesDashboardSpec) GetVariablesOk() (*[]PersesVariableTypes, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PersesDashboardSpec) SetVariables(v []PersesVariableTypes)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *PersesDashboardSpec) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


