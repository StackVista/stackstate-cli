# StackPackConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **int64** |  | [optional] 
**Status** | **string** |  | 
**Error** | Pointer to [**StackPackError**](StackPackError.md) |  | [optional] 
**StackPackVersion** | **string** |  | 
**Config** | **map[string]interface{}** |  | 

## Methods

### NewStackPackConfiguration

`func NewStackPackConfiguration(status string, stackPackVersion string, config map[string]interface{}, ) *StackPackConfiguration`

NewStackPackConfiguration instantiates a new StackPackConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackPackConfigurationWithDefaults

`func NewStackPackConfigurationWithDefaults() *StackPackConfiguration`

NewStackPackConfigurationWithDefaults instantiates a new StackPackConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackPackConfiguration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackPackConfiguration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackPackConfiguration) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StackPackConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *StackPackConfiguration) GetLastUpdateTimestamp() int64`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *StackPackConfiguration) GetLastUpdateTimestampOk() (*int64, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *StackPackConfiguration) SetLastUpdateTimestamp(v int64)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *StackPackConfiguration) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetStatus

`func (o *StackPackConfiguration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StackPackConfiguration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StackPackConfiguration) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *StackPackConfiguration) GetError() StackPackError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StackPackConfiguration) GetErrorOk() (*StackPackError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StackPackConfiguration) SetError(v StackPackError)`

SetError sets Error field to given value.

### HasError

`func (o *StackPackConfiguration) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStackPackVersion

`func (o *StackPackConfiguration) GetStackPackVersion() string`

GetStackPackVersion returns the StackPackVersion field if non-nil, zero value otherwise.

### GetStackPackVersionOk

`func (o *StackPackConfiguration) GetStackPackVersionOk() (*string, bool)`

GetStackPackVersionOk returns a tuple with the StackPackVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPackVersion

`func (o *StackPackConfiguration) SetStackPackVersion(v string)`

SetStackPackVersion sets StackPackVersion field to given value.


### GetConfig

`func (o *StackPackConfiguration) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *StackPackConfiguration) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *StackPackConfiguration) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


