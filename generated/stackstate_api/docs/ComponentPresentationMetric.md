# ComponentPresentationMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MetricQueries** | Pointer to [**[]PresentationMetricQuery**](PresentationMetricQuery.md) |  | [optional] 
**Chart** | Pointer to [**Chart**](Chart.md) |  | [optional] 

## Methods

### NewComponentPresentationMetric

`func NewComponentPresentationMetric(metricId string, ) *ComponentPresentationMetric`

NewComponentPresentationMetric instantiates a new ComponentPresentationMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentPresentationMetricWithDefaults

`func NewComponentPresentationMetricWithDefaults() *ComponentPresentationMetric`

NewComponentPresentationMetricWithDefaults instantiates a new ComponentPresentationMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricId

`func (o *ComponentPresentationMetric) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *ComponentPresentationMetric) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *ComponentPresentationMetric) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.


### GetName

`func (o *ComponentPresentationMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentPresentationMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentPresentationMetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentPresentationMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ComponentPresentationMetric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComponentPresentationMetric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComponentPresentationMetric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComponentPresentationMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetricQueries

`func (o *ComponentPresentationMetric) GetMetricQueries() []PresentationMetricQuery`

GetMetricQueries returns the MetricQueries field if non-nil, zero value otherwise.

### GetMetricQueriesOk

`func (o *ComponentPresentationMetric) GetMetricQueriesOk() (*[]PresentationMetricQuery, bool)`

GetMetricQueriesOk returns a tuple with the MetricQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricQueries

`func (o *ComponentPresentationMetric) SetMetricQueries(v []PresentationMetricQuery)`

SetMetricQueries sets MetricQueries field to given value.

### HasMetricQueries

`func (o *ComponentPresentationMetric) HasMetricQueries() bool`

HasMetricQueries returns a boolean if a field has been set.

### GetChart

`func (o *ComponentPresentationMetric) GetChart() Chart`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *ComponentPresentationMetric) GetChartOk() (*Chart, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *ComponentPresentationMetric) SetChart(v Chart)`

SetChart sets Chart field to given value.

### HasChart

`func (o *ComponentPresentationMetric) HasChart() bool`

HasChart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


