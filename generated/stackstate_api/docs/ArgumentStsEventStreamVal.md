# ArgumentStsEventStreamVal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Parameter** | **int64** |  | 
**Filter** | [**StsEventStreamFilter**](StsEventStreamFilter.md) |  | 

## Methods

### NewArgumentStsEventStreamVal

`func NewArgumentStsEventStreamVal(type_ string, parameter int64, filter StsEventStreamFilter, ) *ArgumentStsEventStreamVal`

NewArgumentStsEventStreamVal instantiates a new ArgumentStsEventStreamVal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgumentStsEventStreamValWithDefaults

`func NewArgumentStsEventStreamValWithDefaults() *ArgumentStsEventStreamVal`

NewArgumentStsEventStreamValWithDefaults instantiates a new ArgumentStsEventStreamVal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ArgumentStsEventStreamVal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArgumentStsEventStreamVal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArgumentStsEventStreamVal) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ArgumentStsEventStreamVal) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArgumentStsEventStreamVal) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArgumentStsEventStreamVal) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ArgumentStsEventStreamVal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *ArgumentStsEventStreamVal) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *ArgumentStsEventStreamVal) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *ArgumentStsEventStreamVal) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *ArgumentStsEventStreamVal) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetParameter

`func (o *ArgumentStsEventStreamVal) GetParameter() int64`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ArgumentStsEventStreamVal) GetParameterOk() (*int64, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ArgumentStsEventStreamVal) SetParameter(v int64)`

SetParameter sets Parameter field to given value.


### GetFilter

`func (o *ArgumentStsEventStreamVal) GetFilter() StsEventStreamFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ArgumentStsEventStreamVal) GetFilterOk() (*StsEventStreamFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ArgumentStsEventStreamVal) SetFilter(v StsEventStreamFilter)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


