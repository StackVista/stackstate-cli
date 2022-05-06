# MetricStreamNoAnomalyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**CheckedInterval** | [**TimeRange**](TimeRange.md) |  | 
**Explanation** | **string** |  | 
**ModelInfo** | **map[string]interface{}** |  | 
**Query** | Pointer to [**AnnotationMetricQuery**](AnnotationMetricQuery.md) |  | [optional] 

## Methods

### NewMetricStreamNoAnomalyData

`func NewMetricStreamNoAnomalyData(type_ string, checkedInterval TimeRange, explanation string, modelInfo map[string]interface{}, ) *MetricStreamNoAnomalyData`

NewMetricStreamNoAnomalyData instantiates a new MetricStreamNoAnomalyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricStreamNoAnomalyDataWithDefaults

`func NewMetricStreamNoAnomalyDataWithDefaults() *MetricStreamNoAnomalyData`

NewMetricStreamNoAnomalyDataWithDefaults instantiates a new MetricStreamNoAnomalyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricStreamNoAnomalyData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricStreamNoAnomalyData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricStreamNoAnomalyData) SetType(v string)`

SetType sets Type field to given value.


### GetCheckedInterval

`func (o *MetricStreamNoAnomalyData) GetCheckedInterval() TimeRange`

GetCheckedInterval returns the CheckedInterval field if non-nil, zero value otherwise.

### GetCheckedIntervalOk

`func (o *MetricStreamNoAnomalyData) GetCheckedIntervalOk() (*TimeRange, bool)`

GetCheckedIntervalOk returns a tuple with the CheckedInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedInterval

`func (o *MetricStreamNoAnomalyData) SetCheckedInterval(v TimeRange)`

SetCheckedInterval sets CheckedInterval field to given value.


### GetExplanation

`func (o *MetricStreamNoAnomalyData) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *MetricStreamNoAnomalyData) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *MetricStreamNoAnomalyData) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.


### GetModelInfo

`func (o *MetricStreamNoAnomalyData) GetModelInfo() map[string]interface{}`

GetModelInfo returns the ModelInfo field if non-nil, zero value otherwise.

### GetModelInfoOk

`func (o *MetricStreamNoAnomalyData) GetModelInfoOk() (*map[string]interface{}, bool)`

GetModelInfoOk returns a tuple with the ModelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelInfo

`func (o *MetricStreamNoAnomalyData) SetModelInfo(v map[string]interface{})`

SetModelInfo sets ModelInfo field to given value.


### GetQuery

`func (o *MetricStreamNoAnomalyData) GetQuery() AnnotationMetricQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetricStreamNoAnomalyData) GetQueryOk() (*AnnotationMetricQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetricStreamNoAnomalyData) SetQuery(v AnnotationMetricQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MetricStreamNoAnomalyData) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


