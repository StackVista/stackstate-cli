# DashboardReadBaseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Identifier** | **string** | The identifier of a dashboard. Either the system/graph ID or URN of the resource. | 
**Id** | **int64** |  | 
**Name** | **string** | Name of the dashboard | 
**Description** | **string** | Description of the dashboard | 
**Scope** | [**DashboardScope**](DashboardScope.md) |  | 
**OwnerId** | Pointer to **int64** | The user id of the dashboard owner. A dashboard was either created by a user or from a StackPack. For a user, the identifier will be the system/graph ID, and for a StackPack, the field will be empty/omitted. | [optional] 
**LastUpdateTimestamp** | **int64** |  | 

## Methods

### NewDashboardReadBaseSchema

`func NewDashboardReadBaseSchema(type_ string, identifier string, id int64, name string, description string, scope DashboardScope, lastUpdateTimestamp int64, ) *DashboardReadBaseSchema`

NewDashboardReadBaseSchema instantiates a new DashboardReadBaseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardReadBaseSchemaWithDefaults

`func NewDashboardReadBaseSchemaWithDefaults() *DashboardReadBaseSchema`

NewDashboardReadBaseSchemaWithDefaults instantiates a new DashboardReadBaseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DashboardReadBaseSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DashboardReadBaseSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DashboardReadBaseSchema) SetType(v string)`

SetType sets Type field to given value.


### GetIdentifier

`func (o *DashboardReadBaseSchema) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DashboardReadBaseSchema) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DashboardReadBaseSchema) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetId

`func (o *DashboardReadBaseSchema) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardReadBaseSchema) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardReadBaseSchema) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *DashboardReadBaseSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardReadBaseSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardReadBaseSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DashboardReadBaseSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardReadBaseSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardReadBaseSchema) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetScope

`func (o *DashboardReadBaseSchema) GetScope() DashboardScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DashboardReadBaseSchema) GetScopeOk() (*DashboardScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DashboardReadBaseSchema) SetScope(v DashboardScope)`

SetScope sets Scope field to given value.


### GetOwnerId

`func (o *DashboardReadBaseSchema) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *DashboardReadBaseSchema) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *DashboardReadBaseSchema) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *DashboardReadBaseSchema) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *DashboardReadBaseSchema) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *DashboardReadBaseSchema) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *DashboardReadBaseSchema) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


