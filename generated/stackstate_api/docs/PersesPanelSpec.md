# PersesPanelSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Display** | Pointer to [**PersesPanelDisplay**](PersesPanelDisplay.md) |  | [optional] 
**Links** | Pointer to [**[]PersesLink**](PersesLink.md) |  | [optional] 
**Plugin** | Pointer to [**PersesPlugin**](PersesPlugin.md) |  | [optional] 
**Queries** | Pointer to [**[]PersesQuery**](PersesQuery.md) |  | [optional] 

## Methods

### NewPersesPanelSpec

`func NewPersesPanelSpec() *PersesPanelSpec`

NewPersesPanelSpec instantiates a new PersesPanelSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersesPanelSpecWithDefaults

`func NewPersesPanelSpecWithDefaults() *PersesPanelSpec`

NewPersesPanelSpecWithDefaults instantiates a new PersesPanelSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplay

`func (o *PersesPanelSpec) GetDisplay() PersesPanelDisplay`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PersesPanelSpec) GetDisplayOk() (*PersesPanelDisplay, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PersesPanelSpec) SetDisplay(v PersesPanelDisplay)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PersesPanelSpec) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetLinks

`func (o *PersesPanelSpec) GetLinks() []PersesLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PersesPanelSpec) GetLinksOk() (*[]PersesLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PersesPanelSpec) SetLinks(v []PersesLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PersesPanelSpec) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPlugin

`func (o *PersesPanelSpec) GetPlugin() PersesPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *PersesPanelSpec) GetPluginOk() (*PersesPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *PersesPanelSpec) SetPlugin(v PersesPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *PersesPanelSpec) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetQueries

`func (o *PersesPanelSpec) GetQueries() []PersesQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *PersesPanelSpec) GetQueriesOk() (*[]PersesQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *PersesPanelSpec) SetQueries(v []PersesQuery)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *PersesPanelSpec) HasQueries() bool`

HasQueries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


