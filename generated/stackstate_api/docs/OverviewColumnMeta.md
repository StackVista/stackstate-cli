# OverviewColumnMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnId** | **string** |  | 
**Title** | **string** |  | 
**Display** | [**OverviewColumnMetaDisplay**](OverviewColumnMetaDisplay.md) |  | 
**Sortable** | **bool** |  | 
**Searchable** | **bool** |  | 
**Order** | **int32** |  | 

## Methods

### NewOverviewColumnMeta

`func NewOverviewColumnMeta(columnId string, title string, display OverviewColumnMetaDisplay, sortable bool, searchable bool, order int32, ) *OverviewColumnMeta`

NewOverviewColumnMeta instantiates a new OverviewColumnMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewColumnMetaWithDefaults

`func NewOverviewColumnMetaWithDefaults() *OverviewColumnMeta`

NewOverviewColumnMetaWithDefaults instantiates a new OverviewColumnMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnId

`func (o *OverviewColumnMeta) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *OverviewColumnMeta) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *OverviewColumnMeta) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.


### GetTitle

`func (o *OverviewColumnMeta) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OverviewColumnMeta) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OverviewColumnMeta) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDisplay

`func (o *OverviewColumnMeta) GetDisplay() OverviewColumnMetaDisplay`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *OverviewColumnMeta) GetDisplayOk() (*OverviewColumnMetaDisplay, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *OverviewColumnMeta) SetDisplay(v OverviewColumnMetaDisplay)`

SetDisplay sets Display field to given value.


### GetSortable

`func (o *OverviewColumnMeta) GetSortable() bool`

GetSortable returns the Sortable field if non-nil, zero value otherwise.

### GetSortableOk

`func (o *OverviewColumnMeta) GetSortableOk() (*bool, bool)`

GetSortableOk returns a tuple with the Sortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortable

`func (o *OverviewColumnMeta) SetSortable(v bool)`

SetSortable sets Sortable field to given value.


### GetSearchable

`func (o *OverviewColumnMeta) GetSearchable() bool`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *OverviewColumnMeta) GetSearchableOk() (*bool, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *OverviewColumnMeta) SetSearchable(v bool)`

SetSearchable sets Searchable field to given value.


### GetOrder

`func (o *OverviewColumnMeta) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OverviewColumnMeta) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OverviewColumnMeta) SetOrder(v int32)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


