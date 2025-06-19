# TraceQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimarySpanFilter** | [**SpanFilter**](SpanFilter.md) |  | 
**SecondarySpanFilter** | Pointer to [**SpanFilter**](SpanFilter.md) |  | [optional] 
**SortBy** | Pointer to [**[]SpanSortOption**](SpanSortOption.md) |  | [optional] 

## Methods

### NewTraceQuery

`func NewTraceQuery(primarySpanFilter SpanFilter, ) *TraceQuery`

NewTraceQuery instantiates a new TraceQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceQueryWithDefaults

`func NewTraceQueryWithDefaults() *TraceQuery`

NewTraceQueryWithDefaults instantiates a new TraceQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimarySpanFilter

`func (o *TraceQuery) GetPrimarySpanFilter() SpanFilter`

GetPrimarySpanFilter returns the PrimarySpanFilter field if non-nil, zero value otherwise.

### GetPrimarySpanFilterOk

`func (o *TraceQuery) GetPrimarySpanFilterOk() (*SpanFilter, bool)`

GetPrimarySpanFilterOk returns a tuple with the PrimarySpanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySpanFilter

`func (o *TraceQuery) SetPrimarySpanFilter(v SpanFilter)`

SetPrimarySpanFilter sets PrimarySpanFilter field to given value.


### GetSecondarySpanFilter

`func (o *TraceQuery) GetSecondarySpanFilter() SpanFilter`

GetSecondarySpanFilter returns the SecondarySpanFilter field if non-nil, zero value otherwise.

### GetSecondarySpanFilterOk

`func (o *TraceQuery) GetSecondarySpanFilterOk() (*SpanFilter, bool)`

GetSecondarySpanFilterOk returns a tuple with the SecondarySpanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySpanFilter

`func (o *TraceQuery) SetSecondarySpanFilter(v SpanFilter)`

SetSecondarySpanFilter sets SecondarySpanFilter field to given value.

### HasSecondarySpanFilter

`func (o *TraceQuery) HasSecondarySpanFilter() bool`

HasSecondarySpanFilter returns a boolean if a field has been set.

### GetSortBy

`func (o *TraceQuery) GetSortBy() []SpanSortOption`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *TraceQuery) GetSortByOk() (*[]SpanSortOption, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *TraceQuery) SetSortBy(v []SpanSortOption)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *TraceQuery) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


