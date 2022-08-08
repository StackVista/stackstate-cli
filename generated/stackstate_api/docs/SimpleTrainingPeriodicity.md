# SimpleTrainingPeriodicity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**FundamentalPeriod** | **int64** |  | 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**TrainingWindowPeriods** | **int64** |  | 

## Methods

### NewSimpleTrainingPeriodicity

`func NewSimpleTrainingPeriodicity(type_ string, fundamentalPeriod int64, trainingWindowPeriods int64, ) *SimpleTrainingPeriodicity`

NewSimpleTrainingPeriodicity instantiates a new SimpleTrainingPeriodicity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleTrainingPeriodicityWithDefaults

`func NewSimpleTrainingPeriodicityWithDefaults() *SimpleTrainingPeriodicity`

NewSimpleTrainingPeriodicityWithDefaults instantiates a new SimpleTrainingPeriodicity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SimpleTrainingPeriodicity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimpleTrainingPeriodicity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimpleTrainingPeriodicity) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *SimpleTrainingPeriodicity) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleTrainingPeriodicity) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleTrainingPeriodicity) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SimpleTrainingPeriodicity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFundamentalPeriod

`func (o *SimpleTrainingPeriodicity) GetFundamentalPeriod() int64`

GetFundamentalPeriod returns the FundamentalPeriod field if non-nil, zero value otherwise.

### GetFundamentalPeriodOk

`func (o *SimpleTrainingPeriodicity) GetFundamentalPeriodOk() (*int64, bool)`

GetFundamentalPeriodOk returns a tuple with the FundamentalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundamentalPeriod

`func (o *SimpleTrainingPeriodicity) SetFundamentalPeriod(v int64)`

SetFundamentalPeriod sets FundamentalPeriod field to given value.


### GetLastUpdateTimestamp

`func (o *SimpleTrainingPeriodicity) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *SimpleTrainingPeriodicity) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *SimpleTrainingPeriodicity) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *SimpleTrainingPeriodicity) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetTrainingWindowPeriods

`func (o *SimpleTrainingPeriodicity) GetTrainingWindowPeriods() int64`

GetTrainingWindowPeriods returns the TrainingWindowPeriods field if non-nil, zero value otherwise.

### GetTrainingWindowPeriodsOk

`func (o *SimpleTrainingPeriodicity) GetTrainingWindowPeriodsOk() (*int64, bool)`

GetTrainingWindowPeriodsOk returns a tuple with the TrainingWindowPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingWindowPeriods

`func (o *SimpleTrainingPeriodicity) SetTrainingWindowPeriods(v int64)`

SetTrainingWindowPeriods sets TrainingWindowPeriods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


