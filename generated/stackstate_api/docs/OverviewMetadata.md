# OverviewMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**[]OverviewColumnMeta**](OverviewColumnMeta.md) |  | 
**Sorting** | [**[]OverviewSorting**](OverviewSorting.md) | The effective sorting applied to the overview results.  | 
**FixedColumns** | **int32** |  | 

## Methods

### NewOverviewMetadata

`func NewOverviewMetadata(columns []OverviewColumnMeta, sorting []OverviewSorting, fixedColumns int32, ) *OverviewMetadata`

NewOverviewMetadata instantiates a new OverviewMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewMetadataWithDefaults

`func NewOverviewMetadataWithDefaults() *OverviewMetadata`

NewOverviewMetadataWithDefaults instantiates a new OverviewMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *OverviewMetadata) GetColumns() []OverviewColumnMeta`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *OverviewMetadata) GetColumnsOk() (*[]OverviewColumnMeta, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *OverviewMetadata) SetColumns(v []OverviewColumnMeta)`

SetColumns sets Columns field to given value.


### GetSorting

`func (o *OverviewMetadata) GetSorting() []OverviewSorting`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *OverviewMetadata) GetSortingOk() (*[]OverviewSorting, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *OverviewMetadata) SetSorting(v []OverviewSorting)`

SetSorting sets Sorting field to given value.


### GetFixedColumns

`func (o *OverviewMetadata) GetFixedColumns() int32`

GetFixedColumns returns the FixedColumns field if non-nil, zero value otherwise.

### GetFixedColumnsOk

`func (o *OverviewMetadata) GetFixedColumnsOk() (*int32, bool)`

GetFixedColumnsOk returns a tuple with the FixedColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedColumns

`func (o *OverviewMetadata) SetFixedColumns(v int32)`

SetFixedColumns sets FixedColumns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


