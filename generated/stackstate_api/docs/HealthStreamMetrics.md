# HealthStreamMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketSizeSeconds** | **int32** |  | 
**LatencySeconds** | Pointer to [**[]MetricBucketValue**](MetricBucketValue.md) |  | [optional] 
**MessagePerSecond** | Pointer to [**[]MetricBucketValue**](MetricBucketValue.md) |  | [optional] 
**CreatesPerSecond** | Pointer to [**[]MetricBucketValue**](MetricBucketValue.md) |  | [optional] 
**UpdatesPerSecond** | Pointer to [**[]MetricBucketValue**](MetricBucketValue.md) |  | [optional] 
**DeletesPerSecond** | Pointer to [**[]MetricBucketValue**](MetricBucketValue.md) |  | [optional] 

## Methods

### NewHealthStreamMetrics

`func NewHealthStreamMetrics(bucketSizeSeconds int32, ) *HealthStreamMetrics`

NewHealthStreamMetrics instantiates a new HealthStreamMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthStreamMetricsWithDefaults

`func NewHealthStreamMetricsWithDefaults() *HealthStreamMetrics`

NewHealthStreamMetricsWithDefaults instantiates a new HealthStreamMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketSizeSeconds

`func (o *HealthStreamMetrics) GetBucketSizeSeconds() int32`

GetBucketSizeSeconds returns the BucketSizeSeconds field if non-nil, zero value otherwise.

### GetBucketSizeSecondsOk

`func (o *HealthStreamMetrics) GetBucketSizeSecondsOk() (*int32, bool)`

GetBucketSizeSecondsOk returns a tuple with the BucketSizeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketSizeSeconds

`func (o *HealthStreamMetrics) SetBucketSizeSeconds(v int32)`

SetBucketSizeSeconds sets BucketSizeSeconds field to given value.


### GetLatencySeconds

`func (o *HealthStreamMetrics) GetLatencySeconds() []MetricBucketValue`

GetLatencySeconds returns the LatencySeconds field if non-nil, zero value otherwise.

### GetLatencySecondsOk

`func (o *HealthStreamMetrics) GetLatencySecondsOk() (*[]MetricBucketValue, bool)`

GetLatencySecondsOk returns a tuple with the LatencySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencySeconds

`func (o *HealthStreamMetrics) SetLatencySeconds(v []MetricBucketValue)`

SetLatencySeconds sets LatencySeconds field to given value.

### HasLatencySeconds

`func (o *HealthStreamMetrics) HasLatencySeconds() bool`

HasLatencySeconds returns a boolean if a field has been set.

### GetMessagePerSecond

`func (o *HealthStreamMetrics) GetMessagePerSecond() []MetricBucketValue`

GetMessagePerSecond returns the MessagePerSecond field if non-nil, zero value otherwise.

### GetMessagePerSecondOk

`func (o *HealthStreamMetrics) GetMessagePerSecondOk() (*[]MetricBucketValue, bool)`

GetMessagePerSecondOk returns a tuple with the MessagePerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagePerSecond

`func (o *HealthStreamMetrics) SetMessagePerSecond(v []MetricBucketValue)`

SetMessagePerSecond sets MessagePerSecond field to given value.

### HasMessagePerSecond

`func (o *HealthStreamMetrics) HasMessagePerSecond() bool`

HasMessagePerSecond returns a boolean if a field has been set.

### GetCreatesPerSecond

`func (o *HealthStreamMetrics) GetCreatesPerSecond() []MetricBucketValue`

GetCreatesPerSecond returns the CreatesPerSecond field if non-nil, zero value otherwise.

### GetCreatesPerSecondOk

`func (o *HealthStreamMetrics) GetCreatesPerSecondOk() (*[]MetricBucketValue, bool)`

GetCreatesPerSecondOk returns a tuple with the CreatesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatesPerSecond

`func (o *HealthStreamMetrics) SetCreatesPerSecond(v []MetricBucketValue)`

SetCreatesPerSecond sets CreatesPerSecond field to given value.

### HasCreatesPerSecond

`func (o *HealthStreamMetrics) HasCreatesPerSecond() bool`

HasCreatesPerSecond returns a boolean if a field has been set.

### GetUpdatesPerSecond

`func (o *HealthStreamMetrics) GetUpdatesPerSecond() []MetricBucketValue`

GetUpdatesPerSecond returns the UpdatesPerSecond field if non-nil, zero value otherwise.

### GetUpdatesPerSecondOk

`func (o *HealthStreamMetrics) GetUpdatesPerSecondOk() (*[]MetricBucketValue, bool)`

GetUpdatesPerSecondOk returns a tuple with the UpdatesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatesPerSecond

`func (o *HealthStreamMetrics) SetUpdatesPerSecond(v []MetricBucketValue)`

SetUpdatesPerSecond sets UpdatesPerSecond field to given value.

### HasUpdatesPerSecond

`func (o *HealthStreamMetrics) HasUpdatesPerSecond() bool`

HasUpdatesPerSecond returns a boolean if a field has been set.

### GetDeletesPerSecond

`func (o *HealthStreamMetrics) GetDeletesPerSecond() []MetricBucketValue`

GetDeletesPerSecond returns the DeletesPerSecond field if non-nil, zero value otherwise.

### GetDeletesPerSecondOk

`func (o *HealthStreamMetrics) GetDeletesPerSecondOk() (*[]MetricBucketValue, bool)`

GetDeletesPerSecondOk returns a tuple with the DeletesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletesPerSecond

`func (o *HealthStreamMetrics) SetDeletesPerSecond(v []MetricBucketValue)`

SetDeletesPerSecond sets DeletesPerSecond field to given value.

### HasDeletesPerSecond

`func (o *HealthStreamMetrics) HasDeletesPerSecond() bool`

HasDeletesPerSecond returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


