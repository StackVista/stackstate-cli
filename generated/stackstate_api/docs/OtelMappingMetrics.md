# OtelMappingMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketSizeSeconds** | **int32** |  | 
**LatencySeconds** | Pointer to [**[]MetricBucketValue**](MetricBucketValue.md) |  | [optional] 

## Methods

### NewOtelMappingMetrics

`func NewOtelMappingMetrics(bucketSizeSeconds int32, ) *OtelMappingMetrics`

NewOtelMappingMetrics instantiates a new OtelMappingMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtelMappingMetricsWithDefaults

`func NewOtelMappingMetricsWithDefaults() *OtelMappingMetrics`

NewOtelMappingMetricsWithDefaults instantiates a new OtelMappingMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketSizeSeconds

`func (o *OtelMappingMetrics) GetBucketSizeSeconds() int32`

GetBucketSizeSeconds returns the BucketSizeSeconds field if non-nil, zero value otherwise.

### GetBucketSizeSecondsOk

`func (o *OtelMappingMetrics) GetBucketSizeSecondsOk() (*int32, bool)`

GetBucketSizeSecondsOk returns a tuple with the BucketSizeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketSizeSeconds

`func (o *OtelMappingMetrics) SetBucketSizeSeconds(v int32)`

SetBucketSizeSeconds sets BucketSizeSeconds field to given value.


### GetLatencySeconds

`func (o *OtelMappingMetrics) GetLatencySeconds() []MetricBucketValue`

GetLatencySeconds returns the LatencySeconds field if non-nil, zero value otherwise.

### GetLatencySecondsOk

`func (o *OtelMappingMetrics) GetLatencySecondsOk() (*[]MetricBucketValue, bool)`

GetLatencySecondsOk returns a tuple with the LatencySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencySeconds

`func (o *OtelMappingMetrics) SetLatencySeconds(v []MetricBucketValue)`

SetLatencySeconds sets LatencySeconds field to given value.

### HasLatencySeconds

`func (o *OtelMappingMetrics) HasLatencySeconds() bool`

HasLatencySeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


