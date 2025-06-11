# DashboardPatchSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the dashboard | [optional] 
**Description** | Pointer to **string** | Description of the dashboard | [optional] 
**Scope** | Pointer to [**DashboardScope**](DashboardScope.md) |  | [optional] 
**Dashboard** | Pointer to [**PersesDashboard**](PersesDashboard.md) |  | [optional] 

## Methods

### NewDashboardPatchSchema

`func NewDashboardPatchSchema() *DashboardPatchSchema`

NewDashboardPatchSchema instantiates a new DashboardPatchSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardPatchSchemaWithDefaults

`func NewDashboardPatchSchemaWithDefaults() *DashboardPatchSchema`

NewDashboardPatchSchemaWithDefaults instantiates a new DashboardPatchSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DashboardPatchSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardPatchSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardPatchSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DashboardPatchSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DashboardPatchSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardPatchSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardPatchSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DashboardPatchSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScope

`func (o *DashboardPatchSchema) GetScope() DashboardScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DashboardPatchSchema) GetScopeOk() (*DashboardScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DashboardPatchSchema) SetScope(v DashboardScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DashboardPatchSchema) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDashboard

`func (o *DashboardPatchSchema) GetDashboard() PersesDashboard`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *DashboardPatchSchema) GetDashboardOk() (*PersesDashboard, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *DashboardPatchSchema) SetDashboard(v PersesDashboard)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *DashboardPatchSchema) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


