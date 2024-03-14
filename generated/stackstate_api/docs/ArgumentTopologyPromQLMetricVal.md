# ArgumentTopologyPromQLMetricVal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Parameter** | **int64** |  | 
**Value** | [**TopologyPromQLMetric**](TopologyPromQLMetric.md) |  | 

## Methods

### NewArgumentTopologyPromQLMetricVal

`func NewArgumentTopologyPromQLMetricVal(type_ string, parameter int64, value TopologyPromQLMetric, ) *ArgumentTopologyPromQLMetricVal`

NewArgumentTopologyPromQLMetricVal instantiates a new ArgumentTopologyPromQLMetricVal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgumentTopologyPromQLMetricValWithDefaults

`func NewArgumentTopologyPromQLMetricValWithDefaults() *ArgumentTopologyPromQLMetricVal`

NewArgumentTopologyPromQLMetricValWithDefaults instantiates a new ArgumentTopologyPromQLMetricVal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ArgumentTopologyPromQLMetricVal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArgumentTopologyPromQLMetricVal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArgumentTopologyPromQLMetricVal) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ArgumentTopologyPromQLMetricVal) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArgumentTopologyPromQLMetricVal) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArgumentTopologyPromQLMetricVal) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ArgumentTopologyPromQLMetricVal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ArgumentTopologyPromQLMetricVal) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ArgumentTopologyPromQLMetricVal) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ArgumentTopologyPromQLMetricVal) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *ArgumentTopologyPromQLMetricVal) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetParameter

`func (o *ArgumentTopologyPromQLMetricVal) GetParameter() int64`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ArgumentTopologyPromQLMetricVal) GetParameterOk() (*int64, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ArgumentTopologyPromQLMetricVal) SetParameter(v int64)`

SetParameter sets Parameter field to given value.


### GetValue

`func (o *ArgumentTopologyPromQLMetricVal) GetValue() TopologyPromQLMetric`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ArgumentTopologyPromQLMetricVal) GetValueOk() (*TopologyPromQLMetric, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ArgumentTopologyPromQLMetricVal) SetValue(v TopologyPromQLMetric)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


