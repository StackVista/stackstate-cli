# OrderedComponentPresentationMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MetricQueries** | Pointer to [**[]PresentationMetricQuery**](PresentationMetricQuery.md) |  | [optional] 
**Chart** | Pointer to [**Chart**](Chart.md) |  | [optional] 
**Order** | Pointer to **float64** |  | [optional] 

## Methods

### NewOrderedComponentPresentationMetric

`func NewOrderedComponentPresentationMetric(metricId string, ) *OrderedComponentPresentationMetric`

NewOrderedComponentPresentationMetric instantiates a new OrderedComponentPresentationMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderedComponentPresentationMetricWithDefaults

`func NewOrderedComponentPresentationMetricWithDefaults() *OrderedComponentPresentationMetric`

NewOrderedComponentPresentationMetricWithDefaults instantiates a new OrderedComponentPresentationMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricId

`func (o *OrderedComponentPresentationMetric) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *OrderedComponentPresentationMetric) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *OrderedComponentPresentationMetric) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.


### GetName

`func (o *OrderedComponentPresentationMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderedComponentPresentationMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderedComponentPresentationMetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderedComponentPresentationMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OrderedComponentPresentationMetric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrderedComponentPresentationMetric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrderedComponentPresentationMetric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrderedComponentPresentationMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetricQueries

`func (o *OrderedComponentPresentationMetric) GetMetricQueries() []PresentationMetricQuery`

GetMetricQueries returns the MetricQueries field if non-nil, zero value otherwise.

### GetMetricQueriesOk

`func (o *OrderedComponentPresentationMetric) GetMetricQueriesOk() (*[]PresentationMetricQuery, bool)`

GetMetricQueriesOk returns a tuple with the MetricQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricQueries

`func (o *OrderedComponentPresentationMetric) SetMetricQueries(v []PresentationMetricQuery)`

SetMetricQueries sets MetricQueries field to given value.

### HasMetricQueries

`func (o *OrderedComponentPresentationMetric) HasMetricQueries() bool`

HasMetricQueries returns a boolean if a field has been set.

### GetChart

`func (o *OrderedComponentPresentationMetric) GetChart() Chart`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *OrderedComponentPresentationMetric) GetChartOk() (*Chart, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *OrderedComponentPresentationMetric) SetChart(v Chart)`

SetChart sets Chart field to given value.

### HasChart

`func (o *OrderedComponentPresentationMetric) HasChart() bool`

HasChart returns a boolean if a field has been set.

### GetOrder

`func (o *OrderedComponentPresentationMetric) GetOrder() float64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OrderedComponentPresentationMetric) GetOrderOk() (*float64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OrderedComponentPresentationMetric) SetOrder(v float64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *OrderedComponentPresentationMetric) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


