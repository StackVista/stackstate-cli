# DurationHistogramBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OkCount** | **int32** | Number of successful traces with a duration between (lowerNanos, upperNanos]. | 
**ErrorCount** | **int32** | Number of erroneous traces with a duration between (lowerNanos, upperNanos]. | 
**UnsetCount** | **int32** | Number of traces with unset status and a duration between (lowerNanos, upperNanos]. | 
**LowerNanos** | Pointer to **int64** | Lower limit on trace duration | [optional] 
**UpperNanos** | Pointer to **int64** | Upper limit on trace duration | [optional] 

## Methods

### NewDurationHistogramBucket

`func NewDurationHistogramBucket(okCount int32, errorCount int32, unsetCount int32, ) *DurationHistogramBucket`

NewDurationHistogramBucket instantiates a new DurationHistogramBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDurationHistogramBucketWithDefaults

`func NewDurationHistogramBucketWithDefaults() *DurationHistogramBucket`

NewDurationHistogramBucketWithDefaults instantiates a new DurationHistogramBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOkCount

`func (o *DurationHistogramBucket) GetOkCount() int32`

GetOkCount returns the OkCount field if non-nil, zero value otherwise.

### GetOkCountOk

`func (o *DurationHistogramBucket) GetOkCountOk() (*int32, bool)`

GetOkCountOk returns a tuple with the OkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOkCount

`func (o *DurationHistogramBucket) SetOkCount(v int32)`

SetOkCount sets OkCount field to given value.


### GetErrorCount

`func (o *DurationHistogramBucket) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *DurationHistogramBucket) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *DurationHistogramBucket) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.


### GetUnsetCount

`func (o *DurationHistogramBucket) GetUnsetCount() int32`

GetUnsetCount returns the UnsetCount field if non-nil, zero value otherwise.

### GetUnsetCountOk

`func (o *DurationHistogramBucket) GetUnsetCountOk() (*int32, bool)`

GetUnsetCountOk returns a tuple with the UnsetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsetCount

`func (o *DurationHistogramBucket) SetUnsetCount(v int32)`

SetUnsetCount sets UnsetCount field to given value.


### GetLowerNanos

`func (o *DurationHistogramBucket) GetLowerNanos() int64`

GetLowerNanos returns the LowerNanos field if non-nil, zero value otherwise.

### GetLowerNanosOk

`func (o *DurationHistogramBucket) GetLowerNanosOk() (*int64, bool)`

GetLowerNanosOk returns a tuple with the LowerNanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerNanos

`func (o *DurationHistogramBucket) SetLowerNanos(v int64)`

SetLowerNanos sets LowerNanos field to given value.

### HasLowerNanos

`func (o *DurationHistogramBucket) HasLowerNanos() bool`

HasLowerNanos returns a boolean if a field has been set.

### GetUpperNanos

`func (o *DurationHistogramBucket) GetUpperNanos() int64`

GetUpperNanos returns the UpperNanos field if non-nil, zero value otherwise.

### GetUpperNanosOk

`func (o *DurationHistogramBucket) GetUpperNanosOk() (*int64, bool)`

GetUpperNanosOk returns a tuple with the UpperNanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperNanos

`func (o *DurationHistogramBucket) SetUpperNanos(v int64)`

SetUpperNanos sets UpperNanos field to given value.

### HasUpperNanos

`func (o *DurationHistogramBucket) HasUpperNanos() bool`

HasUpperNanos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


