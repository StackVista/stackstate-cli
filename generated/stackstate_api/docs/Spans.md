# Spans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spans** | [**[]SpanSummary**](SpanSummary.md) | List of spans | 
**PageSize** | **int32** | Maximum number of the spans in the result. | 
**Page** | **int32** | The requested page. | 
**MatchesTotal** | **int64** | The total number of matching spans. | 

## Methods

### NewSpans

`func NewSpans(spans []SpanSummary, pageSize int32, page int32, matchesTotal int64, ) *Spans`

NewSpans instantiates a new Spans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpansWithDefaults

`func NewSpansWithDefaults() *Spans`

NewSpansWithDefaults instantiates a new Spans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpans

`func (o *Spans) GetSpans() []SpanSummary`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *Spans) GetSpansOk() (*[]SpanSummary, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *Spans) SetSpans(v []SpanSummary)`

SetSpans sets Spans field to given value.


### GetPageSize

`func (o *Spans) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Spans) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Spans) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetPage

`func (o *Spans) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *Spans) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *Spans) SetPage(v int32)`

SetPage sets Page field to given value.


### GetMatchesTotal

`func (o *Spans) GetMatchesTotal() int64`

GetMatchesTotal returns the MatchesTotal field if non-nil, zero value otherwise.

### GetMatchesTotalOk

`func (o *Spans) GetMatchesTotalOk() (*int64, bool)`

GetMatchesTotalOk returns a tuple with the MatchesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchesTotal

`func (o *Spans) SetMatchesTotal(v int64)`

SetMatchesTotal sets MatchesTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


