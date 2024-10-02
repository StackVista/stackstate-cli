# TraceQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpanFilter** | [**SpanFilter**](SpanFilter.md) |  | 
**TraceAttributes** | **map[string][]string** | Filter traces by 1 or more attributes | 
**SortBy** | Pointer to [**[]SpanSortOption**](SpanSortOption.md) |  | [optional] 

## Methods

### NewTraceQuery

`func NewTraceQuery(spanFilter SpanFilter, traceAttributes map[string][]string, ) *TraceQuery`

NewTraceQuery instantiates a new TraceQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceQueryWithDefaults

`func NewTraceQueryWithDefaults() *TraceQuery`

NewTraceQueryWithDefaults instantiates a new TraceQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpanFilter

`func (o *TraceQuery) GetSpanFilter() SpanFilter`

GetSpanFilter returns the SpanFilter field if non-nil, zero value otherwise.

### GetSpanFilterOk

`func (o *TraceQuery) GetSpanFilterOk() (*SpanFilter, bool)`

GetSpanFilterOk returns a tuple with the SpanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanFilter

`func (o *TraceQuery) SetSpanFilter(v SpanFilter)`

SetSpanFilter sets SpanFilter field to given value.


### GetTraceAttributes

`func (o *TraceQuery) GetTraceAttributes() map[string][]string`

GetTraceAttributes returns the TraceAttributes field if non-nil, zero value otherwise.

### GetTraceAttributesOk

`func (o *TraceQuery) GetTraceAttributesOk() (*map[string][]string, bool)`

GetTraceAttributesOk returns a tuple with the TraceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceAttributes

`func (o *TraceQuery) SetTraceAttributes(v map[string][]string)`

SetTraceAttributes sets TraceAttributes field to given value.


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


