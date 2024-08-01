# CheckLeaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** |  | 
**AgentData** | Pointer to [**AgentData**](AgentData.md) |  | [optional] 

## Methods

### NewCheckLeaseRequest

`func NewCheckLeaseRequest(apiKey string, ) *CheckLeaseRequest`

NewCheckLeaseRequest instantiates a new CheckLeaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckLeaseRequestWithDefaults

`func NewCheckLeaseRequestWithDefaults() *CheckLeaseRequest`

NewCheckLeaseRequestWithDefaults instantiates a new CheckLeaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *CheckLeaseRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CheckLeaseRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CheckLeaseRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetAgentData

`func (o *CheckLeaseRequest) GetAgentData() AgentData`

GetAgentData returns the AgentData field if non-nil, zero value otherwise.

### GetAgentDataOk

`func (o *CheckLeaseRequest) GetAgentDataOk() (*AgentData, bool)`

GetAgentDataOk returns a tuple with the AgentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentData

`func (o *CheckLeaseRequest) SetAgentData(v AgentData)`

SetAgentData sets AgentData field to given value.

### HasAgentData

`func (o *CheckLeaseRequest) HasAgentData() bool`

HasAgentData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


