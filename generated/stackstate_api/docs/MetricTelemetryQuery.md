# MetricTelemetryQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Aggregation** | [**DownsamplingMethod**](DownsamplingMethod.md) |  | 
**Baseline** | Pointer to [**Baseline**](Baseline.md) |  | [optional] 
**Conditions** | [**[]TelemetryQueryCondition**](TelemetryQueryCondition.md) |  | 
**GroupBy** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**MetricField** | Pointer to **string** |  | [optional] 

## Methods

### NewMetricTelemetryQuery

`func NewMetricTelemetryQuery(type_ string, aggregation DownsamplingMethod, conditions []TelemetryQueryCondition, ) *MetricTelemetryQuery`

NewMetricTelemetryQuery instantiates a new MetricTelemetryQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricTelemetryQueryWithDefaults

`func NewMetricTelemetryQueryWithDefaults() *MetricTelemetryQuery`

NewMetricTelemetryQueryWithDefaults instantiates a new MetricTelemetryQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricTelemetryQuery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricTelemetryQuery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricTelemetryQuery) SetType(v string)`

SetType sets Type field to given value.


### GetAggregation

`func (o *MetricTelemetryQuery) GetAggregation() DownsamplingMethod`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MetricTelemetryQuery) GetAggregationOk() (*DownsamplingMethod, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MetricTelemetryQuery) SetAggregation(v DownsamplingMethod)`

SetAggregation sets Aggregation field to given value.


### GetBaseline

`func (o *MetricTelemetryQuery) GetBaseline() Baseline`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *MetricTelemetryQuery) GetBaselineOk() (*Baseline, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *MetricTelemetryQuery) SetBaseline(v Baseline)`

SetBaseline sets Baseline field to given value.

### HasBaseline

`func (o *MetricTelemetryQuery) HasBaseline() bool`

HasBaseline returns a boolean if a field has been set.

### GetConditions

`func (o *MetricTelemetryQuery) GetConditions() []TelemetryQueryCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *MetricTelemetryQuery) GetConditionsOk() (*[]TelemetryQueryCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *MetricTelemetryQuery) SetConditions(v []TelemetryQueryCondition)`

SetConditions sets Conditions field to given value.


### GetGroupBy

`func (o *MetricTelemetryQuery) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *MetricTelemetryQuery) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *MetricTelemetryQuery) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *MetricTelemetryQuery) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetId

`func (o *MetricTelemetryQuery) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricTelemetryQuery) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricTelemetryQuery) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MetricTelemetryQuery) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *MetricTelemetryQuery) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *MetricTelemetryQuery) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *MetricTelemetryQuery) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *MetricTelemetryQuery) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetMetricField

`func (o *MetricTelemetryQuery) GetMetricField() string`

GetMetricField returns the MetricField field if non-nil, zero value otherwise.

### GetMetricFieldOk

`func (o *MetricTelemetryQuery) GetMetricFieldOk() (*string, bool)`

GetMetricFieldOk returns a tuple with the MetricField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricField

`func (o *MetricTelemetryQuery) SetMetricField(v string)`

SetMetricField sets MetricField field to given value.

### HasMetricField

`func (o *MetricTelemetryQuery) HasMetricField() bool`

HasMetricField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


