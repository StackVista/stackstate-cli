# QueryMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupingEnabled** | **bool** |  | 
**ShowIndirectRelations** | **bool** |  | 
**MinGroupSize** | **int64** |  | 
**GroupedByLayer** | **bool** |  | 
**GroupedByDomain** | **bool** |  | 
**GroupedByRelation** | **bool** |  | 
**QueryTime** | Pointer to **NullableInt64** | Unix timestamp in milliseconds | [optional] 
**AutoGrouping** | **bool** |  | 
**ConnectedComponents** | **bool** |  | 
**NeighboringComponents** | **bool** |  | 
**ShowFullComponent** | **bool** |  | 

## Methods

### NewQueryMetadata

`func NewQueryMetadata(groupingEnabled bool, showIndirectRelations bool, minGroupSize int64, groupedByLayer bool, groupedByDomain bool, groupedByRelation bool, autoGrouping bool, connectedComponents bool, neighboringComponents bool, showFullComponent bool, ) *QueryMetadata`

NewQueryMetadata instantiates a new QueryMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMetadataWithDefaults

`func NewQueryMetadataWithDefaults() *QueryMetadata`

NewQueryMetadataWithDefaults instantiates a new QueryMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupingEnabled

`func (o *QueryMetadata) GetGroupingEnabled() bool`

GetGroupingEnabled returns the GroupingEnabled field if non-nil, zero value otherwise.

### GetGroupingEnabledOk

`func (o *QueryMetadata) GetGroupingEnabledOk() (*bool, bool)`

GetGroupingEnabledOk returns a tuple with the GroupingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingEnabled

`func (o *QueryMetadata) SetGroupingEnabled(v bool)`

SetGroupingEnabled sets GroupingEnabled field to given value.


### GetShowIndirectRelations

`func (o *QueryMetadata) GetShowIndirectRelations() bool`

GetShowIndirectRelations returns the ShowIndirectRelations field if non-nil, zero value otherwise.

### GetShowIndirectRelationsOk

`func (o *QueryMetadata) GetShowIndirectRelationsOk() (*bool, bool)`

GetShowIndirectRelationsOk returns a tuple with the ShowIndirectRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIndirectRelations

`func (o *QueryMetadata) SetShowIndirectRelations(v bool)`

SetShowIndirectRelations sets ShowIndirectRelations field to given value.


### GetMinGroupSize

`func (o *QueryMetadata) GetMinGroupSize() int64`

GetMinGroupSize returns the MinGroupSize field if non-nil, zero value otherwise.

### GetMinGroupSizeOk

`func (o *QueryMetadata) GetMinGroupSizeOk() (*int64, bool)`

GetMinGroupSizeOk returns a tuple with the MinGroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGroupSize

`func (o *QueryMetadata) SetMinGroupSize(v int64)`

SetMinGroupSize sets MinGroupSize field to given value.


### GetGroupedByLayer

`func (o *QueryMetadata) GetGroupedByLayer() bool`

GetGroupedByLayer returns the GroupedByLayer field if non-nil, zero value otherwise.

### GetGroupedByLayerOk

`func (o *QueryMetadata) GetGroupedByLayerOk() (*bool, bool)`

GetGroupedByLayerOk returns a tuple with the GroupedByLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedByLayer

`func (o *QueryMetadata) SetGroupedByLayer(v bool)`

SetGroupedByLayer sets GroupedByLayer field to given value.


### GetGroupedByDomain

`func (o *QueryMetadata) GetGroupedByDomain() bool`

GetGroupedByDomain returns the GroupedByDomain field if non-nil, zero value otherwise.

### GetGroupedByDomainOk

`func (o *QueryMetadata) GetGroupedByDomainOk() (*bool, bool)`

GetGroupedByDomainOk returns a tuple with the GroupedByDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedByDomain

`func (o *QueryMetadata) SetGroupedByDomain(v bool)`

SetGroupedByDomain sets GroupedByDomain field to given value.


### GetGroupedByRelation

`func (o *QueryMetadata) GetGroupedByRelation() bool`

GetGroupedByRelation returns the GroupedByRelation field if non-nil, zero value otherwise.

### GetGroupedByRelationOk

`func (o *QueryMetadata) GetGroupedByRelationOk() (*bool, bool)`

GetGroupedByRelationOk returns a tuple with the GroupedByRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedByRelation

`func (o *QueryMetadata) SetGroupedByRelation(v bool)`

SetGroupedByRelation sets GroupedByRelation field to given value.


### GetQueryTime

`func (o *QueryMetadata) GetQueryTime() int64`

GetQueryTime returns the QueryTime field if non-nil, zero value otherwise.

### GetQueryTimeOk

`func (o *QueryMetadata) GetQueryTimeOk() (*int64, bool)`

GetQueryTimeOk returns a tuple with the QueryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTime

`func (o *QueryMetadata) SetQueryTime(v int64)`

SetQueryTime sets QueryTime field to given value.

### HasQueryTime

`func (o *QueryMetadata) HasQueryTime() bool`

HasQueryTime returns a boolean if a field has been set.

### SetQueryTimeNil

`func (o *QueryMetadata) SetQueryTimeNil(b bool)`

 SetQueryTimeNil sets the value for QueryTime to be an explicit nil

### UnsetQueryTime
`func (o *QueryMetadata) UnsetQueryTime()`

UnsetQueryTime ensures that no value is present for QueryTime, not even an explicit nil
### GetAutoGrouping

`func (o *QueryMetadata) GetAutoGrouping() bool`

GetAutoGrouping returns the AutoGrouping field if non-nil, zero value otherwise.

### GetAutoGroupingOk

`func (o *QueryMetadata) GetAutoGroupingOk() (*bool, bool)`

GetAutoGroupingOk returns a tuple with the AutoGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoGrouping

`func (o *QueryMetadata) SetAutoGrouping(v bool)`

SetAutoGrouping sets AutoGrouping field to given value.


### GetConnectedComponents

`func (o *QueryMetadata) GetConnectedComponents() bool`

GetConnectedComponents returns the ConnectedComponents field if non-nil, zero value otherwise.

### GetConnectedComponentsOk

`func (o *QueryMetadata) GetConnectedComponentsOk() (*bool, bool)`

GetConnectedComponentsOk returns a tuple with the ConnectedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedComponents

`func (o *QueryMetadata) SetConnectedComponents(v bool)`

SetConnectedComponents sets ConnectedComponents field to given value.


### GetNeighboringComponents

`func (o *QueryMetadata) GetNeighboringComponents() bool`

GetNeighboringComponents returns the NeighboringComponents field if non-nil, zero value otherwise.

### GetNeighboringComponentsOk

`func (o *QueryMetadata) GetNeighboringComponentsOk() (*bool, bool)`

GetNeighboringComponentsOk returns a tuple with the NeighboringComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighboringComponents

`func (o *QueryMetadata) SetNeighboringComponents(v bool)`

SetNeighboringComponents sets NeighboringComponents field to given value.


### GetShowFullComponent

`func (o *QueryMetadata) GetShowFullComponent() bool`

GetShowFullComponent returns the ShowFullComponent field if non-nil, zero value otherwise.

### GetShowFullComponentOk

`func (o *QueryMetadata) GetShowFullComponentOk() (*bool, bool)`

GetShowFullComponentOk returns a tuple with the ShowFullComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFullComponent

`func (o *QueryMetadata) SetShowFullComponent(v bool)`

SetShowFullComponent sets ShowFullComponent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


