# TraceFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimarySpanFilter** | [**SpanFilter**](SpanFilter.md) |  | 
**SecondarySpanFilter** | Pointer to [**SpanFilter**](SpanFilter.md) |  | [optional] 

## Methods

### NewTraceFilter

`func NewTraceFilter(primarySpanFilter SpanFilter, ) *TraceFilter`

NewTraceFilter instantiates a new TraceFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceFilterWithDefaults

`func NewTraceFilterWithDefaults() *TraceFilter`

NewTraceFilterWithDefaults instantiates a new TraceFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimarySpanFilter

`func (o *TraceFilter) GetPrimarySpanFilter() SpanFilter`

GetPrimarySpanFilter returns the PrimarySpanFilter field if non-nil, zero value otherwise.

### GetPrimarySpanFilterOk

`func (o *TraceFilter) GetPrimarySpanFilterOk() (*SpanFilter, bool)`

GetPrimarySpanFilterOk returns a tuple with the PrimarySpanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySpanFilter

`func (o *TraceFilter) SetPrimarySpanFilter(v SpanFilter)`

SetPrimarySpanFilter sets PrimarySpanFilter field to given value.


### GetSecondarySpanFilter

`func (o *TraceFilter) GetSecondarySpanFilter() SpanFilter`

GetSecondarySpanFilter returns the SecondarySpanFilter field if non-nil, zero value otherwise.

### GetSecondarySpanFilterOk

`func (o *TraceFilter) GetSecondarySpanFilterOk() (*SpanFilter, bool)`

GetSecondarySpanFilterOk returns a tuple with the SecondarySpanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySpanFilter

`func (o *TraceFilter) SetSecondarySpanFilter(v SpanFilter)`

SetSecondarySpanFilter sets SecondarySpanFilter field to given value.

### HasSecondarySpanFilter

`func (o *TraceFilter) HasSecondarySpanFilter() bool`

HasSecondarySpanFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


