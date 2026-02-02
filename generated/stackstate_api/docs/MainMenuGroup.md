# MainMenuGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Identifier** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DefaultOpen** | **bool** |  | 
**Iconbase64** | **string** |  | 
**Items** | [**[]MainMenuViewItem**](MainMenuViewItem.md) |  | 

## Methods

### NewMainMenuGroup

`func NewMainMenuGroup(name string, defaultOpen bool, iconbase64 string, items []MainMenuViewItem, ) *MainMenuGroup`

NewMainMenuGroup instantiates a new MainMenuGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainMenuGroupWithDefaults

`func NewMainMenuGroupWithDefaults() *MainMenuGroup`

NewMainMenuGroupWithDefaults instantiates a new MainMenuGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MainMenuGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MainMenuGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MainMenuGroup) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *MainMenuGroup) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MainMenuGroup) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MainMenuGroup) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *MainMenuGroup) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *MainMenuGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MainMenuGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MainMenuGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MainMenuGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultOpen

`func (o *MainMenuGroup) GetDefaultOpen() bool`

GetDefaultOpen returns the DefaultOpen field if non-nil, zero value otherwise.

### GetDefaultOpenOk

`func (o *MainMenuGroup) GetDefaultOpenOk() (*bool, bool)`

GetDefaultOpenOk returns a tuple with the DefaultOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOpen

`func (o *MainMenuGroup) SetDefaultOpen(v bool)`

SetDefaultOpen sets DefaultOpen field to given value.


### GetIconbase64

`func (o *MainMenuGroup) GetIconbase64() string`

GetIconbase64 returns the Iconbase64 field if non-nil, zero value otherwise.

### GetIconbase64Ok

`func (o *MainMenuGroup) GetIconbase64Ok() (*string, bool)`

GetIconbase64Ok returns a tuple with the Iconbase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconbase64

`func (o *MainMenuGroup) SetIconbase64(v string)`

SetIconbase64 sets Iconbase64 field to given value.


### GetItems

`func (o *MainMenuGroup) GetItems() []MainMenuViewItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MainMenuGroup) GetItemsOk() (*[]MainMenuViewItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MainMenuGroup) SetItems(v []MainMenuViewItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


