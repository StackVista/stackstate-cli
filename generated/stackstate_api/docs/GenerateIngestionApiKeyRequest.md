# GenerateIngestionApiKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **int64** |  | [optional] 

## Methods

### NewGenerateIngestionApiKeyRequest

`func NewGenerateIngestionApiKeyRequest(name string, ) *GenerateIngestionApiKeyRequest`

NewGenerateIngestionApiKeyRequest instantiates a new GenerateIngestionApiKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateIngestionApiKeyRequestWithDefaults

`func NewGenerateIngestionApiKeyRequestWithDefaults() *GenerateIngestionApiKeyRequest`

NewGenerateIngestionApiKeyRequestWithDefaults instantiates a new GenerateIngestionApiKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GenerateIngestionApiKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenerateIngestionApiKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenerateIngestionApiKeyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GenerateIngestionApiKeyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenerateIngestionApiKeyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenerateIngestionApiKeyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GenerateIngestionApiKeyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiration

`func (o *GenerateIngestionApiKeyRequest) GetExpiration() int64`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *GenerateIngestionApiKeyRequest) GetExpirationOk() (*int64, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *GenerateIngestionApiKeyRequest) SetExpiration(v int64)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *GenerateIngestionApiKeyRequest) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


