# MetricStreamAnomalyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**CheckedInterval** | [**TimeRange**](TimeRange.md) |  | 
**ElementName** | **string** |  | 
**Explanation** | **string** |  | 
**ModelInfo** | **map[string]interface{}** |  | 
**Query** | Pointer to [**AnnotationMetricQuery**](AnnotationMetricQuery.md) |  | [optional] 
**Severity** | [**AnomalySeverity**](AnomalySeverity.md) |  | 
**SeverityScore** | **float64** |  | 
**StreamName** | **string** |  | 

## Methods

### NewMetricStreamAnomalyData

`func NewMetricStreamAnomalyData(type_ string, checkedInterval TimeRange, elementName string, explanation string, modelInfo map[string]interface{}, severity AnomalySeverity, severityScore float64, streamName string, ) *MetricStreamAnomalyData`

NewMetricStreamAnomalyData instantiates a new MetricStreamAnomalyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricStreamAnomalyDataWithDefaults

`func NewMetricStreamAnomalyDataWithDefaults() *MetricStreamAnomalyData`

NewMetricStreamAnomalyDataWithDefaults instantiates a new MetricStreamAnomalyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricStreamAnomalyData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricStreamAnomalyData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricStreamAnomalyData) SetType(v string)`

SetType sets Type field to given value.


### GetCheckedInterval

`func (o *MetricStreamAnomalyData) GetCheckedInterval() TimeRange`

GetCheckedInterval returns the CheckedInterval field if non-nil, zero value otherwise.

### GetCheckedIntervalOk

`func (o *MetricStreamAnomalyData) GetCheckedIntervalOk() (*TimeRange, bool)`

GetCheckedIntervalOk returns a tuple with the CheckedInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedInterval

`func (o *MetricStreamAnomalyData) SetCheckedInterval(v TimeRange)`

SetCheckedInterval sets CheckedInterval field to given value.


### GetElementName

`func (o *MetricStreamAnomalyData) GetElementName() string`

GetElementName returns the ElementName field if non-nil, zero value otherwise.

### GetElementNameOk

`func (o *MetricStreamAnomalyData) GetElementNameOk() (*string, bool)`

GetElementNameOk returns a tuple with the ElementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementName

`func (o *MetricStreamAnomalyData) SetElementName(v string)`

SetElementName sets ElementName field to given value.


### GetExplanation

`func (o *MetricStreamAnomalyData) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *MetricStreamAnomalyData) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *MetricStreamAnomalyData) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.


### GetModelInfo

`func (o *MetricStreamAnomalyData) GetModelInfo() map[string]interface{}`

GetModelInfo returns the ModelInfo field if non-nil, zero value otherwise.

### GetModelInfoOk

`func (o *MetricStreamAnomalyData) GetModelInfoOk() (*map[string]interface{}, bool)`

GetModelInfoOk returns a tuple with the ModelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelInfo

`func (o *MetricStreamAnomalyData) SetModelInfo(v map[string]interface{})`

SetModelInfo sets ModelInfo field to given value.


### GetQuery

`func (o *MetricStreamAnomalyData) GetQuery() AnnotationMetricQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetricStreamAnomalyData) GetQueryOk() (*AnnotationMetricQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetricStreamAnomalyData) SetQuery(v AnnotationMetricQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MetricStreamAnomalyData) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSeverity

`func (o *MetricStreamAnomalyData) GetSeverity() AnomalySeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *MetricStreamAnomalyData) GetSeverityOk() (*AnomalySeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *MetricStreamAnomalyData) SetSeverity(v AnomalySeverity)`

SetSeverity sets Severity field to given value.


### GetSeverityScore

`func (o *MetricStreamAnomalyData) GetSeverityScore() float64`

GetSeverityScore returns the SeverityScore field if non-nil, zero value otherwise.

### GetSeverityScoreOk

`func (o *MetricStreamAnomalyData) GetSeverityScoreOk() (*float64, bool)`

GetSeverityScoreOk returns a tuple with the SeverityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityScore

`func (o *MetricStreamAnomalyData) SetSeverityScore(v float64)`

SetSeverityScore sets SeverityScore field to given value.


### GetStreamName

`func (o *MetricStreamAnomalyData) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *MetricStreamAnomalyData) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *MetricStreamAnomalyData) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


