# TopologySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupingEnabled** | Pointer to **bool** |  | [optional] 
**ShowIndirectRelations** | Pointer to **bool** |  | [optional] 
**MinimumGroupSize** | Pointer to **int64** |  | [optional] 
**GroupedByLayers** | Pointer to **bool** |  | [optional] 
**GroupedByDomains** | Pointer to **bool** |  | [optional] 
**GroupedByRelations** | Pointer to **bool** |  | [optional] 
**AutoGrouping** | Pointer to **bool** |  | [optional] 
**ConnectedComponents** | Pointer to **bool** |  | [optional] 
**NeighboringComponents** | Pointer to **bool** |  | [optional] 
**Layer** | Pointer to [**TopologyLayer**](TopologyLayer.md) |  | [optional] 
**Domain** | Pointer to [**TopologyDomain**](TopologyDomain.md) |  | [optional] 

## Methods

### NewTopologySettings

`func NewTopologySettings() *TopologySettings`

NewTopologySettings instantiates a new TopologySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologySettingsWithDefaults

`func NewTopologySettingsWithDefaults() *TopologySettings`

NewTopologySettingsWithDefaults instantiates a new TopologySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupingEnabled

`func (o *TopologySettings) GetGroupingEnabled() bool`

GetGroupingEnabled returns the GroupingEnabled field if non-nil, zero value otherwise.

### GetGroupingEnabledOk

`func (o *TopologySettings) GetGroupingEnabledOk() (*bool, bool)`

GetGroupingEnabledOk returns a tuple with the GroupingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingEnabled

`func (o *TopologySettings) SetGroupingEnabled(v bool)`

SetGroupingEnabled sets GroupingEnabled field to given value.

### HasGroupingEnabled

`func (o *TopologySettings) HasGroupingEnabled() bool`

HasGroupingEnabled returns a boolean if a field has been set.

### GetShowIndirectRelations

`func (o *TopologySettings) GetShowIndirectRelations() bool`

GetShowIndirectRelations returns the ShowIndirectRelations field if non-nil, zero value otherwise.

### GetShowIndirectRelationsOk

`func (o *TopologySettings) GetShowIndirectRelationsOk() (*bool, bool)`

GetShowIndirectRelationsOk returns a tuple with the ShowIndirectRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIndirectRelations

`func (o *TopologySettings) SetShowIndirectRelations(v bool)`

SetShowIndirectRelations sets ShowIndirectRelations field to given value.

### HasShowIndirectRelations

`func (o *TopologySettings) HasShowIndirectRelations() bool`

HasShowIndirectRelations returns a boolean if a field has been set.

### GetMinimumGroupSize

`func (o *TopologySettings) GetMinimumGroupSize() int64`

GetMinimumGroupSize returns the MinimumGroupSize field if non-nil, zero value otherwise.

### GetMinimumGroupSizeOk

`func (o *TopologySettings) GetMinimumGroupSizeOk() (*int64, bool)`

GetMinimumGroupSizeOk returns a tuple with the MinimumGroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumGroupSize

`func (o *TopologySettings) SetMinimumGroupSize(v int64)`

SetMinimumGroupSize sets MinimumGroupSize field to given value.

### HasMinimumGroupSize

`func (o *TopologySettings) HasMinimumGroupSize() bool`

HasMinimumGroupSize returns a boolean if a field has been set.

### GetGroupedByLayers

`func (o *TopologySettings) GetGroupedByLayers() bool`

GetGroupedByLayers returns the GroupedByLayers field if non-nil, zero value otherwise.

### GetGroupedByLayersOk

`func (o *TopologySettings) GetGroupedByLayersOk() (*bool, bool)`

GetGroupedByLayersOk returns a tuple with the GroupedByLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedByLayers

`func (o *TopologySettings) SetGroupedByLayers(v bool)`

SetGroupedByLayers sets GroupedByLayers field to given value.

### HasGroupedByLayers

`func (o *TopologySettings) HasGroupedByLayers() bool`

