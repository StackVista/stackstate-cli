# KubernetesLogLevelHistogramBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Total logs record count for a particular log level in the bucket. | 
**Level** | [**LogLevel**](LogLevel.md) |  | [default to LOGLEVEL_UNKNOWN]

## Methods

### NewKubernetesLogLevelHistogramBucket

`func NewKubernetesLogLevelHistogramBucket(count int64, level LogLevel, ) *KubernetesLogLevelHistogramBucket`

NewKubernetesLogLevelHistogramBucket instantiates a new KubernetesLogLevelHistogramBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesLogLevelHistogramBucketWithDefaults

`func NewKubernetesLogLevelHistogramBucketWithDefaults() *KubernetesLogLevelHistogramBucket`

NewKubernetesLogLevelHistogramBucketWithDefaults instantiates a new KubernetesLogLevelHistogramBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *KubernetesLogLevelHistogramBucket) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KubernetesLogLevelHistogramBucket) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KubernetesLogLevelHistogramBucket) SetCount(v int64)`

SetCount sets Count field to given value.


### GetLevel

`func (o *KubernetesLogLevelHistogramBucket) GetLevel() LogLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *KubernetesLogLevelHistogramBucket) GetLevelOk() (*LogLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *KubernetesLogLevelHistogramBucket) SetLevel(v LogLevel)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


