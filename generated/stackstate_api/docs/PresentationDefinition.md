# PresentationDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icon** | Pointer to **string** |  | [optional] 
**Name** | Pointer to [**PresentationName**](PresentationName.md) |  | [optional] 
**Overview** | Pointer to [**PresentationOverview**](PresentationOverview.md) |  | [optional] 

## Methods

### NewPresentationDefinition

`func NewPresentationDefinition() *PresentationDefinition`

NewPresentationDefinition instantiates a new PresentationDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationDefinitionWithDefaults

`func NewPresentationDefinitionWithDefaults() *PresentationDefinition`

NewPresentationDefinitionWithDefaults instantiates a new PresentationDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcon

`func (o *PresentationDefinition) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PresentationDefinition) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PresentationDefinition) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PresentationDefinition) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetName

`func (o *PresentationDefinition) GetName() PresentationName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PresentationDefinition) GetNameOk() (*PresentationName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PresentationDefinition) SetName(v PresentationName)`

SetName sets Name field to given value.

### HasName

`func (o *PresentationDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverview

`func (o *PresentationDefinition) GetOverview() PresentationOverview`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *PresentationDefinition) GetOverviewOk() (*PresentationOverview, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *PresentationDefinition) SetOverview(v PresentationOverview)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *PresentationDefinition) HasOverview() bool`

HasOverview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


