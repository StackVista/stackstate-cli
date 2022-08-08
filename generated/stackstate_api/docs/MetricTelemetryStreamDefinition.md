# MetricTelemetryStreamDefinition

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
**TelemetryQuery** | [**MetricTelemetryQuery**](MetricTelemetryQuery.md) |  | 
**TopologyMapping** | [**[]TopologyMapping**](TopologyMapping.md) |  | 

## Methods

### NewMetricTelemetryStreamDefinition

`func NewMetricTelemetryStreamDefinition(type_ string, bindQuery string, dataSourceId int64, labels []string, metricValuation MetricValuation, name string, priority TelemetryStreamPriority, telemetryQuery MetricTelemetryQuery, topologyMapping []TopologyMapping, ) *MetricTelemetryStreamDefinition`

NewMetricTelemetryStreamDefinition instantiates a new MetricTelemetryStreamDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricTelemetryStreamDefinitionWithDefaults

`func NewMetricTelemetryStreamDefinitionWithDefaults() *MetricTelemetryStreamDefinition`

NewMetricTelemetryStreamDefinitionWithDefaults instantiates a new MetricTelemetryStreamDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricTelemetryStreamDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricTelemetryStreamDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricTelemetryStreamDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetBindQuery

`func (o *MetricTelemetryStreamDefinition) GetBindQuery() string`

GetBindQuery returns the BindQuery field if non-nil, zero value otherwise.

### GetBindQueryOk

`func (o *MetricTelemetryStreamDefinition) GetBindQueryOk() (*string, bool)`

GetBindQueryOk returns a tuple with the BindQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindQuery

`func (o *MetricTelemetryStreamDefinition) SetBindQuery(v string)`

SetBindQuery sets BindQuery field to given value.


### GetDataSourceId

`func (o *MetricTelemetryStreamDefinition) GetDataSourceId() int64`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *MetricTelemetryStreamDefinition) GetDataSourceIdOk() (*int64, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *MetricTelemetryStreamDefinition) SetDataSourceId(v int64)`

SetDataSourceId sets DataSourceId field to given value.


### GetDescription

`func (o *MetricTelemetryStreamDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricTelemetryStreamDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricTelemetryStreamDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricTelemetryStreamDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *MetricTelemetryStreamDefinition) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricTelemetryStreamDefinition) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricTelemetryStreamDefinition) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MetricTelemetryStreamDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *MetricTelemetryStreamDefinition) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MetricTelemetryStreamDefinition) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MetricTelemetryStreamDefinition) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *MetricTelemetryStreamDefinition) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetLabels

`func (o *MetricTelemetryStreamDefinition) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MetricTelemetryStreamDefinition) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MetricTelemetryStreamDefinition) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetLastUpdateTimestamp

`func (o *MetricTelemetryStreamDefinition) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *MetricTelemetryStreamDefinition) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *MetricTelemetryStreamDefinition) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *MetricTelemetryStreamDefinition) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetMetricValuation

`func (o *MetricTelemetryStreamDefinition) GetMetricValuation() MetricValuation`

GetMetricValuation returns the MetricValuation field if non-nil, zero value otherwise.

### GetMetricValuationOk

`func (o *MetricTelemetryStreamDefinition) GetMetricValuationOk() (*MetricValuation, bool)`

GetMetricValuationOk returns a tuple with the MetricValuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValuation

`func (o *MetricTelemetryStreamDefinition) SetMetricValuation(v MetricValuation)`

SetMetricValuation sets MetricValuation field to given value.


### GetName

`func (o *MetricTelemetryStreamDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricTelemetryStreamDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricTelemetryStreamDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetOwnedBy

`func (o *MetricTelemetryStreamDefinition) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *MetricTelemetryStreamDefinition) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *MetricTelemetryStreamDefinition) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *MetricTelemetryStreamDefinition) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetPriority

`func (o *MetricTelemetryStreamDefinition) GetPriority() TelemetryStreamPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MetricTelemetryStreamDefinition) GetPriorityOk() (*TelemetryStreamPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MetricTelemetryStreamDefinition) SetPriority(v TelemetryStreamPriority)`

SetPriority sets Priority field to given value.


### GetTelemetryQuery

`func (o *MetricTelemetryStreamDefinition) GetTelemetryQuery() MetricTelemetryQuery`

GetTelemetryQuery returns the TelemetryQuery field if non-nil, zero value otherwise.

### GetTelemetryQueryOk

`func (o *MetricTelemetryStreamDefinition) GetTelemetryQueryOk() (*MetricTelemetryQuery, bool)`

GetTelemetryQueryOk returns a tuple with the TelemetryQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryQuery

`func (o *MetricTelemetryStreamDefinition) SetTelemetryQuery(v MetricTelemetryQuery)`

SetTelemetryQuery sets TelemetryQuery field to given value.


### GetTopologyMapping

`func (o *MetricTelemetryStreamDefinition) GetTopologyMapping() []TopologyMapping`

GetTopologyMapping returns the TopologyMapping field if non-nil, zero value otherwise.

### GetTopologyMappingOk

`func (o *MetricTelemetryStreamDefinition) GetTopologyMappingOk() (*[]TopologyMapping, bool)`

GetTopologyMappingOk returns a tuple with the TopologyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyMapping

`func (o *MetricTelemetryStreamDefinition) SetTopologyMapping(v []TopologyMapping)`

SetTopologyMapping sets TopologyMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


