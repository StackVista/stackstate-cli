# AnnotationMetricQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | **int64** |  | 
**Conditions** | [**[]TelemetryQueryCondition**](TelemetryQueryCondition.md) |  | 
**AggregationMethod** | [**DownsamplingMethod**](DownsamplingMethod.md) |  | 
**BucketSize** | **int64** |  | 
**MetricField** | Pointer to **string** |  | [optional] 
**QueryHash** | **string** |  | 

## Methods

### NewAnnotationMetricQuery

`func NewAnnotationMetricQuery(dataSourceId int64, conditions []TelemetryQueryCondition, aggregationMethod DownsamplingMethod, bucketSize int64, queryHash string, ) *AnnotationMetricQuery`

NewAnnotationMetricQuery instantiates a new AnnotationMetricQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationMetricQueryWithDefaults

`func NewAnnotationMetricQueryWithDefaults() *AnnotationMetricQuery`

NewAnnotationMetricQueryWithDefaults instantiates a new AnnotationMetricQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *AnnotationMetricQuery) GetDataSourceId() int64`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *AnnotationMetricQuery) GetDataSourceIdOk() (*int64, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *AnnotationMetricQuery) SetDataSourceId(v int64)`

SetDataSourceId sets DataSourceId field to given value.


### GetConditions

`func (o *AnnotationMetricQuery) GetConditions() []TelemetryQueryCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AnnotationMetricQuery) GetConditionsOk() (*[]TelemetryQueryCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AnnotationMetricQuery) SetConditions(v []TelemetryQueryCondition)`

SetConditions sets Conditions field to given value.


### GetAggregationMethod

`func (o *AnnotationMetricQuery) GetAggregationMethod() DownsamplingMethod`

GetAggregationMethod returns the AggregationMethod field if non-nil, zero value otherwise.

### GetAggregationMethodOk

`func (o *AnnotationMetricQuery) GetAggregationMethodOk() (*DownsamplingMethod, bool)`

GetAggregationMethodOk returns a tuple with the AggregationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationMethod

`func (o *AnnotationMetricQuery) SetAggregationMethod(v DownsamplingMethod)`

SetAggregationMethod sets AggregationMethod field to given value.


### GetBucketSize

`func (o *AnnotationMetricQuery) GetBucketSize() int64`

GetBucketSize returns the BucketSize field if non-nil, zero value otherwise.

### GetBucketSizeOk

`func (o *AnnotationMetricQuery) GetBucketSizeOk() (*int64, bool)`

GetBucketSizeOk returns a tuple with the BucketSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketSize

`func (o *AnnotationMetricQuery) SetBucketSize(v int64)`

SetBucketSize sets BucketSize field to given value.


### GetMetricField

`func (o *AnnotationMetricQuery) GetMetricField() string`

GetMetricField returns the MetricField field if non-nil, zero value otherwise.

### GetMetricFieldOk

`func (o *AnnotationMetricQuery) GetMetricFieldOk() (*string, bool)`

GetMetricFieldOk returns a tuple with the MetricField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricField

`func (o *AnnotationMetricQuery) SetMetricField(v string)`

SetMetricField sets MetricField field to given value.

### HasMetricField

`func (o *AnnotationMetricQuery) HasMetricField() bool`

HasMetricField returns a boolean if a field has been set.

### GetQueryHash

`func (o *AnnotationMetricQuery) GetQueryHash() string`

GetQueryHash returns the QueryHash field if non-nil, zero value otherwise.

### GetQueryHashOk

`func (o *AnnotationMetricQuery) GetQueryHashOk() (*string, bool)`

GetQueryHashOk returns a tuple with the QueryHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryHash

`func (o *AnnotationMetricQuery) SetQueryHash(v string)`

SetQueryHash sets QueryHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


