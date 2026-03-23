# OverviewFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagFilters** | Pointer to **[]string** | Tags are grouped by their key.  For each key at least one of the values matches. | [optional] 
**ColumnFilters** | Pointer to [**map[string]OverviewColumnFilter**](OverviewColumnFilter.md) | Map of columnId to filter | [optional] 

## Methods

### NewOverviewFilters

`func NewOverviewFilters() *OverviewFilters`

NewOverviewFilters instantiates a new OverviewFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewFiltersWithDefaults

`func NewOverviewFiltersWithDefaults() *OverviewFilters`

NewOverviewFiltersWithDefaults instantiates a new OverviewFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagFilters

`func (o *OverviewFilters) GetTagFilters() []string`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *OverviewFilters) GetTagFiltersOk() (*[]string, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *OverviewFilters) SetTagFilters(v []string)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *OverviewFilters) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetColumnFilters

`func (o *OverviewFilters) GetColumnFilters() map[string]OverviewColumnFilter`

GetColumnFilters returns the ColumnFilters field if non-nil, zero value otherwise.

### GetColumnFiltersOk

`func (o *OverviewFilters) GetColumnFiltersOk() (*map[string]OverviewColumnFilter, bool)`

GetColumnFiltersOk returns a tuple with the ColumnFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnFilters

`func (o *OverviewFilters) SetColumnFilters(v map[string]OverviewColumnFilter)`

SetColumnFilters sets ColumnFilters field to given value.

### HasColumnFilters

`func (o *OverviewFilters) HasColumnFilters() bool`

HasColumnFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


