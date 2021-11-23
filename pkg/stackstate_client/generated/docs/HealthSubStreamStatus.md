# HealthSubStreamStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]HealthStreamError**](HealthStreamError.md) |  | [optional] 
**Metrics** | [**HealthStreamMetrics**](HealthStreamMetrics.md) |  | 
**SubStreamState** | [**HealthSubStreamConsistencyState**](HealthSubStreamConsistencyState.md) |  | 
**CheckStateCount** | **int32** |  | 

## Methods

### NewHealthSubStreamStatus

`func NewHealthSubStreamStatus(metrics HealthStreamMetrics, subStreamState HealthSubStreamConsistencyState, checkStateCount int32, ) *HealthSubStreamStatus`

NewHealthSubStreamStatus instantiates a new HealthSubStreamStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthSubStreamStatusWithDefaults

`func NewHealthSubStreamStatusWithDefaults() *HealthSubStreamStatus`

NewHealthSubStreamStatusWithDefaults instantiates a new HealthSubStreamStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *HealthSubStreamStatus) GetErrors() []HealthStreamError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *HealthSubStreamStatus) GetErrorsOk() (*[]HealthStreamError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *HealthSubStreamStatus) SetErrors(v []HealthStreamError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *HealthSubStreamStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMetrics

`func (o *HealthSubStreamStatus) GetMetrics() HealthStreamMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *HealthSubStreamStatus) GetMetricsOk() (*HealthStreamMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *HealthSubStreamStatus) SetMetrics(v HealthStreamMetrics)`

SetMetrics sets Metrics field to given value.


### GetSubStreamState

`func (o *HealthSubStreamStatus) GetSubStreamState() HealthSubStreamConsistencyState`

GetSubStreamState returns the SubStreamState field if non-nil, zero value otherwise.

### GetSubStreamStateOk

`func (o *HealthSubStreamStatus) GetSubStreamStateOk() (*HealthSubStreamConsistencyState, bool)`

GetSubStreamStateOk returns a tuple with the SubStreamState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStreamState

`func (o *HealthSubStreamStatus) SetSubStreamState(v HealthSubStreamConsistencyState)`

SetSubStreamState sets SubStreamState field to given value.


### GetCheckStateCount

`func (o *HealthSubStreamStatus) GetCheckStateCount() int32`

GetCheckStateCount returns the CheckStateCount field if non-nil, zero value otherwise.

### GetCheckStateCountOk

`func (o *HealthSubStreamStatus) GetCheckStateCountOk() (*int32, bool)`

GetCheckStateCountOk returns a tuple with the CheckStateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStateCount

`func (o *HealthSubStreamStatus) SetCheckStateCount(v int32)`

SetCheckStateCount sets CheckStateCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


