# GetKubernetesLogsHistogramResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | [**[]KubernetesLogHistogramBucket**](KubernetesLogHistogramBucket.md) |  | 

## Methods

### NewGetKubernetesLogsHistogramResult

`func NewGetKubernetesLogsHistogramResult(buckets []KubernetesLogHistogramBucket, ) *GetKubernetesLogsHistogramResult`

NewGetKubernetesLogsHistogramResult instantiates a new GetKubernetesLogsHistogramResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesLogsHistogramResultWithDefaults

`func NewGetKubernetesLogsHistogramResultWithDefaults() *GetKubernetesLogsHistogramResult`

NewGetKubernetesLogsHistogramResultWithDefaults instantiates a new GetKubernetesLogsHistogramResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *GetKubernetesLogsHistogramResult) GetBuckets() []KubernetesLogHistogramBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *GetKubernetesLogsHistogramResult) GetBucketsOk() (*[]KubernetesLogHistogramBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *GetKubernetesLogsHistogramResult) SetBuckets(v []KubernetesLogHistogramBucket)`

SetBuckets sets Buckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


