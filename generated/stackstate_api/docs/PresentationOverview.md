# PresentationOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**MainMenu** | Pointer to [**PresentationMainMenu**](PresentationMainMenu.md) |  | [optional] 

## Methods

### NewPresentationOverview

`func NewPresentationOverview(title string, ) *PresentationOverview`

NewPresentationOverview instantiates a new PresentationOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationOverviewWithDefaults

`func NewPresentationOverviewWithDefaults() *PresentationOverview`

NewPresentationOverviewWithDefaults instantiates a new PresentationOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PresentationOverview) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PresentationOverview) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PresentationOverview) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMainMenu

`func (o *PresentationOverview) GetMainMenu() PresentationMainMenu`

GetMainMenu returns the MainMenu field if non-nil, zero value otherwise.

### GetMainMenuOk

`func (o *PresentationOverview) GetMainMenuOk() (*PresentationMainMenu, bool)`

GetMainMenuOk returns a tuple with the MainMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainMenu

`func (o *PresentationOverview) SetMainMenu(v PresentationMainMenu)`

SetMainMenu sets MainMenu field to given value.

### HasMainMenu

`func (o *PresentationOverview) HasMainMenu() bool`

HasMainMenu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


