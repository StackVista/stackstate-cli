# LatestTelemetryStreamMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**StreamId** | **int64** |  | 
**Metric** | Pointer to [**LatestTelemetryStreamMetricsMetric**](LatestTelemetryStreamMetricsMetric.md) |  | [optional] 

## Methods

### NewLatestTelemetryStreamMetrics

`func NewLatestTelemetryStreamMetrics(type_ string, streamId int64, ) *LatestTelemetryStreamMetrics`

NewLatestTelemetryStreamMetrics instantiates a new LatestTelemetryStreamMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLatestTelemetryStreamMetricsWithDefaults

`func NewLatestTelemetryStreamMetricsWithDefaults() *LatestTelemetryStreamMetrics`

NewLatestTelemetryStreamMetricsWithDefaults instantiates a new LatestTelemetryStreamMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LatestTelemetryStreamMetrics) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LatestTelemetryStreamMetrics) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LatestTelemetryStreamMetrics) SetType(v string)`

SetType sets Type field to given value.


### GetStreamId

`func (o *LatestTelemetryStreamMetrics) GetStreamId() int64`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *LatestTelemetryStreamMetrics) GetStreamIdOk() (*int64, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *LatestTelemetryStreamMetrics) SetStreamId(v int64)`

SetStreamId sets StreamId field to given value.


### GetMetric

`func (o *LatestTelemetryStreamMetrics) GetMetric() LatestTelemetryStreamMetricsMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *LatestTelemetryStreamMetrics) GetMetricOk() (*LatestTelemetryStreamMetricsMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *LatestTelemetryStreamMetrics) SetMetric(v LatestTelemetryStreamMetricsMetric)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *LatestTelemetryStreamMetrics) HasMetric() bool`

HasMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


