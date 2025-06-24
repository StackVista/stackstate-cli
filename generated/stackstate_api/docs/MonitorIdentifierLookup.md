# MonitorIdentifierLookup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricQuery** | **string** |  | 
**ComponentTypeName** | **string** |  | 
**TopN** | Pointer to **int32** |  | [optional] 
**Overrides** | Pointer to [**MonitorIdentifierLookupOverrides**](MonitorIdentifierLookupOverrides.md) |  | [optional] 

## Methods

### NewMonitorIdentifierLookup

`func NewMonitorIdentifierLookup(metricQuery string, componentTypeName string, ) *MonitorIdentifierLookup`

NewMonitorIdentifierLookup instantiates a new MonitorIdentifierLookup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorIdentifierLookupWithDefaults

`func NewMonitorIdentifierLookupWithDefaults() *MonitorIdentifierLookup`

NewMonitorIdentifierLookupWithDefaults instantiates a new MonitorIdentifierLookup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricQuery

`func (o *MonitorIdentifierLookup) GetMetricQuery() string`

GetMetricQuery returns the MetricQuery field if non-nil, zero value otherwise.

### GetMetricQueryOk

`func (o *MonitorIdentifierLookup) GetMetricQueryOk() (*string, bool)`

GetMetricQueryOk returns a tuple with the MetricQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricQuery

`func (o *MonitorIdentifierLookup) SetMetricQuery(v string)`

SetMetricQuery sets MetricQuery field to given value.


### GetComponentTypeName

`func (o *MonitorIdentifierLookup) GetComponentTypeName() string`

GetComponentTypeName returns the ComponentTypeName field if non-nil, zero value otherwise.

### GetComponentTypeNameOk

`func (o *MonitorIdentifierLookup) GetComponentTypeNameOk() (*string, bool)`

GetComponentTypeNameOk returns a tuple with the ComponentTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentTypeName

`func (o *MonitorIdentifierLookup) SetComponentTypeName(v string)`

SetComponentTypeName sets ComponentTypeName field to given value.


### GetTopN

`func (o *MonitorIdentifierLookup) GetTopN() int32`

GetTopN returns the TopN field if non-nil, zero value otherwise.

### GetTopNOk

`func (o *MonitorIdentifierLookup) GetTopNOk() (*int32, bool)`

GetTopNOk returns a tuple with the TopN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopN

`func (o *MonitorIdentifierLookup) SetTopN(v int32)`

SetTopN sets TopN field to given value.

### HasTopN

`func (o *MonitorIdentifierLookup) HasTopN() bool`

HasTopN returns a boolean if a field has been set.

### GetOverrides

`func (o *MonitorIdentifierLookup) GetOverrides() MonitorIdentifierLookupOverrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *MonitorIdentifierLookup) GetOverridesOk() (*MonitorIdentifierLookupOverrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *MonitorIdentifierLookup) SetOverrides(v MonitorIdentifierLookupOverrides)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *MonitorIdentifierLookup) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


