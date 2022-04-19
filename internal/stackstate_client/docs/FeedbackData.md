# FeedbackData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Query** | Pointer to [**AnnotationMetricQuery**](AnnotationMetricQuery.md) |  | [optional] 
**Subject** | **string** |  | 
**Thumbsdown** | **[]string** |  | 
**Thumbsup** | **[]string** |  | 

## Methods

### NewFeedbackData

`func NewFeedbackData(type_ string, subject string, thumbsdown []string, thumbsup []string, ) *FeedbackData`

NewFeedbackData instantiates a new FeedbackData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackDataWithDefaults

`func NewFeedbackDataWithDefaults() *FeedbackData`

NewFeedbackDataWithDefaults instantiates a new FeedbackData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FeedbackData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeedbackData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeedbackData) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *FeedbackData) GetQuery() AnnotationMetricQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *FeedbackData) GetQueryOk() (*AnnotationMetricQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *FeedbackData) SetQuery(v AnnotationMetricQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *FeedbackData) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSubject

`func (o *FeedbackData) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *FeedbackData) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *FeedbackData) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetThumbsdown

`func (o *FeedbackData) GetThumbsdown() []string`

GetThumbsdown returns the Thumbsdown field if non-nil, zero value otherwise.

### GetThumbsdownOk

`func (o *FeedbackData) GetThumbsdownOk() (*[]string, bool)`

GetThumbsdownOk returns a tuple with the Thumbsdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbsdown

`func (o *FeedbackData) SetThumbsdown(v []string)`

SetThumbsdown sets Thumbsdown field to given value.


### GetThumbsup

`func (o *FeedbackData) GetThumbsup() []string`

GetThumbsup returns the Thumbsup field if non-nil, zero value otherwise.

### GetThumbsupOk

`func (o *FeedbackData) GetThumbsupOk() (*[]string, bool)`

GetThumbsupOk returns a tuple with the Thumbsup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbsup

`func (o *FeedbackData) SetThumbsup(v []string)`

SetThumbsup sets Thumbsup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


