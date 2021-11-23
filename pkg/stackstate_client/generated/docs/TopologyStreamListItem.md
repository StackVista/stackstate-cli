# TopologyStreamListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncIdentifier** | Pointer to **NullableString** |  | [optional] 
**NodeId** | **int64** |  | 
**CreatedRelations** | **int64** |  | 
**DeletedRelations** | **int64** |  | 
**CreatedComponents** | **int64** |  | 
**DeletedComponents** | **int64** |  | 
**Errors** | **int64** |  | 
**Status** | [**TopologySyncStatus**](TopologySyncStatus.md) |  | 

## Methods

### NewTopologyStreamListItem

`func NewTopologyStreamListItem(nodeId int64, createdRelations int64, deletedRelations int64, createdComponents int64, deletedComponents int64, errors int64, status TopologySyncStatus, ) *TopologyStreamListItem`

NewTopologyStreamListItem instantiates a new TopologyStreamListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyStreamListItemWithDefaults

`func NewTopologyStreamListItemWithDefaults() *TopologyStreamListItem`

NewTopologyStreamListItemWithDefaults instantiates a new TopologyStreamListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncIdentifier

`func (o *TopologyStreamListItem) GetSyncIdentifier() string`

GetSyncIdentifier returns the SyncIdentifier field if non-nil, zero value otherwise.

### GetSyncIdentifierOk

`func (o *TopologyStreamListItem) GetSyncIdentifierOk() (*string, bool)`

GetSyncIdentifierOk returns a tuple with the SyncIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncIdentifier

`func (o *TopologyStreamListItem) SetSyncIdentifier(v string)`

SetSyncIdentifier sets SyncIdentifier field to given value.

### HasSyncIdentifier

`func (o *TopologyStreamListItem) HasSyncIdentifier() bool`

HasSyncIdentifier returns a boolean if a field has been set.

### SetSyncIdentifierNil

`func (o *TopologyStreamListItem) SetSyncIdentifierNil(b bool)`

 SetSyncIdentifierNil sets the value for SyncIdentifier to be an explicit nil

### UnsetSyncIdentifier
`func (o *TopologyStreamListItem) UnsetSyncIdentifier()`

UnsetSyncIdentifier ensures that no value is present for SyncIdentifier, not even an explicit nil
### GetNodeId

`func (o *TopologyStreamListItem) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TopologyStreamListItem) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TopologyStreamListItem) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.


### GetCreatedRelations

`func (o *TopologyStreamListItem) GetCreatedRelations() int64`

GetCreatedRelations returns the CreatedRelations field if non-nil, zero value otherwise.

### GetCreatedRelationsOk

`func (o *TopologyStreamListItem) GetCreatedRelationsOk() (*int64, bool)`

GetCreatedRelationsOk returns a tuple with the CreatedRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedRelations

`func (o *TopologyStreamListItem) SetCreatedRelations(v int64)`

SetCreatedRelations sets CreatedRelations field to given value.


### GetDeletedRelations

`func (o *TopologyStreamListItem) GetDeletedRelations() int64`

GetDeletedRelations returns the DeletedRelations field if non-nil, zero value otherwise.

### GetDeletedRelationsOk

`func (o *TopologyStreamListItem) GetDeletedRelationsOk() (*int64, bool)`

GetDeletedRelationsOk returns a tuple with the DeletedRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedRelations

`func (o *TopologyStreamListItem) SetDeletedRelations(v int64)`

SetDeletedRelations sets DeletedRelations field to given value.


### GetCreatedComponents

`func (o *TopologyStreamListItem) GetCreatedComponents() int64`

GetCreatedComponents returns the CreatedComponents field if non-nil, zero value otherwise.

### GetCreatedComponentsOk

`func (o *TopologyStreamListItem) GetCreatedComponentsOk() (*int64, bool)`

GetCreatedComponentsOk returns a tuple with the CreatedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedComponents

`func (o *TopologyStreamListItem) SetCreatedComponents(v int64)`

SetCreatedComponents sets CreatedComponents field to given value.


### GetDeletedComponents

`func (o *TopologyStreamListItem) GetDeletedComponents() int64`

GetDeletedComponents returns the DeletedComponents field if non-nil, zero value otherwise.

### GetDeletedComponentsOk

`func (o *TopologyStreamListItem) GetDeletedComponentsOk() (*int64, bool)`

GetDeletedComponentsOk returns a tuple with the DeletedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedComponents

`func (o *TopologyStreamListItem) SetDeletedComponents(v int64)`

SetDeletedComponents sets DeletedComponents field to given value.


### GetErrors

`func (o *TopologyStreamListItem) GetErrors() int64`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TopologyStreamListItem) GetErrorsOk() (*int64, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TopologyStreamListItem) SetErrors(v int64)`

SetErrors sets Errors field to given value.


### GetStatus

`func (o *TopologyStreamListItem) GetStatus() TopologySyncStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TopologyStreamListItem) GetStatusOk() (*TopologySyncStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TopologyStreamListItem) SetStatus(v TopologySyncStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


