# PresentationOverviewSorting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnId** | **string** |  | 
**Direction** | Pointer to [**PresentationOverviewSortingDirection**](PresentationOverviewSortingDirection.md) |  | [optional] 

## Methods

### NewPresentationOverviewSorting

`func NewPresentationOverviewSorting(columnId string, ) *PresentationOverviewSorting`

NewPresentationOverviewSorting instantiates a new PresentationOverviewSorting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationOverviewSortingWithDefaults

`func NewPresentationOverviewSortingWithDefaults() *PresentationOverviewSorting`

NewPresentationOverviewSortingWithDefaults instantiates a new PresentationOverviewSorting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnId

`func (o *PresentationOverviewSorting) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *PresentationOverviewSorting) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *PresentationOverviewSorting) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.


### GetDirection

`func (o *PresentationOverviewSorting) GetDirection() PresentationOverviewSortingDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PresentationOverviewSorting) GetDirectionOk() (*PresentationOverviewSortingDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PresentationOverviewSorting) SetDirection(v PresentationOverviewSortingDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *PresentationOverviewSorting) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


