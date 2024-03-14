# TopologyPromQLMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**PromQLQuery** | **string** |  | 
**Unit** | **string** |  | 
**AliasTemplate** | **string** |  | 
**TopologyQuery** | **string** |  | 

## Methods

### NewTopologyPromQLMetric

`func NewTopologyPromQLMetric(type_ string, promQLQuery string, unit string, aliasTemplate string, topologyQuery string, ) *TopologyPromQLMetric`

NewTopologyPromQLMetric instantiates a new TopologyPromQLMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyPromQLMetricWithDefaults

`func NewTopologyPromQLMetricWithDefaults() *TopologyPromQLMetric`

NewTopologyPromQLMetricWithDefaults instantiates a new TopologyPromQLMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TopologyPromQLMetric) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopologyPromQLMetric) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopologyPromQLMetric) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *TopologyPromQLMetric) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopologyPromQLMetric) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopologyPromQLMetric) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TopologyPromQLMetric) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPromQLQuery

`func (o *TopologyPromQLMetric) GetPromQLQuery() string`

GetPromQLQuery returns the PromQLQuery field if non-nil, zero value otherwise.

### GetPromQLQueryOk

`func (o *TopologyPromQLMetric) GetPromQLQueryOk() (*string, bool)`

GetPromQLQueryOk returns a tuple with the PromQLQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromQLQuery

`func (o *TopologyPromQLMetric) SetPromQLQuery(v string)`

SetPromQLQuery sets PromQLQuery field to given value.


### GetUnit

`func (o *TopologyPromQLMetric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TopologyPromQLMetric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TopologyPromQLMetric) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetAliasTemplate

`func (o *TopologyPromQLMetric) GetAliasTemplate() string`

GetAliasTemplate returns the AliasTemplate field if non-nil, zero value otherwise.

### GetAliasTemplateOk

`func (o *TopologyPromQLMetric) GetAliasTemplateOk() (*string, bool)`

GetAliasTemplateOk returns a tuple with the AliasTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTemplate

`func (o *TopologyPromQLMetric) SetAliasTemplate(v string)`

SetAliasTemplate sets AliasTemplate field to given value.


### GetTopologyQuery

`func (o *TopologyPromQLMetric) GetTopologyQuery() string`

GetTopologyQuery returns the TopologyQuery field if non-nil, zero value otherwise.

### GetTopologyQueryOk

`func (o *TopologyPromQLMetric) GetTopologyQueryOk() (*string, bool)`

GetTopologyQueryOk returns a tuple with the TopologyQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyQuery

`func (o *TopologyPromQLMetric) SetTopologyQuery(v string)`

SetTopologyQuery sets TopologyQuery field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


