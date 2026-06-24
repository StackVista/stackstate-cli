# Perspectives

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The resolved URN (ComponentPresentation identifier or legacy QueryView/ViewType URN). | 
**TopologyQuery** | **string** | STQL query for this view. Used by the shared timeline, the events perspective, and the topology perspective. | 
**Overview** | Pointer to **map[string]interface{}** |  | [optional] 
**Topology** | Pointer to [**TopologyPerspective**](TopologyPerspective.md) |  | [optional] 
**Events** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPerspectives

`func NewPerspectives(identifier string, topologyQuery string, ) *Perspectives`

NewPerspectives instantiates a new Perspectives object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerspectivesWithDefaults

`func NewPerspectivesWithDefaults() *Perspectives`

NewPerspectivesWithDefaults instantiates a new Perspectives object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *Perspectives) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Perspectives) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Perspectives) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetTopologyQuery

`func (o *Perspectives) GetTopologyQuery() string`

GetTopologyQuery returns the TopologyQuery field if non-nil, zero value otherwise.

### GetTopologyQueryOk

`func (o *Perspectives) GetTopologyQueryOk() (*string, bool)`

GetTopologyQueryOk returns a tuple with the TopologyQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyQuery

`func (o *Perspectives) SetTopologyQuery(v string)`

SetTopologyQuery sets TopologyQuery field to given value.


### GetOverview

`func (o *Perspectives) GetOverview() map[string]interface{}`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *Perspectives) GetOverviewOk() (*map[string]interface{}, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *Perspectives) SetOverview(v map[string]interface{})`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *Perspectives) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetTopology

`func (o *Perspectives) GetTopology() TopologyPerspective`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *Perspectives) GetTopologyOk() (*TopologyPerspective, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *Perspectives) SetTopology(v TopologyPerspective)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *Perspectives) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetEvents

`func (o *Perspectives) GetEvents() map[string]interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Perspectives) GetEventsOk() (*map[string]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Perspectives) SetEvents(v map[string]interface{})`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Perspectives) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


