# HealthStreamStatus1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partition** | **int32** |  | 
**ConsistencyModel** | **string** |  | 
**RecoverMessage** | Pointer to **string** |  | [optional] 
**GlobalErrors** | Pointer to [**[]HealthStreamError**](HealthStreamError.md) |  | [optional] 
**AggregateMetrics** | [**HealthStreamMetrics1**](HealthStreamMetrics1.md) |  | 
**MainStreamStatus** | Pointer to [**HealthSubStreamStatus1**](HealthSubStreamStatus1.md) |  | [optional] 

## Methods

### NewHealthStreamStatus1

`func NewHealthStreamStatus1(partition int32, consistencyModel string, aggregateMetrics HealthStreamMetrics1, ) *HealthStreamStatus1`

NewHealthStreamStatus1 instantiates a new HealthStreamStatus1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthStreamStatus1WithDefaults

`func NewHealthStreamStatus1WithDefaults() *HealthStreamStatus1`

NewHealthStreamStatus1WithDefaults instantiates a new HealthStreamStatus1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartition

`func (o *HealthStreamStatus1) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *HealthStreamStatus1) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *HealthStreamStatus1) SetPartition(v int32)`

SetPartition sets Partition field to given value.


### GetConsistencyModel

`func (o *HealthStreamStatus1) GetConsistencyModel() string`

GetConsistencyModel returns the ConsistencyModel field if non-nil, zero value otherwise.

### GetConsistencyModelOk

`func (o *HealthStreamStatus1) GetConsistencyModelOk() (*string, bool)`

GetConsistencyModelOk returns a tuple with the ConsistencyModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencyModel

`func (o *HealthStreamStatus1) SetConsistencyModel(v string)`

SetConsistencyModel sets ConsistencyModel field to given value.


### GetRecoverMessage

`func (o *HealthStreamStatus1) GetRecoverMessage() string`

GetRecoverMessage returns the RecoverMessage field if non-nil, zero value otherwise.

### GetRecoverMessageOk

`func (o *HealthStreamStatus1) GetRecoverMessageOk() (*string, bool)`

GetRecoverMessageOk returns a tuple with the RecoverMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverMessage

`func (o *HealthStreamStatus1) SetRecoverMessage(v string)`

SetRecoverMessage sets RecoverMessage field to given value.

### HasRecoverMessage

`func (o *HealthStreamStatus1) HasRecoverMessage() bool`

HasRecoverMessage returns a boolean if a field has been set.

### GetGlobalErrors

`func (o *HealthStreamStatus1) GetGlobalErrors() []HealthStreamError`

GetGlobalErrors returns the GlobalErrors field if non-nil, zero value otherwise.

### GetGlobalErrorsOk

`func (o *HealthStreamStatus1) GetGlobalErrorsOk() (*[]HealthStreamError, bool)`

GetGlobalErrorsOk returns a tuple with the GlobalErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalErrors

`func (o *HealthStreamStatus1) SetGlobalErrors(v []HealthStreamError)`

SetGlobalErrors sets GlobalErrors field to given value.

### HasGlobalErrors

`func (o *HealthStreamStatus1) HasGlobalErrors() bool`

HasGlobalErrors returns a boolean if a field has been set.

### GetAggregateMetrics

`func (o *HealthStreamStatus1) GetAggregateMetrics() HealthStreamMetrics1`

GetAggregateMetrics returns the AggregateMetrics field if non-nil, zero value otherwise.

### GetAggregateMetricsOk

`func (o *HealthStreamStatus1) GetAggregateMetricsOk() (*HealthStreamMetrics1, bool)`

GetAggregateMetricsOk returns a tuple with the AggregateMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateMetrics

`func (o *HealthStreamStatus1) SetAggregateMetrics(v HealthStreamMetrics1)`

SetAggregateMetrics sets AggregateMetrics field to given value.


### GetMainStreamStatus

`func (o *HealthStreamStatus1) GetMainStreamStatus() HealthSubStreamStatus1`

GetMainStreamStatus returns the MainStreamStatus field if non-nil, zero value otherwise.

### GetMainStreamStatusOk

`func (o *HealthStreamStatus1) GetMainStreamStatusOk() (*HealthSubStreamStatus1, bool)`

GetMainStreamStatusOk returns a tuple with the MainStreamStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainStreamStatus

`func (o *HealthStreamStatus1) SetMainStreamStatus(v HealthSubStreamStatus1)`

SetMainStreamStatus sets MainStreamStatus field to given value.

### HasMainStreamStatus

`func (o *HealthStreamStatus1) HasMainStreamStatus() bool`

HasMainStreamStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


