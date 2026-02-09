# ComponentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Identifiers** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Tags** | **[]string** |  | 
**Properties** | **map[string]string** |  | 
**LayerName** | Pointer to **string** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**Health** | Pointer to [**HealthStateValue**](HealthStateValue.md) |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Synced** | [**[]ExternalComponent**](ExternalComponent.md) |  | 

## Methods

### NewComponentData

`func NewComponentData(name string, tags []string, properties map[string]string, synced []ExternalComponent, ) *ComponentData`

NewComponentData instantiates a new ComponentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentDataWithDefaults

`func NewComponentDataWithDefaults() *ComponentData`

NewComponentDataWithDefaults instantiates a new ComponentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComponentData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentData) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifiers

`func (o *ComponentData) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *ComponentData) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *ComponentData) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *ComponentData) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetName

`func (o *ComponentData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentData) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ComponentData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComponentData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComponentData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComponentData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *ComponentData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ComponentData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ComponentData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ComponentData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTags

`func (o *ComponentData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ComponentData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ComponentData) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetProperties

`func (o *ComponentData) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ComponentData) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ComponentData) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.


### GetLayerName

`func (o *ComponentData) GetLayerName() string`

GetLayerName returns the LayerName field if non-nil, zero value otherwise.

### GetLayerNameOk

`func (o *ComponentData) GetLayerNameOk() (*string, bool)`

GetLayerNameOk returns a tuple with the LayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerName

`func (o *ComponentData) SetLayerName(v string)`

SetLayerName sets LayerName field to given value.

### HasLayerName

`func (o *ComponentData) HasLayerName() bool`

HasLayerName returns a boolean if a field has been set.

### GetDomainName

`func (o *ComponentData) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ComponentData) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ComponentData) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ComponentData) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetHealth

`func (o *ComponentData) GetHealth() HealthStateValue`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ComponentData) GetHealthOk() (*HealthStateValue, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ComponentData) SetHealth(v HealthStateValue)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ComponentData) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ComponentData) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ComponentData) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ComponentData) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *ComponentData) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetSynced

`func (o *ComponentData) GetSynced() []ExternalComponent`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *ComponentData) GetSyncedOk() (*[]ExternalComponent, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *ComponentData) SetSynced(v []ExternalComponent)`

SetSynced sets Synced field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


