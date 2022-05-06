# HealthSubStreamStatus1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]HealthStreamError**](HealthStreamError.md) |  | [optional] 
**Metrics** | [**HealthStreamMetrics1**](HealthStreamMetrics1.md) |  | 
**SubStreamState** | [**HealthSubStreamConsistencyState**](HealthSubStreamConsistencyState.md) |  | 
**CheckStateCount** | **int32** |  | 

## Methods

### NewHealthSubStreamStatus1

`func NewHealthSubStreamStatus1(metrics HealthStreamMetrics1, subStreamState HealthSubStreamConsistencyState, checkStateCount int32, ) *HealthSubStreamStatus1`

NewHealthSubStreamStatus1 instantiates a new HealthSubStreamStatus1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthSubStreamStatus1WithDefaults

`func NewHealthSubStreamStatus1WithDefaults() *HealthSubStreamStatus1`

NewHealthSubStreamStatus1WithDefaults instantiates a new HealthSubStreamStatus1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *HealthSubStreamStatus1) GetErrors() []HealthStreamError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *HealthSubStreamStatus1) GetErrorsOk() (*[]HealthStreamError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *HealthSubStreamStatus1) SetErrors(v []HealthStreamError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *HealthSubStreamStatus1) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMetrics

`func (o *HealthSubStreamStatus1) GetMetrics() HealthStreamMetrics1`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *HealthSubStreamStatus1) GetMetricsOk() (*HealthStreamMetrics1, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *HealthSubStreamStatus1) SetMetrics(v HealthStreamMetrics1)`

SetMetrics sets Metrics field to given value.


### GetSubStreamState

`func (o *HealthSubStreamStatus1) GetSubStreamState() HealthSubStreamConsistencyState`

GetSubStreamState returns the SubStreamState field if non-nil, zero value otherwise.

### GetSubStreamStateOk

`func (o *HealthSubStreamStatus1) GetSubStreamStateOk() (*HealthSubStreamConsistencyState, bool)`

GetSubStreamStateOk returns a tuple with the SubStreamState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStreamState

`func (o *HealthSubStreamStatus1) SetSubStreamState(v HealthSubStreamConsistencyState)`

SetSubStreamState sets SubStreamState field to given value.


### GetCheckStateCount

`func (o *HealthSubStreamStatus1) GetCheckStateCount() int32`

GetCheckStateCount returns the CheckStateCount field if non-nil, zero value otherwise.

### GetCheckStateCountOk

`func (o *HealthSubStreamStatus1) GetCheckStateCountOk() (*int32, bool)`

GetCheckStateCountOk returns a tuple with the CheckStateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStateCount

`func (o *HealthSubStreamStatus1) SetCheckStateCount(v int32)`

SetCheckStateCount sets CheckStateCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


