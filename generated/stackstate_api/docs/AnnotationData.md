# AnnotationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Severity** | [**AnomalySeverity**](AnomalySeverity.md) |  | 
**SeverityScore** | **float64** |  | 
**CheckedInterval** | [**TimeRange**](TimeRange.md) |  | 
**Explanation** | **string** |  | 
**ModelInfo** | **map[string]interface{}** |  | 
**ElementName** | **string** |  | 
**StreamName** | **string** |  | 
**Query** | Pointer to [**AnnotationMetricQuery**](AnnotationMetricQuery.md) |  | [optional] 
**Annotation** | **map[string]interface{}** |  | 
**Subject** | **string** |  | 
**Thumbsup** | **[]string** |  | 
**Thumbsdown** | **[]string** |  | 
**Comments** | [**[]FeedbackComment**](FeedbackComment.md) |  | 

## Methods

### NewAnnotationData

`func NewAnnotationData(type_ string, severity AnomalySeverity, severityScore float64, checkedInterval TimeRange, explanation string, modelInfo map[string]interface{}, elementName string, streamName string, annotation map[string]interface{}, subject string, thumbsup []string, thumbsdown []string, comments []FeedbackComment, ) *AnnotationData`

NewAnnotationData instantiates a new AnnotationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationDataWithDefaults

`func NewAnnotationDataWithDefaults() *AnnotationData`

NewAnnotationDataWithDefaults instantiates a new AnnotationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AnnotationData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnnotationData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnnotationData) SetType(v string)`

SetType sets Type field to given value.


### GetSeverity

`func (o *AnnotationData) GetSeverity() AnomalySeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AnnotationData) GetSeverityOk() (*AnomalySeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AnnotationData) SetSeverity(v AnomalySeverity)`

SetSeverity sets Severity field to given value.


### GetSeverityScore

`func (o *AnnotationData) GetSeverityScore() float64`

GetSeverityScore returns the SeverityScore field if non-nil, zero value otherwise.

### GetSeverityScoreOk

`func (o *AnnotationData) GetSeverityScoreOk() (*float64, bool)`

GetSeverityScoreOk returns a tuple with the SeverityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityScore

`func (o *AnnotationData) SetSeverityScore(v float64)`

SetSeverityScore sets SeverityScore field to given value.


### GetCheckedInterval

`func (o *AnnotationData) GetCheckedInterval() TimeRange`

GetCheckedInterval returns the CheckedInterval field if non-nil, zero value otherwise.

### GetCheckedIntervalOk

`func (o *AnnotationData) GetCheckedIntervalOk() (*TimeRange, bool)`

GetCheckedIntervalOk returns a tuple with the CheckedInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedInterval

`func (o *AnnotationData) SetCheckedInterval(v TimeRange)`

SetCheckedInterval sets CheckedInterval field to given value.


### GetExplanation

`func (o *AnnotationData) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *AnnotationData) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *AnnotationData) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.


### GetModelInfo

`func (o *AnnotationData) GetModelInfo() map[string]interface{}`

GetModelInfo returns the ModelInfo field if non-nil, zero value otherwise.

### GetModelInfoOk

`func (o *AnnotationData) GetModelInfoOk() (*map[string]interface{}, bool)`

GetModelInfoOk returns a tuple with the ModelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelInfo

`func (o *AnnotationData) SetModelInfo(v map[string]interface{})`

SetModelInfo sets ModelInfo field to given value.


### GetElementName

`func (o *AnnotationData) GetElementName() string`

GetElementName returns the ElementName field if non-nil, zero value otherwise.

### GetElementNameOk

`func (o *AnnotationData) GetElementNameOk() (*string, bool)`

GetElementNameOk returns a tuple with the ElementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementName

`func (o *AnnotationData) SetElementName(v string)`

SetElementName sets ElementName field to given value.


### GetStreamName

`func (o *AnnotationData) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *AnnotationData) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *AnnotationData) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.


### GetQuery

`func (o *AnnotationData) GetQuery() AnnotationMetricQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AnnotationData) GetQueryOk() (*AnnotationMetricQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AnnotationData) SetQuery(v AnnotationMetricQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *AnnotationData) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetAnnotation

`func (o *AnnotationData) GetAnnotation() map[string]interface{}`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *AnnotationData) GetAnnotationOk() (*map[string]interface{}, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *AnnotationData) SetAnnotation(v map[string]interface{})`

SetAnnotation sets Annotation field to given value.


### GetSubject

`func (o *AnnotationData) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AnnotationData) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AnnotationData) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetThumbsup

`func (o *AnnotationData) GetThumbsup() []string`

GetThumbsup returns the Thumbsup field if non-nil, zero value otherwise.

### GetThumbsupOk

`func (o *AnnotationData) GetThumbsupOk() (*[]string, bool)`

GetThumbsupOk returns a tuple with the Thumbsup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbsup

`func (o *AnnotationData) SetThumbsup(v []string)`

SetThumbsup sets Thumbsup field to given value.


### GetThumbsdown

`func (o *AnnotationData) GetThumbsdown() []string`

GetThumbsdown returns the Thumbsdown field if non-nil, zero value otherwise.

### GetThumbsdownOk

`func (o *AnnotationData) GetThumbsdownOk() (*[]string, bool)`

GetThumbsdownOk returns a tuple with the Thumbsdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbsdown

`func (o *AnnotationData) SetThumbsdown(v []string)`

SetThumbsdown sets Thumbsdown field to given value.


### GetComments

`func (o *AnnotationData) GetComments() []FeedbackComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *AnnotationData) GetCommentsOk() (*[]FeedbackComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *AnnotationData) SetComments(v []FeedbackComment)`

SetComments sets Comments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