HasGroupedByLayers returns a boolean if a field has been set.

### GetGroupedByDomains

`func (o *TopologySettings) GetGroupedByDomains() bool`

GetGroupedByDomains returns the GroupedByDomains field if non-nil, zero value otherwise.

### GetGroupedByDomainsOk

`func (o *TopologySettings) GetGroupedByDomainsOk() (*bool, bool)`

GetGroupedByDomainsOk returns a tuple with the GroupedByDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedByDomains

`func (o *TopologySettings) SetGroupedByDomains(v bool)`

SetGroupedByDomains sets GroupedByDomains field to given value.

### HasGroupedByDomains

`func (o *TopologySettings) HasGroupedByDomains() bool`

HasGroupedByDomains returns a boolean if a field has been set.

### GetGroupedByRelations

`func (o *TopologySettings) GetGroupedByRelations() bool`

GetGroupedByRelations returns the GroupedByRelations field if non-nil, zero value otherwise.

### GetGroupedByRelationsOk

`func (o *TopologySettings) GetGroupedByRelationsOk() (*bool, bool)`

GetGroupedByRelationsOk returns a tuple with the GroupedByRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedByRelations

`func (o *TopologySettings) SetGroupedByRelations(v bool)`

SetGroupedByRelations sets GroupedByRelations field to given value.

### HasGroupedByRelations

`func (o *TopologySettings) HasGroupedByRelations() bool`

HasGroupedByRelations returns a boolean if a field has been set.

### GetAutoGrouping

`func (o *TopologySettings) GetAutoGrouping() bool`

GetAutoGrouping returns the AutoGrouping field if non-nil, zero value otherwise.

### GetAutoGroupingOk

`func (o *TopologySettings) GetAutoGroupingOk() (*bool, bool)`

GetAutoGroupingOk returns a tuple with the AutoGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoGrouping

`func (o *TopologySettings) SetAutoGrouping(v bool)`

SetAutoGrouping sets AutoGrouping field to given value.

### HasAutoGrouping

`func (o *TopologySettings) HasAutoGrouping() bool`

HasAutoGrouping returns a boolean if a field has been set.

### GetConnectedComponents

`func (o *TopologySettings) GetConnectedComponents() bool`

GetConnectedComponents returns the ConnectedComponents field if non-nil, zero value otherwise.

### GetConnectedComponentsOk

`func (o *TopologySettings) GetConnectedComponentsOk() (*bool, bool)`

GetConnectedComponentsOk returns a tuple with the ConnectedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedComponents

`func (o *TopologySettings) SetConnectedComponents(v bool)`

SetConnectedComponents sets ConnectedComponents field to given value.

### HasConnectedComponents

`func (o *TopologySettings) HasConnectedComponents() bool`

HasConnectedComponents returns a boolean if a field has been set.

### GetNeighboringComponents

`func (o *TopologySettings) GetNeighboringComponents() bool`

GetNeighboringComponents returns the NeighboringComponents field if non-nil, zero value otherwise.

### GetNeighboringComponentsOk

`func (o *TopologySettings) GetNeighboringComponentsOk() (*bool, bool)`

GetNeighboringComponentsOk returns a tuple with the NeighboringComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighboringComponents

`func (o *TopologySettings) SetNeighboringComponents(v bool)`

SetNeighboringComponents sets NeighboringComponents field to given value.

### HasNeighboringComponents

`func (o *TopologySettings) HasNeighboringComponents() bool`

HasNeighboringComponents returns a boolean if a field has been set.

### GetLayer

`func (o *TopologySettings) GetLayer() TopologyLayer`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *TopologySettings) GetLayerOk() (*TopologyLayer, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *TopologySettings) SetLayer(v TopologyLayer)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *TopologySettings) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### GetDomain

`func (o *TopologySettings) GetDomain() TopologyDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TopologySettings) GetDomainOk() (*TopologyDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TopologySettings) SetDomain(v TopologyDomain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *TopologySettings) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


