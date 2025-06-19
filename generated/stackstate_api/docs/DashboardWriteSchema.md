# DashboardWriteSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the dashboard | 
**Description** | **string** | Description of the dashboard | 
**Scope** | [**DashboardScope**](DashboardScope.md) |  | 
**Dashboard** | [**PersesDashboard**](PersesDashboard.md) |  | 

## Methods

### NewDashboardWriteSchema

`func NewDashboardWriteSchema(name string, description string, scope DashboardScope, dashboard PersesDashboard, ) *DashboardWriteSchema`

NewDashboardWriteSchema instantiates a new DashboardWriteSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWriteSchemaWithDefaults

`func NewDashboardWriteSchemaWithDefaults() *DashboardWriteSchema`

NewDashboardWriteSchemaWithDefaults instantiates a new DashboardWriteSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DashboardWriteSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardWriteSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardWriteSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DashboardWriteSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardWriteSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardWriteSchema) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetScope

`func (o *DashboardWriteSchema) GetScope() DashboardScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DashboardWriteSchema) GetScopeOk() (*DashboardScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DashboardWriteSchema) SetScope(v DashboardScope)`

SetScope sets Scope field to given value.


### GetDashboard

`func (o *DashboardWriteSchema) GetDashboard() PersesDashboard`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *DashboardWriteSchema) GetDashboardOk() (*PersesDashboard, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *DashboardWriteSchema) SetDashboard(v PersesDashboard)`

SetDashboard sets Dashboard field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


