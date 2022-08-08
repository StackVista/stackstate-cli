# ArgumentMetricStreamRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Parameter** | **int64** |  | 
**MaxWindow** | Pointer to **int64** |  | [optional] 
**Stream** | **int64** |  | 
**DownsamplingMethod** | Pointer to [**DownsamplingMethod**](DownsamplingMethod.md) |  | [optional] 
**WindowingMethod** | Pointer to [**WindowingMethod**](WindowingMethod.md) |  | [optional] 

## Methods

### NewArgumentMetricStreamRef

`func NewArgumentMetricStreamRef(type_ string, parameter int64, stream int64, ) *ArgumentMetricStreamRef`

NewArgumentMetricStreamRef instantiates a new ArgumentMetricStreamRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgumentMetricStreamRefWithDefaults

`func NewArgumentMetricStreamRefWithDefaults() *ArgumentMetricStreamRef`

NewArgumentMetricStreamRefWithDefaults instantiates a new ArgumentMetricStreamRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ArgumentMetricStreamRef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArgumentMetricStreamRef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArgumentMetricStreamRef) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ArgumentMetricStreamRef) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArgumentMetricStreamRef) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArgumentMetricStreamRef) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ArgumentMetricStreamRef) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ArgumentMetricStreamRef) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ArgumentMetricStreamRef) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ArgumentMetricStreamRef) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *ArgumentMetricStreamRef) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetParameter

`func (o *ArgumentMetricStreamRef) GetParameter() int64`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ArgumentMetricStreamRef) GetParameterOk() (*int64, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ArgumentMetricStreamRef) SetParameter(v int64)`

SetParameter sets Parameter field to given value.


### GetMaxWindow

`func (o *ArgumentMetricStreamRef) GetMaxWindow() int64`

GetMaxWindow returns the MaxWindow field if non-nil, zero value otherwise.

### GetMaxWindowOk

`func (o *ArgumentMetricStreamRef) GetMaxWindowOk() (*int64, bool)`

GetMaxWindowOk returns a tuple with the MaxWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWindow

`func (o *ArgumentMetricStreamRef) SetMaxWindow(v int64)`

SetMaxWindow sets MaxWindow field to given value.

### HasMaxWindow

`func (o *ArgumentMetricStreamRef) HasMaxWindow() bool`

HasMaxWindow returns a boolean if a field has been set.

### GetStream

`func (o *ArgumentMetricStreamRef) GetStream() int64`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ArgumentMetricStreamRef) GetStreamOk() (*int64, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ArgumentMetricStreamRef) SetStream(v int64)`

SetStream sets Stream field to given value.


### GetDownsamplingMethod

`func (o *ArgumentMetricStreamRef) GetDownsamplingMethod() DownsamplingMethod`

GetDownsamplingMethod returns the DownsamplingMethod field if non-nil, zero value otherwise.

### GetDownsamplingMethodOk

`func (o *ArgumentMetricStreamRef) GetDownsamplingMethodOk() (*DownsamplingMethod, bool)`

GetDownsamplingMethodOk returns a tuple with the DownsamplingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownsamplingMethod

`func (o *ArgumentMetricStreamRef) SetDownsamplingMethod(v DownsamplingMethod)`

SetDownsamplingMethod sets DownsamplingMethod field to given value.

### HasDownsamplingMethod

`func (o *ArgumentMetricStreamRef) HasDownsamplingMethod() bool`

HasDownsamplingMethod returns a boolean if a field has been set.

### GetWindowingMethod

`func (o *ArgumentMetricStreamRef) GetWindowingMethod() WindowingMethod`

GetWindowingMethod returns the WindowingMethod field if non-nil, zero value otherwise.

### GetWindowingMethodOk

`func (o *ArgumentMetricStreamRef) GetWindowingMethodOk() (*WindowingMethod, bool)`

GetWindowingMethodOk returns a tuple with the WindowingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowingMethod

`func (o *ArgumentMetricStreamRef) SetWindowingMethod(v WindowingMethod)`

SetWindowingMethod sets WindowingMethod field to given value.

### HasWindowingMethod

`func (o *ArgumentMetricStreamRef) HasWindowingMethod() bool`

HasWindowingMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


