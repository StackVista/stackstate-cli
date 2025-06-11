# PersesGridLayoutDisplay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The title of the grid layout. | 
**Collapse** | Pointer to [**PersesGridLayoutCollapse**](PersesGridLayoutCollapse.md) |  | [optional] 

## Methods

### NewPersesGridLayoutDisplay

`func NewPersesGridLayoutDisplay(title string, ) *PersesGridLayoutDisplay`

NewPersesGridLayoutDisplay instantiates a new PersesGridLayoutDisplay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersesGridLayoutDisplayWithDefaults

`func NewPersesGridLayoutDisplayWithDefaults() *PersesGridLayoutDisplay`

NewPersesGridLayoutDisplayWithDefaults instantiates a new PersesGridLayoutDisplay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PersesGridLayoutDisplay) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PersesGridLayoutDisplay) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PersesGridLayoutDisplay) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCollapse

`func (o *PersesGridLayoutDisplay) GetCollapse() PersesGridLayoutCollapse`

GetCollapse returns the Collapse field if non-nil, zero value otherwise.

### GetCollapseOk

`func (o *PersesGridLayoutDisplay) GetCollapseOk() (*PersesGridLayoutCollapse, bool)`

GetCollapseOk returns a tuple with the Collapse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapse

`func (o *PersesGridLayoutDisplay) SetCollapse(v PersesGridLayoutCollapse)`

SetCollapse sets Collapse field to given value.

### HasCollapse

`func (o *PersesGridLayoutDisplay) HasCollapse() bool`

HasCollapse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


