# FeedbackWithContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anomaly** | [**Annotation**](Annotation.md) |  | 
**Data** | [**[][]PointInner**]([]PointInner.md) |  | 
**Feedback** | [**FeedbackData**](FeedbackData.md) |  | 

## Methods

### NewFeedbackWithContext

`func NewFeedbackWithContext(anomaly Annotation, data [][]PointInner, feedback FeedbackData, ) *FeedbackWithContext`

NewFeedbackWithContext instantiates a new FeedbackWithContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackWithContextWithDefaults

`func NewFeedbackWithContextWithDefaults() *FeedbackWithContext`

NewFeedbackWithContextWithDefaults instantiates a new FeedbackWithContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomaly

`func (o *FeedbackWithContext) GetAnomaly() Annotation`

GetAnomaly returns the Anomaly field if non-nil, zero value otherwise.

### GetAnomalyOk

`func (o *FeedbackWithContext) GetAnomalyOk() (*Annotation, bool)`

GetAnomalyOk returns a tuple with the Anomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomaly

`func (o *FeedbackWithContext) SetAnomaly(v Annotation)`

SetAnomaly sets Anomaly field to given value.


### GetData

`func (o *FeedbackWithContext) GetData() [][]PointInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FeedbackWithContext) GetDataOk() (*[][]PointInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FeedbackWithContext) SetData(v [][]PointInner)`

SetData sets Data field to given value.


### GetFeedback

`func (o *FeedbackWithContext) GetFeedback() FeedbackData`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *FeedbackWithContext) GetFeedbackOk() (*FeedbackData, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *FeedbackWithContext) SetFeedback(v FeedbackData)`

SetFeedback sets Feedback field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


