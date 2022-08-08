# ArgumentEventStreamRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Parameter** | **int64** |  | 
**Stream** | **int64** |  | 

## Methods

### NewArgumentEventStreamRef

`func NewArgumentEventStreamRef(type_ string, parameter int64, stream int64, ) *ArgumentEventStreamRef`

NewArgumentEventStreamRef instantiates a new ArgumentEventStreamRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgumentEventStreamRefWithDefaults

`func NewArgumentEventStreamRefWithDefaults() *ArgumentEventStreamRef`

NewArgumentEventStreamRefWithDefaults instantiates a new ArgumentEventStreamRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ArgumentEventStreamRef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArgumentEventStreamRef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArgumentEventStreamRef) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ArgumentEventStreamRef) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArgumentEventStreamRef) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArgumentEventStreamRef) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ArgumentEventStreamRef) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ArgumentEventStreamRef) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ArgumentEventStreamRef) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ArgumentEventStreamRef) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *ArgumentEventStreamRef) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetParameter

`func (o *ArgumentEventStreamRef) GetParameter() int64`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ArgumentEventStreamRef) GetParameterOk() (*int64, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ArgumentEventStreamRef) SetParameter(v int64)`

SetParameter sets Parameter field to given value.


### GetStream

`func (o *ArgumentEventStreamRef) GetStream() int64`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ArgumentEventStreamRef) GetStreamOk() (*int64, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ArgumentEventStreamRef) SetStream(v int64)`

SetStream sets Stream field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


