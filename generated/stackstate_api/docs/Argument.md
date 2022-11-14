# Argument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Parameter** | **int64** |  | 
**Value** | **string** |  | 
**ComponentType** | **int64** |  | 
**Stream** | **int64** |  | 
**RelationType** | **int64** |  | 
**Filter** | [**StsEventStreamFilter**](StsEventStreamFilter.md) |  | 
**QueryView** | **int64** |  | 
**MaxWindow** | Pointer to **int64** |  | [optional] 
**DownsamplingMethod** | Pointer to [**DownsamplingMethod**](DownsamplingMethod.md) |  | [optional] 
**WindowingMethod** | Pointer to [**WindowingMethod**](WindowingMethod.md) |  | [optional] 
**Script** | **string** |  | 

## Methods

### NewArgument

`func NewArgument(type_ string, parameter int64, value string, componentType int64, stream int64, relationType int64, filter StsEventStreamFilter, queryView int64, script string, ) *Argument`

NewArgument instantiates a new Argument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgumentWithDefaults

`func NewArgumentWithDefaults() *Argument`

NewArgumentWithDefaults instantiates a new Argument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Argument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Argument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Argument) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *Argument) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Argument) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Argument) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Argument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *Argument) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *Argument) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *Argument) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *Argument) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetParameter

`func (o *Argument) GetParameter() int64`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *Argument) GetParameterOk() (*int64, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *Argument) SetParameter(v int64)`

SetParameter sets Parameter field to given value.


### GetValue

`func (o *Argument) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Argument) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Argument) SetValue(v string)`

SetValue sets Value field to given value.


### GetComponentType

`func (o *Argument) GetComponentType() int64`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *Argument) GetComponentTypeOk() (*int64, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *Argument) SetComponentType(v int64)`

SetComponentType sets ComponentType field to given value.


### GetStream

`func (o *Argument) GetStream() int64`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *Argument) GetStreamOk() (*int64, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *Argument) SetStream(v int64)`

SetStream sets Stream field to given value.


### GetRelationType

`func (o *Argument) GetRelationType() int64`

GetRelationType returns the RelationType field if non-nil, zero value otherwise.

### GetRelationTypeOk

`func (o *Argument) GetRelationTypeOk() (*int64, bool)`

GetRelationTypeOk returns a tuple with the RelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationType

`func (o *Argument) SetRelationType(v int64)`

SetRelationType sets RelationType field to given value.


### GetFilter

`func (o *Argument) GetFilter() StsEventStreamFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Argument) GetFilterOk() (*StsEventStreamFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Argument) SetFilter(v StsEventStreamFilter)`

SetFilter sets Filter field to given value.


### GetQueryView

`func (o *Argument) GetQueryView() int64`

GetQueryView returns the QueryView field if non-nil, zero value otherwise.

### GetQueryViewOk

`func (o *Argument) GetQueryViewOk() (*int64, bool)`

GetQueryViewOk returns a tuple with the QueryView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryView

`func (o *Argument) SetQueryView(v int64)`

SetQueryView sets QueryView field to given value.


### GetMaxWindow

`func (o *Argument) GetMaxWindow() int64`

GetMaxWindow returns the MaxWindow field if non-nil, zero value otherwise.

### GetMaxWindowOk

`func (o *Argument) GetMaxWindowOk() (*int64, bool)`

GetMaxWindowOk returns a tuple with the MaxWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWindow

`func (o *Argument) SetMaxWindow(v int64)`

SetMaxWindow sets MaxWindow field to given value.

### HasMaxWindow

`func (o *Argument) HasMaxWindow() bool`

HasMaxWindow returns a boolean if a field has been set.

### GetDownsamplingMethod

`func (o *Argument) GetDownsamplingMethod() DownsamplingMethod`

GetDownsamplingMethod returns the DownsamplingMethod field if non-nil, zero value otherwise.

### GetDownsamplingMethodOk

`func (o *Argument) GetDownsamplingMethodOk() (*DownsamplingMethod, bool)`

GetDownsamplingMethodOk returns a tuple with the DownsamplingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownsamplingMethod

`func (o *Argument) SetDownsamplingMethod(v DownsamplingMethod)`

SetDownsamplingMethod sets DownsamplingMethod field to given value.

### HasDownsamplingMethod

`func (o *Argument) HasDownsamplingMethod() bool`

HasDownsamplingMethod returns a boolean if a field has been set.

### GetWindowingMethod

`func (o *Argument) GetWindowingMethod() WindowingMethod`

GetWindowingMethod returns the WindowingMethod field if non-nil, zero value otherwise.

### GetWindowingMethodOk

`func (o *Argument) GetWindowingMethodOk() (*WindowingMethod, bool)`

GetWindowingMethodOk returns a tuple with the WindowingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowingMethod

`func (o *Argument) SetWindowingMethod(v WindowingMethod)`

SetWindowingMethod sets WindowingMethod field to given value.

### HasWindowingMethod

`func (o *Argument) HasWindowingMethod() bool`

HasWindowingMethod returns a boolean if a field has been set.

### GetScript

`func (o *Argument) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *Argument) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *Argument) SetScript(v string)`

SetScript sets Script field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


