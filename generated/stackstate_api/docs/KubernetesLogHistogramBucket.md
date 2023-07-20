# KubernetesLogHistogramBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Total logs record count in the bucket. | 
**StartTime** | **int32** | The bucket initial timestamp. | 
**EndTime** | **int32** | The bucket final timestamp. | 
**LogLevelBuckets** | [**[]KubernetesLogLevelHistogramBucket**](KubernetesLogLevelHistogramBucket.md) |  | 

## Methods

### NewKubernetesLogHistogramBucket

`func NewKubernetesLogHistogramBucket(count int64, startTime int32, endTime int32, logLevelBuckets []KubernetesLogLevelHistogramBucket, ) *KubernetesLogHistogramBucket`

NewKubernetesLogHistogramBucket instantiates a new KubernetesLogHistogramBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesLogHistogramBucketWithDefaults

`func NewKubernetesLogHistogramBucketWithDefaults() *KubernetesLogHistogramBucket`

NewKubernetesLogHistogramBucketWithDefaults instantiates a new KubernetesLogHistogramBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *KubernetesLogHistogramBucket) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KubernetesLogHistogramBucket) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KubernetesLogHistogramBucket) SetCount(v int64)`

SetCount sets Count field to given value.


### GetStartTime

`func (o *KubernetesLogHistogramBucket) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *KubernetesLogHistogramBucket) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *KubernetesLogHistogramBucket) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *KubernetesLogHistogramBucket) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *KubernetesLogHistogramBucket) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *KubernetesLogHistogramBucket) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.


### GetLogLevelBuckets

`func (o *KubernetesLogHistogramBucket) GetLogLevelBuckets() []KubernetesLogLevelHistogramBucket`

GetLogLevelBuckets returns the LogLevelBuckets field if non-nil, zero value otherwise.

### GetLogLevelBucketsOk

`func (o *KubernetesLogHistogramBucket) GetLogLevelBucketsOk() (*[]KubernetesLogLevelHistogramBucket, bool)`

GetLogLevelBucketsOk returns a tuple with the LogLevelBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevelBuckets

`func (o *KubernetesLogHistogramBucket) SetLogLevelBuckets(v []KubernetesLogLevelHistogramBucket)`

SetLogLevelBuckets sets LogLevelBuckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


