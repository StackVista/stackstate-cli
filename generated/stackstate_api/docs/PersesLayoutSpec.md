# PersesLayoutSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Display** | Pointer to [**PersesGridLayoutDisplay**](PersesGridLayoutDisplay.md) |  | [optional] 
**Items** | [**[]PersesGridItem**](PersesGridItem.md) |  | 

## Methods

### NewPersesLayoutSpec

`func NewPersesLayoutSpec(items []PersesGridItem, ) *PersesLayoutSpec`

NewPersesLayoutSpec instantiates a new PersesLayoutSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersesLayoutSpecWithDefaults

`func NewPersesLayoutSpecWithDefaults() *PersesLayoutSpec`

NewPersesLayoutSpecWithDefaults instantiates a new PersesLayoutSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplay

`func (o *PersesLayoutSpec) GetDisplay() PersesGridLayoutDisplay`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PersesLayoutSpec) GetDisplayOk() (*PersesGridLayoutDisplay, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PersesLayoutSpec) SetDisplay(v PersesGridLayoutDisplay)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PersesLayoutSpec) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetItems

`func (o *PersesLayoutSpec) GetItems() []PersesGridItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PersesLayoutSpec) GetItemsOk() (*[]PersesGridItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PersesLayoutSpec) SetItems(v []PersesGridItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


