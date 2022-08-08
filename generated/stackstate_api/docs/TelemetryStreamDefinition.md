# TelemetryStreamDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**BindQuery** | **string** |  | 
**DataSourceId** | **int64** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Labels** | **[]string** |  | 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**MetricValuation** | [**MetricValuation**](MetricValuation.md) |  | 
**Name** | **string** |  | 
**OwnedBy** | Pointer to **string** |  | [optional] 
**Priority** | [**TelemetryStreamPriority**](TelemetryStreamPriority.md) |  | 
**TelemetryQuery** | [**EventTelemetryQuery**](EventTelemetryQuery.md) |  | 
**TopologyMapping** | [**[]TopologyMapping**](TopologyMapping.md) |  | 

## Methods

### NewTelemetryStreamDefinition

`func NewTelemetryStreamDefinition(type_ string, bindQuery string, dataSourceId int64, labels []string, metricValuation MetricValuation, name string, priority TelemetryStreamPriority, telemetryQuery EventTelemetryQuery, topologyMapping []TopologyMapping, ) *TelemetryStreamDefinition`

NewTelemetryStreamDefinition instantiates a new TelemetryStreamDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryStreamDefinitionWithDefaults

`func NewTelemetryStreamDefinitionWithDefaults() *TelemetryStreamDefinition`

NewTelemetryStreamDefinitionWithDefaults instantiates a new TelemetryStreamDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryStreamDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryStreamDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryStreamDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetBindQuery

`func (o *TelemetryStreamDefinition) GetBindQuery() string`

GetBindQuery returns the BindQuery field if non-nil, zero value otherwise.

### GetBindQueryOk

`func (o *TelemetryStreamDefinition) GetBindQueryOk() (*string, bool)`

GetBindQueryOk returns a tuple with the BindQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindQuery

`func (o *TelemetryStreamDefinition) SetBindQuery(v string)`

SetBindQuery sets BindQuery field to given value.


### GetDataSourceId

`func (o *TelemetryStreamDefinition) GetDataSourceId() int64`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *TelemetryStreamDefinition) GetDataSourceIdOk() (*int64, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *TelemetryStreamDefinition) SetDataSourceId(v int64)`

SetDataSourceId sets DataSourceId field to given value.


### GetDescription

`func (o *TelemetryStreamDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TelemetryStreamDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TelemetryStreamDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TelemetryStreamDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *TelemetryStreamDefinition) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TelemetryStreamDefinition) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TelemetryStreamDefinition) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TelemetryStreamDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *TelemetryStreamDefinition) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *TelemetryStreamDefinition) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *TelemetryStreamDefinition) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *TelemetryStreamDefinition) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetLabels

`func (o *TelemetryStreamDefinition) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *TelemetryStreamDefinition) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *TelemetryStreamDefinition) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetLastUpdateTimestamp

`func (o *TelemetryStreamDefinition) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *TelemetryStreamDefinition) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *TelemetryStreamDefinition) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *TelemetryStreamDefinition) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetMetricValuation

`func (o *TelemetryStreamDefinition) GetMetricValuation() MetricValuation`

GetMetricValuation returns the MetricValuation field if non-nil, zero value otherwise.

### GetMetricValuationOk

`func (o *TelemetryStreamDefinition) GetMetricValuationOk() (*MetricValuation, bool)`

GetMetricValuationOk returns a tuple with the MetricValuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValuation

`func (o *TelemetryStreamDefinition) SetMetricValuation(v MetricValuation)`

SetMetricValuation sets MetricValuation field to given value.


### GetName

`func (o *TelemetryStreamDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryStreamDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryStreamDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetOwnedBy

`func (o *TelemetryStreamDefinition) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *TelemetryStreamDefinition) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *TelemetryStreamDefinition) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *TelemetryStreamDefinition) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetPriority

`func (o *TelemetryStreamDefinition) GetPriority() TelemetryStreamPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TelemetryStreamDefinition) GetPriorityOk() (*TelemetryStreamPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TelemetryStreamDefinition) SetPriority(v TelemetryStreamPriority)`

SetPriority sets Priority field to given value.


### GetTelemetryQuery

`func (o *TelemetryStreamDefinition) GetTelemetryQuery() EventTelemetryQuery`

GetTelemetryQuery returns the TelemetryQuery field if non-nil, zero value otherwise.

### GetTelemetryQueryOk

`func (o *TelemetryStreamDefinition) GetTelemetryQueryOk() (*EventTelemetryQuery, bool)`

GetTelemetryQueryOk returns a tuple with the TelemetryQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryQuery

`func (o *TelemetryStreamDefinition) SetTelemetryQuery(v EventTelemetryQuery)`

SetTelemetryQuery sets TelemetryQuery field to given value.


### GetTopologyMapping

`func (o *TelemetryStreamDefinition) GetTopologyMapping() []TopologyMapping`

GetTopologyMapping returns the TopologyMapping field if non-nil, zero value otherwise.

### GetTopologyMappingOk

`func (o *TelemetryStreamDefinition) GetTopologyMappingOk() (*[]TopologyMapping, bool)`

GetTopologyMappingOk returns a tuple with the TopologyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyMapping

`func (o *TelemetryStreamDefinition) SetTopologyMapping(v []TopologyMapping)`

SetTopologyMapping sets TopologyMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


