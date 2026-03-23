# OverviewRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentIdentifier** | **string** |  | 
**Cells** | [**map[string]CellValue**](CellValue.md) |  | 

## Methods

### NewOverviewRow

`func NewOverviewRow(componentIdentifier string, cells map[string]CellValue, ) *OverviewRow`

NewOverviewRow instantiates a new OverviewRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewRowWithDefaults

`func NewOverviewRowWithDefaults() *OverviewRow`

NewOverviewRowWithDefaults instantiates a new OverviewRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentIdentifier

`func (o *OverviewRow) GetComponentIdentifier() string`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *OverviewRow) GetComponentIdentifierOk() (*string, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *OverviewRow) SetComponentIdentifier(v string)`

SetComponentIdentifier sets ComponentIdentifier field to given value.


### GetCells

`func (o *OverviewRow) GetCells() map[string]CellValue`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *OverviewRow) GetCellsOk() (*map[string]CellValue, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *OverviewRow) SetCells(v map[string]CellValue)`

SetCells sets Cells field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


