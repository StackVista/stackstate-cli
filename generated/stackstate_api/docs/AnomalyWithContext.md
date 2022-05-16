# AnomalyWithContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anomaly** | [**Annotation**](Annotation.md) |  | 
**Data** | [**[]Point**](Point.md) |  | 
**Feedback** | Pointer to [**FeedbackData**](FeedbackData.md) |  | [optional] 

## Methods

### NewAnomalyWithContext

`func NewAnomalyWithContext(anomaly Annotation, data []Point, ) *AnomalyWithContext`

NewAnomalyWithContext instantiates a new AnomalyWithContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyWithContextWithDefaults

`func NewAnomalyWithContextWithDefaults() *AnomalyWithContext`

NewAnomalyWithContextWithDefaults instantiates a new AnomalyWithContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomaly

`func (o *AnomalyWithContext) GetAnomaly() Annotation`

GetAnomaly returns the Anomaly field if non-nil, zero value otherwise.

### GetAnomalyOk

`func (o *AnomalyWithContext) GetAnomalyOk() (*Annotation, bool)`

GetAnomalyOk returns a tuple with the Anomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomaly

`func (o *AnomalyWithContext) SetAnomaly(v Annotation)`

SetAnomaly sets Anomaly field to given value.


### GetData

`func (o *AnomalyWithContext) GetData() []Point`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnomalyWithContext) GetDataOk() (*[]Point, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnomalyWithContext) SetData(v []Point)`

SetData sets Data field to given value.


### GetFeedback

`func (o *AnomalyWithContext) GetFeedback() FeedbackData`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *AnomalyWithContext) GetFeedbackOk() (*FeedbackData, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *AnomalyWithContext) SetFeedback(v FeedbackData)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *AnomalyWithContext) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


