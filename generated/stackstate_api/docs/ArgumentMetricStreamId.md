# ArgumentMetricStreamId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Parameter** | **int64** |  | 
**Stream** | **int64** |  | 

## Methods

### NewArgumentMetricStreamId

`func NewArgumentMetricStreamId(type_ string, parameter int64, stream int64, ) *ArgumentMetricStreamId`

NewArgumentMetricStreamId instantiates a new ArgumentMetricStreamId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgumentMetricStreamIdWithDefaults

`func NewArgumentMetricStreamIdWithDefaults() *ArgumentMetricStreamId`

NewArgumentMetricStreamIdWithDefaults instantiates a new ArgumentMetricStreamId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ArgumentMetricStreamId) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArgumentMetricStreamId) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArgumentMetricStreamId) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ArgumentMetricStreamId) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArgumentMetricStreamId) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArgumentMetricStreamId) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ArgumentMetricStreamId) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ArgumentMetricStreamId) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ArgumentMetricStreamId) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ArgumentMetricStreamId) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *ArgumentMetricStreamId) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetParameter

`func (o *ArgumentMetricStreamId) GetParameter() int64`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ArgumentMetricStreamId) GetParameterOk() (*int64, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ArgumentMetricStreamId) SetParameter(v int64)`

SetParameter sets Parameter field to given value.


### GetStream

`func (o *ArgumentMetricStreamId) GetStream() int64`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ArgumentMetricStreamId) GetStreamOk() (*int64, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ArgumentMetricStreamId) SetStream(v int64)`

SetStream sets Stream field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


