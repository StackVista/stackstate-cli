# ComponentPresentationRank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Specificity** | **float64** | Determines how much of a \&quot;specialization\&quot; this presentation is. Higher number means more specific. This is used to determine which presentation to use when multiple presentations match. | 

## Methods

### NewComponentPresentationRank

`func NewComponentPresentationRank(specificity float64, ) *ComponentPresentationRank`

NewComponentPresentationRank instantiates a new ComponentPresentationRank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentPresentationRankWithDefaults

`func NewComponentPresentationRankWithDefaults() *ComponentPresentationRank`

NewComponentPresentationRankWithDefaults instantiates a new ComponentPresentationRank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpecificity

`func (o *ComponentPresentationRank) GetSpecificity() float64`

GetSpecificity returns the Specificity field if non-nil, zero value otherwise.

### GetSpecificityOk

`func (o *ComponentPresentationRank) GetSpecificityOk() (*float64, bool)`

GetSpecificityOk returns a tuple with the Specificity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificity

`func (o *ComponentPresentationRank) SetSpecificity(v float64)`

SetSpecificity sets Specificity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


