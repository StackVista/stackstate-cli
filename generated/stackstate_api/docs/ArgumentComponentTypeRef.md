# ArgumentComponentTypeRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Parameter** | **int64** |  | 
**ComponentType** | **int64** |  | 

## Methods

### NewArgumentComponentTypeRef

`func NewArgumentComponentTypeRef(type_ string, parameter int64, componentType int64, ) *ArgumentComponentTypeRef`

NewArgumentComponentTypeRef instantiates a new ArgumentComponentTypeRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgumentComponentTypeRefWithDefaults

`func NewArgumentComponentTypeRefWithDefaults() *ArgumentComponentTypeRef`

NewArgumentComponentTypeRefWithDefaults instantiates a new ArgumentComponentTypeRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ArgumentComponentTypeRef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArgumentComponentTypeRef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArgumentComponentTypeRef) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ArgumentComponentTypeRef) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArgumentComponentTypeRef) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArgumentComponentTypeRef) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ArgumentComponentTypeRef) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ArgumentComponentTypeRef) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ArgumentComponentTypeRef) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ArgumentComponentTypeRef) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *ArgumentComponentTypeRef) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetParameter

`func (o *ArgumentComponentTypeRef) GetParameter() int64`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ArgumentComponentTypeRef) GetParameterOk() (*int64, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ArgumentComponentTypeRef) SetParameter(v int64)`

SetParameter sets Parameter field to given value.


### GetComponentType

`func (o *ArgumentComponentTypeRef) GetComponentType() int64`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *ArgumentComponentTypeRef) GetComponentTypeOk() (*int64, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *ArgumentComponentTypeRef) SetComponentType(v int64)`

SetComponentType sets ComponentType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


