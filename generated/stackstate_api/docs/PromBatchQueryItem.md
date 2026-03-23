# PromBatchQueryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**Query** | **string** |  | 
**Time** | Pointer to **string** |  | [optional] 
**Step** | **string** |  | 
**Timeout** | Pointer to **string** |  | [optional] 
**PostFilter** | Pointer to **[]string** |  | [optional] 
**Start** | **string** |  | 
**End** | **string** |  | 
**Aligned** | Pointer to **bool** |  | [optional] 
**MaxNumberOfDataPoints** | Pointer to **int64** |  | [optional] 

## Methods

### NewPromBatchQueryItem

`func NewPromBatchQueryItem(id string, type_ string, query string, step string, start string, end string, ) *PromBatchQueryItem`

NewPromBatchQueryItem instantiates a new PromBatchQueryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromBatchQueryItemWithDefaults

`func NewPromBatchQueryItemWithDefaults() *PromBatchQueryItem`

NewPromBatchQueryItemWithDefaults instantiates a new PromBatchQueryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PromBatchQueryItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PromBatchQueryItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PromBatchQueryItem) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PromBatchQueryItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromBatchQueryItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromBatchQueryItem) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *PromBatchQueryItem) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PromBatchQueryItem) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PromBatchQueryItem) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTime

`func (o *PromBatchQueryItem) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PromBatchQueryItem) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PromBatchQueryItem) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *PromBatchQueryItem) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetStep

`func (o *PromBatchQueryItem) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *PromBatchQueryItem) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *PromBatchQueryItem) SetStep(v string)`

SetStep sets Step field to given value.


### GetTimeout

`func (o *PromBatchQueryItem) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PromBatchQueryItem) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PromBatchQueryItem) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PromBatchQueryItem) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetPostFilter

`func (o *PromBatchQueryItem) GetPostFilter() []string`

GetPostFilter returns the PostFilter field if non-nil, zero value otherwise.

### GetPostFilterOk

`func (o *PromBatchQueryItem) GetPostFilterOk() (*[]string, bool)`

GetPostFilterOk returns a tuple with the PostFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostFilter

`func (o *PromBatchQueryItem) SetPostFilter(v []string)`

SetPostFilter sets PostFilter field to given value.

### HasPostFilter

`func (o *PromBatchQueryItem) HasPostFilter() bool`

HasPostFilter returns a boolean if a field has been set.

### GetStart

`func (o *PromBatchQueryItem) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PromBatchQueryItem) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PromBatchQueryItem) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *PromBatchQueryItem) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *PromBatchQueryItem) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *PromBatchQueryItem) SetEnd(v string)`

SetEnd sets End field to given value.


### GetAligned

`func (o *PromBatchQueryItem) GetAligned() bool`

GetAligned returns the Aligned field if non-nil, zero value otherwise.

### GetAlignedOk

`func (o *PromBatchQueryItem) GetAlignedOk() (*bool, bool)`

GetAlignedOk returns a tuple with the Aligned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAligned

`func (o *PromBatchQueryItem) SetAligned(v bool)`

SetAligned sets Aligned field to given value.

### HasAligned

`func (o *PromBatchQueryItem) HasAligned() bool`

HasAligned returns a boolean if a field has been set.

### GetMaxNumberOfDataPoints

`func (o *PromBatchQueryItem) GetMaxNumberOfDataPoints() int64`

GetMaxNumberOfDataPoints returns the MaxNumberOfDataPoints field if non-nil, zero value otherwise.

### GetMaxNumberOfDataPointsOk

`func (o *PromBatchQueryItem) GetMaxNumberOfDataPointsOk() (*int64, bool)`

GetMaxNumberOfDataPointsOk returns a tuple with the MaxNumberOfDataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfDataPoints

`func (o *PromBatchQueryItem) SetMaxNumberOfDataPoints(v int64)`

SetMaxNumberOfDataPoints sets MaxNumberOfDataPoints field to given value.

### HasMaxNumberOfDataPoints

`func (o *PromBatchQueryItem) HasMaxNumberOfDataPoints() bool`

HasMaxNumberOfDataPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


