# KubernetesLogSeverityHistogramBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Total logs record count for a particular log severity in the bucket. | 
**Severity** | [**LogSeverity**](LogSeverity.md) |  | [default to LOGSEVERITY_OTHER]

## Methods

### NewKubernetesLogSeverityHistogramBucket

`func NewKubernetesLogSeverityHistogramBucket(count int64, severity LogSeverity, ) *KubernetesLogSeverityHistogramBucket`

NewKubernetesLogSeverityHistogramBucket instantiates a new KubernetesLogSeverityHistogramBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesLogSeverityHistogramBucketWithDefaults

`func NewKubernetesLogSeverityHistogramBucketWithDefaults() *KubernetesLogSeverityHistogramBucket`

NewKubernetesLogSeverityHistogramBucketWithDefaults instantiates a new KubernetesLogSeverityHistogramBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *KubernetesLogSeverityHistogramBucket) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KubernetesLogSeverityHistogramBucket) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KubernetesLogSeverityHistogramBucket) SetCount(v int64)`

SetCount sets Count field to given value.


### GetSeverity

`func (o *KubernetesLogSeverityHistogramBucket) GetSeverity() LogSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *KubernetesLogSeverityHistogramBucket) GetSeverityOk() (*LogSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *KubernetesLogSeverityHistogramBucket) SetSeverity(v LogSeverity)`

SetSeverity sets Severity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


