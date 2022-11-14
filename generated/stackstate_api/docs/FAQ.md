# FAQ

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | **string** |  |
**Answer** | **string** |  |

## Methods

### NewFAQ

`func NewFAQ(question string, answer string, ) *FAQ`

NewFAQ instantiates a new FAQ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFAQWithDefaults

`func NewFAQWithDefaults() *FAQ`

NewFAQWithDefaults instantiates a new FAQ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *FAQ) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *FAQ) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *FAQ) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetAnswer

`func (o *FAQ) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *FAQ) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *FAQ) SetAnswer(v string)`

SetAnswer sets Answer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
