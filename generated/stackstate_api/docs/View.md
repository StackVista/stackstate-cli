# View

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Identifier** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**ViewType** | Pointer to **string** |  | [optional] 
**TopologyQuery** | **string** |  | 
**GroupingEnabled** | **bool** |  | 
**GroupedByLayers** | **bool** |  | 
**GroupedByDomains** | **bool** |  | 
**GroupedByRelations** | **bool** |  | 
**MinimumGroupSize** | **int64** |  | 
**ShowIndirectRelations** | **bool** |  | 
**AutoGrouping** | **bool** |  | 
**ConnectedComponents** | **bool** |  | 
**NeighboringComponents** | **bool** |  | 
**Flags** | [**[]QueryViewFlag**](QueryViewFlag.md) |  | 
**OwnedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewView

`func NewView(id int64, name string, topologyQuery string, groupingEnabled bool, groupedByLayers bool, groupedByDomains bool, groupedByRelations bool, minimumGroupSize int64, showIndirectRelations bool, autoGrouping bool, connectedComponents bool, neighboringComponents bool, flags []QueryViewFlag, ) *View`

NewView instantiates a new View object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewWithDefaults

`func NewViewWithDefaults() *View`

NewViewWithDefaults instantiates a new View object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *View) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *View) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *View) SetId(v int64)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *View) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *View) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *View) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *View) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *View) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *View) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *View) SetName(v string)`

SetName sets Name field to given value.


### GetViewType

`func (o *View) GetViewType() string`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *View) GetViewTypeOk() (*string, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *View) SetViewType(v string)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *View) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### GetTopologyQuery

`func (o *View) GetTopologyQuery() string`

GetTopologyQuery returns the TopologyQuery field if non-nil, zero value otherwise.

### GetTopologyQueryOk

`func (o *View) GetTopologyQueryOk() (*string, bool)`

GetTopologyQueryOk returns a tuple with the TopologyQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyQuery

`func (o *View) SetTopologyQuery(v string)`

SetTopologyQuery sets TopologyQuery field to given value.


### GetGroupingEnabled

`func (o *View) GetGroupingEnabled() bool`

GetGroupingEnabled returns the GroupingEnabled field if non-nil, zero value otherwise.

### GetGroupingEnabledOk

`func (o *View) GetGroupingEnabledOk() (*bool, bool)`

GetGroupingEnabledOk returns a tuple with the GroupingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingEnabled

`func (o *View) SetGroupingEnabled(v bool)`

SetGroupingEnabled sets GroupingEnabled field to given value.


### GetGroupedByLayers

`func (o *View) GetGroupedByLayers() bool`

GetGroupedByLayers returns the GroupedByLayers field if non-nil, zero value otherwise.

### GetGroupedByLayersOk

`func (o *View) GetGroupedByLayersOk() (*bool, bool)`

GetGroupedByLayersOk returns a tuple with the GroupedByLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedByLayers

`func (o *View) SetGroupedByLayers(v bool)`

SetGroupedByLayers sets GroupedByLayers field to given value.


### GetGroupedByDomains

`func (o *View) GetGroupedByDomains() bool`

GetGroupedByDomains returns the GroupedByDomains field if non-nil, zero value otherwise.

### GetGroupedByDomainsOk

`func (o *View) GetGroupedByDomainsOk() (*bool, bool)`

GetGroupedByDomainsOk returns a tuple with the GroupedByDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedByDomains

`func (o *View) SetGroupedByDomains(v bool)`

SetGroupedByDomains sets GroupedByDomains field to given value.


### GetGroupedByRelations

`func (o *View) GetGroupedByRelations() bool`

GetGroupedByRelations returns the GroupedByRelations field if non-nil, zero value otherwise.

### GetGroupedByRelationsOk

`func (o *View) GetGroupedByRelationsOk() (*bool, bool)`

GetGroupedByRelationsOk returns a tuple with the GroupedByRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedByRelations

`func (o *View) SetGroupedByRelations(v bool)`

SetGroupedByRelations sets GroupedByRelations field to given value.


### GetMinimumGroupSize

`func (o *View) GetMinimumGroupSize() int64`

GetMinimumGroupSize returns the MinimumGroupSize field if non-nil, zero value otherwise.

### GetMinimumGroupSizeOk

`func (o *View) GetMinimumGroupSizeOk() (*int64, bool)`

GetMinimumGroupSizeOk returns a tuple with the MinimumGroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumGroupSize

`func (o *View) SetMinimumGroupSize(v int64)`

SetMinimumGroupSize sets MinimumGroupSize field to given value.


### GetShowIndirectRelations

`func (o *View) GetShowIndirectRelations() bool`

GetShowIndirectRelations returns the ShowIndirectRelations field if non-nil, zero value otherwise.

### GetShowIndirectRelationsOk

`func (o *View) GetShowIndirectRelationsOk() (*bool, bool)`

GetShowIndirectRelationsOk returns a tuple with the ShowIndirectRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIndirectRelations

`func (o *View) SetShowIndirectRelations(v bool)`

SetShowIndirectRelations sets ShowIndirectRelations field to given value.


### GetAutoGrouping

`func (o *View) GetAutoGrouping() bool`

GetAutoGrouping returns the AutoGrouping field if non-nil, zero value otherwise.

### GetAutoGroupingOk

`func (o *View) GetAutoGroupingOk() (*bool, bool)`

GetAutoGroupingOk returns a tuple with the AutoGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoGrouping

`func (o *View) SetAutoGrouping(v bool)`

SetAutoGrouping sets AutoGrouping field to given value.


### GetConnectedComponents

`func (o *View) GetConnectedComponents() bool`

GetConnectedComponents returns the ConnectedComponents field if non-nil, zero value otherwise.

### GetConnectedComponentsOk

`func (o *View) GetConnectedComponentsOk() (*bool, bool)`

GetConnectedComponentsOk returns a tuple with the ConnectedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedComponents

`func (o *View) SetConnectedComponents(v bool)`

SetConnectedComponents sets ConnectedComponents field to given value.


### GetNeighboringComponents

`func (o *View) GetNeighboringComponents() bool`

GetNeighboringComponents returns the NeighboringComponents field if non-nil, zero value otherwise.

### GetNeighboringComponentsOk

`func (o *View) GetNeighboringComponentsOk() (*bool, bool)`

GetNeighboringComponentsOk returns a tuple with the NeighboringComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighboringComponents

`func (o *View) SetNeighboringComponents(v bool)`

SetNeighboringComponents sets NeighboringComponents field to given value.


### GetFlags

`func (o *View) GetFlags() []QueryViewFlag`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *View) GetFlagsOk() (*[]QueryViewFlag, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *View) SetFlags(v []QueryViewFlag)`

SetFlags sets Flags field to given value.


### GetOwnedBy

`func (o *View) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *View) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *View) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *View) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


