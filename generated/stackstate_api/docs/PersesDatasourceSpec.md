# PersesDatasourceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** |  | [optional] 
**Display** | Pointer to [**PersesDashboardDisplaySpec**](PersesDashboardDisplaySpec.md) |  | [optional] 
**Plugin** | Pointer to [**PersesPlugin**](PersesPlugin.md) |  | [optional] 

## Methods

### NewPersesDatasourceSpec

`func NewPersesDatasourceSpec() *PersesDatasourceSpec`

NewPersesDatasourceSpec instantiates a new PersesDatasourceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersesDatasourceSpecWithDefaults

`func NewPersesDatasourceSpecWithDefaults() *PersesDatasourceSpec`

NewPersesDatasourceSpecWithDefaults instantiates a new PersesDatasourceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *PersesDatasourceSpec) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PersesDatasourceSpec) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PersesDatasourceSpec) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PersesDatasourceSpec) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDisplay

`func (o *PersesDatasourceSpec) GetDisplay() PersesDashboardDisplaySpec`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PersesDatasourceSpec) GetDisplayOk() (*PersesDashboardDisplaySpec, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PersesDatasourceSpec) SetDisplay(v PersesDashboardDisplaySpec)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PersesDatasourceSpec) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetPlugin

`func (o *PersesDatasourceSpec) GetPlugin() PersesPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *PersesDatasourceSpec) GetPluginOk() (*PersesPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *PersesDatasourceSpec) SetPlugin(v PersesPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *PersesDatasourceSpec) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


