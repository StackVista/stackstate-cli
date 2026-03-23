# ComponentPresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Binding** | [**ComponentPresentationQueryBinding**](ComponentPresentationQueryBinding.md) |  | 
**Rank** | Pointer to [**ComponentPresentationRank**](ComponentPresentationRank.md) |  | [optional] 
**Presentation** | [**PresentationDefinition**](PresentationDefinition.md) |  | 

## Methods

### NewComponentPresentation

`func NewComponentPresentation(identifier string, name string, binding ComponentPresentationQueryBinding, presentation PresentationDefinition, ) *ComponentPresentation`

NewComponentPresentation instantiates a new ComponentPresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentPresentationWithDefaults

`func NewComponentPresentationWithDefaults() *ComponentPresentation`

NewComponentPresentationWithDefaults instantiates a new ComponentPresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *ComponentPresentation) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ComponentPresentation) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ComponentPresentation) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *ComponentPresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentPresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentPresentation) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ComponentPresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComponentPresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComponentPresentation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComponentPresentation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBinding

`func (o *ComponentPresentation) GetBinding() ComponentPresentationQueryBinding`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *ComponentPresentation) GetBindingOk() (*ComponentPresentationQueryBinding, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *ComponentPresentation) SetBinding(v ComponentPresentationQueryBinding)`

SetBinding sets Binding field to given value.


### GetRank

`func (o *ComponentPresentation) GetRank() ComponentPresentationRank`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ComponentPresentation) GetRankOk() (*ComponentPresentationRank, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ComponentPresentation) SetRank(v ComponentPresentationRank)`

SetRank sets Rank field to given value.

### HasRank

`func (o *ComponentPresentation) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetPresentation

`func (o *ComponentPresentation) GetPresentation() PresentationDefinition`

GetPresentation returns the Presentation field if non-nil, zero value otherwise.

### GetPresentationOk

`func (o *ComponentPresentation) GetPresentationOk() (*PresentationDefinition, bool)`

GetPresentationOk returns a tuple with the Presentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentation

`func (o *ComponentPresentation) SetPresentation(v PresentationDefinition)`

SetPresentation sets Presentation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


