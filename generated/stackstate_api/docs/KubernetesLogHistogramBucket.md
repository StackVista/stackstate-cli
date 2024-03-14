# KubernetesLogHistogramBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Total logs record count in the bucket. | 
**StartTime** | **int32** | The bucket initial timestamp. | 
**EndTime** | **int32** | The bucket final timestamp. | 
**LogSeverityBuckets** | [**[]KubernetesLogSeverityHistogramBucket**](KubernetesLogSeverityHistogramBucket.md) |  | 

## Methods

### NewKubernetesLogHistogramBucket

`func NewKubernetesLogHistogramBucket(count int64, startTime int32, endTime int32, logSeverityBuckets []KubernetesLogSeverityHistogramBucket, ) *KubernetesLogHistogramBucket`

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


### GetLogSeverityBuckets

`func (o *KubernetesLogHistogramBucket) GetLogSeverityBuckets() []KubernetesLogSeverityHistogramBucket`

GetLogSeverityBuckets returns the LogSeverityBuckets field if non-nil, zero value otherwise.

### GetLogSeverityBucketsOk

`func (o *KubernetesLogHistogramBucket) GetLogSeverityBucketsOk() (*[]KubernetesLogSeverityHistogramBucket, bool)`

GetLogSeverityBucketsOk returns a tuple with the LogSeverityBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSeverityBuckets

`func (o *KubernetesLogHistogramBucket) SetLogSeverityBuckets(v []KubernetesLogSeverityHistogramBucket)`

SetLogSeverityBuckets sets LogSeverityBuckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


