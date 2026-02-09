# MainMenuViewItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | Either a viewIdentifier or a componentPresentationIdentifier | 
**Title** | **string** |  | 
**Iconbase64** | Pointer to **string** |  | [optional] 

## Methods

### NewMainMenuViewItem

`func NewMainMenuViewItem(identifier string, title string, ) *MainMenuViewItem`

NewMainMenuViewItem instantiates a new MainMenuViewItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainMenuViewItemWithDefaults

`func NewMainMenuViewItemWithDefaults() *MainMenuViewItem`

NewMainMenuViewItemWithDefaults instantiates a new MainMenuViewItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *MainMenuViewItem) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MainMenuViewItem) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MainMenuViewItem) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetTitle

`func (o *MainMenuViewItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MainMenuViewItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MainMenuViewItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIconbase64

`func (o *MainMenuViewItem) GetIconbase64() string`

GetIconbase64 returns the Iconbase64 field if non-nil, zero value otherwise.

### GetIconbase64Ok

`func (o *MainMenuViewItem) GetIconbase64Ok() (*string, bool)`

GetIconbase64Ok returns a tuple with the Iconbase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconbase64

`func (o *MainMenuViewItem) SetIconbase64(v string)`

SetIconbase64 sets Iconbase64 field to given value.

### HasIconbase64

`func (o *MainMenuViewItem) HasIconbase64() bool`

HasIconbase64 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


