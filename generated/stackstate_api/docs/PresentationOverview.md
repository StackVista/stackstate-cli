# PresentationOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**PresentationName**](PresentationName.md) |  | 
**MainMenu** | Pointer to [**PresentationMainMenu**](PresentationMainMenu.md) |  | [optional] 
**Columns** | [**[]OverviewColumnDefinition**](OverviewColumnDefinition.md) |  | 
**FixedColumns** | Pointer to **int32** |  | [optional] 

## Methods

### NewPresentationOverview

`func NewPresentationOverview(name PresentationName, columns []OverviewColumnDefinition, ) *PresentationOverview`

NewPresentationOverview instantiates a new PresentationOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationOverviewWithDefaults

`func NewPresentationOverviewWithDefaults() *PresentationOverview`

NewPresentationOverviewWithDefaults instantiates a new PresentationOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PresentationOverview) GetName() PresentationName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PresentationOverview) GetNameOk() (*PresentationName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PresentationOverview) SetName(v PresentationName)`

SetName sets Name field to given value.


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

### GetColumns

`func (o *PresentationOverview) GetColumns() []OverviewColumnDefinition`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *PresentationOverview) GetColumnsOk() (*[]OverviewColumnDefinition, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *PresentationOverview) SetColumns(v []OverviewColumnDefinition)`

SetColumns sets Columns field to given value.


### GetFixedColumns

`func (o *PresentationOverview) GetFixedColumns() int32`

GetFixedColumns returns the FixedColumns field if non-nil, zero value otherwise.

### GetFixedColumnsOk

`func (o *PresentationOverview) GetFixedColumnsOk() (*int32, bool)`

GetFixedColumnsOk returns a tuple with the FixedColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedColumns

`func (o *PresentationOverview) SetFixedColumns(v int32)`

SetFixedColumns sets FixedColumns field to given value.

### HasFixedColumns

`func (o *PresentationOverview) HasFixedColumns() bool`

HasFixedColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


