# HealthStreamStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partition** | **int32** |  | 
**ConsistencyModel** | **string** |  | 
**RecoverMessage** | Pointer to **string** |  | [optional] 
**GlobalErrors** | Pointer to [**[]HealthStreamError**](HealthStreamError.md) |  | [optional] 
**AggregateMetrics** | [**HealthStreamMetrics**](HealthStreamMetrics.md) |  | 
**MainStreamStatus** | Pointer to [**HealthSubStreamStatus**](HealthSubStreamStatus.md) |  | [optional] 

## Methods

### NewHealthStreamStatus

`func NewHealthStreamStatus(partition int32, consistencyModel string, aggregateMetrics HealthStreamMetrics, ) *HealthStreamStatus`

NewHealthStreamStatus instantiates a new HealthStreamStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthStreamStatusWithDefaults

`func NewHealthStreamStatusWithDefaults() *HealthStreamStatus`

NewHealthStreamStatusWithDefaults instantiates a new HealthStreamStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartition

`func (o *HealthStreamStatus) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *HealthStreamStatus) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *HealthStreamStatus) SetPartition(v int32)`

SetPartition sets Partition field to given value.


### GetConsistencyModel

`func (o *HealthStreamStatus) GetConsistencyModel() string`

GetConsistencyModel returns the ConsistencyModel field if non-nil, zero value otherwise.

### GetConsistencyModelOk

`func (o *HealthStreamStatus) GetConsistencyModelOk() (*string, bool)`

GetConsistencyModelOk returns a tuple with the ConsistencyModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencyModel

`func (o *HealthStreamStatus) SetConsistencyModel(v string)`

SetConsistencyModel sets ConsistencyModel field to given value.


### GetRecoverMessage

`func (o *HealthStreamStatus) GetRecoverMessage() string`

GetRecoverMessage returns the RecoverMessage field if non-nil, zero value otherwise.

### GetRecoverMessageOk

`func (o *HealthStreamStatus) GetRecoverMessageOk() (*string, bool)`

GetRecoverMessageOk returns a tuple with the RecoverMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverMessage

`func (o *HealthStreamStatus) SetRecoverMessage(v string)`

SetRecoverMessage sets RecoverMessage field to given value.

### HasRecoverMessage

`func (o *HealthStreamStatus) HasRecoverMessage() bool`

HasRecoverMessage returns a boolean if a field has been set.

### GetGlobalErrors

`func (o *HealthStreamStatus) GetGlobalErrors() []HealthStreamError`

GetGlobalErrors returns the GlobalErrors field if non-nil, zero value otherwise.

### GetGlobalErrorsOk

`func (o *HealthStreamStatus) GetGlobalErrorsOk() (*[]HealthStreamError, bool)`

GetGlobalErrorsOk returns a tuple with the GlobalErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalErrors

`func (o *HealthStreamStatus) SetGlobalErrors(v []HealthStreamError)`

SetGlobalErrors sets GlobalErrors field to given value.

### HasGlobalErrors

`func (o *HealthStreamStatus) HasGlobalErrors() bool`

HasGlobalErrors returns a boolean if a field has been set.

### GetAggregateMetrics

`func (o *HealthStreamStatus) GetAggregateMetrics() HealthStreamMetrics`

GetAggregateMetrics returns the AggregateMetrics field if non-nil, zero value otherwise.

### GetAggregateMetricsOk

`func (o *HealthStreamStatus) GetAggregateMetricsOk() (*HealthStreamMetrics, bool)`

GetAggregateMetricsOk returns a tuple with the AggregateMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateMetrics

`func (o *HealthStreamStatus) SetAggregateMetrics(v HealthStreamMetrics)`

SetAggregateMetrics sets AggregateMetrics field to given value.


### GetMainStreamStatus

`func (o *HealthStreamStatus) GetMainStreamStatus() HealthSubStreamStatus`

GetMainStreamStatus returns the MainStreamStatus field if non-nil, zero value otherwise.

### GetMainStreamStatusOk

`func (o *HealthStreamStatus) GetMainStreamStatusOk() (*HealthSubStreamStatus, bool)`

GetMainStreamStatusOk returns a tuple with the MainStreamStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainStreamStatus

`func (o *HealthStreamStatus) SetMainStreamStatus(v HealthSubStreamStatus)`

SetMainStreamStatus sets MainStreamStatus field to given value.

### HasMainStreamStatus

`func (o *HealthStreamStatus) HasMainStreamStatus() bool`

HasMainStreamStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


