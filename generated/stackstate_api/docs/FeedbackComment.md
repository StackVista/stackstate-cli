# FeedbackComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | **string** |  | 
**Text** | **string** |  | 
**Timestamp** | **int64** |  | 

## Methods

### NewFeedbackComment

`func NewFeedbackComment(author string, text string, timestamp int64, ) *FeedbackComment`

NewFeedbackComment instantiates a new FeedbackComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackCommentWithDefaults

`func NewFeedbackCommentWithDefaults() *FeedbackComment`

NewFeedbackCommentWithDefaults instantiates a new FeedbackComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *FeedbackComment) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *FeedbackComment) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *FeedbackComment) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetText

`func (o *FeedbackComment) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *FeedbackComment) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *FeedbackComment) SetText(v string)`

SetText sets Text field to given value.


### GetTimestamp

`func (o *FeedbackComment) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FeedbackComment) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FeedbackComment) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


