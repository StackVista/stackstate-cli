# TraceQueryMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceId** | **string** | The id of the trace | 
**Spans** | [**[]Span**](Span.md) | All spans for the trace | 
**Resources** | **map[string]map[string]string** | Map of resource identifiers to their corresponding resource attributes (key/value pairs). Keys are UUIDs used in the spans’ &#x60;resourceId&#x60; fields.  | 
**PrimaryFilterMatches** | **[]string** | Spans matching the primary filter | 
**SecondaryFilterMatches** | **[]string** | Spans matching the secondary filter | 

## Methods

### NewTraceQueryMatch

`func NewTraceQueryMatch(traceId string, spans []Span, resources map[string]map[string]string, primaryFilterMatches []string, secondaryFilterMatches []string, ) *TraceQueryMatch`

NewTraceQueryMatch instantiates a new TraceQueryMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceQueryMatchWithDefaults

`func NewTraceQueryMatchWithDefaults() *TraceQueryMatch`

NewTraceQueryMatchWithDefaults instantiates a new TraceQueryMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceId

`func (o *TraceQueryMatch) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *TraceQueryMatch) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *TraceQueryMatch) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpans

`func (o *TraceQueryMatch) GetSpans() []Span`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *TraceQueryMatch) GetSpansOk() (*[]Span, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *TraceQueryMatch) SetSpans(v []Span)`

SetSpans sets Spans field to given value.


### GetResources

`func (o *TraceQueryMatch) GetResources() map[string]map[string]string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *TraceQueryMatch) GetResourcesOk() (*map[string]map[string]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *TraceQueryMatch) SetResources(v map[string]map[string]string)`

SetResources sets Resources field to given value.


### GetPrimaryFilterMatches

`func (o *TraceQueryMatch) GetPrimaryFilterMatches() []string`

GetPrimaryFilterMatches returns the PrimaryFilterMatches field if non-nil, zero value otherwise.

### GetPrimaryFilterMatchesOk

`func (o *TraceQueryMatch) GetPrimaryFilterMatchesOk() (*[]string, bool)`

GetPrimaryFilterMatchesOk returns a tuple with the PrimaryFilterMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryFilterMatches

`func (o *TraceQueryMatch) SetPrimaryFilterMatches(v []string)`

SetPrimaryFilterMatches sets PrimaryFilterMatches field to given value.


### GetSecondaryFilterMatches

`func (o *TraceQueryMatch) GetSecondaryFilterMatches() []string`

GetSecondaryFilterMatches returns the SecondaryFilterMatches field if non-nil, zero value otherwise.

### GetSecondaryFilterMatchesOk

`func (o *TraceQueryMatch) GetSecondaryFilterMatchesOk() (*[]string, bool)`

GetSecondaryFilterMatchesOk returns a tuple with the SecondaryFilterMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryFilterMatches

`func (o *TraceQueryMatch) SetSecondaryFilterMatches(v []string)`

SetSecondaryFilterMatches sets SecondaryFilterMatches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


