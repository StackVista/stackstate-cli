# PresentationHighlight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Fields** | [**[]PresentationHighlightField**](PresentationHighlightField.md) |  | 
**Provisioning** | Pointer to [**PresentationHighlightProvisioning**](PresentationHighlightProvisioning.md) |  | [optional] 
**RelatedResources** | Pointer to [**[]PresentationRelatedResource**](PresentationRelatedResource.md) |  | [optional] 
**Events** | Pointer to [**PresentationHighlightEvents**](PresentationHighlightEvents.md) |  | [optional] 

## Methods

### NewPresentationHighlight

`func NewPresentationHighlight(title string, fields []PresentationHighlightField, ) *PresentationHighlight`

NewPresentationHighlight instantiates a new PresentationHighlight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationHighlightWithDefaults

`func NewPresentationHighlightWithDefaults() *PresentationHighlight`

NewPresentationHighlightWithDefaults instantiates a new PresentationHighlight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PresentationHighlight) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PresentationHighlight) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PresentationHighlight) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetFields

`func (o *PresentationHighlight) GetFields() []PresentationHighlightField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PresentationHighlight) GetFieldsOk() (*[]PresentationHighlightField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PresentationHighlight) SetFields(v []PresentationHighlightField)`

SetFields sets Fields field to given value.


### GetProvisioning

`func (o *PresentationHighlight) GetProvisioning() PresentationHighlightProvisioning`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *PresentationHighlight) GetProvisioningOk() (*PresentationHighlightProvisioning, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *PresentationHighlight) SetProvisioning(v PresentationHighlightProvisioning)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *PresentationHighlight) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetRelatedResources

`func (o *PresentationHighlight) GetRelatedResources() []PresentationRelatedResource`

GetRelatedResources returns the RelatedResources field if non-nil, zero value otherwise.

### GetRelatedResourcesOk

`func (o *PresentationHighlight) GetRelatedResourcesOk() (*[]PresentationRelatedResource, bool)`

GetRelatedResourcesOk returns a tuple with the RelatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResources

`func (o *PresentationHighlight) SetRelatedResources(v []PresentationRelatedResource)`

SetRelatedResources sets RelatedResources field to given value.

### HasRelatedResources

`func (o *PresentationHighlight) HasRelatedResources() bool`

HasRelatedResources returns a boolean if a field has been set.

### GetEvents

`func (o *PresentationHighlight) GetEvents() PresentationHighlightEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PresentationHighlight) GetEventsOk() (*PresentationHighlightEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PresentationHighlight) SetEvents(v PresentationHighlightEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *PresentationHighlight) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


