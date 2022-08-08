# ArgumentQueryViewRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Parameter** | **int64** |  | 
**QueryView** | **int64** |  | 

## Methods

### NewArgumentQueryViewRef

`func NewArgumentQueryViewRef(type_ string, parameter int64, queryView int64, ) *ArgumentQueryViewRef`

NewArgumentQueryViewRef instantiates a new ArgumentQueryViewRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgumentQueryViewRefWithDefaults

`func NewArgumentQueryViewRefWithDefaults() *ArgumentQueryViewRef`

NewArgumentQueryViewRefWithDefaults instantiates a new ArgumentQueryViewRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ArgumentQueryViewRef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArgumentQueryViewRef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArgumentQueryViewRef) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ArgumentQueryViewRef) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArgumentQueryViewRef) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArgumentQueryViewRef) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ArgumentQueryViewRef) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ArgumentQueryViewRef) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ArgumentQueryViewRef) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ArgumentQueryViewRef) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *ArgumentQueryViewRef) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetParameter

`func (o *ArgumentQueryViewRef) GetParameter() int64`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ArgumentQueryViewRef) GetParameterOk() (*int64, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ArgumentQueryViewRef) SetParameter(v int64)`

SetParameter sets Parameter field to given value.


### GetQueryView

`func (o *ArgumentQueryViewRef) GetQueryView() int64`

GetQueryView returns the QueryView field if non-nil, zero value otherwise.

### GetQueryViewOk

`func (o *ArgumentQueryViewRef) GetQueryViewOk() (*int64, bool)`

GetQueryViewOk returns a tuple with the QueryView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryView

`func (o *ArgumentQueryViewRef) SetQueryView(v int64)`

SetQueryView sets QueryView field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


