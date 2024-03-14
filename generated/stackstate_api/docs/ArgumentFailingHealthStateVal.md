# ArgumentFailingHealthStateVal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Parameter** | **int64** |  | 
**Value** | [**FailingHealthStateValue**](FailingHealthStateValue.md) |  | 

## Methods

### NewArgumentFailingHealthStateVal

`func NewArgumentFailingHealthStateVal(type_ string, parameter int64, value FailingHealthStateValue, ) *ArgumentFailingHealthStateVal`

NewArgumentFailingHealthStateVal instantiates a new ArgumentFailingHealthStateVal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgumentFailingHealthStateValWithDefaults

`func NewArgumentFailingHealthStateValWithDefaults() *ArgumentFailingHealthStateVal`

NewArgumentFailingHealthStateValWithDefaults instantiates a new ArgumentFailingHealthStateVal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ArgumentFailingHealthStateVal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArgumentFailingHealthStateVal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArgumentFailingHealthStateVal) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ArgumentFailingHealthStateVal) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArgumentFailingHealthStateVal) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArgumentFailingHealthStateVal) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ArgumentFailingHealthStateVal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ArgumentFailingHealthStateVal) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ArgumentFailingHealthStateVal) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ArgumentFailingHealthStateVal) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *ArgumentFailingHealthStateVal) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetParameter

`func (o *ArgumentFailingHealthStateVal) GetParameter() int64`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ArgumentFailingHealthStateVal) GetParameterOk() (*int64, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ArgumentFailingHealthStateVal) SetParameter(v int64)`

SetParameter sets Parameter field to given value.


### GetValue

`func (o *ArgumentFailingHealthStateVal) GetValue() FailingHealthStateValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ArgumentFailingHealthStateVal) GetValueOk() (*FailingHealthStateValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ArgumentFailingHealthStateVal) SetValue(v FailingHealthStateValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


