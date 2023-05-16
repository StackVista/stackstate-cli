# KubernetesLogHighlight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartIndex** | **int32** | Index of character in the message to start the highlight (0-based) | 
**EndIndex** | **int32** | Index of character in the message to end the highlight (0-based) | 

## Methods

### NewKubernetesLogHighlight

`func NewKubernetesLogHighlight(startIndex int32, endIndex int32, ) *KubernetesLogHighlight`

NewKubernetesLogHighlight instantiates a new KubernetesLogHighlight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesLogHighlightWithDefaults

`func NewKubernetesLogHighlightWithDefaults() *KubernetesLogHighlight`

NewKubernetesLogHighlightWithDefaults instantiates a new KubernetesLogHighlight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartIndex

`func (o *KubernetesLogHighlight) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *KubernetesLogHighlight) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *KubernetesLogHighlight) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.


### GetEndIndex

`func (o *KubernetesLogHighlight) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *KubernetesLogHighlight) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *KubernetesLogHighlight) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


