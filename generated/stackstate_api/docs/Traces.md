# Traces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Traces** | [**[]TraceIdentifier**](TraceIdentifier.md) | List of traces | 
**PageSize** | **int32** | Maximum number of the traces in the result. | 
**Page** | **int32** | The requested page. | 
**MatchesTotal** | **int64** | The total number of matching traces. | 

## Methods

### NewTraces

`func NewTraces(traces []TraceIdentifier, pageSize int32, page int32, matchesTotal int64, ) *Traces`

NewTraces instantiates a new Traces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTracesWithDefaults

`func NewTracesWithDefaults() *Traces`

NewTracesWithDefaults instantiates a new Traces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraces

`func (o *Traces) GetTraces() []TraceIdentifier`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *Traces) GetTracesOk() (*[]TraceIdentifier, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *Traces) SetTraces(v []TraceIdentifier)`

SetTraces sets Traces field to given value.


### GetPageSize

`func (o *Traces) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Traces) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Traces) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetPage

`func (o *Traces) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *Traces) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *Traces) SetPage(v int32)`

SetPage sets Page field to given value.


### GetMatchesTotal

`func (o *Traces) GetMatchesTotal() int64`

GetMatchesTotal returns the MatchesTotal field if non-nil, zero value otherwise.

### GetMatchesTotalOk

`func (o *Traces) GetMatchesTotalOk() (*int64, bool)`

GetMatchesTotalOk returns a tuple with the MatchesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchesTotal

`func (o *Traces) SetMatchesTotal(v int64)`

SetMatchesTotal sets MatchesTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


