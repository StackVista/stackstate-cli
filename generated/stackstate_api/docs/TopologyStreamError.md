# TopologyStreamError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | [**MessageLevel**](MessageLevel.md) |  | 
**Message** | **string** |  | 
**ExternalId** | Pointer to **string** |  | [optional] 

## Methods

### NewTopologyStreamError

`func NewTopologyStreamError(level MessageLevel, message string, ) *TopologyStreamError`

NewTopologyStreamError instantiates a new TopologyStreamError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyStreamErrorWithDefaults

`func NewTopologyStreamErrorWithDefaults() *TopologyStreamError`

NewTopologyStreamErrorWithDefaults instantiates a new TopologyStreamError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *TopologyStreamError) GetLevel() MessageLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *TopologyStreamError) GetLevelOk() (*MessageLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *TopologyStreamError) SetLevel(v MessageLevel)`

SetLevel sets Level field to given value.


### GetMessage

`func (o *TopologyStreamError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TopologyStreamError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TopologyStreamError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetExternalId

`func (o *TopologyStreamError) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TopologyStreamError) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TopologyStreamError) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *TopologyStreamError) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


