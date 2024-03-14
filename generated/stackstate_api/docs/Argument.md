# Argument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Parameter** | **int64** |  | 
**Value** | [**TopologyPromQLMetric**](TopologyPromQLMetric.md) |  | 
**ComponentType** | **int64** |  | 
**RelationType** | **int64** |  | 
**QueryView** | **int64** |  | 
**ValueMs** | **int64** |  | 
**Query** | **string** |  | 

## Methods

### NewArgument

`func NewArgument(type_ string, parameter int64, value TopologyPromQLMetric, componentType int64, relationType int64, queryView int64, valueMs int64, query string, ) *Argument`

NewArgument instantiates a new Argument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgumentWithDefaults

`func NewArgumentWithDefaults() *Argument`

NewArgumentWithDefaults instantiates a new Argument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Argument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Argument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Argument) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *Argument) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Argument) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Argument) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Argument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *Argument) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *Argument) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *Argument) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *Argument) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetParameter

`func (o *Argument) GetParameter() int64`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *Argument) GetParameterOk() (*int64, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *Argument) SetParameter(v int64)`

SetParameter sets Parameter field to given value.


### GetValue

`func (o *Argument) GetValue() TopologyPromQLMetric`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Argument) GetValueOk() (*TopologyPromQLMetric, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Argument) SetValue(v TopologyPromQLMetric)`

SetValue sets Value field to given value.


### GetComponentType

`func (o *Argument) GetComponentType() int64`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *Argument) GetComponentTypeOk() (*int64, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *Argument) SetComponentType(v int64)`

SetComponentType sets ComponentType field to given value.


### GetRelationType

`func (o *Argument) GetRelationType() int64`

GetRelationType returns the RelationType field if non-nil, zero value otherwise.

### GetRelationTypeOk

`func (o *Argument) GetRelationTypeOk() (*int64, bool)`

GetRelationTypeOk returns a tuple with the RelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationType

`func (o *Argument) SetRelationType(v int64)`

SetRelationType sets RelationType field to given value.


### GetQueryView

`func (o *Argument) GetQueryView() int64`

GetQueryView returns the QueryView field if non-nil, zero value otherwise.

### GetQueryViewOk

`func (o *Argument) GetQueryViewOk() (*int64, bool)`

GetQueryViewOk returns a tuple with the QueryView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryView

`func (o *Argument) SetQueryView(v int64)`

SetQueryView sets QueryView field to given value.


### GetValueMs

`func (o *Argument) GetValueMs() int64`

GetValueMs returns the ValueMs field if non-nil, zero value otherwise.

### GetValueMsOk

`func (o *Argument) GetValueMsOk() (*int64, bool)`

GetValueMsOk returns a tuple with the ValueMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMs

`func (o *Argument) SetValueMs(v int64)`

SetValueMs sets ValueMs field to given value.


### GetQuery

`func (o *Argument) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Argument) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Argument) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


