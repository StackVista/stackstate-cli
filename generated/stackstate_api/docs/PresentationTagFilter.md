# PresentationTagFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**FilterId** | **string** |  | 
**DisplayName** | [**PresentationFilterName**](PresentationFilterName.md) |  | 
**TagKey** | **string** | Tag key to use for filtering. | 

## Methods

### NewPresentationTagFilter

`func NewPresentationTagFilter(type_ string, filterId string, displayName PresentationFilterName, tagKey string, ) *PresentationTagFilter`

NewPresentationTagFilter instantiates a new PresentationTagFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationTagFilterWithDefaults

`func NewPresentationTagFilterWithDefaults() *PresentationTagFilter`

NewPresentationTagFilterWithDefaults instantiates a new PresentationTagFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PresentationTagFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PresentationTagFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PresentationTagFilter) SetType(v string)`

SetType sets Type field to given value.


### GetFilterId

`func (o *PresentationTagFilter) GetFilterId() string`

GetFilterId returns the FilterId field if non-nil, zero value otherwise.

### GetFilterIdOk

`func (o *PresentationTagFilter) GetFilterIdOk() (*string, bool)`

GetFilterIdOk returns a tuple with the FilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterId

`func (o *PresentationTagFilter) SetFilterId(v string)`

SetFilterId sets FilterId field to given value.


### GetDisplayName

`func (o *PresentationTagFilter) GetDisplayName() PresentationFilterName`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PresentationTagFilter) GetDisplayNameOk() (*PresentationFilterName, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PresentationTagFilter) SetDisplayName(v PresentationFilterName)`

SetDisplayName sets DisplayName field to given value.


### GetTagKey

`func (o *PresentationTagFilter) GetTagKey() string`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *PresentationTagFilter) GetTagKeyOk() (*string, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *PresentationTagFilter) SetTagKey(v string)`

SetTagKey sets TagKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


