# RelatedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | Identity key for this related resource. | 
**Title** | **string** | Section heading displayed in the UI. | 
**Stql** | **string** | Resolved STQL query with template variables substituted. | 
**PresentationIdentifier** | **string** | ComponentPresentation identifier whose overview spec is used for rendering. | 

## Methods

### NewRelatedResource

`func NewRelatedResource(resourceId string, title string, stql string, presentationIdentifier string, ) *RelatedResource`

NewRelatedResource instantiates a new RelatedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedResourceWithDefaults

`func NewRelatedResourceWithDefaults() *RelatedResource`

NewRelatedResourceWithDefaults instantiates a new RelatedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *RelatedResource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *RelatedResource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *RelatedResource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetTitle

`func (o *RelatedResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RelatedResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RelatedResource) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStql

`func (o *RelatedResource) GetStql() string`

GetStql returns the Stql field if non-nil, zero value otherwise.

### GetStqlOk

`func (o *RelatedResource) GetStqlOk() (*string, bool)`

GetStqlOk returns a tuple with the Stql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStql

`func (o *RelatedResource) SetStql(v string)`

SetStql sets Stql field to given value.


### GetPresentationIdentifier

`func (o *RelatedResource) GetPresentationIdentifier() string`

GetPresentationIdentifier returns the PresentationIdentifier field if non-nil, zero value otherwise.

### GetPresentationIdentifierOk

`func (o *RelatedResource) GetPresentationIdentifierOk() (*string, bool)`

GetPresentationIdentifierOk returns a tuple with the PresentationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentationIdentifier

`func (o *RelatedResource) SetPresentationIdentifier(v string)`

SetPresentationIdentifier sets PresentationIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


