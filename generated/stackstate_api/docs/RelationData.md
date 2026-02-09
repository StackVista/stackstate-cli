# RelationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | **[]string** |  | 
**Health** | Pointer to [**HealthStateValue**](HealthStateValue.md) |  | [optional] 
**Synced** | [**[]ExternalRelation**](ExternalRelation.md) |  | 
**Identifiers** | **[]string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DependencyDirection** | [**DependencyDirection**](DependencyDirection.md) |  | 

## Methods

### NewRelationData

`func NewRelationData(tags []string, synced []ExternalRelation, identifiers []string, dependencyDirection DependencyDirection, ) *RelationData`

NewRelationData instantiates a new RelationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationDataWithDefaults

`func NewRelationDataWithDefaults() *RelationData`

NewRelationDataWithDefaults instantiates a new RelationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *RelationData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RelationData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RelationData) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetHealth

`func (o *RelationData) GetHealth() HealthStateValue`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *RelationData) GetHealthOk() (*HealthStateValue, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *RelationData) SetHealth(v HealthStateValue)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *RelationData) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetSynced

`func (o *RelationData) GetSynced() []ExternalRelation`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *RelationData) GetSyncedOk() (*[]ExternalRelation, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *RelationData) SetSynced(v []ExternalRelation)`

SetSynced sets Synced field to given value.


### GetIdentifiers

`func (o *RelationData) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *RelationData) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *RelationData) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.


### GetDescription

`func (o *RelationData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RelationData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RelationData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RelationData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *RelationData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationData) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RelationData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *RelationData) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *RelationData) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *RelationData) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *RelationData) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetName

`func (o *RelationData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelationData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelationData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RelationData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDependencyDirection

`func (o *RelationData) GetDependencyDirection() DependencyDirection`

GetDependencyDirection returns the DependencyDirection field if non-nil, zero value otherwise.

### GetDependencyDirectionOk

`func (o *RelationData) GetDependencyDirectionOk() (*DependencyDirection, bool)`

GetDependencyDirectionOk returns a tuple with the DependencyDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyDirection

`func (o *RelationData) SetDependencyDirection(v DependencyDirection)`

SetDependencyDirection sets DependencyDirection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


