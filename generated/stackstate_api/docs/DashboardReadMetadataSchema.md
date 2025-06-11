# DashboardReadMetadataSchema

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

### NewDashboardReadMetadataSchema

`func NewDashboardReadMetadataSchema(type_ string, identifier string, id int64, name string, description string, scope DashboardScope, lastUpdateTimestamp int64, ) *DashboardReadMetadataSchema`

NewDashboardReadMetadataSchema instantiates a new DashboardReadMetadataSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardReadMetadataSchemaWithDefaults

`func NewDashboardReadMetadataSchemaWithDefaults() *DashboardReadMetadataSchema`

NewDashboardReadMetadataSchemaWithDefaults instantiates a new DashboardReadMetadataSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DashboardReadMetadataSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DashboardReadMetadataSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DashboardReadMetadataSchema) SetType(v string)`

SetType sets Type field to given value.


### GetIdentifier

`func (o *DashboardReadMetadataSchema) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DashboardReadMetadataSchema) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DashboardReadMetadataSchema) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetId

`func (o *DashboardReadMetadataSchema) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardReadMetadataSchema) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardReadMetadataSchema) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *DashboardReadMetadataSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardReadMetadataSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardReadMetadataSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DashboardReadMetadataSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardReadMetadataSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardReadMetadataSchema) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetScope

`func (o *DashboardReadMetadataSchema) GetScope() DashboardScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DashboardReadMetadataSchema) GetScopeOk() (*DashboardScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DashboardReadMetadataSchema) SetScope(v DashboardScope)`

SetScope sets Scope field to given value.


### GetOwnerId

`func (o *DashboardReadMetadataSchema) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *DashboardReadMetadataSchema) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *DashboardReadMetadataSchema) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *DashboardReadMetadataSchema) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *DashboardReadMetadataSchema) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *DashboardReadMetadataSchema) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *DashboardReadMetadataSchema) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


