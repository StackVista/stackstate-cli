# OverviewColumnDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnId** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Projection** | Pointer to [**OverviewColumnProjection**](OverviewColumnProjection.md) |  | [optional] 

## Methods

### NewOverviewColumnDefinition

`func NewOverviewColumnDefinition(columnId string, ) *OverviewColumnDefinition`

NewOverviewColumnDefinition instantiates a new OverviewColumnDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewColumnDefinitionWithDefaults

`func NewOverviewColumnDefinitionWithDefaults() *OverviewColumnDefinition`

NewOverviewColumnDefinitionWithDefaults instantiates a new OverviewColumnDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnId

`func (o *OverviewColumnDefinition) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *OverviewColumnDefinition) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *OverviewColumnDefinition) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.


### GetTitle

`func (o *OverviewColumnDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OverviewColumnDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OverviewColumnDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OverviewColumnDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetProjection

`func (o *OverviewColumnDefinition) GetProjection() OverviewColumnProjection`

GetProjection returns the Projection field if non-nil, zero value otherwise.

### GetProjectionOk

`func (o *OverviewColumnDefinition) GetProjectionOk() (*OverviewColumnProjection, bool)`

GetProjectionOk returns a tuple with the Projection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjection

`func (o *OverviewColumnDefinition) SetProjection(v OverviewColumnProjection)`

SetProjection sets Projection field to given value.

### HasProjection

`func (o *OverviewColumnDefinition) HasProjection() bool`

HasProjection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


