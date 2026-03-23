# OverviewSorting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnId** | **string** |  | 
**Direction** | [**OverviewSortingDirection**](OverviewSortingDirection.md) |  | 

## Methods

### NewOverviewSorting

`func NewOverviewSorting(columnId string, direction OverviewSortingDirection, ) *OverviewSorting`

NewOverviewSorting instantiates a new OverviewSorting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewSortingWithDefaults

`func NewOverviewSortingWithDefaults() *OverviewSorting`

NewOverviewSortingWithDefaults instantiates a new OverviewSorting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnId

`func (o *OverviewSorting) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *OverviewSorting) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *OverviewSorting) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.


### GetDirection

`func (o *OverviewSorting) GetDirection() OverviewSortingDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *OverviewSorting) GetDirectionOk() (*OverviewSortingDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *OverviewSorting) SetDirection(v OverviewSortingDirection)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


