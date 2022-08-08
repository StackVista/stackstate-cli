# ArgumentPropagatedHealthStateVal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Parameter** | **int64** |  | 
**Value** | [**PropagatedHealthStateValue**](PropagatedHealthStateValue.md) |  | 

## Methods

### NewArgumentPropagatedHealthStateVal

`func NewArgumentPropagatedHealthStateVal(type_ string, parameter int64, value PropagatedHealthStateValue, ) *ArgumentPropagatedHealthStateVal`

NewArgumentPropagatedHealthStateVal instantiates a new ArgumentPropagatedHealthStateVal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgumentPropagatedHealthStateValWithDefaults

`func NewArgumentPropagatedHealthStateValWithDefaults() *ArgumentPropagatedHealthStateVal`

NewArgumentPropagatedHealthStateValWithDefaults instantiates a new ArgumentPropagatedHealthStateVal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ArgumentPropagatedHealthStateVal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArgumentPropagatedHealthStateVal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArgumentPropagatedHealthStateVal) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ArgumentPropagatedHealthStateVal) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArgumentPropagatedHealthStateVal) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArgumentPropagatedHealthStateVal) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ArgumentPropagatedHealthStateVal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ArgumentPropagatedHealthStateVal) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ArgumentPropagatedHealthStateVal) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ArgumentPropagatedHealthStateVal) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *ArgumentPropagatedHealthStateVal) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetParameter

`func (o *ArgumentPropagatedHealthStateVal) GetParameter() int64`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ArgumentPropagatedHealthStateVal) GetParameterOk() (*int64, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ArgumentPropagatedHealthStateVal) SetParameter(v int64)`

SetParameter sets Parameter field to given value.


### GetValue

`func (o *ArgumentPropagatedHealthStateVal) GetValue() PropagatedHealthStateValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ArgumentPropagatedHealthStateVal) GetValueOk() (*PropagatedHealthStateValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ArgumentPropagatedHealthStateVal) SetValue(v PropagatedHealthStateValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


