# PromBatchRangeQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**Query** | **string** |  | 
**Start** | **string** |  | 
**End** | **string** |  | 
**Step** | **string** |  | 
**Timeout** | Pointer to **string** |  | [optional] 
**Aligned** | Pointer to **bool** |  | [optional] 
**MaxNumberOfDataPoints** | Pointer to **int64** |  | [optional] 
**PostFilter** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPromBatchRangeQuery

`func NewPromBatchRangeQuery(id string, type_ string, query string, start string, end string, step string, ) *PromBatchRangeQuery`

NewPromBatchRangeQuery instantiates a new PromBatchRangeQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromBatchRangeQueryWithDefaults

`func NewPromBatchRangeQueryWithDefaults() *PromBatchRangeQuery`

NewPromBatchRangeQueryWithDefaults instantiates a new PromBatchRangeQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PromBatchRangeQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PromBatchRangeQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PromBatchRangeQuery) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PromBatchRangeQuery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromBatchRangeQuery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromBatchRangeQuery) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *PromBatchRangeQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PromBatchRangeQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PromBatchRangeQuery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetStart

`func (o *PromBatchRangeQuery) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PromBatchRangeQuery) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PromBatchRangeQuery) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *PromBatchRangeQuery) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *PromBatchRangeQuery) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *PromBatchRangeQuery) SetEnd(v string)`

SetEnd sets End field to given value.


### GetStep

`func (o *PromBatchRangeQuery) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *PromBatchRangeQuery) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *PromBatchRangeQuery) SetStep(v string)`

SetStep sets Step field to given value.


### GetTimeout

`func (o *PromBatchRangeQuery) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PromBatchRangeQuery) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PromBatchRangeQuery) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PromBatchRangeQuery) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetAligned

`func (o *PromBatchRangeQuery) GetAligned() bool`

GetAligned returns the Aligned field if non-nil, zero value otherwise.

### GetAlignedOk

`func (o *PromBatchRangeQuery) GetAlignedOk() (*bool, bool)`

GetAlignedOk returns a tuple with the Aligned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAligned

`func (o *PromBatchRangeQuery) SetAligned(v bool)`

SetAligned sets Aligned field to given value.

### HasAligned

`func (o *PromBatchRangeQuery) HasAligned() bool`

HasAligned returns a boolean if a field has been set.

### GetMaxNumberOfDataPoints

`func (o *PromBatchRangeQuery) GetMaxNumberOfDataPoints() int64`

GetMaxNumberOfDataPoints returns the MaxNumberOfDataPoints field if non-nil, zero value otherwise.

### GetMaxNumberOfDataPointsOk

`func (o *PromBatchRangeQuery) GetMaxNumberOfDataPointsOk() (*int64, bool)`

GetMaxNumberOfDataPointsOk returns a tuple with the MaxNumberOfDataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfDataPoints

`func (o *PromBatchRangeQuery) SetMaxNumberOfDataPoints(v int64)`

SetMaxNumberOfDataPoints sets MaxNumberOfDataPoints field to given value.

### HasMaxNumberOfDataPoints

`func (o *PromBatchRangeQuery) HasMaxNumberOfDataPoints() bool`

HasMaxNumberOfDataPoints returns a boolean if a field has been set.

### GetPostFilter

`func (o *PromBatchRangeQuery) GetPostFilter() []string`

GetPostFilter returns the PostFilter field if non-nil, zero value otherwise.

### GetPostFilterOk

`func (o *PromBatchRangeQuery) GetPostFilterOk() (*[]string, bool)`

GetPostFilterOk returns a tuple with the PostFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostFilter

`func (o *PromBatchRangeQuery) SetPostFilter(v []string)`

SetPostFilter sets PostFilter field to given value.

### HasPostFilter

`func (o *PromBatchRangeQuery) HasPostFilter() bool`

HasPostFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


