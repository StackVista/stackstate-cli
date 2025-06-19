# PersesListVariableSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Display** | Pointer to [**PersesVariableDisplaySpec**](PersesVariableDisplaySpec.md) |  | [optional] 
**DefaultValue** | Pointer to [**PersesListVariableDefaultValue**](PersesListVariableDefaultValue.md) |  | [optional] 
**AllowAllValue** | **bool** |  | 
**AllowMultiple** | **bool** |  | 
**CustomAllValue** | Pointer to **string** |  | [optional] 
**CapturingRegexp** | Pointer to **string** |  | [optional] 
**Sort** | Pointer to [**PersesVariableSort**](PersesVariableSort.md) |  | [optional] 
**Plugin** | [**PersesPlugin**](PersesPlugin.md) |  | 

## Methods

### NewPersesListVariableSpec

`func NewPersesListVariableSpec(name string, allowAllValue bool, allowMultiple bool, plugin PersesPlugin, ) *PersesListVariableSpec`

NewPersesListVariableSpec instantiates a new PersesListVariableSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersesListVariableSpecWithDefaults

`func NewPersesListVariableSpecWithDefaults() *PersesListVariableSpec`

NewPersesListVariableSpecWithDefaults instantiates a new PersesListVariableSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PersesListVariableSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PersesListVariableSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PersesListVariableSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDisplay

`func (o *PersesListVariableSpec) GetDisplay() PersesVariableDisplaySpec`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PersesListVariableSpec) GetDisplayOk() (*PersesVariableDisplaySpec, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PersesListVariableSpec) SetDisplay(v PersesVariableDisplaySpec)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PersesListVariableSpec) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDefaultValue

`func (o *PersesListVariableSpec) GetDefaultValue() PersesListVariableDefaultValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *PersesListVariableSpec) GetDefaultValueOk() (*PersesListVariableDefaultValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *PersesListVariableSpec) SetDefaultValue(v PersesListVariableDefaultValue)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *PersesListVariableSpec) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetAllowAllValue

`func (o *PersesListVariableSpec) GetAllowAllValue() bool`

GetAllowAllValue returns the AllowAllValue field if non-nil, zero value otherwise.

### GetAllowAllValueOk

`func (o *PersesListVariableSpec) GetAllowAllValueOk() (*bool, bool)`

GetAllowAllValueOk returns a tuple with the AllowAllValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAllValue

`func (o *PersesListVariableSpec) SetAllowAllValue(v bool)`

SetAllowAllValue sets AllowAllValue field to given value.


### GetAllowMultiple

`func (o *PersesListVariableSpec) GetAllowMultiple() bool`

GetAllowMultiple returns the AllowMultiple field if non-nil, zero value otherwise.

### GetAllowMultipleOk

`func (o *PersesListVariableSpec) GetAllowMultipleOk() (*bool, bool)`

GetAllowMultipleOk returns a tuple with the AllowMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiple

`func (o *PersesListVariableSpec) SetAllowMultiple(v bool)`

SetAllowMultiple sets AllowMultiple field to given value.


### GetCustomAllValue

`func (o *PersesListVariableSpec) GetCustomAllValue() string`

GetCustomAllValue returns the CustomAllValue field if non-nil, zero value otherwise.

### GetCustomAllValueOk

`func (o *PersesListVariableSpec) GetCustomAllValueOk() (*string, bool)`

GetCustomAllValueOk returns a tuple with the CustomAllValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAllValue

`func (o *PersesListVariableSpec) SetCustomAllValue(v string)`

SetCustomAllValue sets CustomAllValue field to given value.

### HasCustomAllValue

`func (o *PersesListVariableSpec) HasCustomAllValue() bool`

HasCustomAllValue returns a boolean if a field has been set.

### GetCapturingRegexp

`func (o *PersesListVariableSpec) GetCapturingRegexp() string`

GetCapturingRegexp returns the CapturingRegexp field if non-nil, zero value otherwise.

### GetCapturingRegexpOk

`func (o *PersesListVariableSpec) GetCapturingRegexpOk() (*string, bool)`

GetCapturingRegexpOk returns a tuple with the CapturingRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturingRegexp

`func (o *PersesListVariableSpec) SetCapturingRegexp(v string)`

SetCapturingRegexp sets CapturingRegexp field to given value.

### HasCapturingRegexp

`func (o *PersesListVariableSpec) HasCapturingRegexp() bool`

HasCapturingRegexp returns a boolean if a field has been set.

### GetSort

`func (o *PersesListVariableSpec) GetSort() PersesVariableSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PersesListVariableSpec) GetSortOk() (*PersesVariableSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PersesListVariableSpec) SetSort(v PersesVariableSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PersesListVariableSpec) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetPlugin

`func (o *PersesListVariableSpec) GetPlugin() PersesPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *PersesListVariableSpec) GetPluginOk() (*PersesPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *PersesListVariableSpec) SetPlugin(v PersesPlugin)`

SetPlugin sets Plugin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


