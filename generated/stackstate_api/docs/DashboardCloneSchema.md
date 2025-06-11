# DashboardCloneSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the dashboard. Required to make it easier so see which dashboard is the new (cloned) dashboard. | 
**Description** | Pointer to **string** | Description of the dashboard | [optional] 
**Scope** | Pointer to [**DashboardScope**](DashboardScope.md) |  | [optional] 
**Dashboard** | Pointer to [**PersesDashboard**](PersesDashboard.md) |  | [optional] 

## Methods

### NewDashboardCloneSchema

`func NewDashboardCloneSchema(name string, ) *DashboardCloneSchema`

NewDashboardCloneSchema instantiates a new DashboardCloneSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardCloneSchemaWithDefaults

`func NewDashboardCloneSchemaWithDefaults() *DashboardCloneSchema`

NewDashboardCloneSchemaWithDefaults instantiates a new DashboardCloneSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DashboardCloneSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardCloneSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardCloneSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DashboardCloneSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardCloneSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardCloneSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DashboardCloneSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScope

`func (o *DashboardCloneSchema) GetScope() DashboardScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DashboardCloneSchema) GetScopeOk() (*DashboardScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DashboardCloneSchema) SetScope(v DashboardScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DashboardCloneSchema) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDashboard

`func (o *DashboardCloneSchema) GetDashboard() PersesDashboard`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *DashboardCloneSchema) GetDashboardOk() (*PersesDashboard, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *DashboardCloneSchema) SetDashboard(v PersesDashboard)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *DashboardCloneSchema) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


