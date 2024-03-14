# SpanQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**SpanFilter**](SpanFilter.md) |  | [optional] 
**SortBy** | Pointer to [**[]SpanSortOption**](SpanSortOption.md) |  | [optional] 

## Methods

### NewSpanQuery

`func NewSpanQuery() *SpanQuery`

NewSpanQuery instantiates a new SpanQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQueryWithDefaults

`func NewSpanQueryWithDefaults() *SpanQuery`

NewSpanQueryWithDefaults instantiates a new SpanQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *SpanQuery) GetFilter() SpanFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SpanQuery) GetFilterOk() (*SpanFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SpanQuery) SetFilter(v SpanFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SpanQuery) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSortBy

`func (o *SpanQuery) GetSortBy() []SpanSortOption`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *SpanQuery) GetSortByOk() (*[]SpanSortOption, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *SpanQuery) SetSortBy(v []SpanSortOption)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *SpanQuery) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


